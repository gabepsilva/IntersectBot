package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {

	tt := []struct {
		Name string
		Cities []string
	}{
		{Name: "Test valid cities", Cities: []string{"California", "Toronto", "Toronto", "Toronto","Texas", "Tokyo", "Lima", "Brasilia",
			"Uruguay", "Paris", "Alaska", "Cairo"}},
	}

	router := serverConfig()

	for _, testCase := range tt {
		for _, city := range testCase.Cities {
			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/v1/weather", strings.NewReader("text="+city))
			if err != nil {
				t.Errorf("Request error: %v", err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			router.ServeHTTP(w, req)

			// test response code
			assert.Equal(t, 200, w.Code)
			// test if returned the result for the right city
			assert.Contains(t, w.Body.String(), city)
			// count number of lines of the response
			assert.Equal(t, strings.Count(w.Body.String(), "\n"), 3)
		}
	}
}

func TestGetWeatherStatus(t *testing.T) {


	router := serverConfig()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/v1/weather-status", nil)
	if err != nil {
		t.Errorf("Request error: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	// test response code
	assert.Equal(t, 200, w.Code)
	// test if returned the result for the right city
	assert.Contains(t, w.Body.String(), "records")

}

func TestGetWeatherInvalidCities(t *testing.T) {

	tt := []string{"California,YY", "Toronto,ZZ", "Toronto,AA", "Toronto,BB",
			"Texas,CC", "Tokyo,DD", "Lima,EE", "Brasilia,FF", "Uruguay,GG", "Paris,HH", "Alaska,II", "Cairo,JJ"}


	router := serverConfig()

	for _, city := range tt {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/v1/weather", strings.NewReader("text="+city))
		if err != nil {
			t.Errorf("Request error: %v", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		router.ServeHTTP(w, req)

		// test response code
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, w.Body.String(), "Nope, this city does not exist or maybe I am broke. Ask admin to check my logs")
	}
}