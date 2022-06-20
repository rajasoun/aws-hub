package integration_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func FromJSONFile(filePath string, ds interface{}) error {
	file, readErr := ioutil.ReadFile(filePath)
	if readErr != nil {
		log.Printf("Err ioutil.ReadFile() = %v ", readErr)
		return readErr
	}
	err := json.Unmarshal(file, &ds)
	if err != nil {
		log.Printf("Err json.Unmarshal() = %v ", err)
		return err
	}
	return nil
}

func ToJSONFile(filePath string, ds interface{}) error {
	jsonData, jsonErr := json.Marshal(&ds)
	if jsonErr != nil {
		log.Printf("Err json.Marshal() = %v ", jsonErr)
		return jsonErr
	}

	fileHandler, fileErr := os.Create(filePath)
	if fileErr != nil {
		log.Printf("Err os.OpenFile() = %v ", fileErr)
		return fileErr
	}
	defer fileHandler.Close()

	_, writeErr := fileHandler.WriteString(string(jsonData))
	if writeErr != nil {
		log.Printf("Err fileHandler.WriteString() = %v ", writeErr)
		return writeErr
	}
	return nil
}
