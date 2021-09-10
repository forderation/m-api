package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

func (idb *InDB) GetAllClassPresence(c *gin.Context) {
	var (
		data []structs.Class
		//err  error
	)

	//limit := c.Param("limit")
	//data, err = GetClass(idb, limit)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetAllTaskPresence(c *gin.Context) {
	var (
		data []structs.Task
		//err  error
	)

	//limit := c.Param("limit")
	//data, err = GetAllTask(idb, limit)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetPresence(c *gin.Context) {
	var (
		data []structs.PresenceStudent
		err  error
	)

	limit := c.Param("limit")
	data, err = GetPresenceStudent(idb, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
