package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func main() {

	connStr := flag.String("conn-str", "", "PostgreSQL connection string")
	table := flag.String("table", "", "Table to extract data from")
	columns := flag.String("columns", "", "Columns to extract")
	outputFile := flag.String("output-file", "", "Output file (CSV)")
	flag.Parse()

	
	if *connStr == "" || *table == "" || *columns == "" {
		fmt.Println("Error: required flags not provided")
		flag.Usage()
		os.Exit(1)
	}

	db, err := pq.Connect(*connStr)
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()


	data, err := extractData(db, *table, *columns)
	if err!= nil {
		log.Fatal(err)
	}

	stats, err := analyzeData(data)
	if err!= nil {
		log.Fatal(err)
	}


	if *outputFile!= "" {
		err = outputCSV(*outputFile, stats)
	} else {
		err = outputTerminal(stats)
	}
	if err!= nil {
		log.Fatal(err)
	}
}
