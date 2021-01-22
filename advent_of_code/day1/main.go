package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	helpers "github.com/thesammy2010/advent-of-code/advent_of_code/helpers"
)

func checkSum(x int, y int) (int, int, int, error) {
	if x+y == 2020 {
		return x, y, x * y, nil
	}
	return 0, 0, 0, errors.New("x & y do not add up to 2020")
}

func castToInt(x string) int {
	y, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	return int(y)
}

func makePairs(data []int) [][]int {
	list := make([][]int, 0)
	// i and j are the pairs
	for _, i := range data {
		for _, j := range data {
			if i != j {
				tempList := []int{i, j}
				sort.Ints(tempList)
				list = append(list, tempList)
			}
		}
	}
	return list
}

func main() {
	inputData := helpers.ReadFileLineByLine("input.txt")
	inputDataIntegers := []int{}
	for _, row := range inputData {
		inputDataIntegers = append(inputDataIntegers, castToInt(row))
	}
	numberPairs := makePairs(inputDataIntegers)
	for _, row := range numberPairs {
		a, b, c, err := checkSum(row[0], row[1])
		if err != nil {
			fmt.Println(a, b, c)
		}
	}
}
