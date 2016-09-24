package readers

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/michalpodeszwa/gowatcher/entities"
)

type fileReader interface {
	AddPeopleToChan(fileData []byte, people chan entities.Person)
}

func getCorrectReader(fileName string) fileReader {
	ext := path.Ext(fileName)
	switch ext {
	case ".json":
		return jsonReader{}
	case ".csv":
		return csvReader{}
	}
	panic(fmt.Sprintf("File reader for %s files does not exist", ext))
}

func openFile(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return data
}

// ReadFile is used to read from file and then to parse it in correct way, returning a Person
func ReadFile(fileName string, people chan entities.Person) {
	log.Println("Reading file:", fileName)
	reader := getCorrectReader(fileName)

	fileData := openFile(fileName)

	reader.AddPeopleToChan(fileData, people)
}
