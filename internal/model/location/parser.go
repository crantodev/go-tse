package location

import (
	"encoding/csv"
	"io"
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

		locations = append(locations, Location{
			ID:       line[0],
			Province: line[1],
			District: line[2],
			Canton:   line[3],
		})
	}

	return locations, nil
}
