package main

import (
	"context"
	"fmt"
	"log"

	"github.com/teilomillet/gollm"
)

func main() {
	// Create a new LLM instance with Ollama provider
	llm, err := gollm.NewLLM(
		gollm.SetProvider("ollama"),
		gollm.SetModel("gemma3"),
		gollm.SetLogLevel(gollm.LogLevelInfo),
	)
	if err != nil {
		log.Fatalf("Failed to create LLM: %v", err)
	}

	// Create a prompt using NewPrompt function
	prompt := gollm.NewPrompt("summarize this blog https://opentelemetry.io/blog/2025/otel-cicd-sig/")

	// Generate a response
	ctx := context.Background()
	response, err := llm.Generate(ctx, prompt)
	if err != nil {
		log.Fatalf("Failed to generate response: %v", err)
	}

	fmt.Printf("Response: %s\n", response)

}

