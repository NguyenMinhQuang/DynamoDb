package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type DataBase struct {
	connection *dynamodb.DynamoDB
}

type IDatabase interface {
	Health() bool
	CreateOrUpdate(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error)
	FindOne(condition map[string]interface{}, tableName string) (*dynamodb.GetItemOutput, error)
	FindAll(condition expression.Expression, tableName string) (*dynamodb.ScanOutput, error)
	Delete(condition map[string]interface{}, tableName string) (*dynamodb.DeleteItemOutput, error)
}

func NewDatabase(conn *dynamodb.DynamoDB) IDatabase {
	return &DataBase{
		connection: conn,
	}
}

func (db *DataBase) Health() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	return err == nil
}

// Input data table
func (db *DataBase) CreateOrUpdate(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error) {
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.PutItem(input)
}

// Select item from table
func (db *DataBase) FindOne(condition map[string]interface{}, tableName string) (*dynamodb.GetItemOutput, error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       conditionParsed,
	}
	return db.connection.GetItem(input)
}

// Select all data table
func (db *DataBase) FindAll(condition expression.Expression, tableName string) (*dynamodb.ScanOutput, error) {
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(tableName),
	}
	return db.connection.Scan(params)
}

// Delete item table
func (db *DataBase) Delete(condition map[string]interface{}, tableName string) (*dynamodb.DeleteItemOutput, error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.DeleteItemInput{
		Key:       conditionParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.DeleteItem(input)
}
