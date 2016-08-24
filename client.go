package goparse

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	defaultEndPoint   = "https://api.parse.com/1"
	warningRestAPIKey = `It seems that the the restAPIKey is not set in the parse server if your intentionally not to set this parameter.
This will lead a data leakage if someone knows the application ID.
We strongly recommend you to set (or ask the server master to set) client keys in the parse server`
)

var defaultClient *ParseClient

// ParseConfig is the configuration for initializing ParseClient
type ParseConfig struct {
	URL              string
	ApplicationID    string
	RESTAPIKey       string
	MasterKey        string
	RevocableSession bool
	TimeOut          time.Duration
}

// ParseClient is the client to access Parse REST API
type ParseClient struct {
	URL              string
	ApplicationID    string
	RESTAPIKey       string
	MasterKey        string
	RevocableSession bool
	TimeOut          time.Duration
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

// NewClient is creating ParseClient
func NewClient() (*ParseClient, error) {
	url := os.Getenv("PARSE_ENDPOINT_URL")
	if url == "" {
		// default URL
		url = defaultEndPoint
	}

	appID := os.Getenv("PARSE_APPLICATION_ID")
	if appID == "" {
		return nil, errors.New("client requires PARSE_APPLICATION_ID")
	}

	apiKey := os.Getenv("PARSE_REST_API_KEY")
	if apiKey == "" && url == defaultEndPoint {
		return nil, errors.New("client requires PARSE_REST_API_KEY")
	} else if apiKey == "" {
		fmt.Fprintf(os.Stderr, warningRestAPIKey)
	}

	return &ParseClient{
		URL:           url,
		ApplicationID: appID,
		RESTAPIKey:    apiKey,
		MasterKey:     os.Getenv("PARSE_MASTER_KEY"),
		TimeOut:       time.Second * 5,
	}, nil
}

// NewClientWithConfig creates Parse Client with configuration
func NewClientWithConfig(config ParseConfig) (*ParseClient, error) {

	if config.URL == "" {
		// use default endpoint
		config.URL = defaultEndPoint
	}

	if config.TimeOut == 0 {
		config.TimeOut = time.Second * 5
	}

	if config.ApplicationID == "" {
		return nil, errors.New("client requires ApplicationId")
	}

	if config.RESTAPIKey == "" && config.URL == defaultEndPoint {
		return nil, errors.New("client requires RESTAPIKey")
	} else if config.RESTAPIKey == "" {
		fmt.Fprintf(os.Stderr, warningRestAPIKey)
	}

	return &ParseClient{
		URL:              config.URL,
		ApplicationID:    config.ApplicationID,
		RESTAPIKey:       config.RESTAPIKey,
		MasterKey:        config.MasterKey,
		TimeOut:          config.TimeOut,
		RevocableSession: config.RevocableSession,
	}, nil
}

// NewSession creates a new session from the client
func (p *ParseClient) NewSession(sessionToken string) *ParseSession {
	return &ParseSession{
		client:       p,
		SessionToken: sessionToken,
	}
}

// NewSession create a new session from default client
func NewSession(sessionToken string) (*ParseSession, error) {
	client, err := getDefaultClient()
	if err != nil {
		return nil, err
	}
	return client.NewSession(sessionToken), nil
}
