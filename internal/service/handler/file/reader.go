package file

import (
	"bufio"
	"encoding/csv"
	"os"
)

func Reader(
	filename string,
	callback func(reader *csv.Reader) ([]interface{}, error),
) ([]interface{}, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bufio.NewReader(file))

	defer file.Close()

	return callback(r)
}
