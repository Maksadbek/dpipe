package main

import (
	"flag"
	"log"

	"github.com/maksadbek/dpipe/config"
)

var (
	configFile = flag.String("config", "config", "config file path")
)

func main() {
	log.Println("dpipe")

	flag.Parse()

	// load config file
	conf, err := config.New(*configFile)
	if err != nil {
		log.Fatalf("failed to load config file: %v", err)
	}

	// log loaded inputs, outpus, filters and aggregators
	log.Printf("loaded inputs: %+v", config.GetAllKeys(conf.Inputs()))
	log.Printf("loaded outputs: %+v", config.GetAllKeys(conf.Outputs()))
	log.Printf("loaded filters: %+v", config.GetAllKeys(conf.Filters()))
	log.Printf("loaded aggregators: %+v", config.GetAllKeys(conf.Aggregators()))

}
