package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type TempScan struct {
	SensorID       int  `json: "sensorid"`
	Temperature    float32 `json: "temperature"`
	Humidity  	   float32 `json: "humidity"`
	Time 		   string  `json: "time"`
}

var tempScanList []TempScan

func getTempScans(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, tempScanList)
	//	c.HTML(http.StatusOK, "/frontend/main.html", nil)
}

func createTempScan(c *gin.Context) {
	var newScan TempScan

	if err := c.BindJSON(&newScan); err != nil {
		return
	}

	tempScanList = append(tempScanList, newScan)
	c.IndentedJSON(http.StatusCreated, newScan)
}

func main() {

	router := gin.Default()
	router.POST("/TempScans", createTempScan)
	router.GET("/TempScans", getTempScans)
	router.Run("192.168.2.11:8080")
}
