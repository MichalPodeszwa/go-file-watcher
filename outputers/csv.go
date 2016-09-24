package outputers

import (
	"fmt"
	"io/ioutil"

	"github.com/michalpodeszwa/gowatcher/entities"
)

type csvFileCreator struct{}

func (creator csvFileCreator) createFile(person entities.Person, fileName string) string {
	fileName = creator.getFileName(fileName)

	fileData := []byte(fmt.Sprintf("%s;%s:%d", person.FirstName, person.LastName, person.Age))

	ioutil.WriteFile(fileName, fileData, 0644)

	return fileName
}

func (creator csvFileCreator) getFileName(fileName string) string {
	return fmt.Sprintf("%s.csv", fileName)
}
