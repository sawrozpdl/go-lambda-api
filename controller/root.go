package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sawrozpdl/go-lambda-api/utils"
	"net/http"
)

func Hello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello from GO Lambda API!",
		"data":    utils.Short("Test this string truncate", 9),
	})
}
