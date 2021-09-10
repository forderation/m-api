package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

func (idb *InDB) GetUserAllClass(c *gin.Context) {
	var (
		data []structs.Class
		err  error
	)

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//limit := c.Param("limit")
	//data, err = UserClasses(idb,limit)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetAllClassTask(c *gin.Context) {
	var (
		data []structs.Task
		err  error
	)

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//limit := c.Param("limit")
	//data, err = AllClassTask(idb,limit)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) ClassTaskDetail(c *gin.Context) {
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

func (idb *InDB) InputGradeStudents(c *gin.Context) {
	var (
		data []structs.UserStudentClass
		err  error
	)

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//limit := c.Param("limit")
	//data, err = GetAllUserStudentClass(idb, limit)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}

	c.JSON(http.StatusOK, data)
	return
}
