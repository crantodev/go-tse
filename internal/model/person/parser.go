package person

import (
	"encoding/csv"
	"io"
)

func Parser(reader *csv.Reader) ([]Person, error) {
	var people []Person

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			return nil, error
		}

		people = append(people, Person{
			ID:         line[0],
			Location:   line[1],
			Gender:     line[2],
			DueDate:    line[3],
			Vote:       line[4],
			FirstName:  line[5],
			MiddleName: line[6],
			LastName:   line[7],
		})
	}

	return people, nil
}
