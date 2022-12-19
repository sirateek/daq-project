package main

import (
	"context"
	"fmt"
	"happyEWater/handler"
	"happyEWater/repository"
	"happyEWater/service"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	clientOption := options.Client().ApplyURI("mongodb+srv://Tdean:12345@daq.fg7tz9x.mongodb.net/?retryWrites=true&w=majority")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		logrus.Info(err)
	}

	fmt.Println("mongo connection established")
	database := client.Database("HappyEWater")
	waterCollection := database.Collection("Water")
	electricCollection := database.Collection("Electric")
	tmdCollection := database.Collection("TMD")
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	defer client.Disconnect(ctx)
	waterRepocitory := repository.NewWaterCollection(waterCollection)
	electricRepocitory := repository.NewElectricRepository(electricCollection)
	tmdRepository := repository.NewTmdRepository(tmdCollection)

	waterService := service.NewWaterService(waterRepocitory)
	electricService := service.NewElectricService(electricRepocitory)
	tmdService := service.NewTmdService(tmdRepository)

	waterHanlder := handler.NewWaterService(waterService)
	electricHandler := handler.NewElectricHandler(electricService)
	tmdHandler := handler.NewTmdHandler(tmdService)
	api := server.Group("/api")
	api.GET("/water", waterHanlder.GetWater)
	api.GET("/water/date", waterHanlder.GetDate)
	api.GET("/water/perday", waterHanlder.GetWaterPerDay)
	api.GET("/water/3hour", waterHanlder.GetWater3Hour)
	api.GET("/elec", electricHandler.GetElectric)
	api.GET("/elec/date", electricHandler.GetDate)
	api.GET("/elec/perday", electricHandler.GetElecPerDay)
	api.GET("/elec/3hour", electricHandler.GetElec3Hour)
	api.GET("/tmd", tmdHandler.GetTmd)
	api.GET("/tmd/perday", tmdHandler.GetTmdPerDay)
	api.GET("/tmd/date", tmdHandler.GetDate)
	api.GET("/tmd/raw/data", tmdHandler.GetRawData)
	server.Run(fmt.Sprint(":", 8080))
}
