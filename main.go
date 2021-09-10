package main

import (
	"crypto/sha1"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"infotech.umm.ac.id/milab/config"
	"infotech.umm.ac.id/milab/engine"
	"infotech.umm.ac.id/milab/security"
	"os"
	"time"
)

func getCookieStore() []byte {
	h := sha1.New()
	return h.Sum([]byte(time.Now().String()))
}

func main() {
	db := config.DBInit()
	InDB := &engine.InDB{DB: db}
	router := gin.Default()
	router.Use(sessions.Sessions("U-ACCESS-SESSX", cookie.NewStore(getCookieStore())))
	router.Use(security.MDRequestBarrier())
	homepage := router.Group("/homepage").Use(security.MDAuthMapper("UAID"))
	profile := router.Group("/profile").Use(security.MDAuthMapper("U"))
	student := router.Group("/student").Use(security.MDAuthMapper("U"))
	assistant := router.Group("/assistant").Use(security.MDAuthMapper("A"))
	class := router.Group("/class").Use(security.MDAuthMapper("UAID"))
	leftStuff := router.Group("/left-stuff").Use(security.MDAuthMapper("UAID"))
	taskDetail := router.Group("/task-detail").Use(security.MDAuthMapper("A"))
	blacklist := router.Group("/report").Use(security.MDAuthMapper("AID"))
	module := router.Group("/module").Use(security.MDAuthMapper("UAID"))

	{
		homepage.GET("/", InDB.GetHomepageData)

		profile.GET("/", InDB.GetUserProfile)

		student.GET("/class", InDB.GetStudentClass)

		assistant.GET("/class", InDB.GetAssistantClass)

		class.GET("/:id/task", InDB.GetClassTask)
		class.GET("/:id/task/student", InDB.GetStudentClassTask)
		class.Use(security.MDAuthMapper("A")).GET("/task/:id", InDB.GetClassTaskWithDetail)
		class.Use(security.MDAuthMapper("A")).POST("/task", InDB.CreateClassTask)

		class.Use(security.MDAuthMapper("A")).Handle("LOCK", "/task/:id", InDB.LockClassTask)

		class.Use(security.MDAuthMapper("A")).PUT("/task/:id", InDB.UpdateClassTask)
		class.Use(security.MDAuthMapper("A")).GET("/task/presence/:id", InDB.GetClassTaskPresenceWithDetail)
		class.Use(security.MDAuthMapper("A")).POST("/task/presence", InDB.CreateClassPresence)
		class.Use(security.MDAuthMapper("A")).PUT("/task/presence/:id/student", InDB.UpdateClassPresenceStudent)

		leftStuff.GET("/", InDB.GetLeftStuff)
		leftStuff.GET("/:id", InDB.GetLeftStuffDetails)

		taskDetail.PUT("/:id", InDB.UpdateTaskDetail)

		blacklist.GET("/category", InDB.GetBlacklistCategory)
		blacklist.POST("/", InDB.CreateBlacklist)

		module.Handle("CHECK", "/:tid", InDB.CheckFile)
		module.GET("/:tid", InDB.DownloadFile)
		module.POST("/:tid", InDB.UploadFile)
	}

	err := router.RunTLS(":6667", os.Getenv("TLS_CERT_FILE"), os.Getenv("TLS_KEY_FILE"))
	if err != nil {
		panic("Error When Running")
	}

	//router.Run(":6667")
}
