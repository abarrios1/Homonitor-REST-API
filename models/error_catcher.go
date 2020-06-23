package models

import "log"

// ErrorCatcher ... Struct for holding error information.
type ErrorCatcher struct {
	ErrorType string
	Error string
}

// CheckErr ... Check if errors exist and log the error caught.
func CheckErr(errType string, err error)  {
	if err != nil {
		var errCatcher ErrorCatcher
		errCatcher.ErrorType = errType
		errCatcher.Error = err.Error()

		log.Printf("REST API ERROR: %v")
	}
}
