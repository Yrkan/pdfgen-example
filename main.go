package main

import (
	"os"

	"github.com/Yrkan/pdfgen"
)
type Data struct {
	Date string
	PG1Name	string
	PG2Name string
	PG1Street	string
	PG2Street string
	PG1CityState string
	PG2CityState string
	PG1DOB string
	PG2DOB string
	PG1SSN string
	PG2SSN string
	PG1Ownership string
	PG2Ownership string

}
func main() {
	t1 := Data{
		Date: "11/11/2020", 
		PG1Name: "name",
		PG2Name: "Mark",
		PG1Street: "Street",
		PG2Street: "13 NY street",
		PG1CityState: "Rabat",
		PG2CityState: "Morocco",
		PG1DOB: "DOB1",
		PG2DOB: "DOB2",
		PG1SSN: "1234",
		PG2SSN: "4567",
		PG1Ownership: "33%",
		PG2Ownership: "93%",
	}

	// Write Example

	f, _ := os.Create("./testIoWriter.pdf")
	pdfgen.Write("./input/Test.docx", t1, f)
	f.Close()

	// Save Example
	pdfgen.Save("./input/Test.docx", t1, "./output/conv.pdf")
	
}
