package goparse

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

const (
	me = "/users/me"
)

type ParseClient struct {
	Url           string
	ApplicationId string
	RESTAPIKey    string
	TimeOut       time.Duration
	NewRequest    func() *gorequest.SuperAgent
}

// Create new parse client
func New(config ParseConfig) (*ParseClient, error) {
	err := checkConfig(config)
	if err != nil {
		return nil, err
	}

	return &ParseClient{
		Url:           config.Url,
		ApplicationId: config.ApplicationId,
		RESTAPIKey:    config.RESTAPIKey,
		TimeOut:       config.TimeOut,
		NewRequest:    gorequest.New,
	}, nil
}

// Get user data for a given session token
//Can use to validate session tokne is already expired
func (pc *ParseClient) GetMe(sessionToken string) (u User, err error) {

	if sessionToken == "" {
		err = errors.New("parse session token is empty")
		return
	}

	req := pc.NewRequest().Get(pc.Url + me)
	pc.initParseHeader(req)
	pc.setParseSessionToken(req, sessionToken)

	resp, body, errs := req.Timeout(pc.TimeOut).End()
	if errs != nil {
		err = errors.New(fmt.Sprintf("%v", errs))
		return
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = errors.New(resp.Status)
		return
	}

	err = json.NewDecoder(strings.NewReader(body)).Decode(&u)

	return
}

func (pc *ParseClient) initParseHeader(req *gorequest.SuperAgent) {
	req.Set("X-Parse-Application-Id", pc.ApplicationId)
	req.Set("X-Parse-REST-API-Key", pc.RESTAPIKey)
}

func (pc *ParseClient) setParseSessionToken(req *gorequest.SuperAgent, token string) {
	req.Set("X-Parse-Session-Token", token)
}
