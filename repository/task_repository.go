package repository

import (
	"database/sql"
	"fmt"
	"tostrike/model"
)

type TaskRepository struct {
	connection *sql.DB
}

func NewTaskRepository(connection *sql.DB) TaskRepository {
	return TaskRepository{connection}
}

func (tr *TaskRepository) GetTask() ([]model.Task, error) {
	query := "SELECT id, title, description, deadline, isComplete FROM task"
	rows, err := tr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Task{}, err
	}

	var taskList []model.Task
	var taskObj model.Task

	for rows.Next() {
		err = rows.Scan(
			&taskObj.ID,
			&taskObj.Title,
			&taskObj.Description,
			&taskObj.Deadline,
			&taskObj.IsCompleted)
		if err != nil {
			fmt.Println(err)
			return []model.Task{}, err
		}

		taskList = append(taskList, taskObj)
	}

	rows.Close()
	return taskList, nil
}

func (tr *TaskRepository) GetTaskById(id int) (*model.Task, error) {
	query, err := tr.connection.Prepare("SELECT * FROM task" + " WHERE id=$1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var taskObj model.Task
	err = query.QueryRow(id).Scan(
		&taskObj.ID,
		&taskObj.Title,
		&taskObj.Description,
		&taskObj.Deadline,
		&taskObj.IsCompleted,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	query.Close()
	return &taskObj, nil
}

func (tr *TaskRepository) CreateTask(task model.Task) (int, error) {

	var id int

	query, err := tr.connection.Prepare("INSERT INTO task" +
		"(title, description, deadline, isComplete) " +
		"VALUES($1,$2,$3,$4) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(task.Title, task.Description, task.Deadline, task.IsCompleted).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
