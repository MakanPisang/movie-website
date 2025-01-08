package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/MakanPisang/movie-website.git/models"
	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	var movies []models.Movie

	// Query semua film dari database
	if err := models.DB.Find(&movies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

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

func CreateMovie(c *gin.Context) {
	var movie models.Movie

	// Validasi input JSON
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log nilai release_date yang diterima
	log.Printf("Received release_date: %v", movie.ReleaseDate)

	// Parsing manual tanggal (release_date) dari input JSON
	parsedDate, err := time.Parse("2006-01-02", movie.ReleaseDate.Format("2006-01-02"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}
	movie.ReleaseDate = parsedDate

	// Simpan data ke database
	if err := models.DB.Create(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}
