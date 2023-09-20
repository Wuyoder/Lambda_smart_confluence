package lib

import (
	"fmt"
)

var slackWebhook = "https://hooks.slack.com/services/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"
var slackChannel = "#slack_web_hook_test"

func HandleErr(err error) {
	if err != nil {
		errorMsg := fmt.Sprintf("Error : %s", err.Error())
		fmt.Println(errorMsg)
		PostToSlack(slackWebhook, slackChannel, errorMsg)
		panic(err)
	}
}
