package ai

import (
	
	"log"
	"os"
	"testing"
	"unicode/utf8"

	"github.com/joho/godotenv"
)

func loadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func TestGpt4Connection(t *testing.T) {
	loadEnv("../.env")
	openAiApiKey := os.Getenv("openAiApiKey")
	openAiApiOrg := os.Getenv("openAiApiOrg")
	envtest := os.Getenv("envtest")
	log.Printf("apiKey=%s, openAiOrg=%s, envTest (should be nil)= %s", openAiApiKey, openAiApiOrg, envtest)
	if utf8.RuneCountInString(envtest) != 0 {
		t.Fatalf("This is  a really unnecessary test. %d", utf8.RuneCountInString(envtest))
	}

    gptConn := &gpt4Connection{openAiApiKey,openAiApiOrg}

    resp := gptConn.getResponse("Hello!!")

    log.Printf("Gpt's response: %s", resp)
}
