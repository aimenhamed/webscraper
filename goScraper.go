package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	f := "data.csv"
	file, err := os.Create(f)
	if err != nil {
		log.Fatalf("Failed to create file, err : %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
}
