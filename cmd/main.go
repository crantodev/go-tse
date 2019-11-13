package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	"padron/internal/model/location"
	"padron/internal/model/person"
)

const (
	fullURL          = "https://www.tse.go.cr/zip/padron/padron_completo.zip"
	zipFilename      = "padron_completo.zip"
	tempFolder       = "./temp"
	locationFilename = "Distelec.txt"
	peopleFilename   = "PADRON_COMPLETO.txt"
)

func main() {

	// if !file.Exists(zipFilename) {
	// 	fmt.Println("Download Started")

	// 	err := file.Download(zipFilename, fullURL)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println("Download Finished")
	// }

	// fmt.Println("Extracting Zip file")
	// _, err := file.Unzip(zipFilename, tempFolder)
	// if err != nil {
	// 	panic(err)
	// }

	locationFile, err := os.Open(fmt.Sprintf("%s/%s", tempFolder, locationFilename))
	if err != nil {
		panic(err)
	}

	defer locationFile.Close()

	r := csv.NewReader(bufio.NewReader(locationFile))

	_, err = location.Parser(r)
	if err != nil {
		panic(fmt.Errorf("Locations could not be parsed %v", err))
	}

	peopleFile, err := os.Open(fmt.Sprintf("%s/%s", tempFolder, peopleFilename))
	if err != nil {
		panic(err)
	}

	defer peopleFile.Close()

	r = csv.NewReader(bufio.NewReader(peopleFile))

	people, err := person.Parser(r)
	if err != nil {
		panic(fmt.Errorf("People could not be parsed %v", err))
	}

	fmt.Println(people[0])
}
