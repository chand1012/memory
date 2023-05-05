package chatsort

import (
	"strconv"

	"github.com/chand1012/memory"
	"github.com/sashabaranov/go-openai"
)

// helper for sorting OpenAI Chat Messages
func SortByRelevance(messages []openai.ChatCompletionMessage, query string) ([]openai.ChatCompletionMessage, error) {
	mem, _, err := memory.New(":memory:")
	if err != nil {
		return nil, err
	}
	defer mem.Close()

	for i, message := range messages {
		err = mem.Add(strconv.Itoa(i), message.Content)
		if err != nil {
			return nil, err
		}
	}

	results, err := mem.Search(query)
	if err != nil {
		return nil, err
	}

	sorted := memory.SortByAverage(results)

	var sortedMessages []openai.ChatCompletionMessage
	for _, item := range sorted {
		index, err := strconv.Atoi(item.ID)
		if err != nil {
			return nil, err
		}

		sortedMessages = append(sortedMessages, messages[index])
	}

	return sortedMessages, nil
}

// gets the most relevant messages from a list of messages using the query.
// will not return more message than the maxTokens allows.
func GetRelevant(message []openai.ChatCompletionMessage, query, model string, maxTokens int) ([]openai.ChatCompletionMessage, int, error) {
	if maxTokens == 0 {
		return message, 0, nil
	}

	sorted, err := SortByRelevance(message, query)
	if err != nil {
		return nil, 0, err
	}

	var relevant []openai.ChatCompletionMessage
	var numTokens int
	for _, message := range sorted {
		messageTokens := PreciseTokensFromMessage(message, model)
		if numTokens+messageTokens > maxTokens {
			break
		}
		relevant = append(relevant, message)
		numTokens += messageTokens
	}

	return relevant, numTokens, nil
}
