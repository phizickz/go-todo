package storage

type TaskRepository interface {
	Create(ctx context.Context, task *task.Task) error
	Update(ctx context.Context, task *task.Task) error
}



type Repository struct {
	Task TaskRepository
}
