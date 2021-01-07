package solutions

import (
	"bufio"
	"os"
	"sort"

	log "github.com/sirupsen/logrus"
)

// Day05solution1 first part
func Day05solution1() {
	file, err := os.Open("puzzles/day05.txt")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var highest int = 0
	var allSeats []int
	for scanner.Scan() {
		boardingpass := scanner.Text()
		r, s := day05convertBoardingpass(boardingpass)
		seat := r*8 + s
		highest = day05max(highest, seat)
		allSeats = append(allSeats, seat)
	}
	log.Printf("Day 05 Solution 1: %v", highest)
	sort.Ints(allSeats)
	for i := 0; i < len(allSeats)-1; i++ {
		if allSeats[i]+1 != allSeats[i+1] {
			log.Printf("Day 05 Solution 2: %v", allSeats[i]+1)
		}
	}
}

func day05max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func day05convertBoardingpass(val string) (int, int) {
	row := val[:7]
	seat := val[7:]
	var r, s byte

	for _, v := range row {
		r = r << 1
		if v == 'B' {
			r = r | 1
		}
	}
	for _, v := range seat {
		s = s << 1
		if v == 'R' {
			s = s | 1
		}
	}
	return int(r), int(s)
}
