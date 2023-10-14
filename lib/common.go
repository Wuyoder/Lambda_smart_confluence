package lib

import (
	"fmt"
	"os"
)

var slackWebhook = os.Getenv("SLACK_WEBHOOK")
var slackChannel = os.Getenv("SLACK_CHANNEL")

func HandleErr(err error) {
	if err != nil {
		errorMsg := fmt.Sprintf("Error : %s", err.Error())
		fmt.Println(errorMsg)
		PostToSlack(slackWebhook, slackChannel, errorMsg)
		panic(err)
	}
}
