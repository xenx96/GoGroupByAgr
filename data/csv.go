package data

import (
	"encoding/csv"
	"log"
	"os"
	"path"
)

func WriteCSV(data [][]string, fileName string) {
	_ = os.Mkdir("Filter_Result", os.ModePerm)

	csvFile, err := os.Create(path.Join("Filter_Result", fileName))
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	for _, Row := range data {
		_ = csvwriter.Write(Row)
	}

	csvwriter.Flush()
	csvFile.Close()
}
