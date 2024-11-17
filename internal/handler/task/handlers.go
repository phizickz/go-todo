package task

import (
	"fmt"
	taskEntity "go-todo/internal/entities/task"
	taskRepo "go-todo/internal/repository/task"
	"go-todo/web/views"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
)

type TaskRepository interface {
	RetrieveAll() []taskEntity.Task
	Create(task taskEntity.Task) error
	Retrieve(id int) (taskEntity.Task, error)
	Update(task taskEntity.Task)
	Delete(id int)
}
type TaskHandler struct {
	TaskRepository *taskRepo.TaskRepository
}

func (th *TaskHandler) Edit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
		return
	}
	if id < 0 {
		log.Fatal("Invalid ID", id)
		return
	}
	task, err := th.TaskRepository.Retrieve(id)
	if err != nil {
		log.Fatal(err)
		return
	}
	handler := templ.Handler(views.TaskEdit(task))
	handler.ServeHTTP(w, r)
}

func (th *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
		return
	}
	if id < 0 {
		log.Fatal("Invalid ID", id)
		return
	}
	task := taskEntity.Task{
		ID:    id,
		Title: r.Form["title"][0],
		Body:  r.Form["body"][0],
	}

	th.TaskRepository.Update(task)
	fmt.Println("Task updated:", id)
	w.Header().Set("HX-Trigger", "task-updated")
	w.WriteHeader(http.StatusOK)
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
	fmt.Println("Task created")
	w.Header().Set("HX-Trigger", "task-updated")
	w.WriteHeader(http.StatusOK)
}

func (th *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	handler := templ.Handler(views.Tasks(th.TaskRepository.RetrieveAll()))
	handler.ServeHTTP(w, r)
}

func (th *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
		return
	}
	if id < 0 {
		log.Fatal("Invalid ID", id)
		return
	}

	th.TaskRepository.Delete(id)

	fmt.Println("Task deleted")
	w.Header().Set("HX-Trigger", "task-updated")
	w.WriteHeader(http.StatusOK)
}

func NewTaskHandler(repo *taskRepo.TaskRepository) *TaskHandler {
	return &TaskHandler{
		TaskRepository: repo,
	}
}
