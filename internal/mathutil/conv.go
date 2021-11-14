package mathutil

func Avg2(a, b, c uint8) uint8 {
	return uint8((uint(a) + uint(b) + uint(c)) / 3)
}
func Avg(a, b, c uint8) uint8 {
	return uint8((uint32(a) + uint32(b) + uint32(c)) / 3)
}
func Lit(a, b, c uint8) uint8 {
	return uint8((19595*uint32(a) + 38470*uint32(b) + 7471*uint32(c) + 32768) >> 16)
}

func Lit2(a, b, c uint8) uint8 {
	return uint8(float64(a)*.299 + float64(b)*.587 + float64(c)*.114)
}
