package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/json-iterator/go"
	"log"
)

type FetchAndQueue struct {
	Fetch
}

func (FetchAndQueue) Name() string {
	return "get_records_and_queue"
}

func (FetchAndQueue) Version() string {
	return "1.0"
}
func (fetch FetchAndQueue) Execute(in step.Context) (interface{}, error) {
	input := FetchAndQueueInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return fetch.executeQueue(input, in)
}

func (fetch FetchAndQueue) executeQueue(input FetchAndQueueInput, in step.Context) (interface{}, error) {
	// set the username and pass on the fetcher
	fetch.username = input.Username
	fetch.password = input.Password
	fetch.url = input.Url
	// first thing we need to do is login to the FAMIS services
	// we require Username, Password, and Url so we do not have to validate the values
	var err error
	fetch.authItem, err = fetch.Login(input.Username, input.Password, input.Url)
	if err != nil {
		return nil, err
	}

	uri, err := fetch.buildUrl(input.FetchInput)
	if err != nil {
		return nil, err
	}

	engine := in.Engine()
	err = fetch.performPagedFetch(uri, input.Endpoint, func(messages []jsoniter.RawMessage) {
		for _, msg := range messages {
			if input.ChildPath == "" {
				err := engine.AddToQueue(input.Workflow, msg)
				if err != nil {
					log.Fatalf("Unable to add item to engine queue: %s", err.Error())
				}
				continue
			}
			parseAndQueueChildItems(msg, input, engine)
		}
	})
	return nil, err
}

/**
This method will attempt to parse the raw json into a map and then look for the record at the input path.
If it is an array, then it will loop through the array and queue each item separately, otherwise it queues
whatever it finds at the path.
 */
func parseAndQueueChildItems(msg jsoniter.RawMessage, input FetchAndQueueInput, engine step.Engine) {
	var record map[string]interface{}
	err := jsoniter.Unmarshal(msg, &record)
	if err != nil {
		log.Fatalf("Unable to parse item from response: %s", err.Error())
		return
	}
	children := record[input.ChildPath]
	if children == nil {
		return
	}
	if arrayVal, ok := children.([]interface{}); ok {
		for _, record := range arrayVal {
			err := engine.AddToQueue(input.Workflow, record)
			if err != nil {
				log.Fatalf("Unable to add item to engine queue: %s", err.Error())					}
		}
		return
	}
	if singleVal, ok := children.(interface{}); ok {
		err := engine.AddToQueue(input.Workflow, singleVal)
		if err != nil {
			log.Fatalf("Unable to add item to engine queue: %s", err.Error())
		}
	}
}
