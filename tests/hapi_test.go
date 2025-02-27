package tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	fhirclient "github.com/jonasrothmann/go-fhir-client"
	"github.com/jonasrothmann/go-fhir-client/query"
	"github.com/jonasrothmann/go-fhir-client/r5"
	"github.com/jonasrothmann/go-fhir-client/r5/valuesets"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestHAPI(t *testing.T) {
	client := HAPIClient

	tests := []struct {
		patient       r5.Patient
		expectedError error
	}{
		{
			patient: r5.Patient{
				Active: fhirclient.Ptr(true),
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

		readPatient, err := query.Read[r5.Patient](ctx, client, *createdPatient.Id, nil)
		require.NoError(t, err)
		require.NotNil(t, readPatient)
		require.Equal(t, readPatient, createdPatient)

		readPatient.Name[0].Text
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

func main() {
	client, err := fhirclient.NewClient(
		"https://hapi.fhir.org/baseR5",
		fhirclient.WithBasicAuthenticator("username", "password"),
	)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	input := r5.Patient{
		Name: []r5.HumanName{
			{
				Use:    valuesets.NameUseOfficial,
				Family: fhirclient.Ptr("Smith"),
				Given:  []string{"Timothy", "James"},
			},
			{
				Use:   valuesets.NameUseUsual,
				Given: []string{"Timmy"},
			},
		},
		Gender: fhirclient.Ptr(valuesets.AdministrativeGenderMale),
	}
	patient, err := query.Create(context.Background(), client, input)

	readPatient, err := query.Read[r5.Patient](context.Background(), client, *patient.Id, nil)

	fmt.Printf("created patient with name %s", *readPatient.Name[0].Text)
}
