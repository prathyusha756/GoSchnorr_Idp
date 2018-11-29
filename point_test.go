package SchnorrIdp

import (
	"fmt"
	"math/big"
	"testing"
)
//Test the given curve is valid or not//
func TestCheckValidCurve(t *testing.T) {
	fmt.Println(CheckValidCurve())
}

//Find a y-coordinate of a point for given x-coordinate//
func TestFindY(t *testing.T) {
	var x big.Int
	x.SetString("3", 10)
	var p Point = FindY(&x)
	fmt.Println("Point p: ", p.x, p.y)
}

//Find point addition//
func TestPointAddition(t *testing.T) {
	var x big.Int
	x.SetString("3", 10)
	var p Point = FindY(&x)
	fmt.Println("Point p: ", p.x, p.y)

	var y big.Int
	y.SetString("10", 10)
	var q Point = FindY(&y)
	fmt.Println("Point q: ", q.x, q.y)

	var r Point = PointAddition(p, q)
	fmt.Println("Point addition: ", r.x, r.y)

}
//Find point doubling//
func TestPointDoubling(t *testing.T) {
	var x big.Int
	x.SetString("3", 10)
	var p Point = FindY(&x)
	fmt.Println("Point p: ", p.x, p.y)

	var r Point = PointDoubling(p)
	fmt.Println("Point doubling", r.x, r.y)

}

//find point multiplication//
func TestPointMultiplication(t *testing.T) {
	var x big.Int
	x.SetString("3", 10)
	var p Point = FindY(&x)
	fmt.Println("Point p: ", p.x, p.y)

	var n = big.NewInt(6)
	var r Point = PointMultiplication(p, n)
	fmt.Println("Point multiplication", r.x, r.y)

}