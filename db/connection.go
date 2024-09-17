package db

import (
	"fmt"

	// _ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api-teste/pkg/commons/models"
	"api-teste/pkg/configs"
)

func OpenConnection() (*gorm.DB, error) {

	config := configs.GetDB()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		config.Host,
		config.User,
		config.Pass,
		config.Database,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}

	// AutoMigrate(db)

	return db, err
}

func AutoMigrate(db *gorm.DB) {

	db.Debug().AutoMigrate(
		&models.Invoice{},
	)

}
