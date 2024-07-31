package usecase

import (
	"tostrike/model"
	"tostrike/repository"
)

type TaskUsecase struct {
	repository repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) TaskUsecase {
	return TaskUsecase{
		repository: repo,
	}
}

func (tu *TaskUsecase) GetTasks() ([]model.Task, error) {
	tasks, err := tu.repository.GetTask()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tu *TaskUsecase) GetTaskById(id int) (*model.Task, error) {
	task, err := tu.repository.GetTaskById(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (tu *TaskUsecase) CreateTask(task model.Task) (model.Task, error) {
	taskId, err := tu.repository.CreateTask(task)
	if err != nil {
		return model.Task{}, err
	}

	task.ID = taskId
	return task, nil
}
