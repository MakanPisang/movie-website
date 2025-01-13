package controllers

import (
	"net/http"
	"strconv"

	"github.com/MakanPisang/movie-website.git/models"
	"github.com/gin-gonic/gin"
)

func GetMovieByID(c *gin.Context) {
	var movie models.Movie

	id := c.Param("id")

	movieID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	if err := models.DB.First(&movie, movieID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}
