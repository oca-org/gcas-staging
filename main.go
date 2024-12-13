package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	runTests()
}

func runTests() {

	expected := "hello"
	actual := "hello"

	if assert.Equal(nil, expected, actual, "Values should be equal") {
		log.Println("Test passed!")
	} else {
		log.Println("Test failed.")
	}
}
