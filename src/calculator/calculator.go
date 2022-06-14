package calculator

import (
	"fmt"
	"strconv"
)

func Divider(query []string) string {
	integerParams := make([]int, len(query))
	for index := range query {
		integerParams[index], _ = strconv.Atoi(query[index])
	}
	fmt.Printf("converted array: %d\n", integerParams)
	result := integerParams[0]
	for index := range integerParams {
		result /= integerParams[index]
	}
	resultObject := strconv.Itoa(result)
	return resultObject
}

func Adder(query []string) string {
	integerParams := make([]int, len(query))
	for index := range query {
		integerParams[index], _ = strconv.Atoi(query[index])
	}
	fmt.Printf("converted array: %d\n", integerParams)
	result := integerParams[0]
	for index := range integerParams {
		result += integerParams[index]
	}
	resultObject := strconv.Itoa(result)
	return resultObject
}

func Substracter(query []string) string {
	integerParams := make([]int, len(query))
	for index := range query {
		integerParams[index], _ = strconv.Atoi(query[index])
	}
	fmt.Printf("converted array: %d\n", integerParams)
	result := integerParams[0]
	for index := range integerParams {
		result -= integerParams[index]
	}
	resultObject := strconv.Itoa(result)
	return resultObject
}

func Multiplier(query []string) string {
	integerParams := make([]int, len(query))
	for index := range query {
		integerParams[index], _ = strconv.Atoi(query[index])
	}
	fmt.Printf("converted array: %d\n", integerParams)
	result := integerParams[0]
	for index := range integerParams {
		result *= integerParams[index]
	}
	resultObject := strconv.Itoa(result)
	return resultObject
}

func Summer(query []string) string {
	integerParams := make([]int, len(query))
	for index := range query {
		integerParams[index], _ = strconv.Atoi(query[index])
	}
	fmt.Printf("converted array: %d\n", integerParams)
	result := 0
	for index := 0; index <= integerParams[0]; index++ {
		result += index
	}
	resultObject := strconv.Itoa(result)
	return resultObject
}
