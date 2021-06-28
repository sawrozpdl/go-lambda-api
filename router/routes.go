package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sawrozpdl/go-lambda-api/controller"
)

func InjectRoutes(r *gin.Engine) {
	r.GET("/", controller.Hello)
}
