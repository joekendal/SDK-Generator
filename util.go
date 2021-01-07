package main

import (
	"os"
)

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func all(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// check if environment variable is set
func isSet(v string) bool {
	_, set := os.LookupEnv("INPUT_" + v)
	return set
}

// check if environment variable is not set
func notSet(v string) bool {
	_, set := os.LookupEnv("INPUT_" + v)
	return !set
}

func getEnv(v string) string {
	return os.Getenv("INPUT_" + v)
}
