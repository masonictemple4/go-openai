package main

// Supporpted models is a map of model keys with the value of eligible endpoints they can hit.
var SupportedModels = map[string][]string{}

// TODO: Should probably come up with a general conversion model so we
// don't have to redo the implementation when openai changes the response
// format.
// Should also probably do this for the input.
var ChatCompletionMessageRoleMap = map[string]string{
	"system":    "system",
	"user":      "user",
	"assistant": "assistant",
}

type ChatCompletionRequestMessage struct {
	// The role of the message.
	// Examples: system, user, assistant
	Role string `json:"role,omitempty"`
	// The actual prompt. This is equivalent of the prompt would be for the
	// regular chat completion.
	Content string `json:"content,omitempty"`
}

type ChatCompletionRequestBody struct {
	// ID of the model to use. You can use the List models API
	// to see all of your available models, or see our Model overview
	// for descriptions of them.
	Model    string                         `json:"model,omitempty"`
	Messages []ChatCompletionRequestMessage `json:"messages,omitempty"`
	// What sampling temperature to use, between 0 and 2.
	// Higher values like 0.8 will make the output more random,
	// while lower values like 0.2 will make it more focused and deterministic.
	// We generally recommend altering this or top_p but not both.
	Temperature float64 `json:"temperature,omitempty"`
	// An alternative to sampling with temperature.
	// An alternative to sampling with temperature, called nucleus sampling,
	// where the model considers the results of the tokens with top_p probability mass.
	// So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	// Generally recommend altering this or temperature but not both.
	TopP float64 `json:"top_p,omitempty"`
	// How many completions to generate with each prompt
	N int64 `json:"n,omitempty"`
	// Whether or not to stream back partial progress. If set, tokens will be sent
	// ass data-only server-sent events as they become available. with the Stream
	// terminated by a data: [Done] message.
	Stream bool `json:"stream,omitempty"`
	// The maximum number of tokens to generate in the completion.
	// The token count of your prompt plus max_tokens cannot exceed
	// the model's context length. Most models have a context length
	// of 2048 tokens (except for the newest models, which support 4096).
	MaxTokens int64 `json:"max_tokens,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear
	// in the text so far, increasing the model's likelihood to talk about new topics.
	PresencePenalty float64 `json:"presence_penalty,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing Frequency
	// in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`
	// Modify the likelihood of specified tokens appearing in the completion.
	// Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer)
	// to an associated bias value from -100 to 100. You can use this tokenizer tool
	// (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically,
	// the bias is added to the logits generated by the model prior to sampling.
	// The exact effect will vary per model, but values between -1 and 1 should decrease or increase
	// likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection
	// of the relevant token.
	// As an example, you can pass {"50256": -100} to prevent the <|endoftext|> token from being generated.
	LogitBias map[string]int64 `json:"logit_bias,omitempty"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.
	User string `json:"user,omitempty"`
}

type ChatCompletionResponse struct {
	// Id of the completion
	Id string `json:"id"`
	// Type of object
	Object string `json:"object"`
	// Unix timestamp
	Created int64 `json:"created"`
	// Model used
	Model string `json:"model"`
	// Responses
	Choices []ChatCompletionChoice `json:"choices"`
	// Usage report
	Usage map[string]int `json:"usage"`
}

type ChatCompletionChoice struct {
	// The text of the response
	Role string `json:"role"`
	// Index?
	Index int64 `json:"index"`
	// The message response in the choice
	Message ChatCompletionRequestMessage
	// Why did it finish i.e length
	FinishReason string `json:"finish_reason"`
}

// Comments and more documentation:
// https://platform.openai.com/docs/api-reference/completions/create
type CompletionRequestBody struct {
	// ID of the model to use. You can use the List models API
	// to see all of your available models, or see our Model overview
	// for descriptions of them.
	Model string `json:"model,omitempty"`
	// The prompt(s) to generate completions for, encoded as a string,
	// array of strings, array of tokens, or array of token arrays.
	Prompt string `json:"prompt,omitempty"`
	// The suffix that comes after a completion of inserted text.
	Suffix string `json:"suffix,omitempty"`
	// The maximum number of tokens to generate in the completion.
	// The token count of your prompt plus max_tokens cannot exceed
	// the model's context length. Most models have a context length
	// of 2048 tokens (except for the newest models, which support 4096).
	MaxTokens int64 `json:"max_tokens,omitempty"`
	// What sampling temperature to use, between 0 and 2.
	// Higher values like 0.8 will make the output more random,
	// while lower values like 0.2 will make it more focused and deterministic.
	// We generally recommend altering this or top_p but not both.
	Temperature float64 `json:"temperature,omitempty"`
	// An alternative to sampling with temperature.
	// An alternative to sampling with temperature, called nucleus sampling,
	// where the model considers the results of the tokens with top_p probability mass.
	// So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	// Generally recommend altering this or temperature but not both.
	TopP float64 `json:"top_p,omitempty"`
	// How many completions to generate with each prompt
	N int64 `json:"n,omitempty"`
	// Whether or not to stream back partial progress. If set, tokens will be sent
	// ass data-only server-sent events as they become available. with the Stream
	// terminated by a data: [Done] message.
	Stream bool `json:"stream,omitempty"`
	// Include the log probabilites on the logprobs most likely tokens, as well the chosen tokens.
	// For example, if logprobs is 5, the api will return a list of the 5 most likely tokens.
	Logprobs int64 `json:"logprobs,omitempty"`
	// Echo back the prompt in addititon to the completion.
	Echo bool `json:"echo,omitempty"`
	// Up to 4 sequences where the API will stop generating further tokens. The returned text
	// will not contain the stop sequence.
	Stop string `json:"stop,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear
	// in the text so far, increasing the model's likelihood to talk about new topics.
	PresencePenalty float64 `json:"presence_penalty,omitempty"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing Frequency
	// in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`
	// Generates best_of completions serverside and returns the "best"
	// (the one with the highest log prob per token). Results cannot be streamed.
	// When used with n, best_of controls the number of candidate completions and n
	// specifies how many to return – best_of must be greater than n.
	// Note: Because this parameter generates many completions, it can quickly consume
	// your token quota. Use carefully and ensure that you have reasonable settings for
	// max_tokens and stop.
	BestOf int64 `json:"best_of,omitempty"`
	// Modify the likelihood of specified tokens appearing in the completion.
	// Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer)
	// to an associated bias value from -100 to 100. You can use this tokenizer tool
	// (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically,
	// the bias is added to the logits generated by the model prior to sampling.
	// The exact effect will vary per model, but values between -1 and 1 should decrease or increase
	// likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection
	// of the relevant token.
	// As an example, you can pass {"50256": -100} to prevent the <|endoftext|> token from being generated.
	LogitBias map[string]int64 `json:"logit_bias,omitempty"`
}

type CompletionResponse struct {
	// Id of the completion
	Id string `json:"id"`
	// Type of object
	Object string `json:"object"`
	// Unix timestamp
	Created int64 `json:"created"`
	// Model used
	Model string `json:"model"`
	// Responses
	Choices []CompletionChoice `json:"choices"`
	// Usage report
	Usage map[string]int `json:"usage"`
}

type CompletionChoice struct {
	// The text of the response
	Text string `json:"text"`
	// Index?
	Index int64 `json:"index"`
	// Log props
	Logprobs int64 `json:"logprobs"`
	// Why did it finish i.e length
	FinishReason string `json:"finish_reason"`
}
