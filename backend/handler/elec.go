package handler

import (
	"happyEWater/repository"
	"happyEWater/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type electricHandler struct {
	electricService service.ElectricService
}

func NewElectricHandler(electricService service.ElectricService) electricHandler {
	return electricHandler{
		electricService: electricService,
	}
}

func (e *electricHandler) GetElectric(g *gin.Context) {
	data, err := e.electricService.GetElectric()
	if err == nil {
		g.JSON(http.StatusOK, data)
		return
	}
	g.JSON(http.StatusBadRequest, gin.H{
		"massage": "Some thing happen",
	})
}

func (e *electricHandler) GetDate(g *gin.Context) {
	data, err := e.electricService.GetDate()
	if err == repository.ErrNotFound {
		g.JSON(http.StatusNotFound, gin.H{
			"massage": "Not found any data",
		})
		return
	} else if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"massage": "Something went wrong",
		})
		return
	}
	g.JSON(http.StatusOK, data)
}

func (e *electricHandler) GetElecPerDay(g *gin.Context) {
	data, err := e.electricService.GetElecPerDay()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"massage": "Something wrong",
		})
		return
	}
	g.JSON(http.StatusOK, data)
}

func (e *electricHandler) GetElec3Hour(g *gin.Context) {
	data, err := e.electricService.GetElec3Hour()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"massage": "Something wrong",
		})
		return
	}
	g.JSON(http.StatusOK, data)
}
