package adpaws

type DynamoDBInterface interface {
	FindByID(table, key, id string, output *any) error
	Create(table, key string, dataDomain any) error
	UpsertAll(table string, dataDomain any) error
	Delete(table, key, id string) error
}

type dynamoDBImpl struct{}

func NewDynamoDB() DynamoDBInterface {
	return new(dynamoDBImpl)
}
