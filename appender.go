package gymat

type GmAppender struct {
	*Gm
	// Gm.Data[ay][ax] to append
	ax int
	ay int
}

func BuildGmAppender(g *Gm, x int, y int) *GmAppender {
	// enter +mode
	// usually x=y=0
	return &GmAppender{g, x, y}
}

func (ga *GmAppender) Close() {
	ga.Gm = nil
}

func (ga *GmAppender) Append(p uint8) {
	ga.Gm.Data[ga.ay][ga.ax] = p
	ga.ax++
	if ga.ax == ga.Gm.X {
		ga.ay++
		ga.ax = 0
	}
}
