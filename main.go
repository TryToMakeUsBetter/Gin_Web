package main

import (
	"fmt"
	"net/http"

	settings "GIN_WEB/pkg"
	"GIN_WEB/routers"
)

func main() {
	/***
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
	***/

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
