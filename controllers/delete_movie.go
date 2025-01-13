package controllers

import (
	"net/http"
	"strconv"

	"github.com/MakanPisang/movie-website.git/models"
	"github.com/gin-gonic/gin"
)

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	movieID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	if err := models.DB.Delete(&models.Movie{}, movieID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
