package part1

import (
	"advent-of-code/advent_of_code/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Policy struct {
	LowerLimit int
	UpperLimit int
	Character  string
	Password   string
	IsValid    bool
}

//  tbd
func makePolicy(line string) *Policy {
	// 1-3 a: abcde
	var policy Policy = Policy{}

	splitLine := strings.Fields(line)
	a := strings.Split(splitLine[0], "-")[0]
	b := strings.Split(splitLine[0], "-")[1]
	c := strings.Split(splitLine[1], ":")[0]
	d := splitLine[2]

	aa, _ := strconv.ParseInt(a, 10, 64)
	bb, _ := strconv.ParseInt(b, 10, 64)

	policy.LowerLimit = int(aa)
	policy.UpperLimit = int(bb)
	policy.Character = c
	policy.Password = d

	return &policy
}

func isPolicyValid(p Policy) bool {
	if p.LowerLimit <= strings.Count(p.Password, p.Character) && strings.Count(p.Password, p.Character) <= p.UpperLimit {
		p.IsValid = true
		return true
	}
	return false
}

func main(filepath string) int {
	lines, err := helpers.ReadFileLineByLine(filepath)
	if err != nil {
		fmt.Println(err)
	}

	var results int

	for _, line := range lines {
		policy := makePolicy(line)
		if isPolicyValid(*policy) {
			results++
		}
	}

	return results

}
