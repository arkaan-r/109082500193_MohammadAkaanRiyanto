package main

import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat  titik
	radius int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c lingkaran, p titik) bool {
	if jarak(c.pusat, p) <= float64(c.radius) {
		return true
	}
	return false
}

func main() {
	var l1, l2 lingkaran
	var p titik
	var d1, d2 bool

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&p.x, &p.y)

	d1 = didalam(l1, p)
	d2 = didalam(l2, p)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}