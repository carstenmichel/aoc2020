package solutions

import (
	"bufio"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Day01solution1 first part
func Day01solution1() {
	file, err := os.Open("puzzles/day01.txt")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var numbers []int
	for _, v := range lines {
		i, _ := strconv.Atoi(v)
		numbers = append(numbers, i)
	}

	for i, v := range numbers {
		for _, w := range numbers[i:] {
			if v+w == 2020 {
				log.Printf("Day 1 Solution 1: %v + %v is 2020, multiplies to %v", v, w, (v * w))
			}
		}
	}
}

// Day01solution2 second part
func Day01solution2() {
	file, err := os.Open("puzzles/day01.txt")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var numbers []int
	for _, v := range lines {
		i, _ := strconv.Atoi(v)
		numbers = append(numbers, i)
	}

	for i, v := range numbers {
		for j, w := range numbers[i:] {
			for _, x := range numbers[(i + j):] {
				if v+w+x == 2020 {
					log.Printf("Day 1 Solution 2: %v + %v + %v is 2020, multiplies to %v", v, w, x, (v * w * x))
				}
			}
		}
	}
}
