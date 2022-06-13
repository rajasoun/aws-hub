package spike

type perimeter struct {
	width  float64
	height float64
}

func (p *perimeter) Perimeter() float64 {
	return 2 * (p.width + p.height)
}
