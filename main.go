package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sawrozpdl/go-lambda-api/router"
	"log"
	"os"
)

var goEnv string
var ginEngine *gin.Engine
var ginLambda *ginadapter.GinLambda

func init() {
	godotenv.Load(".env")

	goEnv = os.Getenv("GO_ENV")

	ginEngine = gin.Default()
	router.InjectRoutes(ginEngine)

	// Initialize  the gin lambda instance for non-local builds.
	if goEnv != "local" {
		ginLambda = ginadapter.New(ginEngine)
	}
}

func Handler(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {

	// Map the AWS generated events with gin injected Routes.
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	port := os.Getenv("PORT")
	if goEnv == "local" {
		err := ginEngine.Run(":" + port)
		if err != nil {
			log.Printf("Error while starting server %+v", err)
		}
	} else {
		lambda.Start(Handler)
	}
}
