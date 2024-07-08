package db

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateDataBase() (*gorm.DB, error) {
	viper.AutomaticEnv()

	pguser := viper.GetString("POSTGRES_USER")
	pgpassword := viper.GetString("POSTGRES_PASSWORD")
	pgbd := viper.GetString("POSTGRES_DB")
	pghostname := viper.GetString("POSTGRES_HOSTNAME")
	pgport := viper.GetString("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", pghostname, pguser, pgpassword, pgbd, pgport)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Init() {
	db, err := CreateDataBase()

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
