package main

import (
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
}