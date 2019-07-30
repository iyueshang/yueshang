package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"yueshang/routes"
)

func main() {
	gin.SetMode("debug")

	routersInit := routes.SetupRouter()
	endPoint := fmt.Sprintf(":%d", 8080)

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    60,
		WriteTimeout:   60,
		MaxHeaderBytes:  1 << 20,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		log.Println("log:", err)
	}
}