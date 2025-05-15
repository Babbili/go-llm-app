# Golang LLM application

This repo demonstrates how to bring Large Language Models to your Go applications and add AI capabilities to Golang which has great scaling capability, impressive performance and concurrency support.

The application uses `gollm` library to create an LLM instance with Ollama as a provider, set up a prompt, and generate a response.


Ollama is running locally and serving Google's Gemma3 model.

```bash
ollama serve
ollama pull gemma3 && ollama run gemma3

```

> I choosed `Gemma3` which is lightweight model from Google, its a multimodal—processing text and images, Available in 1B, 4B, 12B, and 27B parameter sizes, it exceles in tasks like summarization, question answering, content creation, reasoning, translating and visual understanding. 


create an LLM with Ollama provider

`gollm` also streamlines interactions with other LLM providers such as OpenAI, Anthropic, Groq, OpenRouter, ...

```go
import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/teilomillet/gollm"
)

func main() {
	// Create a new LLM instance with Ollama provider
	llm, err := gollm.NewLLM(
		gollm.SetProvider("ollama"),
		gollm.SetModel("gemma3"),
		gollm.SetTimeout(100*time.Second),
		gollm.SetLogLevel(gollm.LogLevelInfo),
		gollm.SetOllamaEndpoint("http://127.0.0.1:11434"),
	)
	if err != nil {
		log.Fatalf("Failed to create LLM: %v", err)
	}
...
```

create a prompt with directives
```go
	// Create a prompt using NewPrompt function
	prompt := gollm.NewPrompt(
		"summarize this blog https://opentelemetry.io/blog/2025/otel-cicd-sig/",
		gollm.WithDirectives(
			"Keep the summary under 300 words",
			"Capture the main points",
			"Focus on the solution provided",
		),
	)
```

generate response
```go
	// Generate a response
	ctx := context.Background()
	response, err := llm.Generate(ctx, prompt)
	if err != nil {
		log.Fatalf("Failed to generate response: %v", err)
	}

	fmt.Printf("Response: %s\n", response)
```
---
**Response**: Okay, here’s a summary of the OpenTelemetry CICD SIG blog post (https://opentelemetry.io/blog/2025/otel-cicd-sig/) under 300 words, focusing on the solution provided:

The OpenTelemetry Community Interest Group (CICD SIG) is tackling a critical challenge: integrating observability seamlessly into Continuous Integration and Continuous Delivery (CI/CD) pipelines. Historically, instrumenting code for tracing within CI/CD has been complex, requiring significant manual effort and often leading to instrumentation being dropped or incomplete.

The CICD SIG’s work addresses this by establishing a standardized, automated approach to OpenTelemetry instrumentation in CI/CD workflows. They’ve developed a set of best practices and tools designed to streamline the process, dramatically reducing the friction for teams to adopt tracing within their deployments.

**Key Solutions & Points:**

*   **Standardized Instrumentation:** The SIG promotes the use of a specific set of OpenTelemetry SDKs and configurations optimized for CI/CD environments.
*   **Automated Instrumentations:** They provide scripts and tooling to automatically instrument key components within your CI/CD pipeline – such as build servers, test runners, and deployment stages.
*   **Pipeline Integration:** The solution focuses on integrating OpenTelemetry directly into the build and test stages, capturing traces throughout the entire pipeline.
*   **Reduced Overhead:** By automating instrumentation, the SIG minimizes the time and effort required to add tracing, encouraging broader adoption of observability practices.

Ultimately, the CICD SIG’s work makes it easier and faster for teams to gain insights into the performance and behavior of their applications during deployment, improving debugging and identifying issues quicker.




