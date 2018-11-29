package SchnorrIdp

import (
	"fmt"
	"testing"
)
// Test key generation//
func TestKeyGeneraton(t *testing.T) {
	result := KeyGeneraton()
	fmt.Println("Point Q is: ", result.x, result.y)
}

//Generate signatures//
func TestSignature(t *testing.T) {
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
}

//verify signatures//
func TestBatchVerification(t *testing.T) {
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
