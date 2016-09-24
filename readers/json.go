package readers

import (
	"encoding/json"

	"github.com/michalpodeszwa/gowatcher/entities"
)

type jsonPersonArray struct {
	People []entities.Person
}

type jsonReader struct{}

func (reader jsonReader) AddPeopleToChan(fileData []byte, people chan entities.Person) {
	var peopleArray *[]entities.Person
	if err := json.Unmarshal(fileData, &peopleArray); err != nil {
		panic(err)
	}
	for _, person := range *peopleArray {
		people <- person
	}
	// peopleArray := jsonPersonArray{}
	// if err := json.Unmarshal(fileData, &peopleArray); err != nil {
	// 	panic(err)
	// }
	//
	// for _, person := range peopleArray.People {
	// 	people <- person
	// }
}
