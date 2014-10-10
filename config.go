package goparse

import (
	"errors"
	"time"
)

type ParseConfig struct {
	Url           string
	ApplicationId string
	RESTAPIKey    string
	MasterKey     string
	TimeOut       time.Duration
}

func checkConfig(pc ParseConfig) error {
	if pc.Url == "" {
		return errors.New("parse url is empty")
	}
	if pc.ApplicationId == "" {
		return errors.New("parse application id is empty")
	}
	if pc.RESTAPIKey == "" && pc.MasterKey == "" {
		return errors.New("parse rest api key and master key are empty")
	}
	if pc.TimeOut <= 0 {
		return errors.New("parse time out is a negative number or zero")
	}
	return nil
}
