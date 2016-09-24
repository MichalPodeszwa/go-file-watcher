package outputers

import (
	"fmt"
	"log"

	"github.com/michalpodeszwa/gowatcher/entities"
)

type fileCreator interface {
	createFile(person entities.Person, fileName string) string
}

func getCorrectFileCreator(person entities.Person, outputType string) fileCreator {
	switch outputType {
	case "json":
		return jsonFileCreator{}
	case "csv":
		return csvFileCreator{}
	default:
		panic(fmt.Sprintf("%s type output is not available", outputType))
	}
}

// CreateFile creates a file based on a person given as an argument and returns the new filename
func CreateFile(person entities.Person, outputDir string, outputType string) {
	fileName := fmt.Sprintf("%s/%s %s", outputDir, person.FirstName, person.LastName)

	fileCreator := getCorrectFileCreator(person, outputType)

	fileName = fileCreator.createFile(person, fileName)
	log.Println("Created file:", fileName)
}
