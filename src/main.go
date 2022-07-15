package main

import (
	config_base "demo1/src/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	CONFIG := config_base.Config()
	r := gin.Default()

	r.Static("/public", "./public")
	r.Static("/images", "./images")
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/public/")
	})
	r.Run(fmt.Sprintf(":%d", CONFIG.Base.Port))
}