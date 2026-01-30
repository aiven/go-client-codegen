// Package aiven provides a client for interacting with the Aiven API.
package aiven

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Error represents an Aiven API Error.
type Error struct {
	// Aiven error response fields https://api.aiven.io/doc/#section/Responses/Failed-responses
	Message string `json:"message"`
	Errors  any    `json:"errors"`

	// Internal fields
	OperationID string `json:"-"`
	Status      int    `json:"-"`
}

// Error concatenates all the fields.
func (e Error) Error() string {
	msg := e.Message
	if e.Errors != nil {
		// Appends Errors if they don't contain the message
		// Otherwise, it will be duplicated
		b, err := json.Marshal(e.Errors)
		errMsg := string(b)
		if err != nil {
			errMsg = err.Error()
		}

		if !strings.Contains(errMsg, msg) {
			msg = fmt.Sprintf("%s (%s)", msg, errMsg)
		}
	}

	// Must not use `%q` here which will escape every quote in the string,
	// which might break external substring checks
	return fmt.Sprintf(`[%d %s]: %s`, e.Status, e.OperationID, msg)
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

func fromBytes(operationID string, statusCode int, b []byte) ([]byte, error) {
	e := Error{OperationID: operationID, Status: statusCode}

	if statusCode < 200 || statusCode >= 300 {
		// According to the documentation,
		// failed responses must have "errors" and "message" fields
		// https://api.aiven.io/doc/#section/Responses/Failed-responses
		err := json.Unmarshal(b, &e)
		if err != nil {
			// 1. The body might contain sensitive data
			// 2. It might fail the unmarshalling into Error and still be a valid json
			if json.Valid(b) {
				e.Message = err.Error()
			} else {
				// If it is not valid json, it shouldn't contain sensitive data
				// Returns the body, because clients might relay on the error message
				e.Message = string(b)
			}
		}
		return nil, e
	}

	return b, nil
}

func fromResponse(operationID string, rsp *http.Response) ([]byte, error) {
	b, err := io.ReadAll(rsp.Body)
	if err != nil {
		e := Error{OperationID: operationID, Status: rsp.StatusCode}
		e.Message = fmt.Sprintf("body read error: %s", err)
		return nil, e
	}

	return fromBytes(operationID, rsp.StatusCode, b)
}
