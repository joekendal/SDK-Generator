package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func execute() {
	cmd := exec.Command("sh", "-c", "swagger-codegen -h")
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()

	go streamPipe(stderr)
	go streamPipe(stdout)

	cmd.Start()
	cmd.Wait()
}

func streamPipe(std io.ReadCloser) {
	scanner := bufio.NewScanner(std)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Printf(m)
	}
}

func main() {
	c := getConfig()
	if &c.aws != (&awsConfig{}) {
		fmt.Print("Grabbing Swagger from AWS... ")
		err := getSwagger(c)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Println("Done!")
		execute()
	}

}
