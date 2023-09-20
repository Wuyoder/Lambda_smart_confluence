package lib

import (
	"fmt"
)

var slackWebhook = ""
var slackChannel = "#slack_web_hook_test"

func HandleErr(err error) {
	if err != nil {
		errorMsg := fmt.Sprintf("Error : %s", err.Error())
		fmt.Println(errorMsg)
		PostToSlack(slackWebhook, slackChannel, errorMsg)
		panic(err)
	}
}
