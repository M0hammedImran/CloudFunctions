package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/m0hammedimran/CloudFunctions/EnterpriseRedirection/utils"
)

func includes(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var httpMethod = event.HTTPMethod
	if !includes([]string{http.MethodPost, http.MethodPut, http.MethodPatch}, httpMethod) {
		log.Println("Invalid HTTP Method")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusMethodNotAllowed, Body: "Method Not Allowed"}, nil
	}

	var AuthHeader = event.Headers["Authorization"]
	bytes, _ := json.MarshalIndent(event, "", "  ")
	log.Println(string(bytes))

	if AuthHeader == "" {
		log.Println("Could not get the token from the header")

		return events.APIGatewayProxyResponse{StatusCode: http.StatusUnauthorized, Body: "Unauthorized"}, nil
	}

	jwt, err := utils.VerifyJWT(AuthHeader, "")
	if err != nil {
		log.Println("Could not decode the JWT")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusUnauthorized, Body: "Unauthorized"}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: fmt.Sprintf("Enterprize: %d", jwt.Id)}, nil
}

func main() {
	lambda.Start(handler)
}
