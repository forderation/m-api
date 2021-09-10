package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

func (idb *InDB) SubmitKRS(c *gin.Context) {
	var (
		data structs.UserStudentClass
		err  error
	)

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = CreateUserStudentClass(idb, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
