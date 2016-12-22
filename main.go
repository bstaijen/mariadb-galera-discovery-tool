package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var host = flag.String("h", "consul:8500", "Host[:port] for consul server")
var serviceName = flag.String("servicename", "", "Name of the service in consul")

func main() {

	version := "0.0.1"
	log.Printf("Starting discovery tool %s ...", version)

	flag.Parse()

	log.Printf("Nice %v \n", *serviceName)
	if len(*serviceName) < 1 {
		log.Fatal(string("servicename is manadatory"))
	}

	// TODO search for tag
	// TODO think about searching in DNS record

	url := fmt.Sprintf("http://%v/v1/catalog/service/%v", *host, *serviceName)
	log.Printf("Get service catalog %v", url)

	request("GET", url, []byte(string("")), func(res *http.Response) {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}

		type itemdata []interface{}
		var datas itemdata

		json.Unmarshal(data, &datas)

		for i := 0; i < len(datas); i++ {
			m := datas[i].(map[string]interface{})
			if i+1 < len(datas) {
				// NOT LAST
				fmt.Printf("%v,", m["ServiceAddress"])
			} else {
				// LAST
				fmt.Printf("%v\n", m["ServiceAddress"])
			}
		}
	})

}

func request(method, url string, body []byte, cb func(*http.Response)) error {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error creating request: " + err.Error())
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error executing request: " + err.Error())
		return err
	}

	// callback
	cb(resp)

	//return
	return nil
}
