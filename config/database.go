package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"time"
)

type ConnectionData struct {
	Username string
	Password string
	DBName   string
	IP       string
	Port     string
}

func parseConnString(cs string) ConnectionData {
	t := strings.Split(cs, "|")

	return ConnectionData{
		Username: t[0],
		Password: t[1],
		DBName:   t[2],
		IP:       t[3],
		Port:     t[4],
	}
}

func DBInit() *gorm.DB {
	//err := godotenv.Load()
	//if err != nil {
	//	panic("GO ENV ERROR : " + err.Error())
	//}

	dbConString, ok := os.LookupEnv("MILAB_CONSTR")
	if !ok {
		panic("Invalid Connection String " + dbConString)
	}

	cs := parseConnString(dbConString)

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", cs.Username, cs.Password, cs.IP, cs.Port, cs.DBName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(0)
	err = db.AutoMigrate()

	if err != nil {
		panic(err)
	}

	return db
}
