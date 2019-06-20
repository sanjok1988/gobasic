package controllers

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "github.com/revel/revel"
    "webapp/app/models"
    "log"
)

var DB *gorm.DB

func InitDB() {
    dbInfo, _ := revel.Config.String("db.info")
    db, err := gorm.Open("mysql", dbInfo) //it reads the DB settings from "db.info" line of the configuration file, and open the database using Gorm library
    if err !=nil {
        log.Panicf("failed gorm.Open: %v\n", err)
    }

    db.DB()
    db.AutoMigrate(&models.Post{})  //creates a table from Post model
    DB = db //assigns the database handle to "DB" variable so that other files can access the database as "controllers.DB".
}

//Next, edit "app/init.go" to call the InitDB() function.