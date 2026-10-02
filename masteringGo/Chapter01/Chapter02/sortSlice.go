package main

import (
	"fmt"
	"sort"
)

func main() {
	sInts := []int{1, 0, 2, -3, 4, -20}
	sFloats := []float64{1.0, 0.2, 0.22, -3, 4.1, -0.1}
	sString := []string{"aa", "a", "A", "Aa", "aab", "AAa"}

	fmt.Println("sInts originals: ", sInts)
	sort.Ints(sInts)
	fmt.Println("sInts: ", sInts)
	sort.Sort(sort.Reverse(sort.IntSlice(sInts)))
	fmt.Println("Reversed: ", sInts)

	fmt.Println("sFloats originals: ", sFloats)
	sort.Float64s(sFloats)
	fmt.Println("sFloats: ", sFloats)
	sort.Sort(sort.Reverse(sort.Float64Slice(sFloats)))
	fmt.Println("Reverse: ", sFloats)

	fmt.Println("sString original: ", sString)
	sort.Strings(sString)
	fmt.Println("sString: ", sString)
	sort.Sort(sort.Reverse(sort.StringSlice(sString)))
	fmt.Println("Reverse: ", sString)

}
