package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	usercontroller "github.com/dynamodb/controller/user"
	"github.com/dynamodb/entities/user"
	"github.com/dynamodb/repository/database"
	"github.com/dynamodb/repository/instance"
	"github.com/google/uuid"
)

func main() {
	createTB()
	connecttion := instance.GetConnection()
	database := database.NewDatabase(connecttion)
	contrller := usercontroller.NewUserController(database)
	user := &user.User{}
	user.ID = uuid.New()
	user.UserName = "Abcd"
	user.Address = "HN"
	err := contrller.InsertUser(user)
	fmt.Print(err)
}

func createTB() {
	connecttion := instance.GetConnection()
	// Create table Movies
	tableName := "User"

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("_id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("_id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := connecttion.CreateTable(input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table", tableName)

}
