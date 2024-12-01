package dynamodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"strings"
)

type TableMode string

const (
	TableModeSingleTable TableMode = "single-table"
	TableModeMultiTable  TableMode = "multi-table"
)

type dynamoDbBackend struct {
	cfg       aws.Config
	tableMode TableMode
	tableName string
	svc       *dynamodb.Client
	schema    *abs.Schema
}

func (d *dynamoDbBackend) GetStatus(ctx context.Context) (connectionAlreadyInitiated bool, testConnection bool, err error) {
	return true, true, nil
}

func (d *dynamoDbBackend) DestroyDataSource(ctx context.Context) {

}

func (d *dynamoDbBackend) AddRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, error) {
	if len(records) == 0 {
		return nil, nil
	}
	if len(records) == 1 {
		// Create the Item to be added
		item := d.recordToAttributeMap(resource, records[0])

		// Create the PutItemInput
		input := &dynamodb.PutItemInput{
			TableName:           &d.tableName,
			Item:                item,
			ConditionExpression: aws.String("attribute_not_exists(#nameProperty)"),
			ExpressionAttributeNames: map[string]string{
				"#nameProperty": "name",
			},
		}

		// Perform the PutItem operation
		_, err := d.svc.PutItem(ctx, input)
		if err != nil {
			return nil, d.handleError(err)
		}
	} else if len(records) > 1 {
		var writeRequests []types.WriteRequest
		for _, record := range records {
			item := d.recordToAttributeMap(resource, record)

			writeRequests = append(writeRequests, types.WriteRequest{
				PutRequest: &types.PutRequest{
					Item: item,
				},
			})
		}

		_, err := d.svc.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				d.tableName: writeRequests,
			},
		})

		if err != nil {
			return nil, d.handleError(err)
		}
	}

	return records, nil
}

func (d *dynamoDbBackend) UpdateRecords(ctx context.Context, resource *model.Resource, records []*model.Record) ([]*model.Record, error) {
	if len(records) == 0 {
		return nil, nil
	}
	if len(records) == 1 {
		// Create the Item to be added
		item := d.recordToAttributeMap(resource, records[0])

		// Create the PutItemInput
		input := &dynamodb.PutItemInput{
			TableName: &d.tableName,
			Item:      item,
		}

		// Perform the PutItem operation
		_, err := d.svc.PutItem(ctx, input)
		if err != nil {
			return nil, d.handleError(err)
		}
	} else if len(records) > 1 {
		var writeRequests []types.WriteRequest
		for _, record := range records {
			item := d.recordToAttributeMap(resource, record)

			writeRequests = append(writeRequests, types.WriteRequest{
				PutRequest: &types.PutRequest{
					Item: item,
				},
			})
		}

		_, err := d.svc.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				d.tableName: writeRequests,
			},
		})

		if err != nil {
			return nil, d.handleError(err)
		}
	}

	return records, nil
}

func (d *dynamoDbBackend) GetRecord(ctx context.Context, resource *model.Resource, id string, resolveReferences []string) (*model.Record, error) {
	input := &dynamodb.GetItemInput{
		TableName: &d.tableName,
		Key: map[string]types.AttributeValue{
			"PK": d.getPKForResource(resource),
			"SK": &types.AttributeValueMemberS{Value: id},
		},
	}

	resp, err := d.svc.GetItem(ctx, input)

	if err != nil {
		return nil, d.handleError(err)
	}

	if resp.Item == nil {
		return nil, nil
	}

	return d.convertAttributeMapToRecord(resource, resp.Item)
}

func (d *dynamoDbBackend) DeleteRecords(ctx context.Context, resource *model.Resource, ids []*model.Record) error {
	if len(ids) == 0 {
		return nil
	}

	if len(ids) == 1 {
		input := &dynamodb.DeleteItemInput{
			TableName: &d.tableName,
			Key: map[string]types.AttributeValue{
				"PK": d.getPKForResource(resource),
				"SK": &types.AttributeValueMemberS{Value: ids[0].Properties["id"].GetStringValue()},
			},
		}

		_, err := d.svc.DeleteItem(ctx, input)

		if err != nil {
			return d.handleError(err)
		}
	} else if len(ids) > 1 {
		var writeRequests []types.WriteRequest
		for _, record := range ids {
			writeRequests = append(writeRequests, types.WriteRequest{
				DeleteRequest: &types.DeleteRequest{
					Key: map[string]types.AttributeValue{
						"PK": d.getPKForResource(resource),
						"SK": &types.AttributeValueMemberS{Value: record.Properties["id"].GetStringValue()},
					},
				},
			})
		}

		_, err := d.svc.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				d.tableName: writeRequests,
			},
		})

		if err != nil {
			return d.handleError(err)
		}
	}

	return nil
}

func (d *dynamoDbBackend) ListRecords(ctx context.Context, resource *model.Resource, params abs.ListRecordParams, resultChan chan<- *model.Record) ([]*model.Record, uint32, error) {
	id, isIdQuery := d.isIdQuery(params)

	if isIdQuery {
		record, err := d.GetRecord(ctx, resource, id, params.ResolveReferences)

		if err != nil {
			return nil, 0, err
		}

		return []*model.Record{record}, 1, nil
	}

	filterExpression, expressionAttributeValues, err := d.prepareQuery(params.Query)

	input := &dynamodb.QueryInput{
		TableName: &d.tableName,
		KeyConditions: map[string]types.Condition{
			"PK": {
				ComparisonOperator: types.ComparisonOperatorEq,
				AttributeValueList: []types.AttributeValue{
					d.getPKForResource(resource),
				},
			},
		},
	}

	if len(expressionAttributeValues) > 0 {
		input.ExpressionAttributeValues = expressionAttributeValues
	}
	if strings.Contains(filterExpression, "name") {
		input.ExpressionAttributeNames = map[string]string{
			"#nameProperty": "name",
		}
		filterExpression = strings.ReplaceAll(filterExpression, "name", "#nameProperty")
	}
	if filterExpression != "" && filterExpression != "()" {
		input.FilterExpression = &filterExpression
	}

	resp, err := d.svc.Query(ctx, input)

	if err != nil {
		return nil, 0, d.handleError(err)
	}

	// Check if the item exists
	var records []*model.Record

	// Iterate over the attributes in the item
	for _, item := range resp.Items {
		record, err := d.convertAttributeMapToRecord(resource, item)

		if err != nil {
			return nil, 0, err
		}

		records = append(records, record)
	}

	return records, uint32(resp.Count), nil
}

func (d *dynamoDbBackend) ListEntities(ctx context.Context) ([]*model.DataSourceCatalog, error) {
	return nil, errors.New("ListEntities not supported")
}

func (d *dynamoDbBackend) PrepareResourceFromEntity(ctx context.Context, catalog, entity string) (*model.Resource, error) {
	resp, err := d.svc.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &d.tableName,
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: "schema"},
			"SK": &types.AttributeValueMemberS{Value: fmt.Sprintf("%s#%s", catalog, entity)},
		},
	})

	if err != nil {
		return nil, d.handleError(err)
	}

	// Check if the item exists
	if resp.Item == nil {
		return nil, nil
	}

	// Iterate over the attributes in the item
	for key, attr := range resp.Item {
		fmt.Printf("Key: %s\n", key)

		// Access the attribute value based on its type
		switch v := attr.(type) {
		case *types.AttributeValueMemberS:
			fmt.Printf("Value (string): %s\n", v.Value)
		case *types.AttributeValueMemberN:
			fmt.Printf("Value (number): %s\n", v.Value)
		case *types.AttributeValueMemberB:
			fmt.Printf("Value (binary): %v\n", v.Value)
		// Add more cases for other attribute types as needed
		default:
			fmt.Printf("Value (unknown type): %v\n", v)
		}
	}

	return nil, nil
}

func (d *dynamoDbBackend) isIdQuery(params abs.ListRecordParams) (string, bool) {
	return "", false
}

func (d *dynamoDbBackend) UpgradeResource(ctx context.Context, params abs.UpgradeResourceParams) error {
	// TODO implement me
	return nil
}

func (d *dynamoDbBackend) SetSchema(schema *abs.Schema) {
	d.schema = schema
}

func (d *dynamoDbBackend) handleError(err error) error {
	return err
}

func (d *dynamoDbBackend) getPKForResource(resource *model.Resource) types.AttributeValue {
	return &types.AttributeValueMemberS{Value: fmt.Sprintf("%s/%s", resource.Namespace, resource.Name)}
}

func NewDynamodbBackend(dataSource abs.DataSource) abs.Backend {
	var options = dataSource.(*resource_model.DataSource).Options

	var opts []func(*config.LoadOptions) error

	if options["region"] != "" {
		opts = append(opts, config.WithRegion(options["region"]))
	}

	if options["accessKey"] != "" {
		opts = append(opts, config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     options["accessKey"],
				SecretAccessKey: options["secretKey"],
			},
		}))
	}

	if options["tableName"] == "" {
		panic("tableName is required")
	}

	if options["tableMode"] == "" {
		panic("tableMode is required")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), opts...)
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	return &dynamoDbBackend{
		svc:       svc,
		cfg:       cfg,
		tableName: options["tableName"],
		tableMode: TableMode(options["tableMode"]),
	}
}
