package main

import (
	"log"

	"github.com/sstehniy/gopix/cmd"
)

func main() {
	log.Println("Starting the application...")
	cmd.Execute()
	log.Println("Application finished executing.")
}
