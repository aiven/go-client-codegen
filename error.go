// Package aiven provides a client for interacting with the Aiven API.
package aiven

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Error represents an Aiven API Error.
type Error struct {
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

// Error concatenates the Status, Message and MoreInfo values.
func (e Error) Error() string {
	return fmt.Sprintf("%d: %s - %s", e.Status, e.Message, e.MoreInfo)
}

// IsNotFound returns true if the specified error has status 404
func IsNotFound(err error) bool {
	var e Error

	return errors.As(err, &e) && e.Status == http.StatusNotFound
}

// IsAlreadyExists returns true if the error message and error code that indicates that entity already exists
func IsAlreadyExists(err error) bool {
	var e Error

	return errors.As(err, &e) && strings.Contains(e.Message, "already exists") && e.Status == http.StatusConflict
}

// IsNilOrNotFound returns true for nil and 404 error.
// This check is quite often used for resource deletion when 404 is not an issue.
func IsNilOrNotFound(err error) bool {
	return err == nil || IsNotFound(err)
}
