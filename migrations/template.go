package migrations

import "fmt"


func versionedTemplate(v string) string {
	return fmt.Sprintf(template,
		v,
		v,
		v,
		v,
	)
}
const template = `package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

const version%s = "%s"

func (s *Syncker) migration_%s() *gormigrate.Migration {

	return &gormigrate.Migration{
		ID: version%s,
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			err := tx.AutoMigrate(nil)
			return err
		},
		Rollback: func(tx *gorm.DB) error {
			err := tx.Migrator(&People).DropTable()
			return err
		},
	}
}
`
