package goparse

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/parnurzeal/gorequest"
)

const (
	headerAppId        = "X-Parse-Application-Id" // Parse Application ID
	headerMasterKey    = "X-Parse-Master-Key"     // Parse Master Key
	headerApiKey       = "X-Parse-REST-API-Key"   // Parse REST API Key
	headerSessionToken = "X-Parse-Session-Token"  // Parse Session Token

	pathMe = "/users/me"
)

type ParseSession struct {
	client       *ParseClient
	SessionToken string
}

// Create a request which is set headers for Parse API
func (s *ParseSession) initRequest(req *gorequest.SuperAgent) {
	req.
		Set(headerAppId, s.client.ApplicationId).
		Set(headerApiKey, s.client.RESTAPIKey).
		Timeout(s.client.TimeOut)

	if s.SessionToken != "" {
		req.Set(headerSessionToken, s.SessionToken)
	}
}

// Create a request which is set mater key
func (s *ParseSession) initMasterRequest(req *gorequest.SuperAgent) {
	req.
		Set(headerAppId, s.client.ApplicationId).
		Set(headerMasterKey, s.client.MasterKey).
		Timeout(s.client.TimeOut)

	if s.SessionToken != "" {
		req.Set(headerSessionToken, s.SessionToken)
	}
}

func (s *ParseSession) get(path string) *gorequest.SuperAgent {
	req := gorequest.New().Get(s.client.Url + path)
	s.initRequest(req)
	return req
}

func (s *ParseSession) post(path string) *gorequest.SuperAgent {
	req := gorequest.New().Post(s.client.Url + path)
	s.initRequest(req)
	return req
}

func (s *ParseSession) del(path string) *gorequest.SuperAgent {
	req := gorequest.New().Delete(s.client.Url + path)
	s.initRequest(req)
	return req
}

// Signup new user
func (s *ParseSession) Signup(data Signup) error {
	return do(s.post("/users").Send(data), nil)
}

// Login with data
func (s *ParseSession) Login(username string, password string) (user User, err error) {

	// Query values
	vals := url.Values{
		"username": []string{username},
		"password": []string{password},
	}

	// Create a user
	err = do(s.get("/login").Query(vals.Encode()), &user)

	if user.SessionToken != "" {
		s.SessionToken = user.SessionToken
	}

	return user, err
}

// Get self user information
func (s *ParseSession) GetMe() (user User, err error) {
	err = s.GetMeInto(&user)
	return user, err
}

// Get self user information into provided object
func (s *ParseSession) GetMeInto(user interface{}) error {
	if user == nil {
		return errors.New("user must not be nil")
	}
	return do(s.get("/users/me"), user)
}

// Delete user
func (s *ParseSession) DeleteUser(userId string) error {
	return do(s.del("/users/"+userId), nil)
}

// Execute a parse request
func do(req *gorequest.SuperAgent, data interface{}) error {

	res, body, errs := req.End()
	if errs != nil {
		return errors.New(fmt.Sprintf("%v", errs))
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		// parse as error model
		reserr := new(Error)
		err := json.NewDecoder(strings.NewReader(body)).Decode(reserr)
		if err != nil {
			return err
		} else {
			return errors.New(reserr.Message + " - code:" + strconv.Itoa(reserr.Code))
		}
	}
	if data == nil {
		return nil
	} else {
		return json.NewDecoder(strings.NewReader(body)).Decode(data)
	}
}
