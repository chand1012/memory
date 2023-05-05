package chatsort

import (
	"fmt"
	"strings"

	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
)

func PreciseTokensFromMessages(messages []openai.ChatCompletionMessage, model string) (num_tokens int) {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf("EncodingForModel: %v", err)
		fmt.Println(err)
		return
	}

	var tokens_per_message int
	var tokens_per_name int
	if model == "gpt-3.5-turbo-0301" || model == "gpt-3.5-turbo" {
		tokens_per_message = 4
		tokens_per_name = -1
	} else if strings.Contains(model, "4") {
		tokens_per_message = 3
		tokens_per_name = 1
	} else {
		tokens_per_message = 3
		tokens_per_name = 1
	}

	for _, message := range messages {
		num_tokens += tokens_per_message
		num_tokens += len(tkm.Encode(message.Content, nil, nil))
		num_tokens += len(tkm.Encode(message.Role, nil, nil))
		if message.Name != "" {
			num_tokens += tokens_per_name
		}
	}
	num_tokens += 3
	return num_tokens
}

func PreciseTokensFromMessage(message openai.ChatCompletionMessage, model string) (num_tokens int) {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf("EncodingForModel: %v", err)
		fmt.Println(err)
		return
	}

	var tokens_per_message int
	var tokens_per_name int
	if model == "gpt-3.5-turbo-0301" || model == "gpt-3.5-turbo" {
		tokens_per_message = 4
		tokens_per_name = -1
	} else if strings.Contains(model, "4") {
		tokens_per_message = 3
		tokens_per_name = 1
	} else {
		tokens_per_message = 3
		tokens_per_name = 1
	}

	num_tokens += tokens_per_message
	num_tokens += len(tkm.Encode(message.Content, nil, nil))
	num_tokens += len(tkm.Encode(message.Role, nil, nil))
	if message.Name != "" {
		num_tokens += tokens_per_name
	}
	return num_tokens
}
