package handler

import (
	"happyEWater/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tmdHandler struct {
	tmdService service.TmdService
}

func NewTmdHandler(tmdService service.TmdService) tmdHandler {
	return tmdHandler{
		tmdService: tmdService,
	}
}

func (t *tmdHandler) GetTmd(g *gin.Context) {
	data, err := t.tmdService.GetTmd()
	if err == nil {
		g.JSON(http.StatusOK, data)
		return
	}
	g.JSON(http.StatusInternalServerError, gin.H{
		"massage": "Server error. Please try it again",
	})
}

func (t *tmdHandler) GetTmdPerDay(g *gin.Context) {
	temp, rain, err := t.tmdService.GetTimePerDay()
	if err == nil {
		g.JSON(http.StatusOK, gin.H{
			"temp": temp,
			"rain": rain,
		})
		return
	}
	g.JSON(http.StatusBadRequest, gin.H{
		"massage": "Something went wrong",
	})
}

func (t *tmdHandler) GetDate(g *gin.Context) {
	data, err := t.tmdService.GetDate()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"massage": "Something went wrong",
		})
		return
	}
	g.JSON(http.StatusOK, data)
}

func (t *tmdHandler) GetAllDate(g *gin.Context) {
	data, err := t.tmdService.GetAllDate()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"massage": "Something went wrong",
		})
		return
	}
	g.JSON(http.StatusOK, data)
}

func (t *tmdHandler) GetRawData(g *gin.Context) {
	data, err := t.tmdService.GetRawData()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"massage": "Something went wrong",
		})
		return
	}
	g.JSON(http.StatusOK, data)
}
