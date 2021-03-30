package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"google.com/sideroff/golang-k8s/config"
	"google.com/sideroff/golang-k8s/utils"
	"google.com/sideroff/golang-k8s/webserver"
)

func main() {
	log.Info(fmt.Sprintf("START - %s", time.Now().Format(time.RFC3339)))
	
	utils.ConfigureLogger()
	log.Info("Logger configured")

	ctx := initializeContext()


	log.Info("Initializing config")
	config := config.Get()
	log.Info("Config initialized")


	log.Info("Starting server")
	go webserver.Start(ctx, config)
	
	log.Info("main blocking on context")
	<-ctx.Done()
}

func initializeContext() context.Context {
	ctx := context.Background()

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx
}