package routes

import (
	"github.com/MakanPisang/movie-website.git/controllers"
	"github.com/MakanPisang/movie-website.git/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/login", controllers.LoginHandler)
	r.GET("/movies", controllers.GetMovies)
	r.GET("/movies/:id", controllers.GetMovieByID)
	r.POST("/movies", controllers.CreateMovie)

	protected := r.Group("/api")
	protected.Use(middlewares.JWTAuthMiddleware())
	protected.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You are authorized"})
	})
}
