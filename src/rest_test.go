package main

import (
	"bytes"
	"com.cosm"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	REST_URL = "https://api.cosm.com/v2/"
	KEY      = "oMyWvwF_rWI8e5ULO1tW9pHAUOqSAKxSUm9nZDFIOGhwWT0g"

	FEEDID = "99624"
	FORMAT = "json"
)

func NoTestNewRequest(t *testing.T) {
	req, e := http.NewRequest("GET", "http://example.com/", nil)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	} else {
		t.Log(string(body))
	}
}

func NoTestGetFeeds(t *testing.T) {
	var feedId = "98339"
	var format = "json"
	var url = REST_URL + "feeds/" + feedId + "." + format + "?key=" + KEY
	t.Log(url)
	req, e := http.NewRequest("GET", url, nil)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(string(bytes))
	var jsonObj = make(map[string]interface{})
	json.Unmarshal(bytes, &jsonObj)
	location := jsonObj["location"].(map[string]interface{})
	t.Log(location["domain"])
}

func NoTestUpdateFeeds(t *testing.T) {
	datestreams := make([]cosm.Datastream, 0)
	feed := &cosm.Feed{Datastreams: datestreams}
	datestream := &cosm.Datastream{Id: "diskUsed", CurrentValue: "2"}
	feed.Datastreams = append(feed.Datastreams, *datestream)

	var url = REST_URL + "feeds/" + FEEDID
	req, e := createRequest(url, "PUT", feed, t)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	body, err := getResponeContent(req, t)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(body)
}

func NoTestUpdateDatastreams(t *testing.T) {
	datestream := &cosm.Datastream{CurrentValue: "4"}

	var datastreamId = "diskUsed"
	var url = REST_URL + "feeds/" + FEEDID + "/datastreams/" + datastreamId
	req, e := createRequest(url, "PUT", datestream, t)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	body, err := getResponeContent(req, t)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(body)
}

func createRequest(url, method string, cosmEntity interface{}, t *testing.T) (*http.Request, error) {
	jsonBytes, err := json.Marshal(cosmEntity)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonBytes)

	req, e := http.NewRequest(method, url, body)
	if e != nil {
		return nil, e
	}
	req.Header.Set("X-ApiKey", KEY)
	return req, nil
}

func getResponeContent(req *http.Request, t *testing.T) (string, error) {
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
