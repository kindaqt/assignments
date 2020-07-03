package retry

import (
	"log"
	"time"

	customErrors "github.com/kindaqt/assignment2/errors"
)

// TODO:
// - Retry backoff
// - Make retryable errors configurable by passing the logic in as a function with an input of the error. Doing so will remove the dependency on the customErrors package.

// Do retries an action based based on maxAttempts and sleep
func Do(maxAttempts int, sleep time.Duration, action func() error) error {
	var err error
	for maxAttempts > 0 {
		log.Printf("Retry: remaining attempts %v", maxAttempts)
		err = action()
		if err == nil {
			return err
		}
		maxAttempts--

		// Check if retryable
		switch err.(type) {
		case customErrors.TemporaryError: // retryable error
			log.Printf("Retry: retryable error: %v", err)
			time.Sleep(sleep)
		default:
			log.Printf("Retry: non-retryable error: %v", err)
			return err
		}
	}
	return err
}
