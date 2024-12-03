package day03

import (
	"testing"

)

const testinput = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
const testinput2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestPart1(t *testing.T) {
	resLine := ScanInput(testinput)
	lines := [][]Mul{
		resLine,
	}
	res := part1(lines)
	if res != 161 {
		t.Errorf("Expeted 161, Got %d", res)
	}
}

func TestPart2(t *testing.T) {
	lines := ScanInput2(testinput2)
	res := part2(lines)
	if res != 48 {
		t.Errorf("Expeted 48, Got %d", res)
	}
}
