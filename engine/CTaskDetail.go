package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

func (idb *InDB) UpdateTaskDetail(c *gin.Context) {
	var (
		data structs.TaskDetail
		err  error
	)

	id := c.Param("id")
	data, err = GetSingleTaskDetail(idb, "id", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = UpdateTaskDetail(idb, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
