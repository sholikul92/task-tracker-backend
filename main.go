package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"task-tracker/db"
	"task-tracker/models"
	"task-tracker/router"

	"github.com/joho/godotenv"
)

func main() {
	dbConfig, err := LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	db, err := db.SetupDB(dbConfig)
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		panic(err)
	}

	router := router.SetupRouter(db)

	fmt.Println("Server runing in http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}

func LoadConfigFromEnv() (*models.DBConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	dbconfig := models.DBConfig{
		HOST:     os.Getenv("DB_HOST"),
		PORT:     port,
		USER:     os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		DBNAME:   os.Getenv("DB_NAME"),
	}

	return &dbconfig, nil
}
