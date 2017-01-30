package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

var address = flag.String("address", "consul:8500", "Address is the address of the Consul server")
var service = flag.String("service", "", "Service is the service we are searching for")
var debug = flag.Bool("debug", false, "Debug is for printing debug log.")
var version = flag.Bool("version", false, "Get version of tool")

const versionString string = "0.3"

func main() {
	flag.Parse()

	// print version and stop
	if *version {
		log.Printf("Discovery tool version %v", versionString)
		return
	}

	// making sure service parameter is provided
	if len(*service) < 1 {
		log.Fatal(string("service is manadatory"))
	}

	// setup the consul API interface
	api, err := consulAPI(*address)
	if err != nil {
		log.Fatalf("[Error] Error creating client: %v", err)
	}

	// query the consul catalog for services
	services, _, err := api.Service(*service, "", nil)
	if err != nil {
		log.Fatalf("[Error] Error querying catalog entries: %v", err)
	}

	// contains results
	var result = ""

	// iterate over each service and create a comma-seperated-string containing all service addresses.
	for i, service := range services {
		if i+1 < len(services) {
			// NOT LAST index
			result += fmt.Sprintf("%v,", service.ServiceAddress)
		} else {
			// LAST index
			result += fmt.Sprintf("%v", service.ServiceAddress)
		}
	}

	// print to standard output
	fmt.Println(result)
}

// consulAPI accepts an address and return the Catalog interface from the consul API
func consulAPI(address string) (*api.Catalog, error) {
	conf := api.DefaultConfig()
	conf.Address = "192.168.99.117:8500"

	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return client.Catalog(), nil
}
