package adpaws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (*dynamoDBImpl) FindByID(table, key, id string, output *any) error {
	ctx := context.TODO()

	cli, err := getDynamoDB()
	if err != nil {
		return err
	}

	result, err := cli.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			key: &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return err
	}
	return attributevalue.UnmarshalMap(result.Item, &output)
}

func (*dynamoDBImpl) Create(table, key string, dataDomain any) error {
	ctx := context.TODO()

	cli, err := getDynamoDB()
	if err != nil {
		return err
	}

	item, err := attributevalue.MarshalMap(dataDomain)
	if err != nil {
		return err
	}
	_, err = cli.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(table),
		Item:                item,
		ConditionExpression: aws.String(fmt.Sprintf("attribute_not_exists(%s)", key)),
	})
	return err
}

func (*dynamoDBImpl) UpsertAll(table string, dataDomain any) error {
	ctx := context.TODO()

	cli, err := getDynamoDB()
	if err != nil {
		return err
	}

	item, err := attributevalue.MarshalMap(dataDomain)
	if err != nil {
		return err
	}
	_, err = cli.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      item,
	})
	return err
}

func (*dynamoDBImpl) Delete(table, key, id string) error {
	ctx := context.TODO()

	cli, err := getDynamoDB()
	if err != nil {
		return err
	}

	_, err = cli.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			key: &types.AttributeValueMemberS{Value: id},
		},
		ConditionExpression: aws.String(fmt.Sprintf("attribute_exists(%s)", key)),
	})
	return err
}
