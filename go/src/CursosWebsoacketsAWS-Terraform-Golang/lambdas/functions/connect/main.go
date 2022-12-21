package main

import (
	"context"
	"fmt"
	"github.com/andreslopez2612/poc-websockets-golang-terraform-aws/functions"
	"github.com/andreslopez2612/poc-websockets-golang-terraform-aws/log"
	websocketmodel "github.com/andreslopez2612/poc-websockets-golang-terraform-aws/model/websocket"
	servicewebsocket "github.com/andreslopez2612/poc-websockets-golang-terraform-aws/services/websocket"
	"github.com/andreslopez2612/poc-websockets-golang-terraform-aws/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func LambdaHandler(cxt context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger.Debug().Msg("Start lambda createDynamoTrigger websocket connection")
	log.Logger.Debug().Msgf("Start lambda createDynamoTrigger claims %v", event.RequestContext.Authorizer)
	websocket := &websocketmodel.WebSocket{
		ID: event.RequestContext.ConnectionID,
	}
	log.Logger.Debug().Msgf("id %v", websocket)
	err := functions.Instances["websocketService"].(servicewebsocket.WebSocketService).Create(cxt, websocket)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on the connect %v", websocket)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
