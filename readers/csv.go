package readers

import (
	"bytes"
	"encoding/csv"
	"strconv"

	"github.com/michalpodeszwa/gowatcher/entities"
)

type csvReader struct{}

func (reader csvReader) ConvertToPerson(fileData []byte) entities.Person {
	r := csv.NewReader(bytes.NewReader(fileData))
	r.Comma = ';'
	csvPerson, _ := r.Read()

	age, _ := strconv.ParseInt(csvPerson[2], 10, 64)

	person := entities.Person{FirstName: csvPerson[0], LastName: csvPerson[1], Age: age}

	return person
}
