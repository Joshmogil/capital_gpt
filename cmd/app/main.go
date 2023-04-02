package main

import  (
	"flag"
	"log"
	"os"
)

type config struct {
	openai_org string
	openai_token string
	motivations map[string]string	
}
 
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	var motivations_location string 


	flag.StringVar(&cfg.openai_org, "org", "None", "organization for openai")
	flag.StringVar(&cfg.openai_token, "token", "None", "token for openai")
	flag.StringVar(&motivations_location, "motofile", "./motivations.json", "list of motivations in json file")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
	config: cfg,
	logger: logger,
	}

	err := app.loadMotivations(motivations_location)
	if err != nil{
		logger.Printf("failed to load motivations from %s with error: %s", motivations_location, err)
	}

	}