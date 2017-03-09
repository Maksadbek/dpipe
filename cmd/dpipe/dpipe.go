package main

import (
	"flag"
	"log"

	"github.com/maksadbek/dpipe/agent"
	"github.com/maksadbek/dpipe/aggregators"
	"github.com/maksadbek/dpipe/config"
	"github.com/maksadbek/dpipe/filters"
	"github.com/maksadbek/dpipe/inputs"
	"github.com/maksadbek/dpipe/outputs"

	_ "github.com/maksadbek/dpipe/aggregators/all"
	_ "github.com/maksadbek/dpipe/filters/all"
	_ "github.com/maksadbek/dpipe/inputs/all"
	_ "github.com/maksadbek/dpipe/outputs/all"
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
	log.Printf("I! registered inputs: %+v", inputs.RegisteredInputs())
	log.Printf("I! registered outputs: %+v", outputs.RegisteredOutputs())
	log.Printf("I! registered filters: %+v", filters.RegisteredFilters())
	log.Printf("I! registered aggregators: %+v", aggregators.RegisteredAggregators())

	// create new agent
	agent := agent.New(conf)

	// initialize agent
	agent.Init()
	defer agent.CloseOutputs()

	// run inputs and outputs
	agent.Run()

	// print stats

	log.Print("I! finished processing, stats:")
	log.Printf("I! failed to write:\t\t %d", agent.Stats.DataWrittenFailed)
	log.Printf("I! succeed to write:\t %d", agent.Stats.DataWrittenOK)
	log.Printf("I! validation fails:\t %d", agent.Stats.DataValidationFailed)
	log.Printf("I! received:\t\t %d", agent.Stats.DataReceived)
	log.Printf("I! aggregated:\t\t %d", agent.Stats.DataAggregatedOK)
	log.Printf("I! failed aggreations:\t %d", agent.Stats.DataAggregatedFailed)
	log.Printf("I! aggreation errors:\t %d", agent.Stats.AggregationErrors)
}
