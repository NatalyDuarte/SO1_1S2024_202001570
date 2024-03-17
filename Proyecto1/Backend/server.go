package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FreeRam struct {
	FreeRam int `json:"freeRam"`
}

type FreeCpu struct {
	CpuTotal      int `json:"cpu_total"`
	CpuPorcentaje int `json:"cpu_porcentaje"`
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

func getCpuInfo() (FreeCpu, error) {
	cmd := exec.Command("cat", "/proc/cpu_so1_1s2024")
	output, err := cmd.Output()
	if err != nil {
		return FreeCpu{}, fmt.Errorf("error executing cat command: %w", err)
	}

	// Decodificar el JSON en un objeto map[string]interface{}
	var data map[string]interface{}
	err = json.Unmarshal(output, &data)
	if err != nil {
		return FreeCpu{}, fmt.Errorf("error decoding JSON: %w", err)
	}

	// Extraer los valores del objeto map
	cpuTotal, ok := data["cpu_total"].(float64)
	if !ok {
		return FreeCpu{}, fmt.Errorf("error getting cpu_total")
	}
	cpuPorcentaje, ok := data["cpu_porcentaje"].(float64)
	if !ok {
		return FreeCpu{}, fmt.Errorf("error getting cpu_porcentaje")
	}

	// Convertir los valores a int
	intCpuTotal := int(cpuTotal)
	intCpuPorcentaje := int(cpuPorcentaje)

	return FreeCpu{CpuTotal: intCpuTotal, CpuPorcentaje: intCpuPorcentaje}, nil
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

func cpuInfoHandler(c *gin.Context) {
	cpuInfo, err := getCpuInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cpuInfo)
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
	router.GET("/cpu_info", cpuInfoHandler)

	fmt.Println("Server listening on port 8080")
	router.Run(":8080")
}
