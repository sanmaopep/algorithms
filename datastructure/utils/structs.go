package utils

/**
 * a implementation of Comparable for test use
 */
type CmpInt struct {
	Num int
}

func (a *CmpInt) Compare(b interface{}) int {
	return a.Num - b.(CmpInt).Num
}

