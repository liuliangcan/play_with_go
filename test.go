package main

import "sort"

func main() {

	p := NewSortedList()
	p.Add(1)
	p.Add(2)
	p.Add(14)
	p.Add(5)
	p.Add(7)
	p.Add(110)
	p.Add(16)

	println(p.Small)
	x := []int{1, 2, 3, 4, 5}
	println(x)
}

type SortedList struct {
	Small []int
	Large []int
}

func NewSortedList() SortedList {
	//small := []int{}
	var large []int
	return SortedList{[]int{}, large}
}

func (u *SortedList) Add(v int) {
	if len(u.Small) > 1500 {
		u.Large = append(u.Large, u.Small...)
		u.Small = []int{}
		sort.Ints(u.Large)
	}
	if len(u.Small) == 0 || u.Small[len(u.Small)-1] <= v {
		u.Small = append(u.Small, v)
	} else if u.Small[0] >= v {
		u.Small = append([]int{v}, u.Small...)
	} else {
		p := sort.Search(len(u.Small), func(i int) bool { return u.Small[i] >= v })
		t := u.Small[p:]
		u.Small = append(u.Small[:p], v)
		u.Small = append(u.Small[:p], t...)
	}
}

func (u SortedList) BisectLeft(v int) int {
	return sort.Search(len(u.Small), func(i int) bool { return u.Small[i] >= v }) + sort.Search(len(u.Large), func(i int) bool { return u.Large[i] >= v })
}

func (u SortedList) Len() int {
	return len(u.Small) + len(u.Large)
}
