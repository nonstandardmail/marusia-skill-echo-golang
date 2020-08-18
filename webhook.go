package main

import (
	"fmt"
	"strings"

	"github.com/antonholmquist/jason"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func createMarusiaResponse(
		phrase string,
		userId string,
		sessionId string,
		messageId int64,
		version string,
	) string {
	escapedPhrase := strings.Replace(phrase, "\"", "\\\"", -1)
	return fmt.Sprintf(`
	{
			"response": {
				"text": "%s",
				"tts": "%s",
				"end_session": false
			},
			"session": {
				"user_id": "%s",
				"session_id": "%s",
				"message_id": %d
			},
			"version": "%s"
	}`, escapedPhrase, escapedPhrase, userId, sessionId, messageId, version)
}

func WebhookHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	v, _ := jason.NewObjectFromBytes([]byte(request.Body))
	phrase, _ := v.GetString("request", "original_utterance")
	version, _ := v.GetString("version")
	sessionId, _ := v.GetString("session", "session_id")
	userId, _ := v.GetString("session", "user_id")
	messageId, _ := v.GetInt64("session", "message_id")
	responseMessage := createMarusiaResponse(phrase, userId, sessionId, messageId, version)
	return events.APIGatewayProxyResponse{Body: responseMessage, StatusCode: 200}, nil
}

func main() {
	lambda.Start(WebhookHandler)
}
