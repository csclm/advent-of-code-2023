package main

// I need to implement all of this manually for some reason?

type SliceHeap[TElement any] struct {
	contents    *[]TElement
	elementLess func(TElement, TElement) bool
}

func (sh SliceHeap[TElement]) Len() int {
	return len(*sh.contents)
}

func (sh SliceHeap[TElement]) Swap(i int, j int) {
	(*sh.contents)[i], (*sh.contents)[j] = (*sh.contents)[j], (*sh.contents)[i]
}

func (sh SliceHeap[TElement]) Less(i int, j int) bool {
	return sh.elementLess((*sh.contents)[i], (*sh.contents)[j])
}

func (sh SliceHeap[TElement]) Push(x interface{}) {
	*sh.contents = append(*sh.contents, x.(TElement))
}

func (sh SliceHeap[TElement]) Pop() interface{} {
	result := (*sh.contents)[len(*sh.contents)-1]
	*sh.contents = (*sh.contents)[:len(*sh.contents)-1]
	return result
}
