package main

import (
	"fmt"

	"padron/internal/service/handler/download"
)

const (
	fullURL = "https://www.tse.go.cr/zip/padron/padron_completo.zip"
)

func main() {

	fmt.Println("Download Started")

	err := download.File("padron_completo.zip", fullURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Download Finished")
	// distelecFile, _ := os.Open("Distelec.txt")
	// distelectReader := csv.NewReader(bufio.NewReader(distelecFile))

	// defer distelecFile.Close()

	// locations, err := models.Parser(distelectReader)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(locations[1].IDSplitter())

	// padronReader, _ := parser.Reader("padron.txt")

	// var people []models.Person

	// for {
	// 	record, err := padronReader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	people = append(people, models.Person{
	// 		ID:         record[0],
	// 		Location:   record[1],
	// 		Gender:     record[2],
	// 		DueDate:    record[3],
	// 		Vote:       record[4],
	// 		FirstName:  strings.TrimSpace(record[5]),
	// 		MiddleName: strings.TrimSpace(record[6]),
	// 		LastName:   strings.TrimSpace(record[7]),
	// 	})
	// }

	// fmt.Println(people[2])
}
