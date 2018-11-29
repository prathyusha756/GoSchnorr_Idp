package SchnorrIdp

import (
	"fmt"
	"math/big"
	"strconv"
)

/*Here we are taking SPEC256k1 curve */
var a = big.NewInt(0)
var b = big.NewInt(7)


type Point struct {
	x, y *big.Int
}

//find order//
func GetOrder() *big.Int{
	var ord big.Int
	ord.SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337",10)
	return &ord
}

//Get prime number(modulo)//
func GetPrime() big.Int{
	var prime big.Int
	prime.SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663",10)
	return prime
}
/*check whether the given curve is valid by calculating 4*a^3+27*b^2 and it should not be equal to zero.
    Indicates that there are no repeated roots*/
func CheckValidCurve() bool {
	s0 := big.NewInt(0).Mul(big.NewInt(0).Mul(b, b), big.NewInt(27))
	s1 := big.NewInt(0).Mul(big.NewInt(0).Mul(big.NewInt(0).Mul(a, a), a), big.NewInt(4))
	s := big.NewInt(0).Add(s0, s1)

	if s == big.NewInt(0) {
		return false
	}
	return true
}

/*Here our curve is in the form  y^2=x^3+ax+b. Given X-coordinate of a point then find Y-coordinate.*/

func FindY(x *big.Int) Point {
	//var prime big.Int
	//prime.SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663",10)
	prime:= GetPrime()
	var y1,x0,x1,x2,x4 big.Int
	//var y *big.Int
	x0.Mul(a, x)
	x1.Add(&x0, b)
	x2.Mul(x, x)
	x2.Mul(&x2, x)
	x4.Add(&x1, &x2)
	y1.ModSqrt(&x4, &prime)
	y1.Sub(&prime, &y1)
	//y=y1
	return Point{x, &y1}
}

/*This method adds two points. Where slope m = (y2-y1)/(x2-x1); new point r = (x3,y3),
    x3 = m^2-(x1+x2), y3 = -m(x3-x1)-y1 */
func PointAddition(p Point, q Point) Point {

	var numerator, denominator, m, x3, y3 big.Int
	prime:= GetPrime()
	zero:=big.NewInt(0)

	if p.x.Cmp(q.x)==0 && p.y.Cmp(q.y)==0 {
		if p.y.Cmp(zero)==0 {
			return Point{big.NewInt(0), big.NewInt(0)}
		}
		return PointDoubling(p)
	}
	if p.x.Cmp(q.x)==0 {
		return Point{big.NewInt(0), big.NewInt(0)}
	}
	if q.x.Cmp(zero)==0 && q.y.Cmp(zero)==0 {
		return Point{p.x, p.y}
	}
	if p.x.Cmp(zero)==0 && p.y.Cmp(zero)==0 {
		return Point{q.x, q.y}
	}

	numerator.Mod(numerator.Sub(q.y, p.y), &prime)

	denominator.Mod(denominator.Sub(q.x, p.x), &prime)
	denominator.ModInverse(&denominator, &prime)

	m.Mul(&numerator, &denominator)
	x3.Mul(&m, &m)
	x3.Mod(x3.Sub(&x3, big.NewInt(0).Add(p.x, q.x)), &prime)
	y3.Mod(y3.Sub(&x3, p.x), &prime)
	y3.Mul(&y3, big.NewInt(0).Mul(&m, big.NewInt(-1)))
	y3.Mod(y3.Sub(&y3, p.y), &prime)
	r := Point{&x3, &y3}
	return r

}

/* For pointDoubling slope m = (3*x1^2 + a)/2y1, new point r=(x3,y3),
    x3 = m^2-(x1+x2), y3 = -m(x3-x1)-y1  */
func PointDoubling(p Point) Point {
	var numerator, denominator, m, x3, y3 big.Int
	prime:= GetPrime()

	numerator.Mul(numerator.Mul(p.x, p.x), big.NewInt(3))
	numerator.Mod(numerator.Add(&numerator, a), &prime)

	denominator.Mod(denominator.Mul(p.y, big.NewInt(2)), &prime)
	denominator.ModInverse(&denominator, &prime)
	m.Mul(&numerator, &denominator)
	x3.Mul(&m, &m)
	x3.Mod(x3.Sub(&x3, big.NewInt(0).Add(p.x, p.x)), &prime)
	y3.Mod(y3.Sub(&x3, p.x), &prime)
	y3.Mul(&y3, big.NewInt(0).Mul(&m, big.NewInt(-1)))
	y3.Mod(y3.Sub(&y3, p.y), &prime)
	r := Point{&x3, &y3}
	return r
}

/*PointMultiplication is the scalar multiplication of the given point. Let say p = (2,3), then find out
[2]p, [3]p etc.
We are using double and add algorithm for point multiplication.
For more details you can visit https://sefiks.com/2016/03/27/double-and-add-method/
* */

func PointMultiplication(p Point, n *big.Int) Point {

	bigStr := fmt.Sprintf("%b", n)
	//fmt.Println(bigStr)
	var i int
	var r Point = p
	for i = 1; i < len(bigStr); i++ {

		currentBit, err := strconv.Atoi(string(bigStr[i : i+1]))
		fmt.Println(currentBit)
		fmt.Println(err)
		r = PointAddition(r, r)
		if r.y == big.NewInt(0) {
			j := 2 * (i + 1)
			if int64(j) == n.Int64() {
				return Point{big.NewInt(0), big.NewInt(0)}
			}
		}
		if currentBit == 1 {
			r = PointAddition(r, p)
			if r.y == big.NewInt(0) {
				j := 2 * (i + 1)
				if int64(j) == n.Int64() {
					return Point{big.NewInt(0), big.NewInt(0)}
				}
			}
		}
	}
	return r
}

//func main() {
//	fmt.Println(CheckValidCurve())
//	var x big.Int
//	x.SetString("3", 10)
//	var p Point = FindY(x)
//	fmt.Println("Point p: ", p.x, p.y)
//
//	var y big.Int
//	y.SetString("10", 10)
//	var q Point = FindY(y)
//	fmt.Println("Point q: ", q.x, q.y)
//
//	var l Point= Point{big.NewInt(0),big.NewInt(0)}
//	var r Point = PointAddition(p, l)
//	fmt.Println("Point addition: ", r.x, r.y)
//	//fmt.Println("Point addition",)
//
//	var r1 Point = PointDoubling(p)
//	fmt.Println("Point doubling", r1.x, r1.y)
//
//	var n = big.NewInt(6)
//	var r2 Point = PintMultiplication(p, n)
//	fmt.Println("Point multiplication", r2.x, r2.y)
//
//}
