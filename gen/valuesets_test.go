package main

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/jonasrothmann/go-fhir-client/custom"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

var valuesets = map[string]ValueSet{}

func init() {
	file, err := os.Open("./fhir/valuesets.json")
	if err != nil {
		log.Fatal().Err(err)
	}

	var bundle Bundle
	err = json.NewDecoder(file).Decode(&bundle)
	if err != nil {
		log.Fatal().Err(err)
	}

	for _, entry := range bundle.Entry {
		var vs ValueSet
		json.Unmarshal(*entry.Resource, &vs)
		valuesets[*entry.FullUrl] = vs
	}

	if len(valuesets) == 0 {
		log.Fatal().Msg("empty valuesets bundle")
	}
}

func TestGetUrl(t *testing.T) {
	tests := []struct {
		has     string
		expects string
	}{
		{"http://hl7.org/fhir/CodeSystem/administration-subpotent-reason", "http://hl7.org/fhir/codesystem-administration-subpotent-reason.json"},
	}

	for _, tst := range tests {
		result, err := getUrl(custom.URI(tst.has))
		require.NoError(t, err)
		require.Equal(t, result, tst.expects)
	}
}

type TestEnum struct {
	Name         string
	OptionValues []string
}

func TestGenerateEnum(t *testing.T) {
	tests := []struct {
		fullUrl   string
		wantsEnum TestEnum
	}{
		{
			"http://hl7.org/fhir/ValueSet/marital-status",
			TestEnum{
				Name:         "MaritalStatusCodes",
				OptionValues: []string{"A", "D", "I", "L", "M", "P", "T", "U", "W", "UNK"},
			},
		},
		{
			"http://hl7.org/fhir/ValueSet/account-type",
			TestEnum{
				Name:         "MaritalStatusCodes",
				OptionValues: []string{"_ActAccountCode", "ACCTRECEIVABLE", "CASH", "CC", "PBILLACCT", "_CreditCard", "AE", "DN", "DV", "MC", "V"},
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	for _, tst := range tests {
		vs, ok := valuesets[tst.fullUrl]
		require.True(t, ok)

		e, err := vs.generateEnum(ctx)
		require.NoError(t, err)

		optionValues := []string{}
		for _, opt := range e.Options {
			optionValues = append(optionValues, opt.Value)
		}
		require.Equal(t, optionValues, tst.wantsEnum.OptionValues)
	}
}

func TestNameAndValue(t *testing.T) {
	tests := []struct {
		value string
		name  string
		wants string
	}{
		{
			value: "recall",
			name:  "Parent/Guardian/Patient Recall",
			wants: "Recall",
		},
		{
			value: "UKN",
			name:  "Unknown",
			wants: "Unknown",
		},
		{
			value: "RCL",
			name:  "Parent/Guardian/Patient Recall",
			wants: "ParentGuardianPatientRecall",
		},
		{
			// Abbreviation: split on spaces.
			value: "DOB",
			name:  "Date of Birth",
			wants: "DateOfBirth",
		},
		{
			// Not an abbreviation: simply capitalize the value.
			value: "gender",
			name:  "Gender",
			wants: "Gender",
		},
		{
			// Abbreviation but the full name is a single word.
			value: "M",
			name:  "Male",
			wants: "Male",
		},
		{
			// Abbreviation: handle multiple delimiters (hyphen and slash).
			value: "RCL",
			name:  "Parent-Guardian/Patient Recall",
			wants: "ParentGuardianPatientRecall",
		},
	}

	for _, tst := range tests {
		name := formatNameAndValue(tst.name, tst.value)
		require.Equal(t, name, tst.wants)
	}
}
