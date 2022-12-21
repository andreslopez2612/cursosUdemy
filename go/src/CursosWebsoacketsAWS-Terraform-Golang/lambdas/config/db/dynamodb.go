package db

import (
	"github.com/andreslopez2612/poc-websockets-golang-terraform-aws/log"
	_ "github.com/andreslopez2612/poc-websockets-golang-terraform-aws/log"
	"github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	_ "github.com/aws/aws-sdk-go/service/dynamodb"
)

type Database interface {
	OpenConnection() error
	GetConnection() interface{}
}

type dynamoDataBase struct {
	databaseConnection *dynamodb.DynamoDB
}

func NewDynamoDBStorage() (Database, error) {
	log.Logger.Debug().Msgf("New Instance Dynamo storage")
	dataBase := &dynamoDataBase{}
	err := dataBase.OpenConnection()
	if err != nil {
		return nil, err
	}

	return dataBase, nil

}

func (db *dynamoDataBase) OpenConnection() error {
	log.Logger.Info().Msg("Starting Dynamo connection")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	db.databaseConnection = dynamodb.New(sess)
	log.Logger.Info().Msg("DynamoDB UP")
	return nil
}

func (db *dynamoDataBase) GetConnection() interface{} {
	return db.databaseConnection
}
