package reader

import (
	"bufio"
	"encoding/csv"
	"os"
)

func File(filename string) (*csv.Reader, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(bufio.NewReader(file))

	defer file.Close()

	return reader, nil
}
