package util

import (
	"encoding/json"
	"fmt"
	"github.com/andreslopez2612/poc-websockets-golang-terraform-aws/log"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type GenericError struct {
	Body Body `json:"body"`
}

type Body struct {
	Code         int32  `json:"code"`
	Message      string `json:"message"`
	MessageError error  `json:"messageError"`
}

func buildGenericError(msg string, err error) []byte {
	genericError := new(GenericError)
	genericError.Body = Body{
		Code:         500,
		Message:      msg,
		MessageError: err,
	}
	data, err := json.Marshal(genericError)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Error Marshal generic error app")
		data = []byte(fmt.Sprintf("500 - Internal server"))
	}
	return data
}

func ResponseErrorFunction(err error, msg string) events.APIGatewayProxyResponse {
	log.Logger.Error().Err(err).Msg(msg)
	data := buildGenericError(msg, err)
	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       string(data),
	}
	return response
}
