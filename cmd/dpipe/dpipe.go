package main

import (
	"flag"
	"log"

	"github.com/maksadbek/dpipe/agent"
	"github.com/maksadbek/dpipe/config"
)

var (
	configFile = flag.String("config", "config", "config file path")
)

func main() {
	log.Println("DPIPE")

	flag.Parse()

	// load config file
	conf, err := config.New(*configFile)
	if err != nil {
		log.Fatalf("failed to load config file: %v", err)
	}

	// log loaded inputs, outpus, filters and aggregators
	log.Printf("I! loaded inputs: %+v", config.GetAllKeys(conf.Inputs()))
	log.Printf("I! loaded outputs: %+v", config.GetAllKeys(conf.Outputs()))
	log.Printf("I! loaded filters: %+v", config.GetAllKeys(conf.Filters()))
	log.Printf("I! loaded aggregators: %+v", config.GetAllKeys(conf.Aggregators()))

	// create new agent
	agent := agent.New(conf)

	// initialize agent
	agent.Init()

	// run inputs and outputs
	agent.Run()
}
