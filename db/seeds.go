package db

import (
	"bufio"
	"embed"
	"reflect"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/kinduff/techq/internal/models"
)

// Seeder type
type Seed struct {
	db *gorm.DB
}

var (
	//go:embed seeds
	seeds embed.FS
)

// ExecuteSeed will executes the given seeder methods.
func ExecuteSeed(db *gorm.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	if len(seedMethodNames) == 0 {
		log.WithField("event", "seeder").Info("Running all seeders")
		for i := 0; i < seedType.NumMethod(); i++ {
			method := seedType.Method(i)
			callSeed(s, method.Name)
		}
	}

	for _, item := range seedMethodNames {
		callSeed(s, item)
	}
}

func callSeed(s Seed, seedMethodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	if !m.IsValid() {
		log.WithFields(log.Fields{
			"seedMethodName": seedMethodName,
			"event":          "seeder",
		}).Info("Undefined seed")
	}

	log.WithFields(log.Fields{
		"seedMethodName": seedMethodName,
		"event":          "seeder",
	}).Info("Seeding")

	m.Call(nil)

	log.WithFields(log.Fields{
		"seedMethodName": seedMethodName,
		"event":          "seeder",
	}).Info("Seeding succeeded")
}

// QuestionSeed method seeds the questions from the txt file.
func (s Seed) QuestionSeed() {
	file, err := seeds.Open("seeds/questions.txt")
	if err != nil {
		log.WithFields(log.Fields{
			"event": "seeder",
		}).Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		s.db.Create(&models.Question{Body: line})
	}
}
