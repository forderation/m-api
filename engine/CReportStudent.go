package engine

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"infotech.umm.ac.id/milab/structs"
	"net/http"
	"strconv"
)

func (idb *InDB) GetReportCategories(c *gin.Context) {
	statusCodeNotFound := "ERR_CHOOSE_REPORT_CATEGORIES"
	statusNotFoundMessage := "Kategori report tidak tersedia"
	statusFoundMessage := "Kategori report ditemukan"
	var (
		dataReportCategories []structs.BlackListCategory
		//err                  error
	)
	//dataReportCategories, err = GetBlacklistCategories(idb)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	if len(dataReportCategories) < 1 {
		c.JSON(http.StatusNotFound, StandardResponse{
			Code:    statusCodeNotFound,
			Message: statusNotFoundMessage,
		})
		return
	}
	c.JSON(http.StatusOK, StandardResponse{
		Code:    CodeOkResponse,
		Message: statusFoundMessage,
		Data:    dataReportCategories,
	})
	return
}

var formFieldReportStudent struct {
	UserId              string `json:"user_id" binding:"required"`
	BlackListCategoryId string `json:"black_list_category_id" binding:"required"`
	UserIdReporter      string `json:"user_id_reporter" binding:"required"`
}

func (idb *InDB) CreateReportStudent(c *gin.Context) {
	statusCodeNotFound := "ERR_CHOOSE_TASK_NOT_FOUND"
	statusNotFoundMessageBlackList := "Category report tidak ditemukan"
	statusNotFoundMessageUserId := "User id tidak ditemukan"
	statusNotFoundMessageUserReporterId := "User id reporter tidak ditemukan"
	statusSuccessMessage := "Berhasil melakukan report student"

	var (
		dataBlackListCategory []structs.BlackListCategory
		dataUser              interface{}
		dataUserReporter      interface{}
		err                   error
	)
	if err := c.ShouldBind(&formFieldReportStudent); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			c.JSON(http.StatusBadRequest, fmt.Sprint(fieldErr))
			return
		}
	}
	userId := formFieldReportStudent.UserId
	blackListCategoryId := formFieldReportStudent.BlackListCategoryId
	userIdReporter := formFieldReportStudent.UserIdReporter
	dataBlackListCategory, err = GetBlacklistCategoryById(idb, blackListCategoryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(dataBlackListCategory) < 1 {
		c.JSON(http.StatusNotFound, StandardResponse{
			Code:    statusCodeNotFound,
			Message: statusNotFoundMessageBlackList,
		})
		return
	}
	dataUser, err = GetSingleUser(idb, "id", userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, StandardResponse{
			Code:    statusCodeNotFound,
			Message: statusNotFoundMessageUserId,
			Data:    dataUser,
		})
		return
	}
	dataUserReporter, err = GetSingleUser(idb, "id", userIdReporter)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, StandardResponse{
			Code:    statusCodeNotFound,
			Message: statusNotFoundMessageUserReporterId,
			Data:    dataUserReporter,
		})
		return
	}
	parsedUserId, err := strconv.ParseUint(userId, 10, 32)
	parsedBlackListCategoryId, err := strconv.ParseUint(blackListCategoryId, 10, 32)
	parsedUserIdReporter, err := strconv.ParseUint(userIdReporter, 10, 32)
	createUserReport := structs.BlackList{
		UserId:              uint(parsedUserId),
		BlackListCategoryId: uint(parsedBlackListCategoryId),
		UserIdReporter:      uint(parsedUserIdReporter),
	}
	err = CreateReportStudent(idb, &createUserReport)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, StandardResponse{
		Code:    CodeOkResponse,
		Message: statusSuccessMessage,
		Data:    formFieldReportStudent,
	})
	return
}
