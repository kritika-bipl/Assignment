package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kritika-bipl/Assignment/Assignment-2/initializers"
	"github.com/kritika-bipl/Assignment/Assignment-2/routes"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r, initializers.DB)

	r.Run()
}
