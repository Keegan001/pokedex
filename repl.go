package main

import (
	"strings"
)

func cleanInput(text string) []string{
	res := strings.Fields(text)
	return res
}