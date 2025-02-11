package tests

import (
	"context"
	"testing"
	"time"

	"github.com/jonasrothmann/go-fhir-client/query"
	"github.com/jonasrothmann/go-fhir-client/resources"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestHAPI(t *testing.T) {
	client := HAPIClient

	tests := []struct {
		patient       resources.Patient
		expectedError error
	}{
		{
			patient: resources.Patient{
				Active: true,
			},
			expectedError: nil,
		},
	}

	for _, tst := range tests {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		createdPatient, err := query.Create(ctx, client, tst.patient)
		require.NoError(t, err)
		require.NotNil(t, createdPatient)

		log.Printf("%+v", createdPatient)

		readPatient, err := query.Read[resources.Patient](ctx, client, createdPatient.ID, nil)
		require.NoError(t, err)
		require.NotNil(t, readPatient)
		require.Equal(t, readPatient, createdPatient)
	}

	// query.Search returns SearchQuery[resources.Patient]

	// Where returns a WhereParams which can be encoded into a url params

	// execute = takes a url params and a resource, queries the api & maps to type

	/*
		patients, err := queryalt.Search[resources.Patient](query).Execute(ctx, client)

		client.Search(
			patientQuery{
				Active: queryalt.Bool{
					Equals: true,
				},
			},
		)

		require.NoError(t, err)
		require.NotNil(t, patients)*/

	//fmt.Printf("%+v", patient)
}
