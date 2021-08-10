package main

import (
	"github.com/gin-gonic/gin"
	"learning_go/configs"
	"learning_go/ip_to_country"
	"learning_go/middlewares"
	"log"
	"net/http"
)

var ipToLocationObj ip_to_country.IpToLocationDb

func findCountry(c *gin.Context) {
	ip, _ := c.GetQuery("ip")
	details := ipToLocationObj.IpToCountry(ip)
	statusCode := http.StatusOK
	if details == nil {
		statusCode = http.StatusNotFound
		details = map[string]string{"error": "ip not found in db"}
	}
	c.JSON(statusCode, details)
}

func main() {
	ipToLocationObj = ip_to_country.GetDbObj(configs.DbConfigs)
	router := gin.Default()
	router.Use(middlewares.Limiter())
	v1 := router.Group("/v1")
	v1.GET("/find-country", findCountry)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatalln("Error running Gin: ", err)
	}
}
