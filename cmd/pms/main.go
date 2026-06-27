package main

import (
	"github.com/glebarez/sqlite"
	"github.com/smtnn-ks/pms/internal/models"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db/dev.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	println("successfully connected to database")

	err = db.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.UserTeam{},
		&models.Task{},
		&models.TaskChangelog{},
		&models.Comment{},
	)
	if err != nil {
		panic(err)
	}

	println("migrations run successfully")
}
