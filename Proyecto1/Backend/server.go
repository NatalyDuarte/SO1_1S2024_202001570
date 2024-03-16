package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FreeRam struct {
	FreeRam int `json:"freeRam"`
}

func getFreeRam() (int, error) {
	cmd := exec.Command("cat", "/proc/ram_so1_1s2024")
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("error executing cat command: %w", err)
	}

	freeRamStr := string(bytes.Split(output, []byte("\n"))[0])
	freeRam, err := strconv.Atoi(freeRamStr)
	if err != nil {
		return 0, fmt.Errorf("error converting free RAM string to int: %w", err)
	}

	return freeRam, nil
}

func ramInfoHandler(c *gin.Context) {
	freeRam, err := getFreeRam()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := FreeRam{FreeRam: freeRam}
	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()

	// Middleware para habilitar CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	router.GET("/ram_info", ramInfoHandler)

	fmt.Println("Server listening on port 8080")
	router.Run(":8080")
}
