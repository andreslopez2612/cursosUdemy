package websocketservice

import (
	"context"
	"errors"
	"github.com/andreslopez2612/poc-websockets-golang-terraform-aws/log"
	websocketmodel "github.com/andreslopez2612/poc-websockets-golang-terraform-aws/model/websocket"
	repositorywebsocket "github.com/andreslopez2612/poc-websockets-golang-terraform-aws/repositories/websocket"
)

type WebSocketService interface {
	Create(context.Context, *websocketmodel.WebSocket) error
	Delete(context.Context, string) error
}

// New creates and returns a new lock service instance
func New(rep repositorywebsocket.WebSocketRepository) WebSocketService {
	return &websocketService{
		repository: rep,
	}
}

type websocketService struct {
	repository repositorywebsocket.WebSocketRepository
}

func (w websocketService) Create(ctx context.Context, socket *websocketmodel.WebSocket) error {
	log.Logger.Debug().Msg("Creating web socket connection")
	err := w.repository.Create(ctx, socket)
	if err != nil {
		return errors.Unwrap(err) //, "Error creating websoacket connection on services [%s]", socket.ID
	}
	return nil
}

func (w websocketService) Delete(ctx context.Context, s string) error {
	log.Logger.Debug().Msg("Deleting a websocket connection on services")
	err := w.repository.Delete(ctx, s)
	if err != nil {
		return errors.Unwrap(err)
	}
	return nil
}
