package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/MakanPisang/movie-website.git/models"
	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
	var movie models.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received release_date: %v", movie.ReleaseDate)

	parsedDate, err := time.Parse("2006-01-02", movie.ReleaseDate.Format("2006-01-02"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}
	movie.ReleaseDate = parsedDate

	if err := models.DB.Create(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}
