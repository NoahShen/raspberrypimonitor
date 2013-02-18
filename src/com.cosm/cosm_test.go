package cosm

import (
	"encoding/json"
	"testing"
)

func NoTestGconf(t *testing.T) {

	datestreams := make([]Datastream, 0)
	feed := &Feed{Id: "99624", Datastreams: datestreams}
	datestream := &Datastream{Id: "diskUsed", CurrentValue: "2"}
	feed.Datastreams = append(feed.Datastreams, *datestream)

	bytes, err := json.Marshal(feed)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(feed)
	t.Log(string(bytes))
}

func TestSocketRequest(t *testing.T) {
	resource := "/feeds/98339/datastreams/commandstream"
	apiKey := "oMyWvwF_rWI8e5ULO1tW9pHAUOqSAKxSUm9nZDFIOGhwWT0g"
	request := &SocketRequest{
		Method:   Subscribe,
		Resource: resource,
		Token:    "token",
	}
	header := make(map[string]string)
	header["X-ApiKey"] = apiKey
	request.Header = header
	bytes, err := json.Marshal(request)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(request)
	t.Log(string(bytes))
}
