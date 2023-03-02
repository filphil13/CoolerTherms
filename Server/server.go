package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)



type Connection struct{
	Name	string `json: "name"`
	Address string `json: "address"`
}

type TempScan struct {
	Name			string	`json: "name"`
	Temperature    	float32 `json: "temperature"`
	Humidity  	   	float32 `json: "humidity"`
	Time		   	int		`json: "time"`
}


var tempScanList []TempScan
var mostRecentScans []TempScan
var addressList []Connection

func inAddressList(addr string) (bool){
	for _, conn := range addressList{
		if  addr == conn.Address {
			return true		
		}
	}
	return false
}


func getTempScans(c *gin.Context) {
	
	c.JSON(200, tempScanList)
}

func getRecentScan(c *gin.Context) {
	c.JSON(200, mostRecentScans)
}


func createTempScan(c *gin.Context) {
	if !(inAddressList(c.ClientIP())){
		c.HTML(400, "Need to Init Sensor First", nil)
		return
	}
	var newScan TempScan

	if err := c.BindJSON(&newScan); err != nil {
		return
	}

	newScan.Time = int(time.Now().Unix())
	tempScanList = append(tempScanList, newScan)
	logTempData()

	for i, scan := range(mostRecentScans){
		if(scan.Name == newScan.Name){
			mostRecentScans[i] = newScan
			c.JSON(http.StatusCreated, newScan)
			return
		}
	}
	mostRecentScans = append(mostRecentScans, newScan)
	c.JSON(http.StatusCreated, newScan)
}

func getHome(c *gin.Context){
	c.HTML(200, "main.html", nil)
}

func initSensor(c *gin.Context){
	if inAddressList(c.ClientIP()){
		//potential security blocks here for unknown addresses
		//for now will remain fully unblocked
		c.JSON(http.StatusOK, nil)

	}else{
		var newConn Connection
		c.BindJSON(&newConn)
		newConn.Address = c.ClientIP()
		addToAddressList(newConn.Name,newConn.Address)
	}
}

func addToAddressList(name, address string){
		newConn := Connection{
			Name: name,
			Address: address}

		addressList = append(addressList, newConn)
}

func logTempData(){
	file, _ := json.Marshal(tempScanList)
 
	_ = ioutil.WriteFile("log.json", file, 0644)
}


func main() {

	router := gin.Default()
	router.LoadHTMLGlob("Front-End/*.html")
	router.POST("/TempScans", createTempScan)
	router.POST("/InitSensor", initSensor)
	router.GET("/TempScans", getTempScans)
	router.GET("/RecentScan", getRecentScan)
	router.GET("/", getHome)
	router.Run("192.168.2.11:8080")
}
