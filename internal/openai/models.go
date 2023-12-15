package openai

const (
	MODEL_GPT4                = "gpt-4"
	MODEL_GPT4_1106_Preview   = "gpt-4-1106-preview"
	MODEL_GPT4_Vision_Preview = "gpt-4-vision-preview"
)

type OpenAIModel struct {
	// Model id which is referenced by the endpoints.
	Id string `json:"id"`
	// Unix timestamp (seconds) when the model was created.
	Created int64 `json:"created"`
	// Always going to be the 'model' object type.
	Object string `json:"object"`
	// Org that owns the model
	OwnedBy string `json:"owned_by"`
}

type OpenAIModelListResponse struct {
	// Will most likely be list.
	Object string        `json:"object"`
	Data   []OpenAIModel `json:"data"`
}

type OpenAIModelInfo struct {
	Model          string   `json:"model"`
	Desc           string   `json:"description"`
	ContextWindow  string   `json:"context_window"`
	TrainingData   string   `json:"training_data"`
	ValidEndpoints []string `json:"valid_endpoints"`
}

type ModelDataStore []OpenAIModelInfo

// TODO: Convert this to Const like i did with
// roles.
var GPT4ModelDataStore = ModelDataStore{
	{
		Model:         "gpt-4-1106-preview",
		Desc:          "GPT-4 TurboNew. The latest GPT-4 model with improved instruction following, JSON mode, reproducible outputs, parallel function calling, and more. Returns a maximum of 4,096 output tokens. This preview model is not yet suited for production traffic.",
		ContextWindow: "128,000 tokens",
		TrainingData:  "Up to Apr 2023",
	},
	{
		Model:         "gpt-4-vision-preview",
		Desc:          "GPT-4 Turbo with visionNew. Ability to understand images, in addition to all other GPT-4 Turbo capabilities. Returns a maximum of 4,096 output tokens. This is a preview model version and not suited yet for production traffic.",
		ContextWindow: "128,000 tokens",
		TrainingData:  "Up to Apr 2023",
	},
	{
		Model:         "gpt-4",
		Desc:          "Currently points to gpt-4-0613. See continuous model upgrades.",
		ContextWindow: "8,192 tokens",
		TrainingData:  "Up to Sep 2021",
	},
	{
		Model:         "gpt-4-32k",
		Desc:          "Currently points to gpt-4-32k-0613. See continuous model upgrades.",
		ContextWindow: "32,768 tokens",
		TrainingData:  "Up to Sep 2021",
	},
	{
		Model:         "gpt-4-0613",
		Desc:          "Snapshot of gpt-4 from June 13th 2023 with improved function calling support.",
		ContextWindow: "8,192 tokens",
		TrainingData:  "Up to Sep 2021",
	},
	{
		Model:         "gpt-4-32k-0613",
		Desc:          "Snapshot of gpt-4-32k from June 13th 2023 with improved function calling support.",
		ContextWindow: "32,768 tokens",
		TrainingData:  "Up to Sep 2021",
	},
	{
		Model:         "gpt-4-0314",
		Desc:          "Legacy. Snapshot of gpt-4 from March 14th 2023 with function calling support. This model version will be deprecated on June 13th 2024.",
		ContextWindow: "8,192 tokens",
		TrainingData:  "Up to Sep 2021",
	},
	{
		Model:         "gpt-4-32k-0314",
		Desc:          "Legacy. Snapshot of gpt-4-32k from March 14th 2023 with function calling support. This model version will be deprecated on June 13th 2024.",
		ContextWindow: "32,768 tokens",
		TrainingData:  "Up to Sep 2021",
	},
}
