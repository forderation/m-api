package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

func (idb *InDB) GetUserLeftStuffCategory(c *gin.Context) {
	var (
		data []structs.LeftStuffCategory
		err  error
	)

	limit := c.Param("limit")
	data, err = GetLeftStuffCategory(idb, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
