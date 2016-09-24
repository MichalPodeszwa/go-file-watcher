package readers

import (
	"bytes"
	"encoding/csv"
	"strconv"

	"github.com/michalpodeszwa/gowatcher/entities"
)

type csvReader struct{}

func (reader csvReader) readSinglePerson(person []string) entities.Person {
	age, _ := strconv.ParseInt(person[2], 10, 64)

	return entities.Person{FirstName: person[0], LastName: person[1], Age: age}
}

func (reader csvReader) AddPeopleToChan(fileData []byte, people chan entities.Person) {
	r := csv.NewReader(bytes.NewReader(fileData))
	r.Comma = ';'
	csvPeople, _ := r.ReadAll()

	for _, person := range csvPeople {
		people <- reader.readSinglePerson(person)
	}
}
