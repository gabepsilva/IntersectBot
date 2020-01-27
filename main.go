package main

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

// TODO: Secret management
// TODO: Organize if code >= too big AND TimeLeft() > 2h

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

	// get bot data
	err := c.Request.ParseForm()
	if err != nil {
		c.String(200, fmt.Sprintf("sorry, I crashed: %v", err))
		return
	}

	city := c.Request.FormValue("text")
	if len(city) <= 0 {
		city = "Toronto,CA"
	}
	c.String(200, getWeather(city))

}

func getWeather(city string) string{
	// initialize API
	var apiKey = os.Getenv("OWM_API_KEY")
	w, err := owm.NewCurrent("C", "EN", apiKey) // (internal - OpenWeatherMap reference for kelvin) with English output
	if err != nil {
		return fmt.Sprintf("sorry, I crashed: %v", err)
	}

	// forecast search
	err = w.CurrentByName(city)
	if err != nil {
		return fmt.Sprintf("sorry, I crashed: %v", err)
	}

	return fmt.Sprintf("%.2f°C %s\nMin:%.2f\nMax:%.2f", w.Main.Temp, w.Weather[0].Main, w.Main.TempMin, w.Main.TempMax)
}