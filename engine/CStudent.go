package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/security"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
	"strconv"
)

func (idb *InDB) GetStudentClass(c *gin.Context) {
	var (
		data structs.UserWithAssociation
		err  error
	)

	accessData := security.GetSessionAccessData(c)

	data, err = GetSingleUserWithPreload(idb, "id", strconv.Itoa(int(accessData.UID)))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userStudentClasses, err := GetUserStudentClasses(idb, strconv.Itoa(int(data.UserStudents[0].ID)))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var activeUSC []structs.UserStudentClassWithAssociation

	for _, e := range userStudentClasses {
		if e.Class.Active == true && e.Class.ClassCategory.Active == 1 {
			activeUSC = append(activeUSC, e)
		}
	}

	c.JSON(http.StatusOK, activeUSC)
	return
}
