package intervalcreation

import (
	_ "embed"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	monitorserialization "github.com/openshift/origin/pkg/monitor/serialization"
)

//go:embed pod_test_01_simple.json
var simplePodLifecyleJSON []byte

func TestIntervalCreation(t *testing.T) {
	inputIntervals, err := monitorserialization.EventsFromJSON(simplePodLifecyleJSON)
	if err != nil {
		t.Fatal(err)
	}
	startTime, err := time.Parse(time.RFC3339, "2022-03-07T12:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	endTime, err := time.Parse(time.RFC3339, "2022-03-07T23:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	result := CreatePodIntervalsFromInstants(inputIntervals, startTime, endTime)

	resultBytes, err := monitorserialization.EventsToJSON(result)
	if err != nil {
		t.Fatal(err)
	}

	expectedJSON := `{
	"items": [
		{
			"level": "Info",
			"locator": "ns/e2e-kubectl-3271 pod/without-label uid/e185b70c-ea3e-4600-850a-b2370a729a73",
			"message": "constructed/true reason/Created ",
			"from": "2022-03-07T18:41:46Z",
			"to": "2022-03-07T18:41:46Z"
		},
		{
			"level": "Info",
			"locator": "ns/e2e-kubectl-3271 pod/without-label uid/e185b70c-ea3e-4600-850a-b2370a729a73",
			"message": "constructed/true reason/Scheduled node/ip-10-0-141-9.us-west-2.compute.internal",
			"from": "2022-03-07T18:41:46Z",
			"to": "2022-03-07T18:41:54Z"
		},
		{
			"level": "Info",
			"locator": "ns/e2e-kubectl-3271 pod/without-label uid/e185b70c-ea3e-4600-850a-b2370a729a73 container/without-label",
			"message": "constructed/true reason/ContainerWait missed real \"ContainerWait\"",
			"from": "2022-03-07T18:41:46Z",
			"to": "2022-03-07T18:41:52Z"
		},
		{
			"level": "Info",
			"locator": "ns/e2e-kubectl-3271 pod/without-label uid/e185b70c-ea3e-4600-850a-b2370a729a73 container/without-label",
			"message": "constructed/true reason/NotReady missed real \"NotReady\"",
			"from": "2022-03-07T18:41:52Z",
			"to": "2022-03-07T18:41:52Z"
		},
		{
			"level": "Info",
			"locator": "ns/e2e-kubectl-3271 pod/without-label uid/e185b70c-ea3e-4600-850a-b2370a729a73 container/without-label",
			"message": "constructed/true reason/ContainerStart cause/ duration/6.00s",
			"from": "2022-03-07T18:41:52Z",
			"to": "2022-03-07T18:41:54Z"
		},
		{
			"level": "Info",
			"locator": "ns/e2e-kubectl-3271 pod/without-label uid/e185b70c-ea3e-4600-850a-b2370a729a73 container/without-label",
			"message": "constructed/true reason/Ready ",
			"from": "2022-03-07T18:41:52Z",
			"to": "2022-03-07T18:41:54Z"
		}
	]
}`

	expectedJSON = strings.ReplaceAll(expectedJSON, "\t", "    ")

	resultJSON := string(resultBytes)
	if expectedJSON != resultJSON {
		t.Fatal(cmp.Diff(expectedJSON, resultJSON))
	}
}

//go:embed pod_test_02_trailing_ready.json
var trailingReadyPodLifecyleJSON []byte

func TestIntervalCreation_TrailingReady(t *testing.T) {
	inputIntervals, err := monitorserialization.EventsFromJSON(trailingReadyPodLifecyleJSON)
	if err != nil {
		t.Fatal(err)
	}
	startTime, err := time.Parse(time.RFC3339, "2022-03-07T12:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	endTime, err := time.Parse(time.RFC3339, "2022-03-10T23:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	result := CreatePodIntervalsFromInstants(inputIntervals, startTime, endTime)

	resultBytes, err := monitorserialization.EventsToJSON(result)
	if err != nil {
		t.Fatal(err)
	}

	expectedJSON := `{
            "items": [
                {
                    "level": "Info",
                    "locator": "ns/openshift-marketplace pod/community-operators-sp6lm uid/efb1885a-1fe1-4f5b-ad41-044e55f806a9",
                    "message": "constructed/true reason/Created ",
                    "from": "2022-03-07T22:47:04Z",
                    "to": "2022-03-07T22:47:04Z"
                },
                {
                    "level": "Info",
                    "locator": "ns/openshift-marketplace pod/community-operators-sp6lm uid/efb1885a-1fe1-4f5b-ad41-044e55f806a9",
                    "message": "constructed/true reason/Scheduled node/ip-10-0-154-151.ec2.internal",
                    "from": "2022-03-07T22:47:04Z",
                    "to": "2022-03-07T22:47:15Z"
                },
                {
                    "level": "Info",
                    "locator": "ns/openshift-marketplace pod/community-operators-sp6lm uid/efb1885a-1fe1-4f5b-ad41-044e55f806a9 container/registry-server",
                    "message": "constructed/true reason/ContainerWait missed real \"ContainerWait\"",
                    "from": "2022-03-07T22:47:04Z",
                    "to": "2022-03-07T22:47:07Z"
                },
                {
                    "level": "Info",
                    "locator": "ns/openshift-marketplace pod/community-operators-sp6lm uid/efb1885a-1fe1-4f5b-ad41-044e55f806a9 container/registry-server",
                    "message": "constructed/true reason/NotReady missed real \"NotReady\"",
                    "from": "2022-03-07T22:47:07Z",
                    "to": "2022-03-07T22:47:14Z"
                },
                {
                    "level": "Info",
                    "locator": "ns/openshift-marketplace pod/community-operators-sp6lm uid/efb1885a-1fe1-4f5b-ad41-044e55f806a9 container/registry-server",
                    "message": "constructed/true reason/ContainerStart cause/ duration/3.00s",
                    "from": "2022-03-07T22:47:07Z",
                    "to": "2022-03-07T22:47:15Z"
                },
                {
                    "level": "Info",
                    "locator": "ns/openshift-marketplace pod/community-operators-sp6lm uid/efb1885a-1fe1-4f5b-ad41-044e55f806a9 container/registry-server",
                    "message": "constructed/true reason/Ready ",
                    "from": "2022-03-07T22:47:14Z",
                    "to": "2022-03-07T22:47:15Z"
                },
                {
                    "level": "Info",
                    "locator": "ns/openshift-marketplace pod/community-operators-sp6lm uid/efb1885a-1fe1-4f5b-ad41-044e55f806a9 container/registry-server",
                    "message": "constructed/true reason/NotReady ",
                    "from": "2022-03-07T22:47:15Z",
                    "to": "2022-03-07T22:47:15Z"
                }
            ]
        }`

	expectedJSON = strings.ReplaceAll(expectedJSON, "\t", "    ")

	resultJSON := string(resultBytes)
	if expectedJSON != resultJSON {
		t.Fatal(resultJSON)
	}
}
