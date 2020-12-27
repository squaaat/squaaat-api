package migrations

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/rs/zerolog/log"
	"github.com/squaaat/squaaat-api/internal/model"
	"gorm.io/gorm"
	"io/ioutil"
	"os"

	"github.com/squaaat/squaaat-api/internal/app"
)

type Syncker struct {
	App *app.Application
	GormMigrator *gormigrate.Gormigrate
}

func New(a *app.Application) *Syncker {
	s := &Syncker{
		App: a,
	}
	s.GormMigrator = gormigrate.New(
		a.ServiceDB.DB,
		gormigrate.DefaultOptions,
		s.load(),
	)
	s.GormMigrator.InitSchema(func(m *gorm.DB) error {
		return m.AutoMigrate(model.Load()...)
	})
	return s
}

func (s *Syncker) Sync() {
	if err := s.GormMigrator.Migrate(); err != nil {
		log.Fatal().Err(err).Msg("Could not migrate")
	}
	log.Info().Msg("Migration did run successfully")
}

func (s *Syncker) Create(v string) {
	tmpl := versionedTemplate(v)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	dest := fmt.Sprintf("%s/migrations/migration_%s.go", pwd, v)
	err = ioutil.WriteFile(dest, []byte(tmpl), 0644)
	if err != nil {
		log.Fatal().Err(err).Send()
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
