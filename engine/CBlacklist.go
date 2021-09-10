package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

func (idb *InDB) GetBlacklistCategory(c *gin.Context) {
	var (
		data []structs.BlackListCategory
		err  error
	)

	o := c.Query("offset")
	data, err = GetBlacklistCategories(idb, o)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) CreateBlacklist(c *gin.Context) {
	var (
		data structs.BlackList
		err  error
	)

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = CreateBlacklist(idb, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
