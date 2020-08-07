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
	"github.com/ExpediaGroup/flyte-graphite/graphite"
	"github.com/ExpediaGroup/flyte-graphite/service"
	"github.com/HotelsDotCom/flyte-client/client"
	"github.com/HotelsDotCom/flyte-client/config"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/HotelsDotCom/go-logger"
	"log"
	"net/url"
	"os"
	"time"
)

func main() {

	graphiteClient, err := newGraphiteClient()
	if err != nil {
		logger.Errorf("Could not initialise Graphite Client: %v", err)
	}

	commandService := service.NewCommandService(graphiteClient)

	envVars := config.FromEnvironment()
	packDef := flyte.PackDef{
		Name:     "Graphite",
		Labels:   envVars.Labels,
		HelpURL:  getUrl("http://github.com/HotelsDotCom/flyte-graphite/README.md"),
		Commands: []flyte.Command{commandService.AddEventCommand()},
	}
	p := flyte.NewPack(packDef, client.NewInsecureClient(envVars.FlyteApiUrl, 10*time.Second))

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

func getUrl(rawUrl string) *url.URL {
	url, err := url.Parse(rawUrl)
	if err != nil {
		log.Fatalf("%s is not a valid url", rawUrl)
	}
	return url
}
