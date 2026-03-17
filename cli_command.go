package main

// type defination
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
	conf *config
}

type config struct {
	nextUrl string
	prevUrl string
}