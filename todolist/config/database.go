package config

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DbConfig {
	dbConfig := DbConfig{
		Host:     "localhost",
		Port:     8989,
		User:     "lester",
		DBName:   "todolist",
		Password: "todolist",
	}
	return &dbConfig
}

func DbURL(dbConfig *DbConfig) string {
	return "lester:todolist@tcp(127.0.0.1:3306)/todolist?charset=utf8&parseTime=True&loc=Local"
	//return fmt.Sprintf(
	//	"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=",
	//	dbConfig.User,
	//	dbConfig.Password,
	//	dbConfig.Host,
	//	dbConfig.Port,
	//	dbConfig.DBName,
	//)
}
