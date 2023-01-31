package main

import (
	"fmt"
	// "math/rand"
	// "sync"
	// "time"
)

func sumInt(m map[string]int64) int64 {
	var total int64

	for _, v := range m {
		total += v
	}

	return total
}

func sumFloat(m map[string]float64) float64 {
	var total float64

	for _, v := range m {
		total += v
	}

	return total
}

func sumIntorFloat[K comparable, V int64 | float64](m map[K]V) V {
	var total V

	for _, v := range m {
		total += v
	}

	return total
}

type Number interface {
	int | int32 | int64 | float32 | float64
}

func sumNumber[K comparable, V Number](m map[K]V) V {
	var total V

	for _, v := range m {
		total += v
	}

	return total
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 32.3,
	}

	fmt.Printf("non generics: %v and %v\n",
		sumInt(ints),
		sumFloat(floats),
	)

	fmt.Printf("generics: %v and %v\n",
		sumIntorFloat(ints),
		sumIntorFloat(floats),
	)

	fmt.Printf("generics with interface: %v and %v\n",
		sumNumber(ints),
		sumNumber(floats),
	)
}
