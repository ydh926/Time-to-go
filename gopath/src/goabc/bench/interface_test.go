package bench

import (
	"goabc/interfacex"
	"testing"
)

func BenchmarkDoDebug(b *testing.B) {
	bob := &interfacex.Bob{}
	for i := 0; i < b.N; i++ {
		interfacex.DoDebug(bob)
	}
}

func BenchmarkDoDebug2(b *testing.B) {
	bob := &interfacex.Bob{}
	for i := 0; i < b.N; i++ {
		interfacex.DoDebug2(bob)
	}
}
