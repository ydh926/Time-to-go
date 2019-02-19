package bench

import (
	"strconv"
	"testing"
)

//未指定容量
func Mapinit(){
	Roster := make(map[string]string,0)
	for i:=0;i<100;i++{
		Roster["bob"+strconv.Itoa(i)] = strconv.Itoa(i)
	}
}
func BenchmarkMapinit(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Mapinit()
	}
}

//指定容量
func Mapinit2(){
	Roster := make(map[string]string,100)
	for i:=0;i<100;i++{
		Roster["bob"+strconv.Itoa(i)] = strconv.Itoa(i)
	}
}
func BenchmarkMapinit2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Mapinit2()
	}
}