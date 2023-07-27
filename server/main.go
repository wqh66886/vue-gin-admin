package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wqh66886/vue-gin-admin/server/server/handler"
)

func main() {
	log.Println("server staring....")
	router := gin.Default()

	// 创建简单的一个Get请求
	handler.NewHandler(&handler.Config{
		R: router,
	})

	server := &http.Server{
		Addr:    ":8083",
		Handler: router,
	}
	// 以下代码作用是 gin优雅地的退出服务端
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()
	log.Printf("Listening on port %v\n", server.Addr)

	// wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// this blocks until a signal is pass into the quit channel
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("Shutting down server ...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v\n", err)
	}
}
