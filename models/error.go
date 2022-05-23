package models

import "time"

type Error struct {
	Message   string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}

var (
	MethodNotAllowed   = Error{Message: "Method not allowed.", Timestamp: time.Now()}
	EndpointNotFound   = Error{Message: "Endpoint not found.", Timestamp: time.Now()}
	AccountFetchError  = Error{Message: "Unable to fetch user account.", Timestamp: time.Now()}
	UserFetchError     = Error{Message: "Unable to fetch user data.", Timestamp: time.Now()}
	AuthorizationError = Error{Message: "User could not be authorized.", Timestamp: time.Now()}
)
