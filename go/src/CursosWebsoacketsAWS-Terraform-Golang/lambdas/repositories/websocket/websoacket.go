package repositorywebsocket

import (
	"context"
	"fmt"
	config "github.com/andreslopez2612/poc-websockets-golang-terraform-aws/config/constans"
	"github.com/andreslopez2612/poc-websockets-golang-terraform-aws/log"
	websocketmodel "github.com/andreslopez2612/poc-websockets-golang-terraform-aws/model/websocket"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// WebSocketRepository describes the lock repository
type WebSocketRepository interface {
	Create(context.Context, *websocketmodel.WebSocket) error
	GetAll(context.Context) ([]websocketmodel.WebSocket, error)
	Delete(context.Context, string) error
}

// NewRepository creates and return a new Websocket repository instance
func NewRepository(database *dynamodb.DynamoDB) WebSocketRepository {
	return &repository{
		database: database,
		table:    config.POCWebSocketConnectionTable,
	}
}

type repository struct {
	database *dynamodb.DynamoDB
	table    string
}

func (r *repository)  Create(ctx context.Context, socket *websocketmodel.WebSocket) (err error) {
	log.Logger.Debug().Msgf("Adding a new notification [%s] ", socket.ID)

	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, config.Timeout)
		defer cancel()
	}

	av, err := dynamodbattribute.MarshalMap(socket)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = r.database.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(r.table),
		Item:      av,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Logger.Debug().Msgf("ID %v inserted. \n", socket.ID)
	return nil
}

func (r *repository) GetAll(ctx context.Context) (connections []websocketmodel.WebSocket, err error) {
	log.Logger.Debug().Msgf("Getting connections")
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, config.Timeout)
		defer cancel()
	}

	input := &dynamodb.ScanInput{
		TableName: aws.String(r.table),
	}

	result, err := r.database.Scan(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, item := range result.Items {
		websocket := &websocketmodel.WebSocket{}
		err = dynamodbattribute.UnmarshalMap(item, &websocket)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		connections = append(connections, *websocket)
	}

	return connections, nil
}

func (r *repository) Delete(ctx context.Context, id string) (err error) {
	log.Logger.Debug().Msgf("Deleting a connection [%s] ", id)
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, config.Timeout)
		defer cancel()
	}
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"connectionId": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(r.table),
	}

	_, err = r.database.DeleteItem(input)
	if err != nil {
		fmt.Println("Hot error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}
	return nil
}
