package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/security"
	"net/http"
	"strconv"
)

func (idb *InDB) GetHomepageData(c *gin.Context) {
	var (
		data interface{}
		err  error
	)

	offset := c.Param("offset")
	if offset == "" {
		offset = "0"
	}

	blacklistData, err := GetBlackLists(idb, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	leftStuffData, err := GetLeftStuff(idb, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	accessData := security.GetSessionAccessData(c)

	user, err := GetSingleUser(idb, "id", strconv.Itoa(int(accessData.UID)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	data = gin.H{
		"BlacklistData": blacklistData,
		"LeftStuffData": leftStuffData,
		"KRSStatus":     1,
		"User":          user,
		"UserRole":      accessData.RO,
	}

	c.JSON(http.StatusOK, data)
	return
}
