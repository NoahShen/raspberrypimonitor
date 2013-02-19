package main

import (
	"bytes"
	"com.cosm"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"model"
	"net/http"
)

func UpdateDatastreams(c *Config, dv *model.DataValue) error {
	ds := &cosm.Datastream{Id: dv.Id, CurrentValue: dv.Value}
	var url = c.RestUrl + "feeds/" + c.FeedId + "/datastreams/" + ds.Id
	req, e := createRequest(url, "PUT", c.ApiKey, ds)
	if e != nil {
		return e
	}

	_, err := getResponeContent(req)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func createRequest(url, method, key string, cosmEntity interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(cosmEntity)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonBytes)

	req, e := http.NewRequest(method, url, body)
	if e != nil {
		return nil, e
	}
	log.Println(url, string(jsonBytes))
	req.Header.Set("X-ApiKey", key)
	return req, nil
}

func getResponeContent(req *http.Request) (body string, e error) {
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e = err
		return
	}

	body = string(bytes)
	if resp.StatusCode != 200 {
		e = errors.New("status code:" + resp.Status + "\ncontent:" + body)
	}
	return
}
