package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tostrike/model"
	"tostrike/usecase"
)

type TaskController struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskController(usecase usecase.TaskUsecase) TaskController {
	return TaskController{
		taskUsecase: usecase,
	}
}

func (t *TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := t.taskUsecase.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (t *TaskController) GetTaskById(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		response := model.Response{
			Message: "id is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "id must be a int",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tasks, err := t.taskUsecase.GetTaskById(taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if tasks == nil {
		response := model.Response{
			Message: "task not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (t *TaskController) CreateTask(ctx *gin.Context) {
	var task model.Task
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	insertedTask, err := t.taskUsecase.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
	}
	ctx.JSON(http.StatusCreated, insertedTask)
}
