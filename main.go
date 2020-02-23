package main

import (
	"net/http"

	"spicyhack/utils"

	"github.com/gin-gonic/gin"
)

var frontends = []string{"React", "Vanila"}
var backends = []string{"Ruby on Rails", ".NET Core", "ExpressJS", "Django", "Flask", "PHP", "Laravel", "Spring Boot"}
var databases = []string{"PostgreSQL", "MySQL", "MongoDB", "DynamoDB", "SQLite", "Maria DB", "CouchDB", "ArangoDB"}

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	router.Run(":3000")
}

type Architecture struct {
	Frontend string `json:"frontend"`
	Backend  string `json:"backend"`
	Database string `json:"database"`
}

func rootHandler(context *gin.Context) {
	result := Architecture{
		Frontend: frontends[utils.GetRandomIndex(len(frontends))],
		Backend:  backends[utils.GetRandomIndex(len(backends))],
		Database: databases[utils.GetRandomIndex(len(databases))],
	}

	context.JSON(http.StatusOK, gin.H{
		"message":      "OK",
		"architecture": result,
	})
}
