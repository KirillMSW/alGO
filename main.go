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
func demoQsort() {
	n := 30
	rand.Seed(time.Now().Unix())
	var arr = rand.Perm(n)
	fmt.Println(arr)
	qsort(arr, 0, n-1)
	fmt.Println(arr)
}

func main() {
	t := Tree{}
	for i := 0; i < 20; i++ {
		toInsert := rand.Intn(100)
		fmt.Println("Insertion", toInsert)
		t.insert(toInsert)
		t.printTree()
	}
	//t.insert(5)
	//t.printTree()
	//t.insert(4)
	//t.printTree()
	//t.insert(10)
	//t.printTree()
	//t.insert(3)
	//t.printTree()
	//t.insert(23)
	//t.printTree()
	//t.insert(1)
	//t.insert(4)
	//t.insert(6)
	//t.insert(2)
	//t.insert(0)
	//t.printTree()
}
