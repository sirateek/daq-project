package handler

import (
	"happyEWater/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type waterHanlder struct {
	waterService service.WaterService
}

func NewWaterService(waterService service.WaterService) waterHanlder {
	return waterHanlder{
		waterService: waterService,
	}
}

func (w *waterHanlder) GetWater(g *gin.Context) {
	data, err := w.waterService.GetWater()
	if err == nil {
		g.JSON(http.StatusOK, data)
		return
	}
	g.JSON(http.StatusInternalServerError, gin.H{
		"massage": "Server error. Please try it again",
	})
}

func (w *waterHanlder) GetWaterPerDay(g *gin.Context) {
	data, err := w.waterService.GetWaterPerDay()
	if err == nil {
		g.JSON(http.StatusOK, data)
		return
	}
	g.JSON(http.StatusInternalServerError, gin.H{
		"massage": "Server error. Please try it again",
	})
}

func (w *waterHanlder) GetDate(g *gin.Context) {
	data, err := w.waterService.GetDate()
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{
			"massage": "Error can't found any data.",
		})
		return
	}
	g.JSON(http.StatusOK, data)
}

func (w *waterHanlder) GetWater3Hour(g *gin.Context) {
	data, err := w.waterService.GetWater3Hour()
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{
			"massage": "Error can't found any data.",
		})
		return
	}
	g.JSON(http.StatusOK, data)
}
