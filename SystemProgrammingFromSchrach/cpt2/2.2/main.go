package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	csv_file, open_err := os.Create("test.csv")
	if open_err != nil {
		log.Fatal(open_err)
	}
	defer csv_file.Close()

	csv_writer := csv.NewWriter(csv_file)
	defer csv_writer.Flush()

	if write_err := csv_writer.Write([]string{"key", "value"}); write_err != nil {
		log.Fatal(write_err)
	}
	if write_err := csv_writer.Write([]string{"hello", "world"}); write_err != nil {
		log.Fatal(write_err)
	}
	if write_err := csv_writer.Write([]string{"im", "hen00af"}); write_err != nil {
		log.Fatal(write_err)
	}
}