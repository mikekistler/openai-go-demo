package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
	client, err = azopenai.NewClientWithKeyCredential(endpoint, cred, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	// turn on logging of the full request and response

}

func chatbot() {
	// Generate Chatbot Response

	prompt := "What is Azure OpenAI?"
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionRequest{
		Prompt: prompt,
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
		"Summarize the following text.\n",
		"\n",
		"Text:\n",
		"\"\"\"\n",
		"Two independent experiments reported their results this morning at CERN,\n",
		"Europe's high-energy physics laboratory near Geneva in Switzerland. Both show\n",
		"convincing evidence of a new boson particle weighing around 125 gigaelectronvolts,\n",
		"which so far fits predictions of the Higgs previously made by theoretical physicists.\n",
		"\n",
		"\"As a layman I would say: 'I think we have it'. Would you agree?\" Rolf-Dieter Heuer,\n",
		"CERN's director-general, asked the packed auditorium. The physicists assembled there\n",
		"burst into applause.\n",
		"\"\"\"\n",
		"\n",
		"Summary: ",
	}
	fmt.Printf("Input: %s\n", prompt)
	request := azopenai.CompletionRequest{
		Prompt: strings.Join(prompt, "\n"),
	}
	response, err := client.GetCompletions(context.TODO(), deploymentId, request, nil)
	if err != nil {
		return
	}
	completion := response.Choices[0].Text
	fmt.Printf("Summarization: %s\n", *completion)
}

func main() {

	//chatbot()

	summarize()
}
