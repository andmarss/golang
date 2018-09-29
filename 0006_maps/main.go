package main

import (
	"fmt"
	"sort"
)

/*
1. Declaring and initializing map
var map[keyType]valueType = make(map[keyType]valueType)

2. Setting a value in the map
map[key] = value

3. Getting a value in the slice
map[key]
*/

func main() {
	var nationCapitals = make(map[string]string)

	nationCapitals["Afganistan"] = "Kabul"
	nationCapitals["Canada"] = "Ottawa"
	nationCapitals["Japan"] = "Tokyo"
	nationCapitals["Kenya"] = "Nairobi"
	nationCapitals["India"] = "New Delhi"
	nationCapitals["Mexico"] = "Mexico City"
	nationCapitals["South Korea"] = "Seoul"
	nationCapitals["United Kingdom"] = "London"
	nationCapitals["USA"] = "Washington D.C."
	nationCapitals["Taiwan"] = "Taipei"
	nationCapitals["Russia"] = "Moscow"

	printNationalCapitalsMap(nationCapitals)

	fmt.Println("Print sorted capitals")

	printSortedNationalCapitalsMap(nationCapitals)
}

func printNationalCapitalsMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("The capital of", key, "is", value)
	}
}

func printSortedNationalCapitalsMap(m map[string]string) {
	keys := make([]string, len(m))

	for key, _ := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, value := range keys {
		if value == "" {
			continue
		}
		fmt.Println("The capital of", value, "is", m[value])
	}
}
