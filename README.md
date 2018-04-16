## Overview

This Graphite pack provides the ability to add events to graphite

### Build and Run

To build and run from the command line:
* Clone this repo
* Run `dep ensure` (must have [dep](https://github.com/golang/dep) installed )
* Run `go build`
* Run `FLYTE_API_URL=http://localhost:8080/ GRAPHITE_HOST=http://localhost:8090`
* Fill in this command with the relevant API url, bamboo host, bamboo user and bamboo password environment variables


### Docker
To build and run from docker
* Run `docker build -t flyte-graphite .`
* Run docker run -e FLYTE=http://localhost:8080/ -e GRAPHITE_HOST=http://localhost:8090`
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

### AddEventsFailre returns
```
"payload": {
    "tags":"tags",
    "data":"data",
    "what":"what"
    "error":"status code 500"
 }
```
