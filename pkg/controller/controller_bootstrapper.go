package controller

import (
	"database/sql"

	"github.com/codedius/imagekit-go"
	"github.com/gorilla/mux"

	// membersController "github.com/kuroyamii/go-backend-learn/internal/members/controller/impl"
	// memberRepository "github.com/kuroyamii/go-backend-learn/internal/members/repository/impl"
	// membersService "github.com/kuroyamii/go-backend-learn/internal/members/service/impl"
	imagekitControllerPkg "github.com/kuroyamii/go-backend-learn/internal/imagekit/controller"
	imagekitServicePkg "github.com/kuroyamii/go-backend-learn/internal/imagekit/service"
	"github.com/kuroyamii/go-backend-learn/internal/ping/controller"
	"github.com/kuroyamii/go-backend-learn/internal/ping/service"
)

func InitializeControllers(router *mux.Router, db *sql.DB, imageKit *imagekit.Client) {

	// webrouter := router.PathPrefix("/web").Subrouter()

	pingRouter := router.PathPrefix(API_PATH_PING).Subrouter()
	pingService := service.NewPingService()
	pingController := controller.NewPingController(pingRouter, &pingService)
	pingController.InitializePingEndPoint()

	imagekitRouter := router.PathPrefix("/imagekit").Subrouter()
	imagekitService := imagekitServicePkg.ProvideImagekitService(imageKit)
	imagekitController := imagekitControllerPkg.ProvideImagekitController(imagekitRouter, &imagekitService)
	imagekitController.InitEndpoints()

	// membersRepository := memberRepository.ProvideMemberRepository(db)
	// membersService := membersService.ProvideMemberService(membersRepository)
	// membersController := membersController.ProvideMembersController(webrouter, membersService)
	// membersController.InitializeController()
}
