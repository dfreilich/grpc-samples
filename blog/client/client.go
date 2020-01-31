package main

import "log"

func main() {
	log.Println("Starting blog client")
	if err := run(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
	log.Println("Succesfully ran")
}

func run() error {
	return nil
}
