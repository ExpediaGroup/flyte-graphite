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
	"encoding/json"
	"fmt"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/HotelsDotCom/flyte-graphite/graphite"
	"github.com/HotelsDotCom/go-logger"
	"github.com/HotelsDotCom/flyte-graphite/event"
)

func (c CommandService) AddEventCommand() flyte.Command {

	return flyte.Command{
		Name:         "AddEvents",
		OutputEvents: []flyte.EventDef{events.AddEventSuccessEventDef, events.AddEventErrorEventDef},
		Handler:      c.AddEventHandler,
	}

}

func (c CommandService) AddEventHandler(input json.RawMessage) flyte.Event {

	handlerInput := &graphite.GraphiteEvent{}

	if err := json.Unmarshal(input, &handlerInput); err != nil {
		err := fmt.Errorf("Could not marshal Tags,What and Data into json: %v\n", err)
		logger.Error(err)
		return events.NewAddEventsFailureEvent(fmt.Sprintf("Fail: %s", err), "unknown", "unknown", "unknown")
	}

	if handlerInput.What == "" {
		logger.Errorf("No What input supplied")
		return events.NewAddEventsFailureEvent(handlerInput.What, handlerInput.Data, handlerInput.Tags, "What must not be null")

	}

	if err := c.graphiteClient.AddEvent(graphite.GraphiteEvent{What: handlerInput.What, Tags: handlerInput.Tags, Data: handlerInput.Data}); err != nil {
		err := fmt.Errorf("Could not apply event: %v", err)
		logger.Error(err)
		return events.NewAddEventsFailureEvent(fmt.Sprintf("Fail: %s", err), handlerInput.What, handlerInput.Data, handlerInput.Tags)
	}

	return events.NewAddEventsSuccessEvent(handlerInput.Tags, handlerInput.What, handlerInput.Data)
}
