package solutions

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Day04solution1 first part
func Day04solution1() {
	file, err := os.Open("puzzles/day04.txt")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	var passports []map[string]string
	passports = make([]map[string]string, 0)
	var tempPassport map[string]string
	tempPassport = make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		e := scanner.Text()
		if e != "" {
			elements := strings.Split(e, " ")
			for _, v := range elements {
				attribute := strings.Split(v, ":")
				tempPassport[attribute[0]] = attribute[1]
			}
		} else {
			// it is an empty line: start new passport
			passports = append(passports, tempPassport)
			tempPassport = make(map[string]string)
		}
	}
	passports = append(passports, tempPassport)
	day04CountValidPassports(passports)
	day04CountValidPassportsWithCheck(passports)
}

func day04CountValidPassports(val []map[string]string) {
	var validPassports int = 0
	for _, p := range val {
		if len(p) == 8 {
			validPassports++
		} else {
			if p["cid"] == "" {
				if len(p) == 7 {
					validPassports++
				}
			}
		}
	}
	log.Printf("Day 4 Solution 1: Number of valid passports %v", validPassports)
}

func day04CountValidPassportsWithCheck(val []map[string]string) {
	var validPassports int = 0
	for _, p := range val {
		if day04syntaxcheck(p) {
			validPassports++
		}
	}
	log.Printf("Day 4 Solution 2: Number of valid checked passports %v", validPassports)
}

var r *regexp.Regexp = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var id *regexp.Regexp = regexp.MustCompile(`^[0-9]{9}$`)
var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

/*
pid (Passport ID) - a nine-digit number, including leading zeroes.
cid (Country ID) - ignored, missing or not.
*/
func day04syntaxcheck(val map[string]string) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002
	byr, err := strconv.Atoi(val["byr"])
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020
	iyr, err := strconv.Atoi(val["iyr"])
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030
	eyr, err := strconv.Atoi(val["eyr"])
	if err != nil {
		return false
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	/*
		hgt (Height) - a number followed by either cm or in:
		If cm, the number must be at least 150 and at most 193.
		If in, the number must be at least 59 and at most 76.
	*/
	hgt := val["hgt"]
	if strings.HasSuffix(hgt, "cm") {
		cm, err := strconv.Atoi(hgt[:len(hgt)-2])
		if err != nil {
			return false
		}
		if cm < 150 || cm > 193 {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		in, err := strconv.Atoi(hgt[:len(hgt)-2])
		if err != nil {
			return false
		}
		if in < 59 || in > 76 {
			return false
		}
	} else {
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hcl := val["hcl"]
	if !r.Match([]byte(hcl)) {

		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	ecl := val["ecl"]
	if !checkEyeColor(ecl) {

		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes
	pid := val["pid"]
	if !id.Match([]byte(pid)) {
		log.Printf("invalid pid %v", pid)
		return false
	}

	return true
}
func checkEyeColor(val string) bool {
	for _, col := range eyeColors {
		if col == val {
			return true
		}
	}
	return false
}
