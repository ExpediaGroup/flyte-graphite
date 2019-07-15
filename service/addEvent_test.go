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

package service

import (
	"github.com/HotelsDotCom/flyte-graphite/event"
	"github.com/HotelsDotCom/flyte-graphite/graphite"
	"reflect"
	"testing"
)

func TestAddEventsHandler_Success(t *testing.T) {

	mockClient := MockGraphiteClient{}
	mockClient.addEvents = func(events graphite.GraphiteEvent) error { return nil }
	service := NewCommandService(mockClient)

	got := service.AddEventHandler([]byte(`{"what":"what","tags":"tags","data":"data"}`))

	if !reflect.DeepEqual("AddEventsSuccess", got.EventDef.Name) {
		t.Errorf("expected event AddEventsSuccess, got %s", got.EventDef.Name)
	}
}

func TestAddEventsHandler_Failure(t *testing.T) {

	mockClient := MockGraphiteClient{}
	mockClient.addEvents = func(events graphite.GraphiteEvent) error { return nil }
	service := NewCommandService(mockClient)

	got := service.AddEventHandler([]byte(`{"what":""}`))

	if !reflect.DeepEqual(got.EventDef.Name, "AddEventsFailure") {
		t.Errorf("expected event AddEventsFailure, got %s", got.EventDef.Name)
	}
}

func TestAddEvents_shouldPushEvents(t *testing.T) {

	mockClient := MockGraphiteClient{}
	mockClient.addEvents = func(events graphite.GraphiteEvent) error { return nil }
	service := NewCommandService(mockClient)

	got := service.AddEventHandler([]byte(`{"what":"what", "tags":"tags", "data":"data"}`))

	expectedPayload := events.AddEventsSuccessPayload{Tags: "tags", What: "what", Data: "data"}

	if reflect.DeepEqual(got.Payload, expectedPayload) {
		t.Logf("expectedPayload:  %+v, got: %+v", expectedPayload, got.Payload)
	} else {
		t.Errorf("expectedPayload:  %+v, got: %+v", expectedPayload, got.Payload)
	}
}

func TestAddEvents_invalidInput(t *testing.T) {

	mockClient := MockGraphiteClient{}
	mockClient.addEvents = func(events graphite.GraphiteEvent) error { return nil }
	service := NewCommandService(mockClient)

	got := service.AddEventHandler([]byte(`{"tags":"tags", "datadata"}`))

	if reflect.DeepEqual(got.EventDef.Name, "AddEventsFailure") {
		t.Logf("expected event:  AddEventsFailure, got %+v ", got.EventDef.Name)
	} else {
		t.Errorf("expected event: AddEventsFailure got %+v", got.EventDef.Name)
	}
}

func TestAddEvents_WhatMustNotBeNull(t *testing.T) {

	mockClient := MockGraphiteClient{}
	mockClient.addEvents = func(events graphite.GraphiteEvent) error { return nil }
	service := NewCommandService(mockClient)

	got := service.AddEventHandler([]byte(`{"what":"","tags":"tags","data":"some data"}`))

	if reflect.DeepEqual(got.EventDef.Name, "AddEventsFailre") {
		t.Errorf("expected event: AddEventsFailure, got %+v", got.EventDef.Name)
	}
}

type MockGraphiteClient struct {
	addEvents func(events graphite.GraphiteEvent) error
}

func (m MockGraphiteClient) AddEvent(events graphite.GraphiteEvent) error {
	return m.addEvents(events)
}
