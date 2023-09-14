package seed

import (
	"log"
	"rest/pkg/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	err := db.AutoMigrate(
		&models.Blog{},
	)

	if err != nil {
		log.Fatalln(err)
		return
	}
}
