/**
 * @date: 2022/7/27
 * @desc:
 */

package core

import (
	"context"
	"fmt"
	"github.com/daniuEvan/mygithub/global"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func ListenAndServe(router *gin.Engine) {

	var host = global.ServerConfig.Host
	var port = global.ServerConfig.Port
	var addr = fmt.Sprintf("%s:%d", host, port)
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen err:%s", err)
		}
	}()
	log.Printf("server listen http://%s:%d", host, port)

	// Wait for interrupt signal to gracefully with shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't catch, so don't need to add it
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error")
	}
	log.Println("server exited")
}
