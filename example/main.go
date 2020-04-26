package main

import (
	"fmt"
	csv_conv "github.com/garupanojisan/csv-conv"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./example.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	conv, err := csv_conv.NewConverter(f)
	if err != nil {
		log.Fatal(err)
	}

	// change column names
	changed, err := conv.ChangeColumnName(map[string]string{
		"a": "A",
		"あ": "ア",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(changed)

	// append a new column
	append, err := conv.AppendColumn("new", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(append)
}
