package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	//"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/cognitiveservices/azopenai"
)

var (
	options        *azopenai.ClientOptions
	endpoint       = os.Getenv("AOAI_ENDPOINT")
	apiKey         = os.Getenv("AOAI_API_KEY")
	openaiEndpoint = "https://api.openai.com"
	openaiApiKey   = os.Getenv("OPENAI_API_KEY")
)

func init() {
	options = &azopenai.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Logging: policy.LogOptions{
				IncludeBody: true,
			},
		},
	}
}

func chatbot() {
	// Generate Chatbot Response
	deploymentID := "text-davinci-003"
	cred := azopenai.KeyCredential{APIKey: apiKey}
	client, err := azopenai.NewClientWithKeyCredential(endpoint, cred, deploymentID, options)
	if err != nil {
		log.Fatalf("%v", err)
	}
	prompt := "What is Azure OpenAI?"
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionsOptions{
		Prompt:      []*string{to.Ptr(prompt)},
		MaxTokens:   to.Ptr(int32(2048 - 127)),
		Temperature: to.Ptr(float32(0.0)),
	}
	response, err := client.GetCompletions(context.TODO(), request, nil)
	if err != nil {
		return
	}
	completion := response.Choices[0].Text
	fmt.Printf("Chatbot: %s\n", *completion)
}

func openaiChatbot() {
	// Generate Chatbot Response
	cred := azopenai.KeyCredential{APIKey: openaiApiKey}
	//openaiEndpoint = "http://localhost:1234"
	client, err := azopenai.NewClientForOpenAI(openaiEndpoint, cred, options)
	if err != nil {
		log.Fatalf("%v", err)
	}
	prompt := "What is Azure OpenAI?"
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionsOptions{
		Model:       to.Ptr("text-davinci-002"),
		Prompt:      []*string{to.Ptr(prompt)},
		MaxTokens:   to.Ptr(int32(1024)),
		Temperature: to.Ptr(float32(0.0)),
	}
	response, err := client.GetCompletions(context.TODO(), request, nil)
	if err != nil {
		log.Fatalf("GetCompletions failed: %v", err)
		return
	}
	completion := response.Choices[0].Text
	fmt.Printf("Chatbot: %s\n", *completion)
}

// Generate Chatbot Response
func crossoverChatbot() {
	// Azure OpenAI client

	deploymentID := "text-davinci-003"
	cred := azopenai.KeyCredential{APIKey: apiKey}
	client, err := azopenai.NewClientWithKeyCredential(endpoint, cred, deploymentID, options)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// OpenAI Completions request

	prompt := "What is Azure OpenAI?"
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionsOptions{
		Model:       to.Ptr("text-davinci-002"),
		Prompt:      []*string{to.Ptr(prompt)},
		MaxTokens:   to.Ptr(int32(1024)),
		Temperature: to.Ptr(float32(0.0)),
	}
	response, err := client.GetCompletions(context.TODO(), request, nil)
	if err != nil {
		log.Fatalf("GetCompletions failed: %v", err)
		return
	}
	completion := response.Choices[0].Text
	fmt.Printf("Chatbot: %s\n", *completion)
}

func summarize() {
	// Summarize Text with Completion
	// Generate Chatbot Response
	deploymentID := "text-davinci-003"
	cred := azopenai.KeyCredential{APIKey: apiKey}
	client, err := azopenai.NewClientWithKeyCredential(endpoint, cred, deploymentID, options)
	if err != nil {
		log.Fatalf("%v", err)
	}
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
	response, err := client.GetCompletions(context.TODO(), request, nil)
	if err != nil {
		return
	}
	completion := response.Choices[0].Text
	fmt.Printf("Summarization: %s\n\n", *completion)
}

func chatbot_sse() error {
	// Generate Chatbot Response
	// Generate Chatbot Response
	deploymentID := "text-davinci-003"
	cred := azopenai.KeyCredential{APIKey: apiKey}
	client, err := azopenai.NewClientWithKeyCredential(endpoint, cred, deploymentID, options)
	if err != nil {
		log.Fatalf("%v", err)
	}
	prompt := "What is Azure OpenAI?"
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionsOptions{
		Prompt:      []*string{to.Ptr(prompt)},
		MaxTokens:   to.Ptr(int32(2048 - 127)),
		Temperature: to.Ptr(float32(0.0)),
	}
	response, err := client.GetCompletionsStream(context.TODO(), request, nil)
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
	// Generate Chatbot Response
	deploymentID := "text-davinci-003"
	cred := azopenai.KeyCredential{APIKey: apiKey}
	client, err := azopenai.NewClientWithKeyCredential(endpoint, cred, deploymentID, options)
	if err != nil {
		log.Fatalf("%v", err)
	}
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
	response, err := client.GetCompletionsStream(context.TODO(), request, nil)
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

func streaming() error {
	// Generate Chatbot Response
	deploymentID := "text-davinci-003"
	cred := azopenai.KeyCredential{APIKey: apiKey}
	client, err := azopenai.NewClientWithKeyCredential(endpoint, cred, deploymentID, options)
	if err != nil {
		log.Fatalf("%v", err)
	}
	prompt := "These are the numbers from 1 to 100: 1,2,3,"
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionsOptions{
		Prompt:      []*string{to.Ptr(prompt)},
		MaxTokens:   to.Ptr(int32(500)),
		Temperature: to.Ptr(float32(0.0)),
	}
	startTime := time.Now()
	response, err := client.GetCompletionsStream(context.TODO(), request, nil)
	if err != nil {
		return err
	}
	r := response.Events
	defer r.Close()

	var eventList []azopenai.Completions = make([]azopenai.Completions, 0, 300)
	var timestampList []time.Time = make([]time.Time, 0, 300)
	for {
		fb, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return err
		}
		eventList = append(eventList, fb)
		timestampList = append(timestampList, time.Now())
		//fmt.Printf("%s", *fb.Choices[0].Text)
	}
	for i := 0; i < len(eventList); i++ {
		latency := timestampList[i].Sub(startTime)
		fmt.Printf("%6.3f: %s\n", latency.Seconds(), *eventList[i].Choices[0].Text)
	}
	return nil
}

func openai_streaming() error {
	// Create client for OpenAI
	cred := azopenai.KeyCredential{APIKey: openaiApiKey}
	//openaiEndpoint = "http://localhost:1234"
	client, err := azopenai.NewClientForOpenAI(openaiEndpoint, cred, options)
	if err != nil {
		log.Fatalf("%v", err)
	}

	prompt := "These are the numbers from 1 to 100: 1,2,3,"
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionsOptions{
		Model:       to.Ptr("text-davinci-002"),
		Prompt:      []*string{to.Ptr(prompt)},
		MaxTokens:   to.Ptr(int32(500)),
		Temperature: to.Ptr(float32(0.0)),
	}
	startTime := time.Now()
	response, err := client.GetCompletionsStream(context.TODO(), request, nil)
	if err != nil {
		return err
	}
	r := response.Events
	defer r.Close()

	var eventList []azopenai.Completions = make([]azopenai.Completions, 0, 300)
	var timestampList []time.Time = make([]time.Time, 0, 300)
	for {
		fb, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return err
		}
		eventList = append(eventList, fb)
		timestampList = append(timestampList, time.Now())
		//fmt.Printf("%s", *fb.Choices[0].Text)
	}
	for i := 0; i < len(eventList); i++ {
		latency := timestampList[i].Sub(startTime)
		fmt.Printf("%6.3f: %s\n", latency.Seconds(), *eventList[i].Choices[0].Text)
	}
	return nil
}

func main() {

	// chatbot()

	//openaiChatbot()

	//crossoverChatbot()

	// summarize()

	//chatbot_sse()

	// summarize_sse()

	// streaming()

	openai_streaming()
}
