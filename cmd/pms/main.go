package main

import (
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/smtnn-ks/pms/internal/controllers"
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

	r := chi.NewRouter()

	r.Mount("/", controllers.NewTaskRouter())

	println("server is running on port 8000...")
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		println("server error:", err)
	}
}
