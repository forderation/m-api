package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

func (idb *InDB) GetLeftStuff(c *gin.Context) {
	var (
		data []structs.LeftStuff
		err  error
	)

	offset := c.Query("offset")
	if offset == "" {
		offset = "0"
	}

	data, err = GetLeftStuff(idb, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetLeftStuffDetails(c *gin.Context) {
	var (
		data []structs.LeftStuffWithAssociation
		err  error
	)

	id := c.Param("id")

	data, err = GetSingleLeftStuffUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
