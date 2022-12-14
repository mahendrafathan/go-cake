package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mahendrafathan/go-cake/common/logger"
	"github.com/mahendrafathan/go-cake/routes"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	router := routes.RegisterRoutes()
	srv := &http.Server{
		Addr:         "0.0.0.0:9090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatalf(ctx, err.Error())
		}
	}()

	logger.Infof(ctx, "Server is Started and Listening on port: 9090")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	logger.Infof(ctx, "server stopping...")
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		logger.Infof(ctx, err.Error())
	}
}
