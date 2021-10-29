package main

import (
	"log"
	"time"

	"github.com/codebycaleb/advent-go/y2020/solutions"
)

func main() {
	start := time.Now()

	log.Println("Advent of Code")
	log.Println("--------------")

	solutions.Day1().Run()
	solutions.Day2().Run()

	elapsed := time.Since(start)
	log.Printf("Runtime: %s", elapsed)
}
