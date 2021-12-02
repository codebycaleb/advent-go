package main

import (
	"flag"
	"log"

	"github.com/codebycaleb/advent-go/common"
)

func createNewDay(year string, day string) {
	// Confirm new year flag is set
	if year == "" {
		log.Fatal("Please specify what year this day should be created in")
	}
	// Validate format of year
	if len(year) != 4 {
		log.Fatal("Year must be in the format YYYY")
	}
	// Confirm new day flag is set
	if day == "" {
		log.Fatal("Please specify a day to create")
	}
	// Validate format of day
	if len(day) != 2 {
		log.Fatal("Day must be in format DD")
	}
	common.CreateNewDayFiles(year, day)
}

func main() {
	// Flags

	// create flag
	var create bool
	flag.BoolVar(&create, "create", false, "create new day")

	// year flag
	var year string
	flag.StringVar(&year, "year", "", "year to create new day (must be in format YYYY)")

	// day flag
	var day string
	flag.StringVar(&day, "day", "", "day to create (must be in format DD)")

	// Parse flags
	flag.Parse()

	// Handle create flag
	if create {
		createNewDay(year, day)
		return
	}

	log.Fatal("Create flag not set; exiting")
}
