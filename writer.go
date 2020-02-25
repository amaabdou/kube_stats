package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strings"
)

func ListAvailableWriters() []string {
	return []string{"console", "json", "csv", "html"}
}

func GetDefaultWriter() string  {
	return "console"
}

func Write(WriterName string, dataset PodsData) error  {
	switch WriterName {
		case "html":
			return writeToHtml(dataset)
		case "console":
			return writeToConsole(dataset)
		case "json":
			return writeJson(dataset)
		case "csv":
			return writeCSV(dataset)
	}

	return errors.New(fmt.Sprintf("Something went very wrong, give sort-by [%s] not exist in [%s]", WriterName, strings.Join(ListAvailableWriters(), ",")))
}

func writeToHtml(data PodsData) error {
	return WriteToHtmlServerResponse(data)
}

func writeCSV(dataset PodsData) error {
	writer := csv.NewWriter(os.Stdout)

	if err := writer.Write(dataset.Headers); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}


	if err := writer.WriteAll(dataset.Data); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}


	writer.Flush()
	if err := writer.Error(); err != nil {

		log.Fatal(err)

	}
	return  nil
}


func writeJson(dataset PodsData) error {
	jsonData,err := json.Marshal(dataset)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
	return  nil
}


func writeToConsole(dataset PodsData) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.SetHeader(dataset.Headers)
	table.AppendBulk(dataset.Data)
	table.Render()
	return  nil
}
