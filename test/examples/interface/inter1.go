package spike

type College interface {
	College() string
}

type Student struct {
	Name   string
	Course string
	id     int
}

func (e *Student) Student() (string, string, int) {
	return e.Name, e.Course, e.id
}
