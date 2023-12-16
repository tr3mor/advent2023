package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	input := strings.Split(string(data), "\n")
	if err != nil {
		fmt.Println(err)
	}
	grid := make([][]rune, len(input))

	for y, line := range input {
		grid[y] = []rune(line)
	}
	var sum int
	gears := map[string][]int{}
	for y, line := range grid {
		for x := 0; x < len(line); x++ {
			if line[x] >= 48 && line[x] <= 57 {
				start := x
				for x++; x < len(line) && line[x] >= 48 && line[x] <= 57; x++ {
				}
				end := x - 1
				number, _ := strconv.Atoi(string(line[start : end+1]))
				if CheckAdjust(grid, start, end, y, number, gears) {
					sum += number
				}
			}
		}
	}
	fmt.Println(sum)
	var sum2 int
	keys := make([]string, len(gears))
	for k := range gears {
		keys = append(keys, k)
	}
	for _, k := range keys {
		if len(gears[k]) > 1 {
			sum2 += gears[k][0] * gears[k][1]
		}
	}
	fmt.Println(sum2)
}

func CheckAdjust(grid [][]rune, start, end, row, number int, gears map[string][]int) bool {
	startRow := math.Max(0, float64(row-1))
	endRow := math.Min(float64(row+1), float64(len(grid)-1))
	startChar := math.Max(0, float64(start-1))
	endChar := math.Min(float64(end+1), float64(len(grid[row])-1))
	adj := false
	for x := startRow; x <= endRow; x++ {
		for c := startChar; c <= endChar; c++ {
			symbol := grid[int(x)][int(c)]
			if symbol != 46 && (symbol < 48 || symbol > 57) {
				if symbol == 42 {
					gears[fmt.Sprintf("%f,%f", x, c)] = append(gears[fmt.Sprintf("%f,%f", x, c)], number)
				}
				adj = true
			}
		}
	}
	return adj
}
