## Overview

![Build Status](https://travis-ci.org/HotelsDotCom/flyte-graphite.svg?branch=master)
[![Docker Stars](https://img.shields.io/docker/stars/hotelsdotcom/flyte-graphite.svg)](https://hub.docker.com/r/hotelsdotcom/flyte-graphite)
[![Docker Pulls](https://img.shields.io/docker/pulls/hotelsdotcom/flyte-graphite.svg)](https://hub.docker.com/r/hotelsdotcom/flyte-graphite)

This Graphite pack provides the ability to add events to graphite

### Build and Run

To build and run from the command line:
* Clone this repo
* Run `dep ensure` (must have [dep](https://github.com/golang/dep) installed )
* Run `go build`
* Run `FLYTE_API=http://localhost:8080/ GRAPHITE_HOST=http://localhost:8090 FLYTE_LABELS="env=lab" ./flyte-graphite`
* Fill in this command with the relevant API url environment variables


### Docker
To build and run from docker
* Run `docker build -t flyte-graphite .`
* Run `docker run -e FLYTE_API=http://localhost:8080/ -e GRAPHITE_HOST=http://localhost:8090 -e FLYTE_LABELS="env=lab"`
* All of these environment variables need to be set


### Commands

### AddEvent command
This command adds a single event to graphite

### Input
```
"input": {
    "tags":"tags",
    "data":"data",
    "what":"what"
 }
```
### Output
This command returns either a `AddEventsSuccess` or `AddEventsFailure`


### AddEventsSuccess returns

```
"payload": {
    "tags":"tags",
    "data":"data",
    "what":"what"
 }
```

### AddEventsFailure returns
```
"payload": {
    "tags":"tags",
    "data":"data",
    "what":"what"
    "error":"status code 500"
 }
```
