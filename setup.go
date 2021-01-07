package main

import (
	"log"
	"strings"
)

var required = []string{"ACCESS_TOKEN", "LANGUAGES"}

var optionalAws = []string{
	"AWS_KEY", "AWS_KEY_SECRET", "AWS_API_ID", "AWS_API_STAGE",
}

type awsConfig struct {
	key    string
	secret string
	api    string
	stage  string
}

type config struct {
	ghToken   string
	languages []string
	aws       awsConfig
}

func getConfig() *config {
	// check required envs
	if !all(required, isSet) {
		missing := filter(required, notSet)
		log.Fatalf("Missing required variables: %v", missing)
	}
	config := &config{
		ghToken:   getEnv(required[0]),
		languages: strings.Split(getEnv(required[1]), ","),
	}
	// all or none of aws envs
	if !all(optionalAws, isSet) {
		if any(optionalAws, isSet) {
			missing := filter(optionalAws, notSet)
			log.Fatalf("Missing AWS variables: %v", missing)
		}
	} else {
		config.aws = awsConfig{
			key:    getEnv(optionalAws[0]),
			secret: getEnv(optionalAws[1]),
			api:    getEnv(optionalAws[2]),
			stage:  getEnv(optionalAws[3]),
		}
	}
	return config
}
