package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Sort Unordered Values In Maps")

	stateMap := make(map[string]string)
	fmt.Println("State Map:", stateMap)

	stateMap["WA"] = "Washington"
	stateMap["OR"] = "Oregon"
	stateMap["CA"] = "California"
	fmt.Println("State Map:", stateMap)

	california := stateMap["CA"]
	fmt.Println("California:", california)

	delete(stateMap, "OR")
	fmt.Println("State Map After Deletion:", stateMap)

	stateMap["NY"] = "New York"
	fmt.Println("State Map:", stateMap)

	for k, v := range stateMap {
		fmt.Println("Key:", k, "Value:", v)
	}

	keySlice := make([]string, len(stateMap))
	i := 0
	for k := range stateMap {
		keySlice[i] = k
		i++
	}

	sort.Strings(keySlice)
	fmt.Println("Sorted Key Slice:", keySlice)

	for i := range keySlice {
		fmt.Println("State Key: ", stateMap[keySlice[i]])
	}
}
