package main

import "strings"

func cleanInput(text string) []string {
	cleaned := strings.TrimSpace(strings.ToLower(text))
	if cleaned == "" {
		return []string{}
	}
	return strings.Fields(cleaned)
}
