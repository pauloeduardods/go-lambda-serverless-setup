package lambda

import (
	"basic-lambda-startup/pkg/logger"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Lambda struct {
	logger *logger.Logger
}

func NewLambda(l *logger.Logger) *Lambda {
	return &Lambda{
		logger: l,
	}
}

func (l *Lambda) handleRequest(ctx context.Context, event interface{}) (string, error) {
	l.logger.Info("Event: %v", event)
	return "Hello, World!", nil
}

func (l *Lambda) Start() {
	l.logger.Info("Starting lambda")
	lambda.Start(l.handleRequest)
}
