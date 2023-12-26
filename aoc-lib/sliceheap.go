package aoc

type SliceHeap[TElement any] struct {
	Contents    []TElement
	ElementLess func(TElement, TElement) bool
}

func (sh *SliceHeap[TElement]) Len() int {
	return len(sh.Contents)
}

func (sh *SliceHeap[TElement]) Swap(i int, j int) {
	sh.Contents[i], sh.Contents[j] = sh.Contents[j], sh.Contents[i]
}

func (sh *SliceHeap[TElement]) Less(i int, j int) bool {
	return sh.ElementLess(sh.Contents[i], sh.Contents[j])
}

func (sh *SliceHeap[TElement]) Push(x interface{}) {
	sh.Contents = append(sh.Contents, x.(TElement))
}

func (sh *SliceHeap[TElement]) Pop() interface{} {
	result := sh.Contents[len(sh.Contents)-1]
	sh.Contents = sh.Contents[:len(sh.Contents)-1]
	return result
}
