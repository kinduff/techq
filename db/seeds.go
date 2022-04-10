package db

import (
	"bufio"
	"embed"
	"log"
	"reflect"

	"github.com/kinduff/tech_qa/models"
	"gorm.io/gorm"
)

// Seeder type
type Seed struct {
	db *gorm.DB
}

var (
	//go:embed seeds
	seeds embed.FS
)

// Execute will executes the given seeder method
func ExecuteSeed(db *gorm.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	// Execute all seeders if no method name is given
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// We are looping over the method on a Seed struct
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			callSeed(s, method.Name)
		}
	}

	// Execute only the given method names
	for _, item := range seedMethodNames {
		callSeed(s, item)
	}
}

func callSeed(s Seed, seedMethodName string) {
	// Get the reflect value of the method
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Execute the method
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succeeded")
}

func (s Seed) QuestionSeed() {
	// Get the file
	file, err := seeds.Open("seeds/questions.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Get the line
		line := scanner.Text()
		s.db.Create(&models.Question{Body: line})
	}
}
