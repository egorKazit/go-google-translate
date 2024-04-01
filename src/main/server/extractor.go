package server

import (
	"fmt"
	"log"
	"regexp"
)

func extractKey(key string, responseBody string) string {
	// create regexp to find expiration and handle error if any
	regex, err := regexp.Compile(fmt.Sprintf("\"%s\":\".*?\"", key))
	if err != nil {
		log.Printf("Error at http client request: %s", err.Error())
		return ""
	}

	// find all findings and handle empty if so
	allFindings := regex.FindAllString(responseBody, 1)
	if len(allFindings) == 0 {
		return ""
	}

	// handle only first occurrence
	result := allFindings[0]
	return result[len(fmt.Sprintf("\"%s\":\"", key)) : len(result)-1]
}
