package controllers

import (
	"currency-exchange/db"
	"currency-exchange/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetForexExchangeList(c *gin.Context) {
	var forexList []models.ForexHistory

	db := db.GetDB()
	fmt.Println(time.Now().Format("2006-01-02"))
	db.Where("forex_date BETWEEN ? AND  ?", time.Now().AddDate(0, 0, -7), time.Now()).Find(&forexList)

	c.JSON(200, &forexList)

}

func CreateForexExchange(c *gin.Context) {
	var forexList models.ForexHistory
	var forexRequest models.ForexHistoryRequest
	var db = db.GetDB()

	if err := c.BindJSON(&forexRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	forexList.ForexDate, _ = time.Parse("2006-01-02", forexRequest.ForexDate)
	forexList.CurrencyFrom = forexRequest.CurrencyFrom
	forexList.CurrencyTo = forexRequest.CurrencyTo
	forexList.ExchangeRate = forexRequest.ExchangeRate

	db.Create(&forexList)
	c.JSON(http.StatusOK, &forexList)
}

func UpdateForexExchange(c *gin.Context) {

	id := c.Param("id")
	var forexList models.ForexHistory
	var db = db.GetDB()

	if err := db.Where("id = ?", id).First(&forexList).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&forexList)
	db.Save(&forexList)
	c.JSON(http.StatusOK, &forexList)
}

func DeleteForexExchange(c *gin.Context) {
	id := c.Param("id")
	var forexList models.ForexHistory
	var db = db.GetDB()

	if err := db.Where("id = ?", id).First(&forexList).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	db.Delete(&forexList)
}
