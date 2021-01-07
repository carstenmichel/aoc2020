package solutions

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type entry struct {
	minimum   int
	maximum   int
	character string
	password  string
}

func getEntry(in string) (*entry, error) {
	r := regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+)\ (.): (.+)`)
	res := r.FindStringSubmatch(in)
	min, err := strconv.Atoi(res[1])
	if err != nil {
		return nil, err
	}
	max, err := strconv.Atoi(res[2])
	if err != nil {
		return nil, err
	}

	return &entry{minimum: min, maximum: max, character: res[3], password: res[4]}, nil

}

// Day02solution1 first part
func Day02solution1() {
	file, err := os.Open("puzzles/day02.txt")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	var counter int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		e, err := getEntry(scanner.Text())
		if err != nil {
			log.Fatalf("Failed to read puzzle 2 : %v", err)
		}
		numberOfChars := strings.Count(e.password, e.character)
		if numberOfChars >= e.minimum && numberOfChars <= e.maximum {
			counter++
		}
	}
	log.Printf("Day 02 Solution 1 Result: %v", counter)
}

// Day02solution2 first part
func Day02solution2() {
	file, err := os.Open("puzzles/day02.txt")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	var counter int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		e, err := getEntry(scanner.Text())
		if err != nil {
			log.Fatalf("Failed to read puzzle 2 : %v", err)
		}
		if (string(e.password[e.minimum-1]) == e.character) != (string(e.password[e.maximum-1]) == e.character) {
			counter++
		}
	}
	log.Printf("Day 02 Solution 2 Result: %v", counter)
}
