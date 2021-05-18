package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"

	"github.com/konstantinfarrell/go-example/pkg/util/log"	
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", healthCheck)
	return e
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

type Config struct {
	Port 			string
	LogLevel		string
	TimeoutSeconds	int
}

func Start(e *echo.Echo, conf *Config){
	s := &http.Server{
		Addr:			conf.Port,
		ReadTimeout: 	time.Duration(conf.TimeoutSeconds) * time.Second,
		WriteTimeout: 	time.Duration(conf.TimeoutSeconds) * time.Second,
	}
	e.Debug = log.IsDebug(conf.LogLevel)
	e.Logger.SetLevel(log.GetLogLevel(conf.LogLevel))
	e.Logger.Debug("Set Log Levels")

	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("Server shutting down, %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}