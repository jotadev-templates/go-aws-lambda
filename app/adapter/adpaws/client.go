package adpaws

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

/**********************************************************
* SINGLETON
**********************************************************/
var (
	mu      sync.RWMutex
	clients awsClients
)

type awsClients struct {
	dynamoDB *dynamodb.Client
}

/**********************************************************
* CLIENTS
**********************************************************/

func getConfig() (aws.Config, error) {
	const envAWSRegion string = "PROVIDE_AWS_REGION"
	region, ok := os.LookupEnv(envAWSRegion)
	if !ok {
		return aws.Config{}, fmt.Errorf("aws region is invalid")
	}

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, func(opts *config.LoadOptions) error {
		opts.Region = region
		return nil
	})
	if err != nil {
		return cfg, fmt.Errorf("aws config is invalid: %v", err)
	}
	return cfg, err
}

func getDynamoDB() (*dynamodb.Client, error) {
	mu.Lock()
	defer mu.Unlock()

	if clients.dynamoDB == nil {
		cfg, err := getConfig()
		if err != nil {
			return nil, err
		}
		clients.dynamoDB = dynamodb.NewFromConfig(cfg)
	}
	return clients.dynamoDB, nil
}
