package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFilters(t *testing.T) {
	vs, ok := valuesets["http://hl7.org/fhir/ValueSet/account-type"]
	require.True(t, ok)
	require.NotNil(t, vs.Name)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	codesystems, err := vs.getCodeSystems(ctx)
	require.NoError(t, err)
	require.Greater(t, len(codesystems), 0)

	for _, codesystem := range codesystems {
		codesystem.Concept = filterConcepts(codesystem.Concept, codesystem.ValueSetFilters)

	}
}
