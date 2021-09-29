package core

// 点
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p *Point) Set(x, y int) {
	p.X = x
	p.Y = y
}
