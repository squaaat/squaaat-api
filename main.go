package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//import (
//	"github.com/squaaat/squaaat-api/cmd"
//	_ "github.com/squaaat/squaaat-api/docs"
//)
//
//func main() {
//	cmd.Start()
//}


// https://github.com/awslabs/aws-lambda-go-api-proxy
func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "text/plain; charset=utf-8"},
		Body:       fmt.Sprintf("Hello, %s!\n", req.Path),
	}
	return res, nil
}

func main() {
	lambda.Start(handleRequest)
}