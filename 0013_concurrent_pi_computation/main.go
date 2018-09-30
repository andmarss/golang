package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

const ANSIClearScreenSequence = "\033[H\033[2J"
const ANSIFirstSlotScreenSequence = "\033[2;0H"
const ANSISecondSlotScreenSequence = "\033[3;0H"

var pichan chan float64 = make(chan float64)

var computationDone chan bool = make(chan bool, 1)

var termsCount int

func printCalculationSummary() {
	fmt.Print(ANSIClearScreenSequence)
	fmt.Println(ANSIFirstSlotScreenSequence, "Computed value of Pi:\t\t", <-pichan)
	fmt.Println(ANSISecondSlotScreenSequence, "# of Nilakantha Terms:\t\t", termsCount)
}

func pi(n int) float64 {
	ch := make(chan float64)
	for k := 1; k <= n; k++ {
		go nilakanthaTerm(ch, float64(k))
	}
	f := 3.0

	for k := 1; k <= n; k++ {
		termsCount++
		f += <-ch
		pichan <- f
	}

	computationDone <- true
	return f
}

func nilakanthaTerm(ch chan float64, k float64) {
	j := 2 * k
	if int64(k)%2 == 1 {
		ch <- 4.0 / (j * (j + 1) * (j + 2))
	} else {
		ch <- -4.0 / (j * (j + 1) * (j + 2))
	}
}

func main() {
	ticker := time.NewTicker(time.Millisecond * 108)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go pi(5000)

	go func() {
		for range ticker.C {
			printCalculationSummary()
		}
	}()

	for {
		select {
		case <-computationDone:
			ticker.Stop()
			fmt.Println("Program done calculaiting pi")

		case <-interrupt:
			ticker.Stop()
			fmt.Println("Program interrupted by the user")
			return
		}
	}
}
