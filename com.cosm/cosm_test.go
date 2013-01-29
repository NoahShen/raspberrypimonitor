package cosm

import (
	"encoding/json"
	"testing"
)

func TestGconf(t *testing.T) {

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
