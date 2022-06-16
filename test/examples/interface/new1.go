package spike

import (
	"fmt"
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

	NewFile, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)

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
	Name     string //`json:"name"`
	Email    string //`json:"email"`
	UserName string //`json:"user_name"`
}

func NewUser(u User) error {

	avilable := UserAvilable(u.Email)
	if avilable {
		return fmt.Errorf("email is already '%s'avilable", u.Email)
	}
	log.Println(u.Name)
	return nil

}
