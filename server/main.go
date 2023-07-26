package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("server staring....")
	router := gin.Default()
	fmt.Println(router)
}
