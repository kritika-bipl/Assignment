package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kritika-bipl/Assignment/Assignment-2/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	
	r.Run()
}
