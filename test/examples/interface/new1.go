package spike

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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
func NewAudio(lang string) Audio {
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

func (p *player) PlayingGames() string {
	mk := player{}
	mk.age = 25
	mk.name = "ajit kumar"
	mk.perimeter.height = 20
	mk.perimeter.width = 30
	mk.langu = "English"

	new := NewAudio(mk.langu)
	return new.SayHello()
}

/* writeing a function which will write in a file
and if the file is not created
it will create the file
*/

func CreatingFile() *os.File {
	filename := "Demo.txt"

	NewFile, error1 := os.Create(filename)
	if error1 != nil {
		log.Fatal(error1)
	}

	/*if err != nil {
		log.Println("Error whlie creating the file %s Err = %v", NewFile)
		return nil, err
	}
	*/
	//info := os.Stat(NewFile)

	//data := []byte(val)
	/*err1 := ioutil.WriteFile(filename, data, 0644)
	if err1 != nil {
		log.Fatal(err1)
	}*/

	/*
		leng, err := NewFile.Read()
		if err != nil {
			log.Println("Error in Writing inside the file ")
			return nil, err
		}
	*/
	return NewFile
}

// moking

type User struct {
	Name     string
	Email    string
	UserName string
}
type precheck interface {
	userExists(string) bool
}
type regiCheck struct{}

func (r regiCheck) userExists(email string) bool {
	return UserAvilable(email)
}

var regCond precheck

func init() {
	regCond = regiCheck{}
}
func NewUser(u User) error {

	avilable := regCond.userExists(u.Email)
	if avilable {
		return fmt.Errorf("email is already '%s'avilable", u.Email)
	}
	log.Println(u.Name)
	return nil

}

// working with json

type Salary struct {
	Basic float32
	tax   float64
	total float64
}

type Employee struct {
	FirstName, LastName, Email string
	age                        int
	MonthlySalary              []Salary
}

func EmployeeSalary() (*os.File, bool) {
	set := false
	filename := "demo.json"
	newfile, _ := os.Create(filename)

	data := Employee{
		FirstName: "ajit",
		LastName:  "kumar",
		Email:     "ajithkumar.sinha@srsconsultinginc.com",
		age:       26,
		MonthlySalary: []Salary{
			{
				Basic: 10000,
				tax:   1000,
				total: 11000,
			},
		},
	}
	file, _ := json.MarshalIndent(data, "", "")
	erro := ioutil.WriteFile(newfile.Name(), file, 0644)
	if erro != nil {
		log.Fatal(erro)
	}

	/*_, err := os.ReadFile(newfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	*/
	fileinfo, _ := os.Stat(newfile.Name())
	log.Println(fileinfo)
	lenghtOfFile := fileinfo.Size()
	if lenghtOfFile != 0 {
		set = true

	}
	return newfile, set

}
