/*
Copyright (C) 2018 Expedia Group.

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

package graphite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HotelsDotCom/go-logger"
	"net/http"
	"net/url"
	"time"
)

type GraphiteEvent struct {
	What string `json:"what"`
	Tags string `json:"tags"`
	Data string `json:"data"`
}

type GraphiteClient interface {
	AddEvent(event GraphiteEvent) error
}

func NewGraphiteClient(httpClient *http.Client, baseUrl string) (GraphiteClient, error) {

	if httpClient == nil {
		httpClient = &http.Client{Timeout: time.Second * 10}
	}

	passedBaseUrl, err := url.Parse(baseUrl)
	if err != nil {
		logger.Errorf("Could parse Url: %v", err)
	}

	return graphiteClient{
		baseUrl:    passedBaseUrl,
		httpClient: httpClient,
	}, err
}

func DefaultGraphiteClient(baseUrl string) (GraphiteClient, error) {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	return NewGraphiteClient(httpClient, baseUrl)
}

type graphiteClient struct {
	baseUrl    *url.URL
	httpClient *http.Client
}

func (c graphiteClient) AddEvent(events GraphiteEvent) error {

	eventData, err := json.Marshal(events)
	if err != nil {
		return err
	}

	request, err := c.constructPostRequest("/events/", eventData)

	if err != nil {
		return err
	}

	return c.sendRequest(request)
}

func (c graphiteClient) constructPostRequest(path string, data []byte) (*http.Request, error) {

	urlStr := fmt.Sprintf("%s%s", c.baseUrl, path)

	request, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer([]byte(data)))

	if err != nil {
		return request, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	return request, err
}

func (c graphiteClient) sendRequest(request *http.Request) error {

	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	return nil
}
