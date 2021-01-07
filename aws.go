package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func getSwagger(c *config) error {
	// create session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
		Credentials: credentials.NewStaticCredentials(c.aws.key, c.aws.secret, ""),
	})
	if err != nil {
		log.Fatalf("Error authenticating with AWS: %v", err)
	}

	svc := apigateway.New(sess)

	// get export
	res, err := svc.GetExport(&apigateway.GetExportInput{
		Accepts:    aws.String("application/json"),
		ExportType: aws.String("oas30"),
		RestApiId:  &c.aws.api,
		StageName:  &c.aws.stage,
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	ioutil.WriteFile("swagger.json", res.Body, os.FileMode(0644))

	return nil
}
