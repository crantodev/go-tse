package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"padron/internal/model/location"
	"padron/internal/model/person"
	"padron/internal/service/handler/file"
)

const (
	fullURL          = "https://www.tse.go.cr/zip/padron/padron_completo.zip"
	zipFilename      = "padron_completo.zip"
	tempFolder       = "./temp"
	locationFilename = "Distelec.txt"
	peopleFilename   = "PADRON_COMPLETO.txt"
)

func main() {
	var sf string = fmt.Sprintf("%s/%s", tempFolder, zipFilename)

	if !file.Exists(sf) {
		if err := file.Download(sf, fullURL); err != nil {
			panic(err)
		}
	}

	var lf string = fmt.Sprintf("%s/%s", tempFolder, locationFilename)
	var pf string = fmt.Sprintf("%s/%s", tempFolder, peopleFilename)

	if !file.Exists(lf) || !file.Exists(pf) {
		fmt.Println("Extracting Zip file")
		_, err := file.Unzip(sf, tempFolder)
		if err != nil {
			panic(err)
		}
	}

	file, err := os.Open(lf)
	if err != nil {
		panic(err)
	}

	lfr := csv.NewReader(bufio.NewReader(file))
	defer file.Close()

	locations, err := location.Parser(lfr)
	if err != nil {
		panic(err)
	}

	file, err = os.Open(pf)
	if err != nil {
		panic(err)
	}

	pfr := csv.NewReader(bufio.NewReader(file))

	people, err := person.Parser(pfr)
	if err != nil {
		panic(err)
	}

	// Process parsed data
	fmt.Println(locations[0], people[0])
}
