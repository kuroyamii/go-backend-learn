package main

import (
	"database/sql"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kuroyamii/go-backend-learn/pkg/controller"
	"github.com/kuroyamii/go-backend-learn/pkg/database"
	"github.com/kuroyamii/go-backend-learn/pkg/middleware"
	"github.com/kuroyamii/go-backend-learn/pkg/server"
)

func getEnvVariables() map[string]string {
	envVariables := make(map[string]string)

	envVariables["SERVER_ADDRESS"] = os.Getenv("SERVER_ADDRESS")

	envVariables["DB_ADDRESS"] = os.Getenv("DB_ADDRESS")
	envVariables["DB_USERNAME"] = os.Getenv("DB_USERNAME")
	envVariables["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	envVariables["DB_NAME"] = os.Getenv("DB_NAME")

	envVariables["WHITELISTED_URLS"] = os.Getenv("WHITELISTED_URLS")
	return envVariables
}

func initializeDatabase(envVariables map[string]string) *sql.DB {
	return database.GetDatabase(
		envVariables["DB_USERNAME"],
		envVariables["DB_PASSWORD"],
		envVariables["DB_ADDRESS"],
		envVariables["DB_NAME"],
	)
}

func initializeGlobalRouter(envVariables map[string]string) *mux.Router {
	r := mux.NewRouter()

	arrayWhitelistedUrls := strings.Split(envVariables["WHITELISTED_URLS"], ",")

	whitelistedUrls := make(map[string]bool)

	for _, v := range arrayWhitelistedUrls {
		whitelistedUrls[v] = true
	}
	r.Use(middleware.ContentTypeMiddleware)
	r.Use(middleware.CorsMiddlerware(whitelistedUrls))
	return r
}

func main() {

	godotenv.Load()
	envVariables := getEnvVariables()
	db := initializeDatabase(envVariables)

	r := initializeGlobalRouter(envVariables)

	controller.InitializeControllers(r, db)

	server := server.ProvideServer(envVariables["SERVER_ADDRESS"], r)
	server.ListenAndServe()
}
