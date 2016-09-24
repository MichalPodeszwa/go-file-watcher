package readers

import (
	"encoding/json"

	"github.com/michalpodeszwa/gowatcher/entities"
)

type jsonReader struct{}

func (reader jsonReader) ConvertToPerson(fileData []byte) entities.Person {
	person := entities.Person{}
	if err := json.Unmarshal(fileData, &person); err != nil {
		panic(err)
	}
	return person
}
