package person

import (
	"encoding/csv"
	"io"
	"regexp"
	"strings"
)

// Parser get the CSV content and creates the person struct
func Parser(reader *csv.Reader) ([]Person, error) {
	var people []Person

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			return nil, error
		}

		ids := splitLocation(line[1])

		people = append(people, Person{
			ID:         line[0],
			IDProvince: ids[0],
			IDDistrict: ids[1],
			IDCanton:   ids[2],
			IDLocation: line[1],
			Gender:     line[2],
			DueDate:    line[3],
			Vote:       line[4],
			FirstName:  format(line[5]),
			MiddleName: format(line[6]),
			LastName:   format(line[7]),
		})
	}

	return people, nil
}

func format(s string) string {
	return strings.Title(strings.ToLower(strings.TrimSpace(s)))
}

func splitLocation(id string) []string {
	re := regexp.MustCompile(`^(\d{1})(\d{2})(\d{3})$`)
	match := re.FindStringSubmatch(id)

	return match[1:]
}
