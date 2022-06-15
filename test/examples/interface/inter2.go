package spike

type SalaryCalc interface {
	Salary() int
}

type Company struct {
	Name     string
	Location string
	basicpay int
	pf       int
}

func (c *Company) SalaryCalc() (string, string, int, int) {
	return c.Name, c.Location, c.basicpay, c.pf
}
