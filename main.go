package main

import (
	"prism/cmd"
	"prism/internal/service"
)

func main() {
	service := service.NewService()
	cmd.Execute(service)
}
