package bench

import (
	"goabc/routine"
	"math/rand"
	"testing"
)
const (
	size1000 = 1000
	size1w = 10000
	size10w = 100000
	size100w = 1000000
)

func BenchmarkQuickSort1000(b *testing.B) {
	array := make([]int,size1000)
	for i:= range array{
		array[i] = rand.Intn(100)
	}
	for i := 0; i < b.N; i++ {
		routine.QuickSort01(array)
	}
}

func BenchmarkQuickSort10000(b *testing.B) {
	array := make([]int,size1w)
	for i:= range array{
		array[i] = rand.Intn(100)
	}
	for i := 0; i < b.N; i++ {
		routine.QuickSort01(array)
	}
}

func BenchmarkQuickSort100000(b *testing.B) {
	array := make([]int,size10w)
	for i:= range array{
		array[i] = rand.Intn(100)
	}
	for i := 0; i < b.N; i++ {
		routine.QuickSort01(array)
	}
}

func BenchmarkQuickSort1000000(b *testing.B) {
	array := make([]int,size100w)
	for i:= range array{
		array[i] = rand.Intn(100)
	}
	for i := 0; i < b.N; i++ {
		routine.QuickSort01(array)
	}
}

func BenchmarkQuickSortS1000(b *testing.B) {
	array := make([]int,size1000)
	for i:= range array{
		array[i] = rand.Intn(100)
	}
	for i := 0; i < b.N; i++ {
		routine.Waitgroup.Add(1)
		routine.QuickSort02(array)
		routine.Waitgroup.Wait()
	}
}

func BenchmarkQuickSortS10000(b *testing.B) {
	array := make([]int,size1w)
	for i:= range array{
		array[i] = rand.Intn(100)
	}
	for i := 0; i < b.N; i++ {
		routine.Waitgroup.Add(1)
		routine.QuickSort02(array)
		routine.Waitgroup.Wait()
	}
}

func BenchmarkQuickSortS100000(b *testing.B) {
	array := make([]int,size10w)
	for i:= range array{
		array[i] = rand.Intn(100)
	}
	for i := 0; i < b.N; i++ {
		routine.Waitgroup.Add(1)
		routine.QuickSort02(array)
		routine.Waitgroup.Wait()
	}
}

func BenchmarkQuickSortS1000000(b *testing.B) {
	array := make([]int,size100w)
	for i:= range array{
		array[i] = rand.Intn(100)
	}
	for i := 0; i < b.N; i++ {
		routine.Waitgroup.Add(1)
		routine.QuickSort02(array)
		routine.Waitgroup.Wait()
	}
}