package solutions_test

import (
	"testing"

	"github.com/codebycaleb/advent-go/y2020/solutions"
)

const input04 = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

func TestDay04Part1(t *testing.T) {
	const want = "4"
	result := solutions.Day04().Part1(input04)
	if result != want {
		t.Fatalf("Part 1 did not execute successfully (expected %v but got %v).", want, result)
	}
}

func TestDay04Part2(t *testing.T) {
	const want = "4"
	result := solutions.Day04().Part2(input04)
	if result != want {
		t.Fatalf("Part 2 did not execute successfully (expected %v but got %v).", want, result)
	}
}
