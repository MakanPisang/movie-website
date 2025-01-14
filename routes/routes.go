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

	admin := r.Group("/admin")
	admin.Use(middlewares.JWTAuthMiddleware(), middlewares.RoleMiddleware("admin"))
	admin.POST("/movies", controllers.CreateMovie)
	admin.PUT("/movies/:id", controllers.UpdateMovie)
	admin.DELETE("/movies/:id", controllers.DeleteMovie)
}
