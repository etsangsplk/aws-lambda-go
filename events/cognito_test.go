// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
package events

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-lambda-go/events/test"
)

func TestCognitoEventMarshaling(t *testing.T) {

	// read json from file
	inputJSON, err := ioutil.ReadFile("./testdata/cognito-event.json")
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	// de-serialize into CognitoEvent
	var inputEvent CognitoEvent
	if err := json.Unmarshal(inputJSON, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	// serialize to json
	outputJSON, err := json.Marshal(inputEvent)
	if err != nil {
		t.Errorf("could not marshal event. details: %v", err)
	}

	test.AssertJsonsEqual(t, inputJSON, outputJSON)
}

func TestCognitoMarshalingMalformedJson(t *testing.T) {
	test.TestMalformedJson(t, CognitoEvent{})
}
