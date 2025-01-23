package slices

var s byte

func SliceIterator(n []byte) byte {
	for _, elem := range n {
		s = elem
		println(s)
	}
	return s
}
