package solutions

import (
	"bufio"
	"os"
	"sort"
	"strings"

	"github.com/mpvl/unique"
	log "github.com/sirupsen/logrus"
)

type bagRule struct {
	master string
	inner  []string
}

var rulesmap map[string]bagRule = make(map[string]bagRule)

func (b bagRule) printContent() {
	log.Printf("Bag %v contains", b.master)
	for _, v := range b.inner {
		log.Printf("\t :%v", v)
	}
}

func (b bagRule) containsColor(c string) bool {
	for _, v := range b.inner {
		if strings.Contains(v, c) {
			return true
		}
	}
	return false

}

// Day07solution1 first part
func Day07solution1() {
	file, err := os.Open("puzzles/day07.txt")
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()
	var rules []bagRule
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		iind := strings.Index(line, " bags contain ")
		masterbag := line[:iind]
		master := bagRule{master: masterbag}
		rest := line[(iind + 14) : len(line)-1]
		innerBags := strings.Split(rest, ", ")
		for _, v := range innerBags {
			master.inner = append(master.inner, v)
		}
		rules = append(rules, master)
	}

	// Rules are all parsed now

	for _, v := range rules {
		rulesmap[v.master] = v
	}
	t := day07RecurseBags("shiny gold", 0)
	log.Printf("Day 07 Solution 1: %v", day07GetTreeElements(*t))
	flatrules = make([]string, 0)
	day07FlattenTheTree(*t)
	sort.Strings(flatrules)
	unique.Strings(&flatrules)
	log.Printf("Day 07 Solution 1: %v", len(flatrules))

}

var flatrules []string

func day07FlattenTheTree(t Tree) {
	flatrules = append(flatrules, t.color)
	for _, v := range t.leafs {
		day07FlattenTheTree(*v)
	}
}

func day07GetTreeElements(t Tree) int {

	res := 1
	for _, v := range t.leafs {
		res += day07GetTreeElements(*v)
		//log.Printf("Leaf color %v", v.color)
	}
	return res
}

// Tree holds the bags
type Tree struct {
	color string
	rules []bagRule
	leafs []*Tree
}

func day07RecurseBags(color string, level int) *Tree {
	level++
	t := &Tree{color: color}
	for _, v := range rulesmap {
		if v.containsColor(color) {
			t.rules = append(t.rules, v)
			t.leafs = append(t.leafs, day07RecurseBags(v.master, level))
		}
	}
	return t
}
