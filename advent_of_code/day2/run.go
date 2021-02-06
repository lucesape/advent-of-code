package main

import (
	part1 "advent-of-code/advent_of_code/day2/part1"
	// "advent-of-code/advent_of_code/day1/part2"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Part 1")
	c := part1.Main("advent_of_code/day2/input.txt")
	fmt.Println(c)

	// fmt.Println("Part 2")
	// a, b, c, d, err := part2.Main("advent_of_code/day1/input.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(2)
	// }
	// fmt.Println(a, b, c, d)

	os.Exit(0)

}
