package db

import (
    "log"

    "github.com/ajit-kerle/go-gin-api-medium/pkg/common/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init(url string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(url))

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Book{})
    db.AutoMigrate(&models.Author{})
    db.AutoMigrate(&models.Publisher{})

    return db
}