package gymat

import (
	"github.com/zenyusy/gymat/internal/mathutil"
	"sync"
)

// /SpDifRes/
type DifRes struct {
	X int
	Y int
	V int
}

func SqrDif(tpl *Gm, src *Gm, thres int) DifRes {
	if tpl.X > src.X || tpl.Y > src.Y || (tpl.X == src.X && tpl.Y == src.Y) {
		return DifRes{-1, -1, -1} // SpDifRes: invalid input
	}
	if thres < 0 {
		thres = 65025*tpl.Y*tpl.X + 1
	}
	xMov := src.X - tpl.X + 1
	yMov := src.Y - tpl.Y + 1

	argSlot := make([]int, 4)

	xLarger := xMov > yMov // assume false
	movArg := yMov
	movArgIdx := 2
	movArgIdx2 := 3
	argSlot[1] = xMov // +[0]=0 fixed
	if xLarger {
		movArg = xMov
		movArgIdx = 0
		movArgIdx2 = 1
		argSlot[3] = yMov // +[2]=0 fixed
	}

	var wg sync.WaitGroup
	wg.Add(mathutil.CNUM)
	resSlot := make([]DifRes, mathutil.CNUM)

	for i := range resSlot {
		argSlot[movArgIdx] = i * movArg / mathutil.CNUM
		argSlot[movArgIdx2] = (i + 1) * movArg / mathutil.CNUM
		if argSlot[movArgIdx] == argSlot[movArgIdx2] {
			wg.Done()
			continue
		}

		go func(lft int, rgt int, top int, btm int, i int) {
			res := DifRes{lft, top, thres}
			for srcY := top; srcY < btm; srcY++ {
				for srcX := lft; srcX < rgt; srcX++ {
					diff := 0

					// square diff begin
				Lp:
					for y := 0; y < tpl.Y; y++ {
						for x := 0; x < tpl.X; x++ {
							v1 := tpl.Data[y][x]
							v2 := src.Data[y+srcY][x+srcX]
							if v1 > v2 {
								d := int(v1 - v2)
								diff += d * d
							} else {
								d := int(v2 - v1)
								diff += d * d
							}
							if diff >= res.V {
								break Lp
							}
						}
					} // Lp end

					if diff < res.V {
						res.X = srcX
						res.Y = srcY
						res.V = diff
					}
				}
			}
			resSlot[i] = res
			wg.Done()
		}(argSlot[0], argSlot[1], argSlot[2], argSlot[3], i)
	}
	wg.Wait()

	ret := DifRes{-9, -9, thres} // SpDifRes: all fail thres
	for _, r := range resSlot {
		if r.V < ret.V {
			ret.X = r.X
			ret.Y = r.Y
			ret.V = r.V
		}
	}
	return ret
}
