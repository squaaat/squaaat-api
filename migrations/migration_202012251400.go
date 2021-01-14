package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

const version202012251400 = "202012251400"

func (s *Syncker) migration_202012251400() *gormigrate.Migration {

	return &gormigrate.Migration{
		ID: version202012251400,
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type Person struct {
				gorm.Model
				Name string
			}

			return tx.AutoMigrate(&Person{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("people")
		},
	}
}
