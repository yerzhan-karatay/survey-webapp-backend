package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/yerzhan-karatay/survey-webapp-backend/config"
	"github.com/yerzhan-karatay/survey-webapp-backend/errors"
	"github.com/yerzhan-karatay/survey-webapp-backend/services"
)

func main() {
	r := gin.New()
	r.Use(errors.HandleHTTPError())

	r = services.MakeHTTPHandler(r)

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
