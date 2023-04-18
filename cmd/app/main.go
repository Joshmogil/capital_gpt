package main

import  (
	"flag"
	"log"
	"os"
	"capital_gpt/models"
)

type config struct {
	openai_org string
	openai_token string	
}
 
type environment struct {
	config config
	resources models.Resources
	motivations models.Motivations
	logger *log.Logger
}

func main() {
	var cfg config
	var motivations_location string 


	flag.StringVar(&cfg.openai_org, "org", "None", "organization for openai")
	flag.StringVar(&cfg.openai_token, "token", "None", "token for openai")
	flag.StringVar(&motivations_location, "motofile", "./inputs/motivations.json", "list of motivations in json file")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	env := &environment{
	config: cfg,
	logger: logger,
	}

	err := env.loadMotivations(motivations_location)
	if err != nil{
		logger.Printf("failed to load motivations from %s with error: %s", motivations_location, err)
	}

	}