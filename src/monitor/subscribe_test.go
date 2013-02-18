package main

import (
	"testing"
)

func TestSubResource(t *testing.T) {
	resource := "/feeds/98339/datastreams/commandstream"
	hostip := "api.cosm.com:8081"
	//hostip := "173.203.98.29:8081"
	apiKey := "oMyWvwF_rWI8e5ULO1tW9pHAUOqSAKxSUm9nZDFIOGhwWT0g"
	err := SubscribeResource(resource, hostip, apiKey)
	if err != nil {
		t.Log(err)
	}
}
