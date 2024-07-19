package main

import "fmt"

type DefaultPlugin struct{}

func (p *DefaultPlugin) Init() error {
	fmt.Printf("Init plugin")
	return nil
}

func (p *DefaultPlugin) Execute(args []string) error {
	fmt.Printf("Execution")
	fmt.Printf("Args : %+v", args)
	return nil
}

func main() {
	
}