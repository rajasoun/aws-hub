package manu

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*var (
	file string
)*/

func Createfile() {
	file, err := os.Create("flow.md")
	if err != nil {
		fmt.Println("error creating file", err)
	}
	fmt.Println("File created successfully", file)
	defer Writefile()
}

func Writefile() {
	val := "hello, world\n"
	data := []byte(val)

	file := ioutil.WriteFile("flow.md", data, 0)

	log.Printf("%v", file)

}
