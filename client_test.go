package fhirclient

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHostNames(t *testing.T) {
	tests := []struct {
		host    string
		expects string
		errors  bool
	}{
		{"", "", true},
		{"https://hapi.fhir.org/baseR5", "https://hapi.fhir.org/baseR5", false},
		{"https://hapi.fhir.org/baseR5/", "https://hapi.fhir.org/baseR5", false},
		{"hapi.fhir.org/baseR5/", "https://hapi.fhir.org/baseR5", false},
	}

	for _, tst := range tests {
		client, err := NewClient(tst.host)
		require.NoError(t, err)
		require.Equal(t, tst.expects, client.HttpClient.Host)
	}
}
