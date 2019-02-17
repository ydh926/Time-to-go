package bench

import "testing"

var size = 10000

func Loops1() {
	for i := 0; i < size; i++ {
		print(i)
	}
}

func Loops2(s []string) {
	for i := range s {
		print(i)
	}
}

func BenchmarkLoops1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loops1()
	}
}

func BenchmarkLoops2(b *testing.B) {
	slice1 := make([]string, size)
	for i := 0; i < b.N; i++ {
		Loops2(slice1)
	}
}
