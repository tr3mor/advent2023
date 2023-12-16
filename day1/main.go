package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var part2data = map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	data, err := os.ReadFile("input.txt")
	input := strings.Split(string(data), "\n")
	if err != nil {
		fmt.Println(err)
	}
	var part2Re = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	var part2Re2 = regexp.MustCompile(".*(" + part2Re.String() + ")")
	var sum int

	for _, x := range input {
		first2 := part2data[part2Re.FindString(x)]
		last2 := part2data[part2Re2.FindStringSubmatch(x)[1]]
		code := 10*first2 + last2
		sum += code
	}
	fmt.Println(sum)

}
