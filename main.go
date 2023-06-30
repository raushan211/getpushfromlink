package main

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type CatBreed struct {
	Breed   string `json:"breed"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}
type WordCountRequest struct {
	Str string `json:"str"`
}

func main() {
	router := gin.Default()
	router.POST("/", CheckWordCount)
	router.Run(":8080")
}
func CheckWordCount(c *gin.Context) {
	var request WordCountRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Using regex to count words
	regex := regexp.MustCompile(`\w+`)
	words := regex.FindAllString(request.Str, -1)
	if len(words) >= 8 {
		c.JSON(http.StatusOK, gin.H{"message": "200 OK"})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Not Acceptable"})
	}
}
