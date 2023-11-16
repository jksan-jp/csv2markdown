package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a CSV file name as an argument.")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	mdFile, err := os.Create("output.md")
	if err != nil {
		log.Fatal(err)
	}
	defer mdFile.Close()

	fmt.Fprintf(mdFile, "| %s |\n", strings.Join(header, " | "))

	fmt.Fprintf(mdFile, "| --- ")
	for i := 1; i < len(header); i++ {
		fmt.Fprintf(mdFile, "| --- ")
	}
	fmt.Fprintf(mdFile, "|\n")

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		markdownRow := fmt.Sprintf("| %s |\n", strings.Join(record, " | "))
		fmt.Fprintf(mdFile, "%s", markdownRow)
	}
}
