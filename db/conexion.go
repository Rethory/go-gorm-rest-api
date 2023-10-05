package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=javiernajera password=javiernajera dbname=javiernajera port=5432"
var DB *gorm.DB

func DBConexion() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("BD conectada")
	}
}