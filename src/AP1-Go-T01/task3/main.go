package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntersectionLists struct {
	first  []int
	second []int
	result []int
}

func parseLine() ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ", ")
		var result []int
		for _, num := range parts {
			intNum, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("invalid input")
				return nil, fmt.Errorf("invalid input")
			}
			result = append(result, intNum)
		}
		return result, nil
	} else {
		return nil, fmt.Errorf("no input")
	}
}

func getUserInput(interList *IntersectionLists) {
	fmt.Println("Enter first list:")
	first, err := parseLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	interList.first = first

	second, err := parseLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	interList.second = second

}

func getIntersection(interList *IntersectionLists) {
	for _, num1 := range interList.first {
		for _, num2 := range interList.second {
			if num1 == num2 {
				interList.result = append(interList.result, num1)
				break
			}
		}
	}
}

func main() {
	var interList IntersectionLists
	getUserInput(&interList)
	getIntersection(&interList)
	fmt.Println(interList.result)
}
