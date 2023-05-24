package structural

import "fmt"

// data formats example
// this example uses interface Source and one method for add csv data to db
// need to create an adapter to add data from a source that returns a different format

type Client struct {
}

func (c *Client) AddDataToDB(source Source) {
	source.ParseCsv()
	fmt.Println("Adding data to DB.")
}

type Source interface {
	ParseCsv()
}

type Csv struct{}

func (c *Csv) ParseCsv() {
	fmt.Println("Parsing CSV data.")
}

type Json struct{}

func (j *Json) ParseJson() {
	fmt.Println("Parsing JSON data.")
}

type FormatAdapter struct {
	JsonParser *Json
}

func (a *FormatAdapter) ParseCsv() {
	a.JsonParser.ParseJson()
	fmt.Println("Adapting JSON data to CSV format.")
}

func RunAdapter() {
	client := &Client{}
	csvSource := &Csv{}

	JsonSource := &FormatAdapter{
		JsonParser: &Json{},
	}

	fmt.Println("Adding data from CSV to DB:")
	client.AddDataToDB(csvSource)

	fmt.Println()

	fmt.Println("Adding data from JSON to DB (using adapter):")
	client.AddDataToDB(JsonSource)
}
