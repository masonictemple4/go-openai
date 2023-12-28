package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type ImageGenerationResponseDataItem struct {
	// The url of the image
	Url string `json:"url,omitempty"`
}

type ImageGenerationResponseBody struct {
	// When was this response created.
	Created int64 `json:"created,omitempty"`
	// The response data.
	Data []ImageGenerationResponseDataItem `json:"data,omitempty"`
}

type ImageGenerationRequestBody struct {
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string `json:"prompt,omitempty"`
	// The number of images to generate. Must be between 1 and 10.
	N int64 `json:"n,omitempty"`
	// The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.
	Size string `json:"size,omitempty"`
	// The format the generated images should return in. Must be one of url or b64_json
	ResponseFormat string `json:"response_format,omitempty"`
	// The unique identifier representing your end-user, which can help OpenAI to monitor
	// and detect abuse.
	User string `json:"user,omitempty"`
}

func (o *OpenAIClient) RequestImageGeneration(body ImageGenerationRequestBody) (*ImageGenerationResponseBody, error) {
	client := &http.Client{}

	byts, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	data, _ := json.MarshalIndent(body, "", "  ")
	fmt.Printf("\n%s\n", data)

	reqBody := bytes.NewBuffer(byts)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", reqBody)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if o.Organization != "" {
		req.Header.Add("OpenAI-Organization", o.Organization)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+o.APIKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("\nThe response error is: %+v", res)
		return nil, errors.New("Invalid request.")
	}

	var resp ImageGenerationResponseBody
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil

}

// TODO: Needs to be an interface function for the Response models.
func (r *ImageGenerationResponseBody) CleanText() string {
	var urls []string
	for i := range r.Data {
		urls = append(urls, r.Data[i].Url)
	}
	text := strings.Join(urls, "\n")
	text = strings.TrimLeft(text, "?\n")
	return text
}
