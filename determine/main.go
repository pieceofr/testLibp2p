package main

import (
	"fmt"
)

func main() {
	determineConnections(1, 2)
	fmt.Println(" ---- New Test ----")
	determineConnections(0, 2)
	fmt.Println(" ---- New Test ----")
	determineConnections(0, 3)
	fmt.Println(" ---- New Test ----")
	determineConnections(1, 3)
	fmt.Println(" ---- New Test ----")
	determineConnections(1, 5)
	fmt.Println(" ---- New Test ----")
}

func determineConnections(index, count int) {
	// various increment values
	e := count / 8
	q := count / 4
	h := count / 2

	jump := 3      // to deal with N3/P3 and too few nodes
	if count < 4 { // if insufficient
		jump = 1 // just duplicate N1/P1
	}

	names := [11]string{
		"N1",
		"N3",
		"X1",
		"X2",
		"X3",
		"X4",
		"X5",
		"X6",
		"X7",
		"P1",
		"P3",
	}

	// compute all possible offsets
	// if count is too small then there will be duplicate offsets
	var n [11]int
	n[0] = index + 1             // N1 (+1)
	n[1] = index + jump          // N3 (+3)
	n[2] = e + index             // X⅛
	n[3] = q + index             // X¼
	n[4] = q + e + index         // X⅜
	n[5] = h + index             // X½
	n[6] = h + e + index         // X⅝
	n[7] = h + q + index         // X¾
	n[8] = h + q + e + index     // X⅞
	n[9] = index + count - 1     // P1 (-1)
	n[10] = index + count - jump // P3 (-3)

	u := -1
deduplicate:
	for i, v := range n {
		if v == index || v == u {
			//fmt.Printf("deduplicate[%d]:index:%d v:%d u:%d\n", i, index, v, u)
			continue deduplicate
		}
		u = v
		if v >= count {
			v -= count
		}
		fmt.Printf(" Node Index:%d  Tree Index :%d name:%s\n", index, v, names[i])

	}
}
