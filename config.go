package goparse

import (
	"errors"
	"time"
)

type parseConfig struct {
	Url           string
	ApplicationId string
	RESTAPIKey    string
	TimeOut       time.Duration
}

var config parseConfig

func InitConfig(url string, applicationId string, restApiKey string, timeOut time.Duration) error {
	config = parseConfig{url, applicationId, restApiKey, timeOut * time.Millisecond}
	return checkConfig()
}

func checkConfig() error {
	if config.Url == "" {
		return errors.New("parse url is empty")
	}
	if config.ApplicationId == "" {
		return errors.New("parse application id is empty")
	}
	if config.RESTAPIKey == "" {
		return errors.New("parse rest api key is empty")
	}
	if config.TimeOut <= 0 {
		return errors.New("parse time out is a negative number or zero")
	}
	return nil
}
