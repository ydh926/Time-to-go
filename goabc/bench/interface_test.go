package bench

import (
	"testing"
)

func BenchmarkDoDebug(b *testing.B) {
	bob := &Bob{}
	for i := 0; i < b.N; i++ {
		DoDebug(bob)
	}
}

func BenchmarkDoDebug2(b *testing.B) {
	bob := &Bob{}
	for i := 0; i < b.N; i++ {
		DoDebug2(bob)
	}
}

type Person interface {
	Say()string
}

type Programer interface {
	Person
	Debug()
}

type Bob struct{
}
func (*Bob)Say()string{
	return "I am Bob"
}
func (*Bob)Debug(){
	println("I can debug")
}

func DoDebug(p Programer){
	switch m:= p.(type){
	case *Bob:
		print("This is Bob")
		m.Debug()
	default:
		print("I don't konw who")
	}
}

func DoDebug2(p Programer){
	switch p.(type){
	case *Bob:
		print("This is Bob")
	default:
		print("I don't konw who")
	}
	p.Debug()
}