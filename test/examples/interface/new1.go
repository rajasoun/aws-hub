package spike

type perimeter struct {
	width  float64
	height float64
}

type Audio interface {
	SayHello() string
	//perimeter()
}
type player struct {
	name      string
	age       int64
	perimeter perimeter
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

func PlayerDetails() *player {
	new := player{
		name: "ajit kumar",
		age:  25,
		perimeter: perimeter{
			height: 10,
			width:  20,
		},
	}

	return &new
}
