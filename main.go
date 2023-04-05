package main

import (
	"fmt"
	"math/rand"
	"time"
)

func qsort(arr []int, beg, end int) {
	if beg >= end {
		return
	}
	sep := arr[beg+rand.Intn(end-beg)]
	lscan, rscan := beg, end
	for lscan < rscan {
		for arr[lscan] < sep {
			lscan++
		}
		for arr[rscan] > sep {
			rscan--
		}
		arr[lscan], arr[rscan] = arr[rscan], arr[lscan]
	}
	qsort(arr, beg, rscan)
	qsort(arr, rscan+1, end)
}
func main() {
	n := 30
	rand.Seed(time.Now().Unix())
	var arr = rand.Perm(n)
	fmt.Println(arr)
	qsort(arr, 0, n-1)
	fmt.Println(arr)
}
