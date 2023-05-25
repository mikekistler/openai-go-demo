package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	//"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/cognitiveservices/azopenai"
)

var (
	client       *azopenai.Client
	endpoint     = os.Getenv("OPENAI_ENDPOINT")
	apiKey       = os.Getenv("OPENAI_APIKEY")
	deploymentId = os.Getenv("OPENAI_DEPLOYMENT_ID")
)

func init() {
	var err error
	cred := azopenai.KeyCredential{APIKey: apiKey}
	options := azopenai.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Logging: policy.LogOptions{
				IncludeBody: true,
			},
		},
	}
	client, err = azopenai.NewClientWithKeyCredential(endpoint, cred, &options)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func chatbot() {
	// Generate Chatbot Response

	prompt := "What is Azure OpenAI?"
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionsOptions{
		Prompt:      []*string{to.Ptr(prompt)},
		MaxTokens:   to.Ptr(int32(2048 - 127)),
		Temperature: to.Ptr(float32(0.0)),
	}
	response, err := client.GetCompletions(context.TODO(), deploymentId, request, nil)
	if err != nil {
		return
	}
	completion := response.Choices[0].Text
	fmt.Printf("Chatbot: %s\n", *completion)
}

func summarize() {
	// Summarize Text with Completion

	prompt := []string{
		"Summarize the following text.",
		"",
		"Text:",
		"\"\"\"",
		"Two independent experiments reported their results this morning at CERN,",
		"Europe's high-energy physics laboratory near Geneva in Switzerland. Both show",
		"convincing evidence of a new boson particle weighing around 125 gigaelectronvolts,",
		"which so far fits predictions of the Higgs previously made by theoretical physicists.",
		"",
		"\"As a layman I would say: 'I think we have it'. Would you agree?\" Rolf-Dieter Heuer,",
		"CERN's director-general, asked the packed auditorium. The physicists assembled there",
		"burst into applause.",
		"\"\"\"",
		"",
		"Summary: ",
	}
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionsOptions{
		Prompt:      []*string{to.Ptr(strings.Join(prompt, " "))},
		MaxTokens:   to.Ptr(int32(2048 - 127)),
		Temperature: to.Ptr(float32(0.0)),
	}
	response, err := client.GetCompletions(context.TODO(), deploymentId, request, nil)
	if err != nil {
		return
	}
	completion := response.Choices[0].Text
	fmt.Printf("Summarization: %s\n\n", *completion)
}

func chatbot_sse() error {
	// Generate Chatbot Response

	prompt := "What is Azure OpenAI?"
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionsOptions{
		Prompt:      []*string{to.Ptr(prompt)},
		MaxTokens:   to.Ptr(int32(2048 - 127)),
		Temperature: to.Ptr(float32(0.0)),
	}
	response, err := client.GetCompletionEvents(context.TODO(), deploymentId, request, nil)
	if err != nil {
		return err
	}
	reader := response.Events
	defer reader.Close()

	for {
		fb, err := reader.Read()
		if err == io.EOF {
			//fmt.Println("End of stream")
			break
		}
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return err
		}
		fmt.Printf("%s", *fb.Choices[0].Text)
	}
	return nil
}

func summarize_sse() error {
	// Summarize Text with Completion

	prompt := []string{
		"Summarize the following text.",
		"",
		"Text:",
		"\"\"\"",
		"Two independent experiments reported their results this morning at CERN, ",
		"Europe's high-energy physics laboratory near Geneva in Switzerland. Both show",
		"convincing evidence of a new boson particle weighing around 125 gigaelectronvolts,",
		"which so far fits predictions of the Higgs previously made by theoretical physicists.",
		"",
		"\"As a layman I would say: 'I think we have it'. Would you agree?\" Rolf-Dieter Heuer,",
		"CERN's director-general, asked the packed auditorium. The physicists assembled there",
		"burst into applause.",
		"\"\"\"",
		"",
		"Summary: ",
	}
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionsOptions{
		Prompt:      []*string{to.Ptr(strings.Join(prompt, " "))},
		MaxTokens:   to.Ptr(int32(2048 - 127)),
		Temperature: to.Ptr(float32(0.0)),
	}
	response, err := client.GetCompletionEvents(context.TODO(), deploymentId, request, nil)
	if err != nil {
		log.Fatal(err)
	}
	reader := response.Events
	defer reader.Close()

	for {
		fb, err := reader.Read()
		if err == io.EOF {
			//fmt.Println("End of stream")
			break
		}
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return err
		}
		fmt.Printf("%s", *fb.Choices[0].Text)
	}
	return nil
}

func main() {

	// chatbot()

	// summarize()

	chatbot_sse()

	// summarize_sse()
}
