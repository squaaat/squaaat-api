package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"

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

func (s *Syncker) load() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		s.migration_202012251400(),
	}
}
