package spike

type perimeter struct {
	wi float64
	he float64
}

func (p *perimeter) Perimeter() float64 {
	return 2 * (p.wi + p.he)
}
