package test

import (
	"data/lodash"
	"testing"
)

func BenchmarkNewSliceInt(b *testing.B)  {
	a:= []int{1,2,3,4,5}
	for i:=0;i<b.N;i++ {
		lodash.NewSliceInt(a)
	}
}

func BenchmarkIsEmpty(b *testing.B)  {
	a:= []int{1,2,3,4,5}
	sliceInt := lodash.NewSliceInt(a)
	sliceInt.IsEmpty()
}

func BenchmarkChunk(b *testing.B)  {
	a:= []int{1,2,3,4,5}
	sliceInt := lodash.NewSliceInt(a)
	sliceInt.Chunk(2)
}

func BenchmarkIsIn(b *testing.B)  {
	a:= []int{1,2,3,4,5}
	sliceInt := lodash.NewSliceInt(a)
	sliceInt.In(2)
}
