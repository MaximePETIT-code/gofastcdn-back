package handlers

import (
	"net/http"
	"os"
	"runtime"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

type HealthResponse struct {
	Status    string        `json:"status"`
	Timestamp time.Time     `json:"timestamp"`
	Uptime    time.Duration `json:"uptime"`
	GoVersion string        `json:"go_version"`
}

func DebugEnv(c *gin.Context) {
	env := os.Environ()
	sort.Strings(env)
	c.JSON(http.StatusOK, gin.H{
		"env": env,
	})
}

func HealthCheck(c *gin.Context) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Uptime:    time.Since(startTime).Round(time.Second),
		GoVersion: runtime.Version(),
	}

	c.JSON(http.StatusOK, response)
}
