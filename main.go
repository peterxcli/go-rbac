package main

import (
	"easy-rbac/config"
	"easy-rbac/models"
	"easy-rbac/routers"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile("config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading config file:", err)
	}
	if err := viper.Unmarshal(&config.Config); err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return
	}
}

func main() {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName,
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})

	r := routers.SetupRouter(db)

	addr := fmt.Sprintf("%s:%s", config.Config.HttpServer.Hostname, config.Config.HttpServer.Port)
	fmt.Println("Server is running at", addr)

	r.Run(addr)
}
