package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	log "github.com/sirupsen/logrus"
)

// Day03solution1 first part
func Day03solution1() {
	file, err := os.Open("puzzles/day03.txt")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	var day03Matrix [][31]bool
	day03Matrix = make([][31]bool, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		e := scanner.Text()
		var tempArray [31]bool
		for i, v := range e {
			if v == 35 {
				tempArray[i] = true
			} else {
				tempArray[i] = false
			}
		}
		day03Matrix = append(day03Matrix, tempArray)
	}
	day03countTrees(day03Matrix)
	day03countMultipleSlopes(day03Matrix)
}

func day03countTrees(val [][31]bool) {
	var numberOfTrees int = 0
	var position int = 1
	for i, v := range val {

		if i != 0 {
			//log.Printf("%v", day03DumpRow(v, position))
			if v[position-1] {
				numberOfTrees++
			}
		}
		position += 3
		if position > 31 {
			position -= 31
		}
	}
	log.Printf("Day 3 Solution 1 %v", numberOfTrees)
}

func day03countMultipleSlopes(val [][31]bool) {
	var numberOfTrees int = 0
	var position int = 1
	otherSlopes := [5][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	var result int = 1
	for _, slopeDetails := range otherSlopes {
		numberOfTrees = 0
		position = 1
		for i, v := range val {
			if (i%2 != 0) && (slopeDetails[1] == 2) {
				continue
			}
			if i != 0 {
				if v[position-1] {
					numberOfTrees++
				}
			}
			position += slopeDetails[0]
			if position > 31 {
				position -= 31
			}
		}
		result *= numberOfTrees
	}
	log.Printf("Day 3 Solution 2 %v", result)
}

func day03DumpRow(val [31]bool, position int) string {
	var sb strings.Builder
	red := color.New(color.FgHiRed).SprintFunc()
	for i, w := range val {
		if i == (position - 1) {
			if w {
				sb.WriteString(red("#"))
			} else {
				sb.WriteString(red("."))
			}
		} else {
			if w {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
	}
	return sb.String()
}

func day03DumpMatrix(val [][31]bool) {
	for _, v := range val {
		fmt.Println()
		for _, w := range v {
			if w {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println("\n---")
}
