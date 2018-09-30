package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var greetings string
var howdyDone chan bool
var mutex = &sync.Mutex{}

func howdyGreetings() {
	mutex.Lock()
	greetings = "Howdy Mars!"
	mutex.Unlock()
	howdyDone <- true
}

func fetchUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done() // defer служит для отсрочки выполнения функции wg.Done() после выполнения окружающей его функции, в данном примере -  после выполнения функции fetchUrl

	response, err := http.Get(url)

	if err != nil {
		log.Fatal("Failed to fetch the URL ", url, " and encountered this error", err)
	} else {
		fmt.Println("Contents of url ", url, " is:\n")
		contents, err := ioutil.ReadAll(response.Body)

		response.Body.Close()

		if err != nil {
			log.Fatal("Failed to read the response body and encountered this error: ", err)
		}
		fmt.Println(string(contents), "\n")
	}
}

func main() {
	howdyDone = make(chan bool, 1)

	go howdyGreetings()
	// mutex
	mutex.Lock()

	greetings = "Hello Mars!"
	fmt.Println(greetings)

	mutex.Unlock()

	<-howdyDone

	//wait groups

	var wg sync.WaitGroup
	var urlList = []string{"http://www.golang.org/", "http://www.google.com/", "http://www.youtube.com/"}

	for _, url := range urlList {
		wg.Add(1)

		go fetchUrl(url, &wg)
	}

	wg.Wait()
}
