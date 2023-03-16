package main

import (
	"fmt"
	"log"
	"manage_address/pkg"
	"manage_address/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(pkg.RunMode)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", pkg.HTTPPort),
		Handler:        routers.InitRouter(),
		ReadTimeout:    pkg.ReadTimeout,
		WriteTimeout:   pkg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println(fmt.Sprintf("manage address start success, port: %d", pkg.HTTPPort))
	s.ListenAndServe()
}
