package solutions

import (
	"strconv"
	"strings"

	"github.com/codebycaleb/advent-go/common"
)

func Day04() common.Day {

	// passport struct
	type Passport struct {
		byr string
		iyr string
		eyr string
		hgt string
		hcl string
		ecl string
		pid string
	}

	// parse input into a list of passport structs
	parsePassports := func(input string) []Passport {
		passports := []Passport{}
		for _, entry := range strings.Split(input, "\n\n") {
			passport := Passport{}
			for _, field := range strings.Fields(entry) {
				kv := strings.Split(field, ":")
				switch kv[0] {
				case "byr":
					passport.byr = kv[1]
				case "iyr":
					passport.iyr = kv[1]
				case "eyr":
					passport.eyr = kv[1]
				case "hgt":
					passport.hgt = kv[1]
				case "hcl":
					passport.hcl = kv[1]
				case "ecl":
					passport.ecl = kv[1]
				case "pid":
					passport.pid = kv[1]
				}
			}
			passports = append(passports, passport)
		}
		return passports
	}

	// check all passport fields exist
	allFields := func(passport Passport) bool {
		return passport.byr != "" &&
			passport.iyr != "" &&
			passport.eyr != "" &&
			passport.hgt != "" &&
			passport.hcl != "" &&
			passport.ecl != "" &&
			passport.pid != ""
	}

	// validate all passport fields
	validate := func(passport Passport) bool {
		// parse int values
		byr, _ := strconv.Atoi(passport.byr)
		iyr, _ := strconv.Atoi(passport.iyr)
		eyr, _ := strconv.Atoi(passport.eyr)
		var hgt int
		if strings.HasSuffix(passport.hgt, "cm") {
			hgt, _ = strconv.Atoi(strings.TrimSuffix(passport.hgt, "cm"))
		} else if strings.HasSuffix(passport.hgt, "in") {
			hgt, _ = strconv.Atoi(strings.TrimSuffix(passport.hgt, "in"))
		} else {
			return false
		}

		// byr (Birth Year) - four digits; at least 1920 and at most 2002.
		if byr < 1920 || byr > 2002 {
			return false
		}

		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		if iyr < 2010 || iyr > 2020 {
			return false
		}

		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		if eyr < 2020 || eyr > 2030 {
			return false
		}

		// hgt (Height) - a number followed by either cm or in:
		// If cm, the number must be at least 150 and at most 193.
		// If in, the number must be at least 59 and at most 76.
		if strings.HasSuffix(passport.hgt, "cm") {
			if hgt < 150 || hgt > 193 {
				return false
			}
		} else if strings.HasSuffix(passport.hgt, "in") {
			if hgt < 59 || hgt > 76 {
				return false
			}
		} else {
			return false
		}

		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		if !strings.HasPrefix(passport.hcl, "#") || len(passport.hcl) != 7 {
			return false
		}
		for _, c := range passport.hcl[1:] {
			if c < '0' || c > '9' && c < 'a' || c > 'f' {
				return false
			}
		}

		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		if !(passport.ecl == "amb" || passport.ecl == "blu" || passport.ecl == "brn" || passport.ecl == "gry" || passport.ecl == "grn" || passport.ecl == "hzl" || passport.ecl == "oth") {
			return false
		}

		// pid (Passport ID) - a nine-digit number, including leading zeroes.
		if !(len(passport.pid) == 9 && strings.ContainsAny(passport.pid, "0123456789")) {
			return false
		}

		return true
	}

	part1 := func(input string) string {
		passports := parsePassports(input)
		valid := 0
		for _, passport := range passports {
			if allFields(passport) {
				valid++
			}
		}
		return strconv.Itoa(valid)
	}

	part2 := func(input string) string {
		passports := parsePassports(input)
		valid := 0
		for _, passport := range passports {
			if allFields(passport) && validate(passport) {
				valid++
			}
		}
		return strconv.Itoa(valid)
	}

	return common.Day{ID: "04", Part1: part1, Part2: part2}
}
