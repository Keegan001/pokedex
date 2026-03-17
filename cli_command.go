package main

// type defination
type cliCommand struct {
	name        string
	description string
	callback    func() error
}