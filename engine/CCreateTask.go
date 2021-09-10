package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

func (idb *InDB) SubmitTask(c *gin.Context) {
	var (
		data structs.Task
		err  error
	)

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//data, err = CreateTask(idb, data)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetTaskDetail(c *gin.Context) {
	var (
		data interface{}
		err  error
	)

	id := c.Param("q1")
	data, err = GetSingleTaskDetail(idb, "id", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
