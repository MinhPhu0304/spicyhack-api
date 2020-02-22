package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var frontends = []string{"React", "Vanila"}
var backends = []string{"Ruby on Rails", ".NET Core", "ExpressJS", "Django", "Flask", "PHP", "Laravel", "Spring Boot"}
var databases = []string{"PostgreSQL", "MySQL", "MongoDB", "DynamoDB", "SQLite", "Maria DB", "CouchDB", "ArangoDB"}

func main() {
	router := gin.Default()
	router.GET("/spicy", serveGETPing)
	fmt.Println("Finish set up router")
	router.Run(":3000")
}

type Architecture struct {
	Frontend string `json:"frontend"`
	Backend  string `json:"backend"`
	Database string `json:"database"`
}

func serveGETPing(context *gin.Context) {
	result := Architecture{
		Frontend: frontends[getRandomIndex(len(frontends))],
		Backend:  backends[getRandomIndex(len(backends))],
		Database: databases[getRandomIndex(len(databases))],
	}

	context.JSON(http.StatusOK, gin.H{
		"message":      "OK",
		"architecture": result,
	})
}

func getRandomIndex(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
