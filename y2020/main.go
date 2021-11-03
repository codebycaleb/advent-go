package main

import (
	"flag"
	"log"
	"sort"
	"time"

	"github.com/codebycaleb/advent-go/y2020/solutions"
)

var completedSolutions = map[string]solutions.Day{
	"01": solutions.Day01(),
	"02": solutions.Day02(),
	"03": solutions.Day03(),
	"04": solutions.Day04(),
}

func createNewDay(day string) {
	// Confirm new day flag is set
	if day == "" {
		log.Fatal("Please specify a day to create")
	}
	// Validate format of day
	if len(day) != 2 {
		log.Fatal("Day must be in format DD")
	}
	solutions.CreateNewDayFiles(day)
}

func main() {
	// Flags

	// create flag
	var create bool
	flag.BoolVar(&create, "create", false, "create new day")

	// day flag
	var day string
	flag.StringVar(&day, "day", "", "day to create / run (must be in format DD)")

	// Parse flags
	flag.Parse()

	// Handle create flag
	if create {
		createNewDay(day)
		return
	}

	// Handle day flag
	var daysToRun []solutions.Day
	if day != "" {
		daysToRun = append(daysToRun, completedSolutions[day])
	} else {
		ids := []string{}
		for id := range completedSolutions {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		for _, id := range ids {
			daysToRun = append(daysToRun, completedSolutions[id])
		}
	}

	start := time.Now()

	log.Println("Advent of Code")
	log.Println("--------------")

	for _, day := range daysToRun {
		day.Run()
	}

	elapsed := time.Since(start)
	log.Printf("Runtime: %s", elapsed)
}
