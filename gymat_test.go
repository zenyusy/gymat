package gymat

import "testing"

type bftype func(int, int) *Gm

func TestBuildGm(t *testing.T) {
	const (
		X = 2
		Y = 3
	)
	for _, bf := range []bftype{
		BuildGm,
		BuildGm2,
		BuildGm3,
	} {
		g := bf(X, Y)
		for i := 0; i < X; i++ {
			for j := 0; j < Y; j++ {
				if g.Data[j][i] != 0 {
					t.Errorf("%d, %d not 0 but %d", i, j, g.Data[i][j])
				}
			}
		}
		t.Log(g)
	}
}

const (
	buildGmBcX = 4567
	buildGmBcY = 5678
)

var resulttmp uint8

func benchmarkBuildGm(bf bftype, x int, y int, b *testing.B) {
	var r *Gm
	for n := 0; n < b.N; n++ {
		r = bf(x, y)
	}
	resulttmp = r.Data[0][0]
}
func BenchmarkBuildGm0(b *testing.B) { benchmarkBuildGm(BuildGm, buildGmBcX, buildGmBcY, b) }
func BenchmarkBuildGm2(b *testing.B) { benchmarkBuildGm(BuildGm2, buildGmBcX, buildGmBcY, b) }
func BenchmarkBuildGm3(b *testing.B) { benchmarkBuildGm(BuildGm3, buildGmBcX, buildGmBcY, b) }

func TestConvC4(t *testing.T) {
	const (
		GX = 17
		GY = 13
	)
	u8Line := make([]uint8, GX*GY) // 221
	for i := range u8Line {
		u8Line[i] = uint8(i) + 1
	} // 1->221
	c4Line := make([]uint8, GX*GY*4)
	for ig, ic := 0, 0; ig < GX*GY; ig, ic = ig+1, ic+4 {
		c4Line[ic] = u8Line[ig] - 1
		c4Line[ic+1] = u8Line[ig]
		c4Line[ic+2] = u8Line[ig] + 1
		c4Line[ic+3] = 0
	}

	gm := BuildGmFromC4(c4Line, GX*4, true, 1, 2, 3, 4)
	if gm.X != GX {
		t.Fatal("Xnum=", gm.X, "expect:", GX)
	}
	if gm.Y != GY {
		t.Fatal("Ynum=", gm.Y, "expect:", GY)
	}
	k := 0
	for j := 0; j < GY; j++ {
		for i := 0; i < GX; i++ {
			if gm.Data[j][i] != u8Line[k] {
				t.Fatal("y", j, "x", i, "(", gm.Data[j][i], ")!=k", k, "(", u8Line[k])
			}
			k++
		}
	}

	for tx1 := 0; tx1 < GX-1; tx1++ {
		for tx2 := tx1 + 1; tx2 < GX; tx2++ {
			for ty1 := 0; ty1 < GY-1; ty1++ {
				for ty2 := ty1 + 1; ty2 < GY; ty2++ {
					gm2 := BuildGmFromC4(c4Line, GX*4, false, tx1, ty1, tx2, ty2)
					gmXSpan := tx2 - tx1 + 1
					gmYSpan := ty2 - ty1 + 1
					if gm2.X != gmXSpan {
						t.Fatal(tx1, tx2, ty1, ty2, "Xnum=", gm2.X, "expect:", gmXSpan)
					}
					if gm2.Y != gmYSpan {
						t.Fatal(tx1, tx2, ty1, ty2, "Ynum=", gm2.Y, "expect:", gmYSpan)
					}
					for j := 0; j < gmYSpan; j++ {
						for i := 0; i < gmXSpan; i++ {
							if gm2.Data[j][i] != gm.Data[j+ty1][i+tx1] {
								t.Fatal(tx1, tx2, ty1, ty2, "y", j, "x", i, "(", gm2.Data[j][i], ")!=k", "(", gm.Data[j+ty1][i+tx1])
							}
						}
					}

				}
			}
		}
	}
}
