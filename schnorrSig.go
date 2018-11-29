package SchnorrIdp

import (

	"fmt"
	"math/big"
	"crypto/sha1"
	"crypto/rand"
)

var p,q Point
var c =big.NewInt(0)
/* step 1)Choose a random point p on elliptical curve
      *     2)Choose a random integer 'a' from range[1, r],
      *       where r is the order of point p.
      *     3) calculate Q=[a]p
      *     4) Output: public key Pk:(p,[a]p) , secret key Sk=(a, Pk) */
func KeyGeneraton() Point {
	var x big.Int
	//var Reader io.Reader
	x.SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)
	P := FindY(&x)
	Q := Point{big.NewInt(0), big.NewInt(0)}
	a:=big.NewInt(0)
	var err error
	curveOrder := GetOrder()
	for Q.x.Cmp(big.NewInt(0))==0 && Q.y.Cmp(big.NewInt(0))==0 {
		if !(P.x == big.NewInt(0)) || !(P.x == big.NewInt(0)) {

			a, err = rand.Int(rand.Reader, curveOrder)
			if err != nil {
				//error handling
			}

			fmt.Println("Random value a: ", a)

			Q = PointMultiplication(P, a)

		}
	}
	p = P
	q = Q
	c = a
	return Q
}

type PreSigObj struct {
	R Point
	k *big.Int
}

type SigObj struct{
	R Point
	S *big.Int
	m string
}

/* Calculating Point R :
    * 1)Choose a random integer 'k' from range[1, r],where r is the order of point p.
      2) Calculate Point R= [k]p */
  func OffLineCalculation() PreSigObj {
	  curveOrder := GetOrder()
	  R := Point{big.NewInt(0), big.NewInt(0)}
	   k:= big.NewInt(0)
	  var err error

	  for R.x.Cmp(big.NewInt(0))==0 && R.y.Cmp(big.NewInt(0))==0 {
		  if !(p.x == big.NewInt(0)) || !(p.x == big.NewInt(0)) {

			  k,err = rand.Int(rand.Reader, curveOrder)
			  if err != nil {
				//error handling
			  }


			  fmt.Println("Random value k: ", k)
			  R = PointMultiplication(p, k)

		  }
	  }
	  result:= PreSigObj{R,k}

	  return result
  }


 func Signature(obj PreSigObj, m string) SigObj {
    ae:=big.NewInt(0)
    s:=big.NewInt(0)
    e:=big.NewInt(0)
    R:=obj.R
 	rx:= obj.R.x.String()
    ry:=obj.R.y.String()
    concatenation:=m+rx+ry
    h:=sha1.New()
    h.Write([]byte(concatenation))
    bs:=h.Sum([]byte{})
    e.SetBytes(bs)
	 e.Mod(e, GetOrder())
    ae.Mul(c,e)
   // ae.Mod(ae,getOrder())
    s.Add(ae,obj.k)
    s.Mod(s,GetOrder())
    result:=SigObj{R,s,m}
     fmt.Println("e at sig",e)
    return result
 }

/*In batch verification we have given list of signatures. In order to validate all signatures at once
* i)Get s value from each signature in the list and add them to get 'S'
* ii)Calculate  e = H(message || R) from each signature and sum to get 'E'
* iii)Get point R from each signature and add to get Rs
 *iv) Finally calculate, if Rs+[E]Q = [S]p then output true else output false.*/
func BatchVerification(sigList [500]SigObj) bool{
//func batchVerification(sigList [1]SigObj) bool{
	var sig SigObj
	var m string
	var R Point

	Rs:=Point{big.NewInt(0),big.NewInt(0)}
	S:=big.NewInt(0)
	E:=big.NewInt(0)
	e:= big.NewInt(0)
	for i:=0; i<len(sigList); i++{
		sig=sigList[i]
		m=sig.m
		R=sig.R
		S.Add(S,sig.S)
		S.Mod(S, GetOrder())

		Rs=PointAddition(Rs,R)
		rx:= sig.R.x.String()
		ry:= sig.R.y.String()
		concatenation:=m+rx+ry
		h:=sha1.New()
		h.Write([]byte(concatenation))
		bs:=h.Sum([]byte{})
		e.SetBytes(bs)
		e.Mod(e, GetOrder())
		E.Add(E, e)
		E.Mod(E, GetOrder())
	}
    EtimesQ:=PointMultiplication(q, E)
    leftSide:=PointAddition(Rs,EtimesQ)
    rightSide:=PointMultiplication(p,S)

    if leftSide.x.Cmp(rightSide.x)==0 && leftSide.y.Cmp(rightSide.y)==0{
    	return true
	} else {
		return false
	}
}



func main() {
	result := KeyGeneraton()
	fmt.Println("Point Q is: ", result.x, result.y)
    const size int=500
	var arr [size]PreSigObj
	var arr1 [size]SigObj
	for i:=0; i<size; i++ {
		result1 := OffLineCalculation()
		arr[i]=result1
	}

	for i:=0; i<size;i++ {
		result1:=arr[i]
		result2 := Signature(result1, "hello world")
		arr1[i]=result2
		fmt.Println("Point R", result2.R.x, result2.R.y)
		fmt.Println("s, m are: ", result2.S, result2.m)
	}

	fmt.Println("All signatures are valid: ", BatchVerification(arr1))

}
