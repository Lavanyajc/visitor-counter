package main

import (
	"encoding/json"
	"io/ioutil"

	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

var mu sync.Mutex
const filePath = "counter.json"

type Counter struct {
	Visits int `json:"visits"`
}

func readCounter() Counter {
	var counter Counter

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		// If file doesn't exist, initialize with 0
		counter.Visits = 0
	} else {
		json.Unmarshal(data, &counter)
	}
	return counter
}

func writeCounter(counter Counter) {
	data, _ := json.Marshal(counter)
	_ = ioutil.WriteFile(filePath, data, 0644)
}

func main() {
	r := gin.Default()
	 r.Use(func(c *gin.Context) {
	       c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	       c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	       c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	       if c.Request.Method == "OPTIONS" {
		    c.AbortWithStatus(204)
	            return
	       }
	       c.Next()
       })



	r.GET("/visits", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()

		counter := readCounter()
		counter.Visits++
		writeCounter(counter)

		c.JSON(http.StatusOK, gin.H{
			"visits": counter.Visits,
		})
	})

	// Optional health check
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Visitor Counter API is running Lavanya JC!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
