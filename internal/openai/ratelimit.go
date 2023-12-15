package openai

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type RateLimitHeaders struct {
	RequestsLimit     int64 `json:"x-ratelimit-limit-requests"`
	TokenLimit        int64 `json:"x-ratelimit-limit-tokens"`
	RemainingRequests int64 `json:"x-ratelimit-remaining-requests"`
	RemainingTokens   int64 `json:"x-ratelimit-remaining-tokens"`
	// Time until the requests rate limit resets.
	// Ex: 1s
	ResetRequests string `json:"x-ratelimit-reset-requests"`
	// Time until the token rate limit resets.
	// Ex: 6m0s
	ResetTokens string `json:"x-ratelimit-reset-tokens"`
}

func GetRateLimitHeaders(resp *http.Response) *RateLimitHeaders {
	reqLimit, _ := strconv.Atoi(resp.Header.Get("x-ratelimit-limit-requests"))
	tokenLimit, _ := strconv.Atoi(resp.Header.Get("x-ratelimit-limit-tokens"))
	remainingRequests, _ := strconv.Atoi(resp.Header.Get("x-ratelimit-remaining-requests"))
	remainingTokens, _ := strconv.Atoi(resp.Header.Get("x-ratelimit-remaining-tokens"))

	return &RateLimitHeaders{
		RequestsLimit:     int64(reqLimit),
		TokenLimit:        int64(tokenLimit),
		RemainingRequests: int64(remainingRequests),
		RemainingTokens:   int64(remainingTokens),
		ResetRequests:     resp.Header.Get("x-ratelimit-reset-requests"),
		ResetTokens:       resp.Header.Get("x-ratelimit-reset-tokens"),
	}
}

func (r *RateLimitHeaders) TokenResetTime() *time.Time {
	dur, err := time.ParseDuration(r.ResetTokens)
	if err != nil {
		log.Printf("Error parsing token reset time: %s", err)
		return nil
	}
	newTime := time.Now().Add(dur)
	return &newTime
}

func (r *RateLimitHeaders) RequestResetTime() *time.Time {
	dur, err := time.ParseDuration(r.ResetRequests)
	if err != nil {
		log.Printf("Error parsing requests reset time: %s", err)
		return nil
	}
	newTime := time.Now().Add(dur)
	return &newTime
}
