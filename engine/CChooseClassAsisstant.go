package engine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

var formFieldChooseClassAssistant struct {
	UserId string `json:"user_id" binding:"required"`
}

func (idb *InDB) GetChooseClassAssistantByUserId(c *gin.Context) {
	statusCodeNotFound := "ERR_CHOOSE_CLASS_ASSISTANT_NOT_FOUND"
	statusNotFoundMessage := "Kelas tidak ditemukan"
	statusFoundMessage := "Kelas ditemukan"
	var (
		dataUserAssistant      []structs.UserAssistantWithAssociation
		dataAssistantClass     []structs.ClassWithAssociation
		dataUserAssistantClass []structs.UserAssistantClassWithAssociation
		err                    error
	)
	if err := c.ShouldBind(&formFieldChooseClassAssistant); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			c.JSON(http.StatusBadRequest, fmt.Sprint(fieldErr))
			return
		}
	}
	userId := formFieldChooseClassAssistant.UserId
	dataUserAssistant, err = GetUserAssistant(idb, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(dataUserAssistant) < 1 {
		c.JSON(http.StatusNotFound, StandardResponse{
			Code: statusCodeNotFound,
			Message: statusNotFoundMessage,
		})
		return
	}
	userAssistantId := []int{int(dataUserAssistant[0].ID)}
	dataUserAssistantClass, err = GetUserAssistantClass(idb, userAssistantId)
	if len(dataUserAssistantClass) < 1 {
		c.JSON(http.StatusNotFound, StandardResponse{
			Code: statusCodeNotFound,
			Message: statusNotFoundMessage,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	classesId := make([]int, len(dataUserAssistantClass))
	for i, c := range dataUserAssistantClass {
		classesId[i] = int(c.ClassID)
	}
	dataAssistantClass, err = GetClassByUserAssistant(idb, classesId)
	if len(dataAssistantClass) < 1 {
		c.JSON(http.StatusNotFound, StandardResponse{
			Code: statusCodeNotFound,
			Message: statusNotFoundMessage,
		})
		return
	}
	dataOutputs := make([]structs.ClassWithAssociation, 0)
	for _, c := range dataAssistantClass {
		if c.ClassCategory.Active == 1  {
			dataOutputs = append(dataOutputs, c)
		}
	}
	if len(dataOutputs) < 1 {
		c.JSON(http.StatusNotFound, StandardResponse{
			Code: statusCodeNotFound,
			Message: statusNotFoundMessage,
		})
		return
	}
	c.JSON(http.StatusOK, StandardResponse{
		Code: CodeOkResponse,
		Message: statusFoundMessage,
		Data: dataOutputs,
	})
	return
}
