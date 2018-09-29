package main

import "fmt"

type EDocument struct {
	Number       string
	Date         string
	NumberOfPage int
	FooterDoc    Footer
}

type Footer struct {
	Copyright string
}

func (doc EDocument) showCopyright() {
	fmt.Println(doc.FooterDoc.Copyright)
}

// лучше использовать методы с указателями * для экономии памяти

func (doc *EDocument) showPageNumber() {
	fmt.Println(doc.NumberOfPage)
}

func main() {
	doc1 := EDocument{
		Number:       "1",
		Date:         "29.09.2018",
		NumberOfPage: 1,
		FooterDoc: Footer{
			Copyright: "(c)",
		},
	}

	var doc2 EDocument
	doc2.NumberOfPage = 2
	doc2.Date = "29.09.2018"
	doc2.Number = "2"
	doc2.FooterDoc.Copyright = "(cc)"

	doc3 := new(EDocument)
	doc3.NumberOfPage = 3
	doc3.Date = "29.09.2018"
	doc3.Number = "3"
	doc3.FooterDoc.Copyright = "(ccc)"

	fmt.Printf("%T - %v\n", doc1, doc1)
	fmt.Printf("%T - %v\n", doc2, doc2)
	fmt.Printf("%T - %v\n", doc3, doc3)

	doc1.showPageNumber()
	doc2.showPageNumber()
	doc3.showPageNumber()
}
