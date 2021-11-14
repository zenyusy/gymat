package gymat

import (
	"sync"

	"github.com/zenyusy/gymat/internal/mathutil"
)

type Gm struct {
	X    int
	Y    int
	Data [][]uint8 // [height] [width], numpy style
}

// When built, Data[*][*] = 0
func BuildGm(x int, y int) *Gm {
	r := &Gm{
		X:    x,
		Y:    y,
		Data: make([][]uint8, 0, y),
	}
	for i := 0; i < y; i++ {
		r.Data = append(r.Data, make([]uint8, x))
	}
	return r
}

func BuildGm2(x int, y int) *Gm {
	r := &Gm{
		X:    x,
		Y:    y,
		Data: make([][]uint8, y),
	}
	for i := range r.Data {
		r.Data[i] = make([]uint8, x)
	}
	return r
}

func BuildGm3(x int, y int) *Gm {
	r := &Gm{
		X:    x,
		Y:    y,
		Data: make([][]uint8, y),
	}
	for i := 0; i < y; i++ {
		r.Data[i] = make([]uint8, x)
	}
	return r
}

func BuildGmFromC4(
	cPx []uint8, cStride int,
	cFull bool, cX1 int, cY1 int, cX2 int, cY2 int) *Gm {
	// NOTE: cX1, cY1, cX2, cY2 are INCLUDED & LEGAL
	// cPx = { r1,g1,b1,a1,  r2,g2,b2,a2,  ... }
	// cStride = cWidth * 4 (rgba)

	if cFull {
		cX1 = 0                    // INCLU
		cY1 = 0                    // INCLU
		cX2 = cStride / 4          // EXCLU
		cY2 = len(cPx)/cStride - 1 // INCLU
	} else {
		cX2++ // EXCLU
	}
	gX := cX2 - cX1
	gY := cY2 - cY1 + 1

	ret := BuildGm(gX, gY)

	step := mathutil.CalcStep(gY)

	var wg sync.WaitGroup
	wg.Add(mathutil.DivUp(gY, step))

	lftInclu := 0
	rgtExclu := step
	for lftInclu < gY {
		if rgtExclu > gY {
			rgtExclu = gY
		}

		go func(cStart int, gy1 int, gy2 int) {
			defer wg.Done()
			for j := gy1; j < gy2; j++ {
				c := cStart
				for i := 0; i < gX; i++ {
					ret.Data[j][i] = mathutil.Avg(cPx[c], cPx[c+1], cPx[c+2])
					c += 4
				}
				cStart += cStride
			}
		}(cStride*(cY1+lftInclu)+4*cX1, lftInclu, rgtExclu)

		lftInclu = rgtExclu
		rgtExclu = lftInclu + step
	}
	wg.Wait()

	return ret
}
