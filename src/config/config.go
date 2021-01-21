package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT      = 0
	SECRETKEY []byte
	DBURL     = ""
	// DBDRIVER = ""
)

func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 8080
	}

	// fmt.Printf("hello %s %q %s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	//DBURL = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	DBURL = fmt.Sprintf("root:Suchil$123@tcp(localhost:3306)/golang?charset=utf8&parseTime=True&loc=Local")
	SECRETKEY = []byte(os.Getenv("API_SECRET"))

}
