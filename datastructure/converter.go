package ds

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fatih/structs"
)

type DataStructure struct{}

func (data *DataStructure) StructToJSON(result *struct{}) error {
	json, err := json.Marshal(&result)
	if err != nil {
		log.Printf("struct to JSON err = %v", err)
		return err
	}
	fileHandler, err := os.Create("data.json")
	if err != nil {
		log.Printf("file creation failed err = %v", err)
		return err
	}
	defer fileHandler.Close()
	_, writeErr := fileHandler.WriteString(string(json))
	if writeErr != nil {
		log.Printf("file write failed err = %v", err)
		return err
	}
	return nil
}

// StructToMap to convert struct into a map[string]interface{}.
func (data *DataStructure) StructToMap(ds interface{}) map[string]interface{} {
	s := structs.New(ds)
	m := s.Map()
	return m
}
