package function

type Institution interface {
	Institution() string
}

type Students struct {
	Name    string
	Class   string
	Section string
}

func (a *Students) Students() (string, string, string) {
	return a.Name, a.Class, a.Section
}
