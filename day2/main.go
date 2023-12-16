package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	input := strings.Split(string(data), "\n")
	if err != nil {
		fmt.Println(err)
	}
	idRegexp := regexp.MustCompile(`\d+`)
	greenRegexp := regexp.MustCompile(`(\d+) green`)
	redRegexp := regexp.MustCompile(`(\d+) red`)
	blueRegexp := regexp.MustCompile(`(\\d+) blue`)
	var sum1, sum2 int
	for _, line := range input {
		id, _ := strconv.Atoi(idRegexp.FindAllString(strings.Split(line, ":")[0], 1)[0])
		green := FindBigestNumberOfCubes(greenRegexp.FindAllStringSubmatch(line, -1))
		red := FindBigestNumberOfCubes(redRegexp.FindAllStringSubmatch(line, -1))
		blue := FindBigestNumberOfCubes(blueRegexp.FindAllStringSubmatch(line, -1))
		if red <= 12 && green <= 13 && blue <= 14 {
			sum1 += id
		}
		sum2 = sum2 + (green * red * blue)
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func FindBigestNumberOfCubes(a [][]string) int {
	maxNumber := 0
	for _, x := range a {
		d, _ := strconv.Atoi(x[1])
		if d > maxNumber {
			maxNumber = d
		}
	}
	return maxNumber
}
