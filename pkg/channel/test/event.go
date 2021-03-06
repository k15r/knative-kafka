package test

import (
	"encoding/json"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/kyma-incubator/knative-kafka/pkg/channel/constants"
)

// Test Data
var EventDataJson, _ = json.Marshal(map[string]string{EventDataKey: EventDataValue})

// Utility Function For Creating A Test CloudEvent
func CreateCloudEvent(cloudEventVersion string) *event.Event {
	cloudEvent := event.New(cloudEventVersion)
	cloudEvent.SetID(EventId)
	cloudEvent.SetType(EventType)
	cloudEvent.SetSource(EventSource)
	cloudEvent.SetDataContentType(EventDataContentType)
	cloudEvent.SetSubject(EventSubject)
	cloudEvent.SetDataSchema(EventDataSchema)
	cloudEvent.SetExtension(constants.ExtensionKeyPartitionKey, PartitionKey)
	_ = cloudEvent.SetData(EventDataContentType, EventDataJson)
	return &cloudEvent
}
