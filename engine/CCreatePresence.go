package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

func (idb *InDB) AddPresence(c *gin.Context) {
	var (
		data structs.Presence
		err  error
	)

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = CreatePresence(idb, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
