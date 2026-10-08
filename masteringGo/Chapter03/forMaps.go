package main

import (
	"fmt"
)

func main() {
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "A value"

	// range также работает с картами
	for key, v := range aMap {
		fmt.Println("key: ", key, "value: ", v)
	}

}
