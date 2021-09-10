package engine

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/security"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
	"strconv"
	"time"
)

func (idb *InDB) GetStudentClassTask(c *gin.Context) {
	var (
		data structs.StudentClassWithAssociation
		err  error
	)

	accessData := security.GetSessionAccessData(c)
	studentUser, err := GetSingleUserWithPreload(idb, "id", strconv.Itoa(int(accessData.UID)))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	data, err = GetSingleStudentClassUsingID(idb, id, studentUser.UserStudents[0].ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetClassTask(c *gin.Context) {
	var (
		data structs.ClassWithAssociation
		err  error
	)

	id := c.Param("id")
	data, err = GetSingleClassUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetClassTaskWithDetail(c *gin.Context) {
	var (
		data structs.TaskWithAssociation
		err  error
	)

	id := c.Param("id")
	data, err = GetSingleAssociatedTaskUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetClassTaskPresenceWithDetail(c *gin.Context) {
	var (
		data structs.PresenceWithAssociation
		err  error
	)

	id := c.Param("id")
	data, err = GetSingleAssociatedPresenceUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) CreateClassPresence(c *gin.Context) {
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

func (idb *InDB) CreateClassTask(c *gin.Context) {
	var (
		data structs.Task
		err  error
	)

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = CreateTask(idb, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

type GradePoint struct {
	Category string `json:"category"`
	Weight   string `json:"weight"`
}

type GradeEachPoint struct {
	Category string `json:"category"`
	Grade    string `json:"grade"`
}

func (idb *InDB) LockClassTask(c *gin.Context) {
	var (
		data structs.StudentTaskWithAssociation
		gp   []GradePoint
		gep  []GradeEachPoint
		err  error
	)

	id := c.Param("id")
	data, err = GetSingleStudentAssociatedTaskUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//if data.LockDate != nil {
	//	c.JSON(http.StatusLoopDetected, "Already Locked")
	//	return
	//}

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal([]byte(data.GradePoints), &gp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	for _, v := range gp {
		nGep := GradeEachPoint{
			Category: v.Category,
			Grade:    "0",
		}
		gep = append(gep, nGep)
	}

	gradeJSON, err := json.Marshal(gep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	nTD, err := ScrambleAssistant(idb, data.TaskDetails, data.ClassID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	now := time.Now()
	data.LockDate = &now
	data.TaskDetails = nTD

	for i, _ := range data.TaskDetails {
		data.TaskDetails[i].GradeEachPoint = string(gradeJSON)
	}

	data, err = UpdateAssociatedTask(idb, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) UpdateClassTask(c *gin.Context) {
	var (
		data structs.Task
		err  error
	)

	id := c.Param("id")
	data, err = GetSingleTaskUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = UpdateTask(idb, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) UpdateClassPresenceStudent(c *gin.Context) {
	var (
		data structs.PresenceStudent
		err  error
	)

	id := c.Param("id")
	data, err = GetSinglePresenceStudentUsingID(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	data.ID = uint(_id)

	data, err = UpdatePresenceStudent(idb, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
