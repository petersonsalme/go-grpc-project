package db

import (
	"log"

	"github.com/petersonsalme/go-grpc-project/go-grpc-product-svc/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.StockDecreaseLog{})

	return Handler{db}
}
