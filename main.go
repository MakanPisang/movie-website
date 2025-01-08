package main

import (
	"github.com/MakanPisang/movie-website.git/models"
	"github.com/MakanPisang/movie-website.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
