package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type dats struct {
	Carnet string `json: "carnet"`
	Nombre string `json: "nombre"`
}

var datos = "Carnet: 202001570 , Nombre: Nataly Saraí Guzmán Duarte"

func getDatos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, datos)
}

func main() {
	router := gin.Default()
	router.GET("/data", getDatos)

	router.Run("localhost:3001")
}
