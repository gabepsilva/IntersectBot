package weather

import (
	owm "github.com/briandowns/openweathermap"
	"log"
)

type Weather struct {
	City string
	Main string
	Temp float64
	Min  float64
	Max float64

}

var OWM_API_KEY string

func GetByCity(city string) *Weather {
	// initialize API
	if len(OWM_API_KEY) <= 10 {
		log.Fatalf("API key is not right")
	}
	w, err := owm.NewCurrent("C", "EN", OWM_API_KEY)
	if err != nil {
		log.Printf("failed to initialize api")
		return nil
	}

	// forecast search
	err = w.CurrentByName(city)
	if err != nil {
		log.Printf("can not get data about %s", city)
		return nil
	}

	we := new(Weather)
	we.City = w.Name
	we.Main = w.Weather[0].Main
	we.Temp = w.Main.Temp
	we.Min = w.Main.TempMin
	we.Max = w.Main.TempMax

	return we
}
