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

// ParseConfig is the configuration for initializing ParseClient
type ParseConfig struct {
	URL           string
	ApplicationID string
	RESTAPIKey    string
	MasterKey     string
	TimeOut       time.Duration
}

// ParseClient is the client to access Parse REST API
type ParseClient struct {
	URL           string
	ApplicationID string
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
	if apiKey == "" {
		return nil, errors.New("client requires PARSE_REST_API_KEY")
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
		config.TimeOut = time.Millisecond * 5
	}

	if config.ApplicationID == "" {
		return nil, errors.New("client requires ApplicationId")
	}

	if config.RESTAPIKey == "" {
		return nil, errors.New("client requires RESTAPIKey")
	}

	return &ParseClient{
		URL:           config.URL,
		ApplicationID: config.ApplicationID,
		RESTAPIKey:    config.RESTAPIKey,
		MasterKey:     config.MasterKey,
		TimeOut:       config.TimeOut,
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
