package main

import (
	"learning_go/configs"
	"learning_go/ip_to_country"
	"learning_go/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestValidIpResponse(t *testing.T) {
	ipToLocationObj = ip_to_country.GetDbObj(configs.DbConfigs)
	router := getRouter()

	v1 := router.Group("/v1")
	v1.GET("/find-country", findCountry)

	req, _ := http.NewRequest("GET", "/v1/find-country?ip=1.1.1.1", nil)

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}

func TestNotValidIpResponse(t *testing.T) {
	ipToLocationObj = ip_to_country.GetDbObj(configs.DbConfigs)
	router := getRouter()

	v1 := router.Group("/v1")
	v1.GET("/find-country", findCountry)

	req, _ := http.NewRequest("GET", "/v1/find-country?ip=5555.555.555.555", nil)

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusNotFound
		return statusOK
	})
}

func TestNValidIpResponse(t *testing.T) {
	ipToLocationObj = ip_to_country.GetDbObj(configs.DbConfigs)
	router := getRouter()
	router.Use(middlewares.Limiter())

	v1 := router.Group("/v1")
	v1.GET("/find-country", findCountry)

	req, _ := http.NewRequest("GET", "/v1/find-country?ip=1.1.1.1", nil)

	for i := 0; i < configs.LimiterMaxCredits; i++ {
		testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
			statusOK := w.Code == http.StatusOK
			return statusOK
		})
	}
	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusTooManyRequests
		return statusOK
	})

	time.Sleep(time.Duration(configs.LimiterTimeFrame) * time.Second)

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}
