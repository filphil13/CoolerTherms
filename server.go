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

var mostRecentScan TempScan

func getTempScans(c *gin.Context) {
	c.JSON(200, tempScanList)
	//	c.HTML(http.StatusOK, "/frontend/main.html", nil)
}

func getRecentScan(c *gin.Context) {
	c.JSON(200, mostRecentScan)
	//	c.HTML(http.StatusOK, "/frontend/main.html", nil)
}

func createTempScan(c *gin.Context) {
	var newScan TempScan

	if err := c.BindJSON(&newScan); err != nil {
		return
	}

	tempScanList = append(tempScanList, newScan)
	mostRecentScan = newScan
	c.JSON(http.StatusCreated, newScan)
}

func getHome(c *gin.Context){
	c.HTML(200, "main.html", nil)
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("Front-End/*.html")
	router.POST("/TempScans", createTempScan)
	router.GET("/TempScans", getTempScans)
	router.GET("/RecentScan", getRecentScan)
	router.GET("/", getHome)
	router.Run("192.168.2.11:8080")
}
