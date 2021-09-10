package engine

import (
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
)

var formFieldGetSchedule struct {
	ClassCategoryId string `json:"class_category_id" binding:"required"`
}

var ScheduleTimeArray = [...]string{"", "07:00", "07:50", "08:40", "09:30", "10:20", "12:10", "13:00", "13:50", "15:15", "16:05", "16:55", "18:15", "19:05", "19:55", "20:45"}
var ScheduleDay = [...]string{"", "Senin", "Selasa", "Rabu", "Kamis", "Jum\"at", "Sabtu"}
var ScheduleRoom = [...]string{"", "Lab A", "Lab B", "Lab C", "Lab D", "Lab E", "Lab F"}

type ClassScheduleResponse struct {
	ClassID   uint
	FullName  string
	Location  string
	TimeFrom  string
	TimeUntil string
	Day       string
	Week      int
}

func (idb *InDB) GetScheduleActiveByClassCategory(c *gin.Context) {
	statusCodeNotFound := "ERR_GET_SCHEDULE_ALL_CLASS"
	statusNotFoundMessage := "Jadwal kategori kelas tidak ditemukan"
	statusFoundMessage := "Jadwal kategori kelas ditemukan"
	var (
		dataClass []structs.ClassWithAssociation
		err       error
	)
	if err = c.ShouldBind(&formFieldGetSchedule); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	classCategoryId := formFieldGetSchedule.ClassCategoryId
	dataClass, err = GetAllClassActive(idb, &classCategoryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(dataClass) < 1 {
		c.JSON(http.StatusNotFound, StandardResponse{
			Code:    statusCodeNotFound,
			Message: statusNotFoundMessage,
		})
		return
	}
	response := make(map[string][]ClassScheduleResponse)
	for _, c := range dataClass {
		if c.ScheduleLocation != 0 {
			r := ClassScheduleResponse{
				ClassID:   c.ID,
				FullName:  c.FullName,
				Location:  ScheduleRoom[c.ScheduleLocation],
				Week:      c.ScheduleWeeks,
				Day:       ScheduleDay[c.ScheduleWeek],
				TimeFrom:  ScheduleTimeArray[c.ScheduleTimeFrom],
				TimeUntil: ScheduleTimeArray[c.ScheduleTimeUntil+1],
			}
			key := ScheduleDay[c.ScheduleWeek]
			response[key] = append(response[key], r)
		}
	}
	c.JSON(http.StatusOK, StandardResponse{
		Code:    CodeOkResponse,
		Message: statusFoundMessage,
		Data:    response,
	})
	return
}

func (idb *InDB) GetScheduleActiveClass(c *gin.Context) {
	statusCodeNotFound := "ERR_GET_SCHEDULE_ALL_CLASS"
	statusNotFoundMessage := "Jadwal semua kelas aktif tidak ditemukan"
	statusFoundMessage := "Jadwal semua kelas aktif ditemukan"
	var (
		dataClass []structs.ClassWithAssociation
		err       error
	)
	dataClass, err = GetAllClassActive(idb, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(dataClass) < 1 {
		c.JSON(http.StatusNotFound, StandardResponse{
			Code:    statusCodeNotFound,
			Message: statusNotFoundMessage,
		})
		return
	}
	response := make(map[string][]ClassScheduleResponse)
	for _, c := range dataClass {
		if c.ScheduleLocation != 0 {
			r := ClassScheduleResponse{
				ClassID:   c.ID,
				FullName:  c.FullName,
				Location:  ScheduleRoom[c.ScheduleLocation],
				Week:      c.ScheduleWeeks,
				Day:       ScheduleDay[c.ScheduleWeek],
				TimeFrom:  ScheduleTimeArray[c.ScheduleTimeFrom],
				TimeUntil: ScheduleTimeArray[c.ScheduleTimeUntil+1],
			}
			key := ScheduleDay[c.ScheduleWeek]
			response[key] = append(response[key], r)
		}
	}
	c.JSON(http.StatusOK, StandardResponse{
		Code:    CodeOkResponse,
		Message: statusFoundMessage,
		Data:    response,
	})
	return
}
