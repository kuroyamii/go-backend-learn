package controller

import (
	"net/http"

	"github.com/kuroyamii/go-backend-learn/pkg/entity/response"

	"github.com/gorilla/mux"
	"github.com/kuroyamii/go-backend-learn/internal/ping/service"
)

type PingController struct {
	router *mux.Router
	ps     service.PingService
}

func (pc *PingController) HandlePing(w http.ResponseWriter, r *http.Request) {
	pingData := pc.ps.GetPingData()

	w.WriteHeader(http.StatusOK)
	response.NewBaseResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		response.NewErrorResponseData(
			response.NewErrorResponseValue("key1", "tes error 1"),
			response.NewErrorResponseValue("key2", "tes error 2"),
		),
		pingData,
	).ToJSON(w)
}

func (pc PingController) InitializePingEndPoint() {
	pc.router.HandleFunc("", pc.HandlePing)
}

func NewPingController(router *mux.Router, ps service.PingService) PingController {
	return PingController{
		router: router,
		ps:     ps,
	}
}
