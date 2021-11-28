package utils

import (
	"log"
	"os"

	"github.com/DouglasSerena/solid-go/src/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Erro ao carregar o .env: %v", err)
		panic(err)
	}

	dialect := os.Getenv("DB_TYPE")
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(dialect, dsn)

	if os.Getenv("DB_LOG") == "true" {
		db.LogMode(true)
	}

	if err != nil {
		log.Fatalf("Erro ao se conectar no banco: %v", err)
		panic(err)
	}

	db.AutoMigrate(&domain.User{})

	return db
}
