package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var count int
var mu sync.Mutex

func main() {
	r := gin.Default()

	r.GET("/visits", func(c *gin.Context) {
		mu.Lock()
		count++
		current := count
		mu.Unlock()

		c.JSON(http.StatusOK, gin.H{
			"visits": current,
		})
	})

	r.Run() // defaults to port 8080
}
