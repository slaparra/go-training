package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	csvFile, err := os.Open("file-to-read.csv")

	defer csvFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(csvFile)

	var aSliceOfMaps []map[string]string

	for {
		values, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		element := make(map[string]string)
		element["policyID"] = values[0]
		element["statecode"] = values[1]
		element["county"] = values[2]
		element["eq_site_limit"] = values[3]
		element["hu_site_limit"] = values[4]
		element["fl_site_limit"] = values[5]
		element["fr_site_limit"] = values[6]
		element["tiv_2011"] = values[7]
		element["tiv_2012"] = values[8]
		element["eq_site_deductible"] = values[9]
		element["hu_site_deductible"] = values[10]
		element["fl_site_deductible"] = values[11]
		element["fr_site_deductible"] = values[12]
		element["point_latitude"] = values[13]
		aSliceOfMaps = append(aSliceOfMaps, element)
	}

	for _, value := range aSliceOfMaps {
		fmt.Println(value["policyID"], "-", value["county"])
	}
}
