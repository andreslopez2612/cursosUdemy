package functions

import (
	"github.com/andreslopez2612/poc-websockets-golang-terraform-aws/config/db"
	"github.com/andreslopez2612/poc-websockets-golang-terraform-aws/log"
	repositorywebsocket "github.com/andreslopez2612/poc-websockets-golang-terraform-aws/repositories/websocket"
	servicewebsocket "github.com/andreslopez2612/poc-websockets-golang-terraform-aws/services/websocket"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var Instances = MakeDependencyInjection()

func MakeDependencyInjection() map[string]interface{} {
	log.Logger.Debug().Msgf("Start boostrap app objects")
	instances := make(map[string]interface{})

	database, err := db.NewDynamoDBStorage()
	if err != nil {
		return nil
	}
	instances["dataBase"] = database

	instances["websocketRepository"] = repositorywebsocket.NewRepository(database.GetConnection().(*dynamodb.DynamoDB))

	instances["websoacketService"] = servicewebsocket.New(
		instances["websocketRepository"].(repositorywebsocket.WebSocketRepository))

	log.Logger.Debug().Msg("End bootstrap app objects")
	return instances

}
