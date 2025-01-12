package main

import (
	"fmt"
	"orderFood-server-cus/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouters(r)

	err := r.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}
