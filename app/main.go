package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"crm-lambda/bootstrap"
	"crm-lambda/port/portin"
)

func main() {
	i := bootstrap.NewInject()
	portin.NewHandler(i)

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
