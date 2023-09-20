package main

import (
	"context"
	"fmt"
	"smart_confluence_label/confluence"
	"smart_confluence_label/constant"
	"smart_confluence_label/openai"

	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func main() {
	lambda.Start(Handler)
}

// GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap -tags lambda.norpc main.go
// zip myFunction.zip bootstrap

type PageID struct {
	PageID string `json:"page_id"`
}

type Message struct {
	Message string `json:"message"`
}

func Handler(event PageID) (Message, error) {
	timeStart := time.Now()
	ctx := context.Background()

	var output Message
	id := event.PageID
	if id == "" {
		fmt.Println("page_id is empty")
		output.Message = "page_id is empty"
		return output, nil
	}

	isChatMode := true // true for chat, false for completion
	isGPT4 := false    // true for GPT-4, false for GPT-3 // use GPT-3 for more token per minute

	getCompletionObj := openai.CompletionObj{
		Payload:    confluence.GetPageContentByID(id).Body.View.Value,
		Query:      constant.Query,
		IsChatMode: isChatMode,
		IsGPT4:     isGPT4,
	}

	result, _ := getCompletionObj.GetOpenAIResp(ctx)
	output.Message = "Labels : [ " + strings.Join(result, ",") + " ]"

	var tag string
	for _, v := range result { // set options = 1
		fmt.Println(v)
		tag = v
	}

	confluence.UpdateTagsToPage(id, tag)

	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(timeStart))

	return output, nil
}

func init() {
	godotenv.Load()
}
