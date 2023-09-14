package point

type Point struct {
	x float64
	y float64
}

func New() *Point {
	return &Point{}
}

func (p *Point) SetX(val float64) {
	p.x = val
}
func (p *Point) SetY(val float64) {
	p.y = val
}
func (p *Point) GetX() float64 {
	return p.x
}
func (p *Point) GetY() float64 {
	return p.y
}
