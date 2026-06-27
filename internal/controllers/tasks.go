package controllers

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/smtnn-ks/pms/internal/models"
	"github.com/smtnn-ks/pms/internal/views/base"
	tasks_views "github.com/smtnn-ks/pms/internal/views/tasks"
)

type TaskController struct{}

func NewTaskRouter() chi.Router {
	taskController := &TaskController{}

	r := chi.NewRouter()

	r.Get("/", taskController.Get)

	return r
}

func (c *TaskController) Get(w http.ResponseWriter, r *http.Request) {
	tasks := []models.Task{
		{
			ID:    1,
			Title: "task 1",
		},
		{
			ID:    2,
			Title: "task 2",
		},
		{
			ID:    3,
			Title: "task 3",
		},
	}

	ctx := templ.WithChildren(r.Context(), tasks_views.List(tasks))

	err := base.
		AppIndex(base.AppIndexData{TitlePrefix: "Welcome"}).
		Render(ctx, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to render template: %s", err)
	}
}
