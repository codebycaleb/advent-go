package solutions

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Day struct {
	ID    string
	Part1 func(string) string
	Part2 func(string) string
}

func (d Day) Run() {
	// load the file
	location := fmt.Sprintf("inputs/%v.txt", d.ID)
	data, err := os.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day %v Part 1: %v\n", d.ID, d.Part1(string(data)))
	log.Printf("Day %v Part 2: %v\n", d.ID, d.Part2(string(data)))
}

// Write boilerplate solution and test files for the day
func CreateNewDayFiles(id string) {
	day := Day{ID: id}

	// read template files
	templates, err := template.ParseFiles("utils/templates/id.txt", "utils/templates/id_test.txt")
	if err != nil {
		panic(err)
	}

	// open writer file for day
	file, err := os.Create(fmt.Sprintf("solutions/%v.go", id))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// open writer file for tests
	fileTest, err := os.Create(fmt.Sprintf("solutions/tests/%v_test.go", id))
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
	_, err = os.Create(fmt.Sprintf("inputs/%v.txt", id))
	if err != nil {
		panic(err)
	}
}
