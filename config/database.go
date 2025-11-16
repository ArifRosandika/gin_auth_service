package config

import (
	"fmt"
	"learn_clean_architecture/internal/entity"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDB() *gorm.DB {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	host := viper.GetString("HOST")
	user := viper.GetString("USER")
	password := viper.GetString("PASSWORD")
	dbname := viper.GetString("DBNAME")
	port := viper.GetString("PORT")

	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("missing env variables")
		os.Exit(1)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect db", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&entity.User{})

	if err != nil {
		log.Fatal("failed to migrate db", err)
		os.Exit(1)
	}

	log.Println("database connected successfully")
	return db
}