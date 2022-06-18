package ds

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fatih/structs"
)

type DataStructure struct {
	fileName    string
	fileCreator func(name string) (*os.File, error)
}

func New() DataStructure {
	ds := DataStructure{
		fileName:    "data.json",
		fileCreator: os.Create,
	}
	return ds
}

func (data *DataStructure) StructToJSON(result interface{}) error {
	json, err := json.Marshal(&result)
	if err != nil {
		log.Printf("struct to JSON err = %v", err)
		return err
	}
	fileHandler, err := data.fileCreator(data.fileName)
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
