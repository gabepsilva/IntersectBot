package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"intersectBot/appDB"
	"intersectBot/weather"
	"log"
	"os"
)

func main(){

	// initialize database
	connectionUrl := os.Getenv("DB_URL")
	appDB.Connect(connectionUrl)
	appDB.MySQL.Init()

	// initialize weather api
	weather.OWM_API_KEY = os.Getenv("OWM_API_KEY")

	// initialize web server
	router := gin.New()

	v1 := router.Group("/v1")
	{
		v1.POST("/weather", getWeather)
		v1.POST("/weather-status", getWeatherStatus)
	}

	err := router.Run("0.0.0.0:8000")
	if err != nil {
		log.Panicf("%v", err)
	}

}

func getWeather(c *gin.Context){

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

	botReturn := "Nope, this city does not exist or maybe I am broke. Ask admin to check my logs"
	w := weather.GetByCity(city)
	if w == nil {
		c.String(200, botReturn)
		return
	}

	botReturn = fmt.Sprintf("%.2f°C in %s\n Min:%.2f\nMax:%.2f\n%s", w.Temp,w.City, w.Min, w.Max, w.Main)
	c.String(200, botReturn)
	query := fmt.Sprintf("INSERT INTO weather (city, weather, temperature) VALUES ('%s', '%s', %.2f)", w.City, w.Main, w.Temp)
	appDB.MySQL.Exec(query)
}

func getWeatherStatus(c *gin.Context) {

	var totalRows int
	res := appDB.MySQL.Query("SELECT COUNT(*) FROM weather")
	for res.Next() {
		err := res.Scan(&totalRows)
		if err != nil {
			panic(err.Error())
		}
	}

	var botReturn string
	var weatherStatus string
	var weatherCount int

	res = appDB.MySQL.Query("SELECT weather, COUNT(*) FROM weather GROUP BY weather")
	for res.Next() {
		err := res.Scan(&weatherStatus, &weatherCount)
		if err != nil {
			panic(err.Error())
		}
		botReturn += fmt.Sprintf("%s - %d/%d records (%d%%)\n", weatherStatus, weatherCount, totalRows, (weatherCount * 100 / totalRows))
	}

	c.String(200, botReturn)
}
