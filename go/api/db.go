package api

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable",
		"paungram",
		"paungram",
		"paungram",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Ошибка подключения к БД: ", err)
	}

	// Автомиграция
	if err := AutoMigrate(db); err != nil {
		log.Fatal("❌ Ошибка миграции: ", err)
	}

	return db
}
