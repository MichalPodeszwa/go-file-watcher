package outputers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/michalpodeszwa/gowatcher/entities"
)

type jsonFileCreator struct{}

func (creator jsonFileCreator) createFile(person entities.Person, fileName string) string {
	fileName = creator.getFileName(fileName)

	fileData, _ := json.Marshal(person)

	ioutil.WriteFile(fileName, fileData, 0644)

	return fileName
}

func (creator jsonFileCreator) getFileName(fileName string) string {
	return fmt.Sprintf("%s.csv", fileName)
}
