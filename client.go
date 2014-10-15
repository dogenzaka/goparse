package goparse

import (
	"errors"
	"os"
	"time"
)

const (
	defaultEndPoint = "https://api.parse.com/1"
)

var defaultClient *ParseClient

type ParseConfig struct {
	Url           string
	ApplicationId string
	RESTAPIKey    string
	MasterKey     string
	TimeOut       time.Duration
}

type ParseClient struct {
	Url           string
	ApplicationId string
	RESTAPIKey    string
	MasterKey     string
	TimeOut       time.Duration
}

// Get an default client.
func getDefaultClient() (*ParseClient, error) {
	if defaultClient == nil {
		client, err := NewClient()
		if err != nil {
			return nil, err
		}
		defaultClient = client
	}
	return defaultClient, nil
}

// Create a new parse client with environment variables
func NewClient() (*ParseClient, error) {
	url := os.Getenv("PARSE_ENDPOINT_URL")
	if url == "" {
		// default URL
		url = defaultEndPoint
	}

	appId := os.Getenv("PARSE_APPLICATION_ID")
	if appId == "" {
		return nil, errors.New("client requires PARSE_APPLICATION_ID")
	}

	apiKey := os.Getenv("PARSE_REST_API_KEY")
	if apiKey == "" {
		return nil, errors.New("client requires PARSE_REST_API_KEY")
	}

	return &ParseClient{
		Url:           url,
		ApplicationId: appId,
		RESTAPIKey:    apiKey,
		MasterKey:     os.Getenv("PARSE_MASTER_KEY"),
		TimeOut:       time.Second * 5,
	}, nil
}

// Create new parse client with a parse configuration
func NewClientWithConfig(config ParseConfig) (*ParseClient, error) {

	if config.Url == "" {
		// use default endpoint
		config.Url = defaultEndPoint
	}

	if config.TimeOut == 0 {
		config.TimeOut = time.Millisecond * 5
	}

	if config.ApplicationId == "" {
		return nil, errors.New("client requires ApplicationId")
	}

	if config.RESTAPIKey == "" {
		return nil, errors.New("client requires RESTAPIKey")
	}

	return &ParseClient{
		Url:           config.Url,
		ApplicationId: config.ApplicationId,
		RESTAPIKey:    config.RESTAPIKey,
		MasterKey:     config.MasterKey,
		TimeOut:       config.TimeOut,
	}, nil
}

// Create a new session from the client
func (p *ParseClient) NewSession(sessionToken string) *ParseSession {
	return &ParseSession{
		client:       p,
		SessionToken: sessionToken,
	}
}

// Create a new session from default client
func NewSession(sessionToken string) (*ParseSession, error) {
	client, err := getDefaultClient()
	if err != nil {
		return nil, err
	} else {
		return client.NewSession(sessionToken), nil
	}
}
