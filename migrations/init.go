package migrations

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"io/ioutil"
	"log"
	"os"

	"github.com/squaaat/squaaat-api/internal/app"
)

type Syncker struct {
	App *app.Application
}

func New(a *app.Application) *Syncker {
	return &Syncker{
		App: a,
	}
}

func (s *Syncker) Sync() {
	m := gormigrate.New(
		s.App.ServiceDB.DB,
		gormigrate.DefaultOptions,
		s.load(),
	)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}

func (s *Syncker) Create(v string) {
	tmpl := versionedTemplate(v)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dest := fmt.Sprintf("%s/migrations/migration_%s.go", pwd, v)
	err = ioutil.WriteFile(dest, []byte(tmpl), 0644)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	msg := `
Completed to create migrations file

* Migration File: 
Link => file://%s/migrations/migration_%s.go

* Add Migration method
Link => file://%s/migrations/init.go

exmaple: 
func (s *Syncker) load() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		s.migration_202012251400(),
		// Write your codes here
		... ,
	}
}

`
	fmt.Printf(
		msg,
		pwd,
		v,
		pwd,
		)
}

func (s *Syncker) load() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		s.migration_202012251400(),

		// migration script
	}
}
