package main

import (
	"github.com/kritika-bipl/Assignment/Assignment-2/initializers"
	
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	
}
