package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type FreeRam struct {
	FreeRam int `json:"freeRam"`
}

type FreeCpuva struct {
	CpuPorcentaje int `json:"cpu_porcentaje"`
	CpuBound      int `json:"cpu_bound"`
}

type HistRams struct {
	HistRams []HistRam `json:"histrams"`
}

type HistRam struct {
	HistRam int    `json:"histram"`
	Fecha   string `json:"fech"`
}

type HistCpus struct {
	HistCpus []HistCpu `json:"histcpus"`
}

type HistCpu struct {
	HistCpu int    `json:"histcpu"`
	Fecha   string `json:"fech"`
}

type FreeCpu struct {
	CpuTotal      int `json:"cpu_total"`
	CpuPorcentaje int `json:"cpu_porcentaje"`
}

func mandarsqlram(freram int) {
	db, err := sql.Open("mysql", "root:2000@tcp(192.168.0.17:3306)/Proyecto")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into raminfo(freeram,boundram,time) values (?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	bound := 16000000 - freram

	_, err = stmt.Exec(freram, bound, time.Now())
	if err != nil {
		panic(err)
	}

	println("Ram insertado correctamente")
	db.Close()
}

func mandarsqlcpu(cputotal int, cpuusado int) {
	db, err := sql.Open("mysql", "root:2000@tcp(192.168.0.17:3306)/Proyecto")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("insert into cpuinfo(freecpu,boundcpu,time) values (?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	bound := cputotal - cpuusado
	_, err = stmt.Exec(cpuusado, bound, time.Now())
	if err != nil {
		panic(err)
	}
	println("Cpu insertado correctamente")
	db.Close()
}

func getFreeRam() FreeRam {
	cmd := exec.Command("cat", "/proc/ram_so1_1s2024")
	output, err := cmd.Output()
	if err != nil {
		return FreeRam{}
	}
	freeRamStr := string(bytes.Split(output, []byte("\n"))[0])
	freeRam, err := strconv.Atoi(freeRamStr)
	if err != nil {
		return FreeRam{}
	}
	db, err := sql.Open("mysql", "root:2000@tcp(192.168.0.17:3306)/Proyecto")
	if err != nil {
		return FreeRam{}
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into raminfo(freeram,boundram,time) values (?,?,?)")
	if err != nil {
		return FreeRam{}
	}
	defer stmt.Close()
	bound := 8000000 - freeRam

	_, err = stmt.Exec(freeRam, bound, time.Now())
	if err != nil {
		panic(err)
	}

	println("Ram insertado correctamente")
	db.Close()

	return FreeRam{FreeRam: freeRam}
}

func getFreeCpu() FreeCpu {
	cmd := exec.Command("cat", "/proc/cpu_so1_1s2024")
	output, err := cmd.Output()
	if err != nil {
		return FreeCpu{}
	}

	// Decodificar el JSON en un objeto map[string]interface{}
	var data map[string]interface{}
	err = json.Unmarshal(output, &data)
	if err != nil {
		return FreeCpu{}
	}

	// Extraer los valores del objeto map
	cpuTotal, ok := data["cpu_total"].(float64)
	if !ok {
		return FreeCpu{}
	}
	cpuPorcentaje, ok := data["cpu_porcentaje"].(float64)
	if !ok {
		return FreeCpu{}
	}

	// Convertir los valores a int
	intCpuTotal := int(cpuTotal)
	//	println(intCpuTotal)
	intCpuPorcentaje := int(cpuPorcentaje)
	db, err := sql.Open("mysql", "root:2000@tcp(192.168.0.17:3306)/Proyecto")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("insert into cpuinfo(freecpu,boundcpu,time) values (?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	bound := intCpuTotal - intCpuPorcentaje
	_, err = stmt.Exec(intCpuPorcentaje, bound, time.Now())
	if err != nil {
		panic(err)
	}
	println("Cpu insertado correctamente")
	db.Close()
	return FreeCpu{CpuTotal: intCpuTotal, CpuPorcentaje: intCpuPorcentaje}

}

func cpuhistoricos() (HistCpus, error) {
	db, err := sql.Open("mysql", "root:2000@tcp(192.168.0.17:3306)/Proyecto")
	if err != nil {
		return HistCpus{}, fmt.Errorf("error opening database connection: %w", err)
	}
	defer db.Close()
	var cpus HistCpus
	stmt, err := db.Query("Select boundcpu, time from cpuinfo")
	if err != nil {
		return HistCpus{}, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var cpuInfo HistCpu
		err := stmt.Scan(&cpuInfo.HistCpu, &cpuInfo.Fecha)
		if err != nil {
			return HistCpus{}, err
		}
		cpus.HistCpus = append(cpus.HistCpus, cpuInfo)
	}
	if len(cpus.HistCpus) == 0 {
		fmt.Println("No hay datos")
	}
	return cpus, nil

}

func ramhistoricos() (HistRams, error) {
	db, err := sql.Open("mysql", "root:2000@tcp(192.168.0.17:3306)/Proyecto")
	if err != nil {
		return HistRams{}, fmt.Errorf("error opening database connection: %w", err)
	}
	defer db.Close()

	var rams HistRams
	stmt, err := db.Query("Select boundram, time from raminfo")
	if err != nil {
		return HistRams{}, err
	}
	defer stmt.Close()

	// Check if there are any rows available before scanning
	for stmt.Next() {
		var ramInfo HistRam
		err := stmt.Scan(&ramInfo.HistRam, &ramInfo.Fecha)
		if err != nil {
			return HistRams{}, err
		}
		rams.HistRams = append(rams.HistRams, ramInfo)
		//rams.Fecha = append(rams.Fecha, ramInfo)
	}
	if len(rams.HistRams) == 0 {
		fmt.Println("No hay datos historicos en la tabla")
	}
	return rams, nil
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
	//	println(intCpuTotal)
	intCpuPorcentaje := int(cpuPorcentaje)
	mandarsqlcpu(intCpuTotal, intCpuPorcentaje)
	return FreeCpu{CpuTotal: intCpuTotal, CpuPorcentaje: intCpuPorcentaje}, nil
}

func cpuhistorico(c *gin.Context) {
	histcpu, err := cpuhistoricos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(histcpu.HistCpus) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay datos historicos"})
		return
	}
	c.JSON(http.StatusOK, histcpu)
}

func ramhistorico(c *gin.Context) {
	histram, err := ramhistoricos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(histram.HistRams) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay datos historicos"})
		return
	}
	c.JSON(http.StatusOK, histram)
}

func cpuInfoHandler(c *gin.Context) {
	cpuInfo, err := getCpuInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cpuInfo)
}

func recevram(db *sql.DB) (HistRam, error) {
	stmt, err := db.Query("Select boundram, time from raminfo")
	if err != nil {
		return HistRam{}, err
	}
	defer stmt.Close()

	var ramInfo HistRam
	// Check if there are any rows available before scanning
	if !stmt.Next() {
		// Handle no rows case (optional)
		return HistRam{}, nil // or return an error indicating no data found
	}
	err = stmt.Scan(&ramInfo.HistRam, &ramInfo.Fecha)
	if err != nil {
		return HistRam{}, err
	}
	return HistRam{HistRam: ramInfo.HistRam, Fecha: ramInfo.Fecha}, nil
}

func ramget() (int, error) {
	db, err := sql.Open("mysql", "root:2000@tcp(192.168.0.17:3306)/Proyecto")
	if err != nil {
		return 0, fmt.Errorf("error opening database connection: %w", err)
	}
	defer db.Close()

	var freeRam int
	stmt, err := db.Query("Select freeram from raminfo order by idraminfo desc limit 1;")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Verifica si hay resultados en la consulta (opcional)
	if !stmt.Next() {
		return 0, nil // No hay datos (manejar según sea necesario)
	}

	err = stmt.Scan(&freeRam)
	if err != nil {
		return 0, fmt.Errorf("error scanning freeram: %w", err)
	}

	return freeRam, nil
}

func getRam(c *gin.Context) {
	freeram, err := ramget()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := FreeRam{FreeRam: freeram}
	c.JSON(http.StatusOK, response)
}

func cpuget() (FreeCpu, error) {
	db, err := sql.Open("mysql", "root:2000@tcp(192.168.0.17:3306)/Proyecto")
	if err != nil {
		return FreeCpu{}, fmt.Errorf("error opening database connection: %w", err)
	}
	defer db.Close()

	var freeCpu FreeCpuva
	stmt, err := db.Query("Select freecpu , boundcpu from cpuinfo order by idcpuinfo desc limit 1;")
	if err != nil {
		return FreeCpu{}, err
	}
	defer stmt.Close()

	if !stmt.Next() {
		return FreeCpu{}, nil // No hay datos (manejar según sea necesario)
	}

	err = stmt.Scan(&freeCpu.CpuPorcentaje, &freeCpu.CpuBound)
	if err != nil {
		return FreeCpu{}, fmt.Errorf("error scanning cpu info: %w", err)
	}

	totalcpu := freeCpu.CpuPorcentaje + freeCpu.CpuBound

	return FreeCpu{CpuTotal: totalcpu, CpuPorcentaje: freeCpu.CpuPorcentaje}, nil
}

func getCpu(c *gin.Context) {
	cpuInfo, err := cpuget()
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

	go func() {
		for {
			getFreeRam()
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			getFreeCpu()
			time.Sleep(5 * time.Second)
		}
	}()

	router.GET("/tiempor/ram", getRam)
	router.GET("/tiempor/cpu", getCpu)
	router.GET("/tiempohis/ram", ramhistorico)
	router.GET("/tiempohis/cpu", cpuhistorico)

	fmt.Println("Server listening on port 8080")
	router.Run(":8080")
}
