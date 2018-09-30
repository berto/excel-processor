package main

import "log"

func main() {
	err := RunMigrations()
	if err != nil {
		log.Fatal(err)
	}
}
