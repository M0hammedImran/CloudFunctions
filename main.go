package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/m0hammedimran/CloudFunctions/EnterpriseRedirection/utils"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var AuthHeader = event.Headers["authorization"]

	if AuthHeader == "" {
		log.Println("Could not get the token from the header")

		return events.APIGatewayProxyResponse{StatusCode: http.StatusUnauthorized, Body: "Unauthorized"}, nil
	}
	AuthHeader = strings.Replace(AuthHeader, "Bearer ", "", 1)
	jwt, err := utils.DecodeJwtClaim(AuthHeader)
	if err != nil {
		log.Println("Could not decode the JWT", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusUnauthorized, Body: "Unauthorized"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: fmt.Sprintf("Valid Enterprize: %d", jwt.Id)}, nil
}

func main() {
	lambda.Start(handler)
}
