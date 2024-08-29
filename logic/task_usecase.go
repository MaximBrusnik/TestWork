package logic

import (
	"context"
	"log"
	"time"

	"ApiRest/handlers"
	"ApiRest/model"
	"ApiRest/repository"
)

type taskUsecase struct {
	taskRepo       repository.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(t repository.TaskRepository, timeout time.Duration) handlers.TaskUsecase {
	return &taskUsecase{
		taskRepo:       t,
		contextTimeout: timeout,
	}
}

func (t *taskUsecase) GetTaskByID(c context.Context, id uint) (res []model.Task, err error) {
	task, err := t.taskRepo.GetTaskByID(c, id)
	if err == nil {
		log.Printf("Task %s is get", task)
	} else {
		log.Printf("Task is not get1")
	}
	return task, nil
}

func (t *taskUsecase) GetAllTask(ctx context.Context) (res []model.Task, err error) {
	tasks, err := t.taskRepo.GetAllTask(ctx)
	if err == nil {
		log.Printf("Tasks %s is get", tasks)
	} else {
		log.Printf("Tasks %s is not get", tasks)
	}
	return tasks, nil
}

func (t *taskUsecase) UpdateTask(ctx context.Context, id uint, task *model.Task) error {
	err := t.taskRepo.UpdateTask(ctx, id, task)
	if err == nil {
		log.Printf("Task %s is update", task)
	} else {
		log.Fatalf("Task %s is not update", task)
	}
	return nil
}

func (t *taskUsecase) DeleteTask(ctx context.Context, id uint) error {
	err := t.taskRepo.DeleteTask(ctx, id)
	if err == nil {
		log.Printf("Task with id = %d is deleted", id)
	} else {
		log.Printf("Task with id = %d is not deleted", id)
	}
	return nil
}

func (t *taskUsecase) AddTask(ctx context.Context, task *model.Task) error {
	err := t.taskRepo.AddTask(ctx, task)
	if err == nil {
		log.Printf("Task %s is added", task)
	} else {
		log.Fatalf("Task %s is not added", task)
	}
	return nil
}
