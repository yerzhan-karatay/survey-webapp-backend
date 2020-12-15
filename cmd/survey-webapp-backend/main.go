package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/yerzhan-karatay/survey-webapp-backend/config"
	"github.com/yerzhan-karatay/survey-webapp-backend/db"
	"github.com/yerzhan-karatay/survey-webapp-backend/errors"
	"github.com/yerzhan-karatay/survey-webapp-backend/services/auth"
	"github.com/yerzhan-karatay/survey-webapp-backend/services/user"

	_ "github.com/yerzhan-karatay/survey-webapp-backend/docs"
)

// @title Swagger Survey service API
// @version 1.0
// @description Survey service.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl http://localhost:8080/login
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

func main() {
	r := gin.Default()
	r.Use(errors.HandleHTTPError())
	db := db.Get()

	// User service
	userRepo := user.GetRepository(db)
	userService := user.GetService(userRepo)

	r = user.MakeHTTPHandler(r, userService)

	// Auth service
	authRepo := auth.GetRepository(db)
	authService := auth.GetService(authRepo)

	r = auth.MakeHTTPHandler(r, authService)

	// Swagger settings
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	errs := make(chan error)
	cfg := config.Get()
	httpPort := fmt.Sprintf(":%d", cfg.HTTP.Port)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Println("transport:", "HTTP", "port:", httpPort)
		errs <- http.ListenAndServe(httpPort, r)
	}()

	log.Fatal(<-errs, "exit")
}
