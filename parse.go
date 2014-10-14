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
	MasterKey     string
	TimeOut       time.Duration
	NewRequest    func() *gorequest.SuperAgent
}

// Create new parse client
func New(config ParseConfig) (*ParseClient, error) {
	if err := checkConfig(config); err != nil {
		return nil, err
	}

	return &ParseClient{
		Url:           config.Url,
		ApplicationId: config.ApplicationId,
		RESTAPIKey:    config.RESTAPIKey,
		MasterKey:     config.MasterKey,
		TimeOut:       config.TimeOut,
		NewRequest:    gorequest.New,
	}, nil
}

// Get user data for a given session token
// Can use to validate session token is already expired
func (pc *ParseClient) GetMe(sessionToken string) (u User, err error) {
	err = pc.GetCustomMe(sessionToken, &u)
	return
}

func (pc *ParseClient) GetCustomMe(sessionToken string, u interface{}) error {

	if sessionToken == "" {
		return errors.New("parse session token is empty")
	}

	if u == nil {
		return errors.New("user interface is nil")
	}

	req := pc.NewRequest().Get(pc.Url + me)
	pc.initParseHeader(req, false)
	pc.setParseSessionToken(req, sessionToken)

	resp, body, errs := req.Timeout(pc.TimeOut).End()
	if errs != nil {
		return errors.New(fmt.Sprintf("%v", errs))
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New(resp.Status)
	}

	return json.NewDecoder(strings.NewReader(body)).Decode(u)
}

func (pc *ParseClient) initParseHeader(req *gorequest.SuperAgent, useMasterKey bool) {
	req.Set("X-Parse-Application-Id", pc.ApplicationId)
	if useMasterKey {
		req.Set("X-Parse-Master-Key", pc.MasterKey)
	} else {
		req.Set("X-Parse-REST-API-Key", pc.RESTAPIKey)
	}
}

func (pc *ParseClient) setParseSessionToken(req *gorequest.SuperAgent, token string) {
	req.Set("X-Parse-Session-Token", token)
}
