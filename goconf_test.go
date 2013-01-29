package main

import (
	"code.google.com/p/goconf/conf"
	"testing"
)

func NoTestGconf(t *testing.T) {

	c, err := conf.ReadConfigFile("config/config.test")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(c.GetString("default", "url"))
	t.Log(c.GetString("default", "path"))
	t.Log(c.GetString("key", "all"))

}
