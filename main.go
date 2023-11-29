package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/m0hammedimran/CloudFunctions/EnterpriseRedirection/utils"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Println("JWT_SECRET is not set")
		jwtSecret = "secret"
	}
	BLACKLIST_ID := os.Getenv("BLACKLIST_ID")
	if BLACKLIST_ID == "" {
		log.Println("BLACKLIST_ID is not set")
		BLACKLIST_ID = "1234"
	}

	var AuthHeader = event.Headers["authorization"]
	if AuthHeader == "" {
		log.Println("Could not get the token from the header")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusUnauthorized, Body: "Unauthorized"}, nil
	}

	AuthHeader = strings.Replace(AuthHeader, "Bearer ", "", 1)
	claims, err := utils.VerifyJWT(AuthHeader, jwtSecret)
	if err != nil {
		log.Println("Could not decode the JWT", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusUnauthorized, Body: "Unauthorized"}, nil
	}

	var responseHeaders = map[string]string{
		"Location": "https://www.google.com",
	}

	if strconv.Itoa(claims.UserId) == BLACKLIST_ID {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusMovedPermanently, Headers: responseHeaders}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: fmt.Sprintf("Valid Enterprize: %d", claims.UserId)}, nil
}

func main() {
	lambda.Start(handler)
}
