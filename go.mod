module openai-go-demo

go 1.18

replace github.com/Azure/azure-sdk-for-go/sdk/cognitiveservices/azopenai => /Users/mikekistler/Projects/Azure/azure-sdk-for-go/sdk/cognitiveservices/azopenai

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.6.0
	github.com/Azure/azure-sdk-for-go/sdk/cognitiveservices/azopenai v0.1.0
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.3.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)
