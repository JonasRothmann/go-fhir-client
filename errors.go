package fhirclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type FhirError struct {
	//Message   string
	OpOutcome fmt.Stringer
}

func NewError(r io.Reader) (err FhirError) {
	marshalErr := json.NewDecoder(r).Decode(&err.OpOutcome)
	if marshalErr != nil {
		log.Debug().Err(marshalErr).Msg("unexpected marshal err")
	}

	return err
}

func (e FhirError) Error() string {
	return e.OpOutcome.String()
}

var (
	ErrNotFound = errors.New("not found")
	ErrGone     = errors.New("gone")
	ErrUnknown  = errors.New("unknown error")
)

var statusCodeToError = map[int]error{
	http.StatusNotFound: ErrNotFound,
	http.StatusGone:     ErrGone,
}

func IsUnexpectedError(statusCode int, whitelist ...int) error {
	for _, whitelisted := range whitelist {
		if whitelisted == statusCode {
			return nil
		}
	}

	if err, ok := statusCodeToError[statusCode]; ok {
		return err
	}

	if statusCode > 200 || 300 >= statusCode {
		return ErrUnknown
	}

	return nil
}
