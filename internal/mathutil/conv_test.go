package mathutil

import (
	"math/rand"
	"testing"
)

var resulttmp uint8

func TestLit(t *testing.T) {
	for i := 0; i < 0x12345; i++ {
		x := uint8(rand.Intn(255))
		y := uint8(rand.Intn(255))
		z := uint8(rand.Intn(255))
		r1 := Lit(x, y, z)
		r2 := Lit2(x, y, z)
		gg := false
		if r1 > r2 {
			if r1-r2 != 1 {
				gg = true
			}
		} else if r1 < r2 {
			if r2-r1 != 1 {
				gg = true
			}
		}
		if gg {
			t.Fatal("x", x, "y", y, "z", z)
		}
	}
}

func BenchmarkAvg0(b *testing.B) {
	var r uint8 = 10
	for n := 0; n < b.N; n++ {
		r = Avg(11, 83, 201)
	}
	resulttmp = r
}

func BenchmarkAvg2(b *testing.B) {
	var r uint8 = 10
	for n := 0; n < b.N; n++ {
		r = Avg2(11, 83, 201)
	}
	resulttmp = r
}

func BenchmarkLit0(b *testing.B) {
	var r uint8 = 10
	for n := 0; n < b.N; n++ {
		r = Lit(11, 83, 201)
	}
	resulttmp = r
}

func BenchmarkLit2(b *testing.B) {
	var r uint8 = 10
	for n := 0; n < b.N; n++ {
		r = Lit2(11, 83, 201)
	}
	resulttmp = r
}
