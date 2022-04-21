package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/codedius/imagekit-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	imagekitServicePkg "github.com/kuroyamii/go-backend-learn/internal/imagekit/service"
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
	envVariables["IMAGEKIT_PUBLIC_KEY"] = os.Getenv("IMAGEKIT_PUBLIC_KEY")
	envVariables["IMAGEKIT_PRIVATE_KEY"] = os.Getenv("IMAGEKIT_PRIVATE_KEY")

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

func initializeImageKit(envVariables map[string]string) (*imagekit.Client, error) {
	return imagekitServicePkg.CreateNewClient(envVariables["IMAGEKIT_PUBLIC_KEY"], envVariables["IMAGEKIT_PRIVATE_KEY"])
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
	imageKit, err := initializeImageKit(envVariables)
	if err != nil {
		log.Println("ERROR (Imagekit): Error while creating client")
	}

	r := initializeGlobalRouter(envVariables)

	controller.InitializeControllers(r, db, imageKit)

	server := server.ProvideServer(envVariables["SERVER_ADDRESS"], r)
	server.ListenAndServe()
}
