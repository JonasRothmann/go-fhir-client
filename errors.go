package fhirclient

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	datatypes "github.com/jonasrothmann/go-fhir-client/data-types"
	"github.com/rs/zerolog/log"
)

type FhirError struct {
	//Message   string
	OpOutcome datatypes.OperationOutcome
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
