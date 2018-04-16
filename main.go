/*
Copyright (C) 2016-2017 Expedia Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"log"
	"net/url"
	"os"
	"github.com/HotelsDotCom/flyte-client/client"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/HotelsDotCom/flyte-graphite/graphite"
	"github.com/HotelsDotCom/flyte-graphite/service"
	"github.com/HotelsDotCom/go-logger"
	"time"
)

func main() {

	graphiteClient, err := newGraphiteClient()
	if err != nil {
		logger.Errorf("Could not initialise Graphite Client: %v", err)
	}

	commandService := service.NewCommandService(graphiteClient)

	packDef := flyte.PackDef{
		Name:     "Graphite",
		HelpURL:  getUrl("http://github.com/HotelsDotCom/flyte-graphite/README.md"),
		Commands: []flyte.Command{commandService.AddEventCommand()},
	}

	p := flyte.NewPack(packDef, client.NewClient(hostUrl(), 10*time.Second))

	p.Start()

	select {}

}

func newGraphiteClient() (graphite.GraphiteClient, error) {

	graphiteHost := os.Getenv("GRAPHITE_HOST")
	if graphiteHost == "" {
		log.Fatal("GRAPHITE_HOST environment variable is not set")
	}

	return graphite.DefaultGraphiteClient(graphiteHost)

}

func hostUrl() *url.URL {
	flyteHost := os.Getenv("FLYTE_API_URL")
	if flyteHost == "" {
		log.Fatal("FLYTE_API_URL environment variable is not set")
	}

	return getUrl(flyteHost)
}

func getUrl(rawUrl string) *url.URL {
	url, err := url.Parse(rawUrl)
	if err != nil {
		log.Fatalf("%s is not a valid url", rawUrl)
	}
	return url
}
