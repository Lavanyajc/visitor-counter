# visitor-counter


## ✅ Phase 3 – Visitor Counter API: Full Project Files & Explanation
link - https://visitor-counter-4.onrender.com
---

### 📁 Project Structure

```
visitor-counter/
├── main.go           # The Go backend logic
├── go.mod            # Module definition and dependencies
├── go.sum            # Auto-generated dependency checksums
├── counter.json      # Created automatically to store visit counts
```

---

## 📄 1. `go.mod`

```go
module github.com/Lavanyajc/visitor-counter

go 1.24

require (
    github.com/gin-gonic/gin v1.9.1
)
```

✅ Describes your Go module
✅ Includes the Gin framework
✅ Required for dependency management and building in the cloud (like on Render)

---

## 📄 2. `main.go`

```go
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

// Reads the counter value from file or initializes to 0
func readCounter() Counter {
	var counter Counter
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		counter.Visits = 0
	} else {
		json.Unmarshal(data, &counter)
	}
	return counter
}

// Writes the updated visit count back to file
func writeCounter(counter Counter) {
	data, _ := json.Marshal(counter)
	_ = ioutil.WriteFile(filePath, data, 0644)
}

func main() {
	r := gin.Default()

	// Main visitor endpoint
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

	// Simple health check
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Visitor Counter API is running")
	})

	// Use Render’s dynamic port or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
```

---

## 📄 3. `counter.json` (created automatically)

This stores the **visitor count** persistently, like:

```json
{
  "visits": 53
}
```

✅ Gets updated on every `/visits` call
✅ Prevents losing data after restart (unlike in-memory counters)

---

## 🟢 Deployment Guide (Render)

### 🌐 Hosting Your API for Free:

1. Login to [https://render.com](https://render.com)
2. Click **“New Web Service”**
3. Choose your `visitor-counter` GitHub repo
4. Use:

   * **Environment:** Go
   * **Build Command:** `go build -o app .`
   * **Start Command:** `./app`
5. Render detects `main.go`, builds & deploys
6. App is live at `https://<your-app>.onrender.com`

---

## ✅ API Testing

### `/` endpoint:

```http
GET / → returns "Visitor Counter API is running"
```

### `/visits` endpoint:

```http
GET /visits → returns JSON { "visits": 1 }
```

✅ Increments on every call
✅ Returns total count

---

## 🔗 Frontend Integration Example (resume)

Add this to your `index.html` before `</body>`:

```html
<p>Total Visits: <span id="visit-count">Loading...</span></p>

<script>
  fetch("https://<your-app>.onrender.com/visits")
    .then(response => response.json())
    .then(data => {
      document.getElementById("visit-count").innerText = data.visits;
    })
    .catch(error => {
      console.error("Error fetching visits:", error);
    });
</script>
```

---

## 🧠 Why It Matters in CRC

| CRC Checklist Item                                      | ✅ Completed        |
| ------------------------------------------------------- | ------------------ |
| Build and deploy a backend API                          | ✅ Go + Gin         |
| Use a database/file to store visits                     | ✅ `counter.json`   |
| Deploy it in the cloud                                  | ✅ Render (free)    |
| Connect it to your resume site                          | ✅ Done |
| Use good practices (locking, error handling, port mgmt) | ✅ ✔️✔️✔️           |

This **proves your ability to:**

* Write backend logic
* Manage persistent state
* Deploy services to the cloud
* Integrate with frontend
* Follow DevOps best practices (PORT var, lightweight deps, JSON APIs)

---


```
# Visitor Counter API – Cloud Resume Challenge

This is a lightweight Go-based visitor counter API used in my [Cloud Resume Challenge](https://luffyjc.xyz).

## Features
- Written in Go using Gin
- Persists visits in a local JSON file
- Deployed for free on Render
- Integrated with my HTML resume site

## Endpoint
`GET /visits` → `{ "visits": <int> }`

## Deployment
Free on Render using Go build and start commands.
```

