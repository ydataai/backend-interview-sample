package common

import (
	"time"
)

// UID defines UID type represents a ID for a specific resource
type UID string

// Namespace defines the namespace type that represents the cluster namespace
type Namespace string

// UserID defines the user id type
type UserID string

// RESTRequest defines a struct for a request
type RESTRequest struct {
	URL         string
	Body        interface{}
	Headers     map[string]string
	QueryParams map[string]string
	Timeout     time.Duration
}

func (n Namespace) String() string {
	return string(n)
}
