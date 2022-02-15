package main

import (
	"ali-tariff/internal/utils"
	"fmt"
)

func main() {

	fmt.Println(fmt.Printf("CHUNKS\nbefore: %d", getSlice(22)))
	fmt.Println(fmt.Printf("after: %d \n\n", utils.GetChunkSlices(getSlice(22), 3)))

	fmt.Println("CONVERTER\nbefore: ")
	fmt.Println(getMap())
	fmt.Println("after: ")
	fmt.Println(utils.ConverterKeyValue(getMap()))
}

func getSlice(sliceSize int) []int {
	slice := make([]int, sliceSize)

	for i := 0; i < sliceSize; i++ {
		slice[i] = i + 1
	}

	return slice
}

func getMap() map[string]int {

	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	return m
}
