package main

import (
	"com.cosm"
	"encoding/json"
	"fmt"
	"net"
	"utils"
)

func SubscribeResource(resource, socketHost, apiKey string) error {
	client, err := net.Dial("tcp", socketHost)
	if err != nil {
		return err
	}
	defer client.Close()
	token := utils.RandomString(10)
	subRequest := createSubRequest(resource, token, apiKey)
	jsonBytes, err2 := json.Marshal(subRequest)
	if err2 != nil {
		return nil
	}

	client.Write(jsonBytes)
	buf := make([]byte, 1024)
	c, err := client.Read(buf)
	if err != nil {
		return err
	}
	fmt.Println(string(buf[0:c]))
	return nil
}

func createSubRequest(resource, token, apiKey string) *cosm.SocketRequest {
	request := &cosm.SocketRequest{
		Method:   cosm.Subscribe,
		Resource: resource,
		Token:    token,
	}
	header := make(map[string]string)
	header["X-ApiKey"] = apiKey
	request.Header = header
	return request
}
