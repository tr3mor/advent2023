package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	input := strings.Split(string(data), "\n")
	if err != nil {
		fmt.Println(err)
	}
	deck := make(map[int]int)
	var sum float64
	var wins int
	for id, card := range input {
		cardId := id + 1
		cardData := strings.Split(strings.Split(card, ":")[1], " | ")
		winNumbers := strings.Fields(cardData[0])
		cardNumbers := strings.Fields(cardData[1])
		wins = 0
		deck[cardId] += 1
		for _, x := range winNumbers {
			for _, y := range cardNumbers {
				if x == y {
					wins++
				}
			}
		}
		if wins > 0 {
			sum += math.Pow(2.0, float64(wins-1))
		}
		for c := 0; c < deck[cardId]; c++ {
			for k := 0; k < wins; k++ {
				deck[cardId+k+1] += 1
			}
		}
	}
	fmt.Println(int(sum))
	var sum2 int
	for _, v := range deck {
		sum2 += v
	}
	fmt.Println(sum2)
}
