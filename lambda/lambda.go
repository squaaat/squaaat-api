package main

import (
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	adapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"

	"github.com/squaaat/squaaat-api/internal/app"
	"github.com/squaaat/squaaat-api/internal/config"
	serverhttp "github.com/squaaat/squaaat-api/internal/server/http"
)

func main() {
	env := os.Getenv("SQ_ENV")
	sqcicd := os.Getenv("SQ_CICD")
	cicd, _ := strconv.ParseBool(sqcicd)

	cfg := config.MustInit(env, cicd)

	app := app.New(cfg)
	http := serverhttp.New(app)
	lambdaApp := adapter.New(http)

	lambda.Start(lambdaApp.Proxy)
}
