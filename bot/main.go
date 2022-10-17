package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CasperDev394/goClient/getinfo"
	"github.com/CasperDev394/goClient/getinfo/types"
	"github.com/CasperDev394/goClient/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server...")

	router := gin.Default()

	handler.NewHandler(&handler.Config{
		R: router,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	/*
		httpClient := &http.Client{
			Timeout: time.Second * 10,
		}

		r := getinfo.NewClient(httpClient)
		result, err := r.Ping()
		if err != nil {
			fmt.Printf("Fail %s", err)
		}
		fmt.Println(result)
		fmt.Println(reflect.TypeOf(result))
	*/
}

func get() *types.Ping {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	r := getinfo.NewClient(httpClient)
	result, err := r.Ping()
	if err != nil {
		log.Printf("Fail %s", err)
	}
	//log.Println(result)
	//log.Println(reflect.TypeOf(result))
	return result
}
