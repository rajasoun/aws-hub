package spike

type perimeter struct {
	width  float64
	height float64
}

type Audio interface {
	SayHello() string
	//perimeter()
}

func (p *perimeter) Perimeter() float64 {
	return 2 * (p.width + p.height)
}
func newaudio(lang string) Audio {
	if lang == "Hello" {
		return English{}
	}
	return Spanish{}

}
