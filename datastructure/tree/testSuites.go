package tree

import "algorithms/datastructure/interfaces"

type ComparableInt int

func (a *ComparableInt) GreaterThan(b interfaces.Comparable) bool {
	return (*a) > b.(ComparableInt)
}

func (a *ComparableInt) LessThan(b interfaces.Comparable) bool {
	return (*a) < b
}


func testSet(setPtr *interfaces.Set){
	set := *(setPtr)

	ints := []ComparableInt{1,3,5,7,9,10}
	for _,v := range ints {
		set.Insert(v)
	}
}


