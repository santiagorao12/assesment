package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Docket struct {
	OrderNo       string  `json:"OrderNo"`
	Customer      string  `json:"Customer"`
	PickUpPoint   string  `json:"PickUpPoint"`
	DeliveryPoint string  `json:"DeliveryPoint"`
	Quantity      int     `json:"Quantity"`
	Volume        float64 `json:"Volume"`
	Status        string  `json:"Status"`
	TruckNo       string  `json:"TruckNo"`
	LogsheetNo    string  `json:"LogsheetNo"`
}

type Logsheet struct {
	LogsheetNo string   `json:"LogsheetNo"`
	Dockets    []string `json:"Dockets"`
	TruckNo    string   `json:"TruckNo"`
}

var (
	dockets      = make(map[string]Docket) // In-memory storage for dockets
	orderNoMu    sync.Mutex                // Mutex to synchronize access to orderNo counter
	orderNo      int                       // Global counter for unique OrderNo
	logsheetNoMu sync.Mutex                // Mutex to synchronize access to logsheetNo counter
	logsheetNo   int                       // Global counter for unique LogsheetNo
)

func generateUniqueOrderNo() string {
	orderNoMu.Lock()
	defer orderNoMu.Unlock()
	orderNo++
	return fmt.Sprintf("TDN%04d", orderNo)
}

func generateUniqueLogsheetNo() string {
	logsheetNoMu.Lock()
	defer logsheetNoMu.Unlock()
	logsheetNo++
	return fmt.Sprintf("DT%04d", logsheetNo)
}

func createDocketHandler(c *gin.Context) {
	var newDocket Docket
	if err := c.ShouldBindJSON(&newDocket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Generate a unique OrderNo for the new Docket
	newDocket.OrderNo = generateUniqueOrderNo()

	// Set the status as "Created" for the new Docket
	newDocket.Status = "Created"

	// Store the new Docket in the dockets map
	dockets[newDocket.OrderNo] = newDocket

	// Return the response
	c.JSON(http.StatusCreated, newDocket)
}

func getDocketHandler(c *gin.Context) {
	orderNo := c.Param("orderNo")
	docket, found := dockets[orderNo]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Docket not found"})
		return
	}
	c.JSON(http.StatusOK, docket)
}

func listDocketsHandler(c *gin.Context) {
	var filteredDockets []Docket
	for _, docket := range dockets {
		if docket.Status == "Created" {
			filteredDockets = append(filteredDockets, docket)
		}
	}
	c.JSON(http.StatusOK, filteredDockets)
}

func createLogsheetHandler(c *gin.Context) {
	var newLogsheet Logsheet
	if err := c.ShouldBindJSON(&newLogsheet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Generate a unique LogsheetNo for the new Logsheet
	newLogsheet.LogsheetNo = generateUniqueLogsheetNo()

	// Update the dockets' TruckNo and LogsheetNo fields based on the new Logsheet
	for _, docketNo := range newLogsheet.Dockets {
		docket, found := dockets[docketNo]
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "Docket not found"})
			return
		}
		docket.TruckNo = newLogsheet.TruckNo
		docket.LogsheetNo = newLogsheet.LogsheetNo
		dockets[docketNo] = docket
	}

	// Return the response with the updated dockets
	c.JSON(http.StatusCreated, newLogsheet)
}

func getLogsheetHandler(c *gin.Context) {
	logsheetNo := c.Param("logsheetNo")
	logsheet, found := logsheets[logsheetNo]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Logsheet not found"})
		return
	}
	c.JSON(http.StatusOK, logsheet)
}

var logsheets = make(map[string]Logsheet) // In-memory storage for logsheets

func main() {
	r := gin.Default()

	// Create a new Docket (Question 1)
	r.POST("/docket", createDocketHandler)

	// Fetch a Docket based on the order number (Question 2)
	r.GET("/docket/:orderNo", getDocketHandler)

	// Fetch a list of dockets (Question 3)
	r.GET("/docket", listDocketsHandler)

	// Create a new Logsheet and update the dockets accordingly (Question 4)
	r.POST("/logsheet", createLogsheetHandler)

	// Fetch a Logsheet based on the LogsheetNo (Question 5)
	r.GET("/logsheet/:logsheetNo", getLogsheetHandler)

	// Run the server on port 8080
	r.Run(":8080")
}
