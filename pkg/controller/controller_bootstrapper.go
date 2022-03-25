package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/kuroyamii/go-backend-learn/internal/ping/controller"
	"github.com/kuroyamii/go-backend-learn/internal/ping/service"
)

func InitializeControllers(router *mux.Router, db *sql.DB) {
	pingRouter := router.PathPrefix(API_PATH_PING).Subrouter()
	pingService := service.NewPingService()
	pingController := controller.NewPingController(pingRouter, &pingService)
	pingController.InitializePingEndPoint()

}
