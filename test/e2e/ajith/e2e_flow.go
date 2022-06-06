package ajith

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)
func CreateMarkdown()(*os.File,error){
	name := "e2e.md"
	fileOpetion  :=os.O_RDWR | os.O_CREATE
	filePermission := 0666
	logFile ,err :=os.OpenFile(name,fileOpetion,os.FileMode(filePermission))
	if err !=nil {

		log.Println("error in oepaning and creation %s err= %v",name,err)
		return nil ,err
	}
	return logFile, nil

}func flowmanager()
