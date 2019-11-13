package location

import (
	"encoding/csv"
	"io"
	"regexp"
	"strings"
)

func Parser(reader *csv.Reader) ([]Location, error) {
	var locations []Location

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			return nil, error
		}

		ids := splitID(line[0])

		locations = append(locations, Location{
			ID:         line[0],
			IDProvince: ids[0],
			IDDistrict: ids[1],
			IDCanton:   ids[2],
			Province:   format(line[1]),
			District:   format(line[2]),
			Canton:     format(line[3]),
		})
	}

	return locations, nil
}

func format(s string) string {
	return strings.Title(strings.ToLower(strings.TrimSpace(s)))
}

func splitID(id string) []string {
	re := regexp.MustCompile(`^(\d{1})(\d{2})(\d{3})$`)
	match := re.FindStringSubmatch(id)

	return match[1:]
}
