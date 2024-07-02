package main

import (
	"api-gin/migrate"
	"github.com/gin-contrib/pprof"
	"log"
	"net/http"

	"api-gin/controllers"
	"api-gin/initializers"
	"api-gin/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	migrate.Migrate()

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

	gin.SetMode(gin.ReleaseMode) // [ReleaseMode, DebugMode, TestMode]
	server = gin.Default()

}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	// http://localhost:8080/api/healthchecker
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	// ------ в Gin очень просто подлючаются миддлвары ------
	//swaggerURL := ginSwagger.URL("http://localhost:" + config.ServerPort + "/swagger/doc.json")
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	// http://localhost:8080/api/debug/pprof/
	pprof.Register(router)

	PostRouteController.PostRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
