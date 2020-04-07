package main

import (
	"flag"
	"log"
)

func main() {
	var config string
	flag.StringVar(&config, "c", "config.txt", "config file with rules")
	flag.Parse()

	routes, err := readConfig(config)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	if len(routes) == 0 {
		log.Fatalln("No routes, refusing to start")
	}
	log.Printf("Total Routes: %d\n", len(routes))
	for _, r := range routes {
		go r.listen()
	}
	<-chan bool(nil)
}
