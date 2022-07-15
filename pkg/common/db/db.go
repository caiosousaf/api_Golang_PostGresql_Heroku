package db

import (
    "log"

    "github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init(url string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Pessoa{})
    db.AutoMigrate(&models.Projeto{})
    db.AutoMigrate(&models.Equipe{})
    db.AutoMigrate(&models.Task{})
    db.AutoMigrate(&models.User{})
    return db
}