package main

import (
	"github.com/wantedly/apig-sample/db"
	"github.com/wantedly/apig-sample/server"
)

// main ...
func main() {
	database := db.Connect()
	s := server.Setup(database)
	s.Run(":8080")
}
