package main

import (
	"basic-lambda-startup/cmd/lambda"
	"basic-lambda-startup/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	lambda := lambda.NewLambda(logger)
	lambda.Start()
}
