package db

import (
	"log"

	"github.com/Brun0Nasc/sys-projetos/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	//Migração dos modelos definidos na API para o banco de dados PostgreSQL
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Equipe{})
	db.AutoMigrate(&models.Pessoa{})
	db.AutoMigrate(&models.Projeto{})
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.User{})

	return db
}