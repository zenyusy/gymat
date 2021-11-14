package gymat

import "testing"

func badAvg(a *[]uint8) uint8 {
	return ((*a)[0] + (*a)[1] + (*a)[2] + (*a)[3]) / 4
}

func TestGmAppender(t *testing.T) {
	g := BuildGm(2, 3)
	ga := BuildGmAppender(g, 0, 0)
	s := []uint8{ // width 2, height 3
		7, 9, 11, 5, 12, 16, 14, 18,
		15, 17, 19, 21, 24, 26, 28, 30,
		31, 33, 35, 37, 52, 54, 56, 58}
	take := s[0:4:4]
	left := s[4:]
	for len(left) > 0 {
		ga.Append(badAvg(&take))
		take = left[0:4:4]
		left = left[4:]
	}
	if len(take) > 0 {
		ga.Append(badAvg(&take))
	}

	ga.Close()

	for _, tt := range []struct {
		y int
		x int
		v uint8
	}{
		{0, 0, 8},
		{0, 1, 15},
		{1, 0, 18},
		{1, 1, 27},
		{2, 0, 34},
		{2, 1, 55},
	} {
		if g.Data[tt.y][tt.x] != tt.v {
			t.Errorf("x=%d, y=%d not %d but %d", tt.x, tt.y, tt.v, g.Data[tt.y][tt.x])
		}
	}
}
