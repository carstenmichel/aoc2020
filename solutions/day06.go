package solutions

import (
	"bufio"
	"os"

	log "github.com/sirupsen/logrus"
)

// Day06solution2 first part
func Day06solution2() {
	file, err := os.Open("puzzles/day06.txt")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()
	//var lines []string
	var groups [][26]bool
	scanner := bufio.NewScanner(file)
	// All Groups
	for scanner.Scan() {
		group := [26]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
		// Single Group
		for {
			t := scanner.Text()
			if t == "" {
				//End of group reached
				break
			}
			var tempGroup [26]bool
			for _, c := range t {
				tempGroup[day06RuneToNum(c)] = true
			}
			for i := range group {
				group[i] = group[i] && tempGroup[i]
			}

			scanner.Scan()
		}
		groups = append(groups, group)
	}
	var allSums int = 0
	for _, v := range groups {
		allSums += day06CountBits(v)
	}
	log.Printf("Day 06 Solution 2: %v", allSums)
}

// Day06solution1 first part
func Day06solution1() {
	file, err := os.Open("puzzles/day06.txt")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()
	//var lines []string
	var groups [][26]bool
	scanner := bufio.NewScanner(file)
	// All Groups
	for scanner.Scan() {
		var group [26]bool
		// Single Group
		for {
			t := scanner.Text()
			if t == "" {
				//End of group reached
				break
			}
			for _, c := range t {
				group[day06RuneToNum(c)] = true
			}
			scanner.Scan()
		}
		groups = append(groups, group)
	}
	var allSums int = 0
	for _, v := range groups {
		allSums += day06CountBits(v)
	}
	log.Printf("Day 06 Solution 1: %v", allSums)
}

func day06RuneToNum(r rune) int {
	return int(r - 'a')
}

func day06CountBits(group [26]bool) int {
	res := 0
	for _, v := range group {
		if v {
			res++
		}
	}
	return res
}
