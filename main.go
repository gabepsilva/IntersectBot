package main

import (
	"github.com/gin-gonic/gin"
	"log"
)



func main(){


	router := gin.New()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/hi", commandHi)
	}

	err := router.Run("0.0.0.0:8000")
	if err != nil {
		log.Panicf("%v", err)
	}

}

func commandHi(c *gin.Context){

	greetings := []string{
		"Enjoy your special day.",
		"The day is all yours - have fun!",
	}


	err := c.Request.ParseForm()
	if err != nil {
		log.Printf("invalid request from server: %v", err)
	}

	log.Println(c.Request.FormValue("text"))
	c.String(200, "%s", greetings[1])

}