package main

import "sort"

func main() {

	a := []int{-1, 2, 3, 4, 5, 8}
	p := sort.Search(len(a), func(i int) bool { return a[i] >= 3 })
	println(p)
}
