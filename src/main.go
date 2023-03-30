package main

import (
	"fmt"
	"os"
	"service-api/src/core/system"

	"github.com/gin-gonic/gin"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
	system.Start(gin.Default())
}
