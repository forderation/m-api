package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/security"
	"net/http"
	"strconv"
)

func (idb *InDB) GetUserProfile(c *gin.Context) {
	var (
		data interface{}
		err  error
	)

	accessData := security.GetSessionAccessData(c)

	data, err = GetSingleUserWithPreload(idb, "id", strconv.Itoa(int(accessData.UID)))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
