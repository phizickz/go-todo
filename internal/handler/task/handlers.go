package task

import (
	"fmt"
	taskEntity "go-todo/internal/entities/task"
	taskRepo "go-todo/internal/repository/task"
	"go-todo/web/views"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

type TaskRepository interface {
	RetrieveAll() []taskEntity.Task
	Create(task taskEntity.Task) error
	Retrieve(id int) (taskEntity.Task, error)
	// Update(id int, task taskEntity.Task)
	Delete(id int)
}
type TaskHandler struct {
	TaskRepository *taskRepo.TaskRepository
}

func (th *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if err := th.TaskRepository.Create(taskEntity.Task{
		Title: r.Form["title"][0],
		Body:  r.Form["body"][0],
	},
	); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("HX-Trigger", "task-created")
	w.WriteHeader(http.StatusOK)
}

func (th *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	handler := templ.Handler(views.Tasks(th.TaskRepository.RetrieveAll()))
	handler.ServeHTTP(w, r)
}

func (th *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(r.Body)
	fmt.Println(r.FormValue("value"))
}

func NewTaskHandler(repo *taskRepo.TaskRepository) *TaskHandler {
	return &TaskHandler{
		TaskRepository: repo,
	}
}
