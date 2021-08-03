package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todolist/config"
	"todolist/models"
	"todolist/routes"
)

var err error

func main(){

	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("status: ", err)
	}
	config.DB.AutoMigrate(&models.Todo{})

	r := routes.SetupRouter()
	r.Run()

}
