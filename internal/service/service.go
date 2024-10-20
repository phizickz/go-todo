package service

type TaskService interface {
	Create(ctx context.Context, task *task.Task) error
	#Update(ctx context.Context, task *task.Task) error
	#Delete()
	#FindByID()
}

type Service struct {
	TaskService TaskService
}
