package main

import (
	Controller "currency-exchange/controllers"
	"currency-exchange/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server.....")

	db.Init()

	r := gin.Default()

	r.LoadHTMLGlob("view/*")
	r.GET("/getForexList", Controller.GetForexExchangeList)
	r.POST("/addForexList", Controller.CreateForexExchange)
	r.PUT("/updateForexList/:id", Controller.UpdateForexExchange)
	r.DELETE("/deleteForexList/:id", Controller.DeleteForexExchange)

	r.Run()
}
