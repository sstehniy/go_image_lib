package main

import (
	"log"
	"time"

	"github.com/sstehniy/gopix/cmd"
)

func main() {
	now := time.Now()
	log.Println("Starting the application...")
	cmd.Execute()
	log.Println("Application took", time.Since(now))
	log.Println("Application finished executing.")
}
