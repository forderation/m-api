package engine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

var formFieldChooseTask struct {
	ClassId string `json:"class_id" binding:"required"`
}

func (idb *InDB) GetChooseTaskByClassId(c *gin.Context) {
	statusCodeNotFound := "ERR_CHOOSE_TASK_NOT_FOUND"
	statusNotFoundMessage := "Tugas kelas tidak ditemukan"
	statusFoundMessage := "Tugas kelas ditemukan"
	var (
		dataTask []structs.TaskWithAssociation
		err      error
	)
	if err := c.ShouldBind(&formFieldChooseTask); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			c.JSON(http.StatusBadRequest, fmt.Sprint(fieldErr))
			return
		}
	}
	classId := formFieldChooseTask.ClassId
	dataTask, err = GetClassTaskByClassId(idb, classId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(dataTask) < 1 {
		c.JSON(http.StatusNotFound, StandardResponse{
			Code:    statusCodeNotFound,
			Message: statusNotFoundMessage,
		})
		return
	}
	c.JSON(http.StatusOK, StandardResponse{
		Code:    CodeOkResponse,
		Message: statusFoundMessage,
		Data:    dataTask,
	})
	return
}
