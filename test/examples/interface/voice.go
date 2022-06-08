package spike

type Voice interface {
	SayHello() string
}

func NewVoice(lang string) Voice {
	switch lang {
	case "Spanish":
		return Spanish{}
	default:
		return English{}
	}
}
