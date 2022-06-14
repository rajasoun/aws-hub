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
	langu     string
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

func (p *player) playingGames() string {
	mk := player{}
	mk.age = 25
	mk.name = "ajit kumar"
	mk.perimeter.height = 20
	mk.perimeter.width = 30
	mk.langu = "English"

	new := newaudio(mk.langu)
	return new.SayHello()
}
