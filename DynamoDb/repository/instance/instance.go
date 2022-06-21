package instance

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	RegionName = endpoints.ApNortheast1RegionID
	URL        = "http://localhost:8000"
	ACCESSKEY  = "test"
	SECRECTKEY = "test"
	TOKEN      = ""
)

func GetConnection() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:    aws.String(URL),
		Region:      aws.String(RegionName),
		Credentials: credentials.NewStaticCredentials(ACCESSKEY, SECRECTKEY, TOKEN),
	}))
	return dynamodb.New(sess)
}
