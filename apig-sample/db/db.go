package db

import (
	"log"
	"os"
	//"path/filepath"
	"strings"

	"github.com/wantedly/apig-sample/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/serenize/snaker"
)

func Connect() *gorm.DB {
	// dir := filepath.Dir("db/database.db")
	// db, err := gorm.Open("sqlite3", dir+"/database.db")
	db, err := gorm.Open("postgres", "host=localhost user=apig-sample dbname=apig-sample sslmode=disable password=mypassword")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	db.LogMode(false)
	if gin.IsDebugging() {
		db.LogMode(true)
	}

	if os.Getenv("AUTOMIGRATE") == "1" {
		db.AutoMigrate(
			&models.Email{},
			&models.User{},
		)
	}
	return db
}

func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}

func SetPreloads(preloads string, db *gorm.DB) *gorm.DB {
	if preloads == "" {
		return db
	}

	for _, preload := range strings.Split(preloads, ",") {
		var a []string

		for _, s := range strings.Split(preload, ".") {
			a = append(a, snaker.SnakeToCamel(s))
		}

		db = db.Preload(strings.Join(a, "."))
	}

	return db
}
