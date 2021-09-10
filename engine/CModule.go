package engine

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jlaffaye/ftp"
	"infotech.umm.ac.id/milab/config"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func (idb *InDB) CheckFile(c *gin.Context) {
	taskID := c.Param("tid")
	folderName := fmt.Sprintf("./file_task/%x%s", md5.Sum([]byte(taskID+"abc")), taskID)
	con, err := ftp.Dial(config.FS_IP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer con.Quit()

	err = con.Login(config.FS_USER, config.FS_PASSWORD)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	es, err := con.List(folderName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(es) <= 0 {
		c.JSON(http.StatusBadRequest, "Not File Found")
		return
	}
	c.JSON(http.StatusOK, es)
	return
}

func (idb *InDB) DownloadFile(c *gin.Context) {
	tS := config.TEMP_STORAGE

	taskID := c.Param("tid")
	extQuery := c.Query("ext")
	ext := ""
	if extQuery == "" {
		ext = "pdf"
	}

	folderName := fmt.Sprintf("./file_task/%x%s", md5.Sum([]byte(taskID+"abc")), taskID)
	fileName := fmt.Sprintf("%x%smodule.%s", md5.Sum([]byte(taskID+"module_upload")), taskID, ext)

	con, err := ftp.Dial(config.FS_IP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer con.Quit()

	err = con.Login(config.FS_USER, config.FS_PASSWORD)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	es, err := con.List(folderName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(es) <= 0 {
		c.JSON(http.StatusBadRequest, "File Not Found")
		return
	}

	file, err := con.Retr(folderName + "/" + fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error()+" :: "+folderName+"/"+fileName)
		return
	}

	timestamp := time.Now()
	d := path.Join(tS, fmt.Sprint(timestamp.Nanosecond()))
	f, err := os.Create(d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.FileAttachment(d, "try1.png")
	_ = os.Remove(d)
	return
}

func (idb *InDB) UploadFile(c *gin.Context) {
	taskID := c.Param("tid")
	tS := config.TEMP_STORAGE
	folderName := fmt.Sprintf("./file_task/%x%s", md5.Sum([]byte(taskID+"abc")), taskID)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	fileExtension := strings.Split(file.Filename, ".")
	fileName := fmt.Sprintf("%x%smodule.%s", md5.Sum([]byte(taskID+"module_upload")), taskID, fileExtension[len(fileExtension)-1])

	tFN := path.Join(tS, file.Filename)
	err = c.SaveUploadedFile(file, tFN)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "3 :: "+err.Error())
		return
	}
	defer os.Remove(tFN)

	f, err := os.Open(tFN)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "4 :: "+err.Error())
		return
	}

	//Save via FTP

	con, err := ftp.Dial(config.FS_IP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "1 :: "+err.Error())
		return
	}

	err = con.Login(config.FS_USER, config.FS_PASSWORD)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "2 :: "+err.Error())
		return
	}
	defer con.Quit()

	err = con.MakeDir(folderName)
	if err != nil && err.Error() != "550 Create directory operation failed." {
		c.JSON(http.StatusInternalServerError, "6 :: "+err.Error())
		return
	}

	err = con.Stor(fmt.Sprint(folderName, "/", fileName), f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "5 :: "+err.Error()+" || "+fmt.Sprint(folderName, "/", fileName))
		return
	}

	defer os.Remove(tFN)

	c.JSON(http.StatusInternalServerError, fmt.Sprint("Success Save :: ", fmt.Sprint(folderName, "/", fileName)))
	return
}
