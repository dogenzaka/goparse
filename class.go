package goparse

import (
	"errors"
)

// ParseClass is an object that contains the ParseSession
type ParseClass struct {
	Session   *ParseSession
	Name      string
	ClassURL  string
	UseMaster bool
}

// Select gets class data information
func (c *ParseClass) Select(objectID string, result interface{}) error {
	path := c.ClassURL
	if objectID != "" {
		path = c.ClassURL + "/" + objectID
	}
	return do(c.Session.get(path, c.UseMaster), &result)
}

// Create creates class from data
func (c *ParseClass) Create(data interface{}, result interface{}) error {
	return do(c.Session.post(c.ClassURL, c.UseMaster).Send(data), &result)
}

// Update updates class by ID
func (c *ParseClass) Update(objectID string, data interface{}, result interface{}) error {
	if objectID == "" {
		return errors.New("ObjectID must not be empty")
	}
	return do(c.Session.put(c.ClassURL+"/"+objectID, c.UseMaster).Send(data), &result)
}

// Delete deletes class by ID
func (c *ParseClass) Delete(objectID string) error {
	if objectID == "" {
		return errors.New("ObjectID must not be empty")
	}
	return do(c.Session.del(c.ClassURL+"/"+objectID, c.UseMaster), nil)
}
