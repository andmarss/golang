package main

import "fmt"

type EDocument struct {
	Number       string
	Date         string
	NumberOfPage int
	Footer
}

type Footer struct {
	Copyright string
}

func main() {
	doc1 := EDocument{
		Number:       "1",
		Date:         "29.09.2018",
		NumberOfPage: 1,
		Footer: Footer{
			Copyright: "(c)",
		},
	}

	var doc2 EDocument
	doc2.NumberOfPage = 2
	doc2.Date = "29.09.2018"
	doc2.Number = "2"
	doc2.Copyright = "(cc)"

	fmt.Printf("%T - %v\n", doc1, doc1)
	fmt.Printf("%T - %v\n", doc2, doc2)
	fmt.Printf("%T - %v\n", doc2.Footer, doc2.Footer)
}
