package db

import (
	"log"
	"metabeat/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func InitDB() *DB {
	dsn := "host=localhost user=postgres password=password dbname=metabeat port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Database connection established")
	return &DB{Db: database}
}

func (DB *DB) Migrate() {
	log.Println("Running migration")
	err := DB.Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("An error occurred during migration: %v", err)
	} else {
		log.Println("Migration completed successfully")
	}
}
