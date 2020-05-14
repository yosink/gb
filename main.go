package main

import (
	"blog/config"
	"blog/logging"
	"blog/pkg/gredis"
	"blog/routers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.Setup()
	logging.Setup()
	_ = gredis.Setup()
	r := routers.Setup()

	srv := &http.Server{
		Addr:           ":8089",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen err: %v", err)
		}
	}()

	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	log.Println("shutting down server .....")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown error:%v", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout 5 seconds")

	}
	log.Println("server exiting")
}
