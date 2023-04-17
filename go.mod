module openai-go-demo

go 1.18

replace github.com/Azure/azure-sdk-for-go/sdk/cognitiveservices/azopenai => /Users/mikekistler/Projects/Azure/azure-sdk-for-go/sdk/cognitiveservices/azopenai

require github.com/Azure/azure-sdk-for-go/sdk/cognitiveservices/azopenai v0.1.0

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.4.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.1.2 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/text v0.7.0 // indirect
)
