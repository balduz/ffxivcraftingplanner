package main

import (
	"log"
	"os"

	"ffxivcraftingplanner/server"
)

func setUpLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func main() {
	setUpLogger()
	port := getPort()
	log.Fatal(server.StartServer(port))
}
