package main

import (
	"app.modules/aws-lambda/lambdautils"
	"app.modules/core"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type CheckBotAccessResponseStruct struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

func CheckBotAccess(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("CheckLiveStream()")
	
	ctx := context.Background()
	clientOption, err := lambdautils.FirestoreClientOption()
	if err != nil {
		return lambdautils.ErrorResponse(err)
	}
	_system, err := core.NewSystem(ctx, clientOption)
	if err != nil {
		return lambdautils.ErrorResponse(err)
	}
	defer _system.CloseFirestoreClient()
	
	// TODO: 変更
	err = _system.CheckLiveStreamStatus(ctx)
	if err != nil {
		_ = _system.LineBot.SendMessageWithError("failed to check live stream", err)
		return lambdautils.ErrorResponse(err)
	}
	
	return CheckBotAccessResponse()
}

func CheckBotAccessResponse() (events.APIGatewayProxyResponse, error) {
	var apiResp CheckBotAccessResponseStruct
	apiResp.Result = lambdautils.OK
	jsonBytes, _ := json.Marshal(apiResp)
	return lambdautils.Response(jsonBytes)
}

func main() {
	lambda.Start(CheckBotAccess)
}
