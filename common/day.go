package common

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Day struct {
	YEAR  string
	ID    string
	Part1 func(string) string
	Part2 func(string) string
}

func (d Day) Run() {
	// load the file
	location := fmt.Sprintf("inputs/%v.txt", d.ID)
	data, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}

	log.Printf("Day %v Part 1: %v\n", d.ID, d.Part1(string(data)))
	log.Printf("Day %v Part 2: %v\n", d.ID, d.Part2(string(data)))
}

// Write boilerplate solution and test files for the day
func CreateNewDayFiles(year string, id string) {
	day := Day{YEAR: year, ID: id}

	// read template files
	templates, err := template.ParseFiles("common/utils/templates/id.txt", "common/utils/templates/id_test.txt")
	if err != nil {
		panic(err)
	}

	// open writer file for day
	file, err := os.Create(fmt.Sprintf("y%v/solutions/%v.go", year, id))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// open writer file for tests
	fileTest, err := os.Create(fmt.Sprintf("y%v/solutions/tests/%v_test.go", year, id))
	if err != nil {
		panic(err)
	}
	defer fileTest.Close()

	// write to files
	err = templates.ExecuteTemplate(file, "id.txt", day)
	if err != nil {
		panic(err)
	}
	err = templates.ExecuteTemplate(fileTest, "id_test.txt", day)
	if err != nil {
		panic(err)
	}

	// also touch input file
	_, err = os.Create(fmt.Sprintf("y%v/inputs/%v.txt", year, id))
	if err != nil {
		panic(err)
	}

	log.Printf("Created new day files for %v %v\n", year, id)
}
