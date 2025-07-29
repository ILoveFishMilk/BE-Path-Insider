package main

import (
	"context"
	"log"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"
)

func main() {
	srv := &http.Server{Addr: ":8080", Handler: nil} //router missing

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	//interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("forced to shutdown: %s", err)
	}

	log.Println("exiting")
}
