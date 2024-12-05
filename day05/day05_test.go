package day05

import (
	"testing"
)

var input = []string{
	"47|53",
	"97|13",
	"97|61",
	"97|47",
	"75|29",
	"61|13",
	"75|53",
	"29|13",
	"97|29",
	"53|29",
	"61|53",
	"97|53",
	"61|29",
	"47|13",
	"75|47",
	"97|75",
	"47|61",
	"75|61",
	"47|29",
	"75|13",
	"53|13",
	"",
	"75,47,61,53,29",
	"97,61,53,29,13",
	"75,29,13",
	"75,97,47,61,53",
	"61,13,29",
	"97,13,75,29,47",
}

func TestPart1(t *testing.T) {
	rules := Parse(input)

	res := part1(rules)
	if (res != 143) {
		t.Errorf("Expected 143, Got %d", res)
	}
}

func TestPart2(t *testing.T) {
	rules := Parse(input)

	res := part2(rules)
	if (res != 123) {
		t.Errorf("Expected 123, Got %d", res)
	}
}
