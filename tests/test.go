package tests

import (
	fhirclient "github.com/jonasrothmann/go-fhir-client"
)

var HAPIClient = fhirclient.Must(
	fhirclient.NewClient(
		"https://hapi.fhir.org/baseR5",
		//fhirclient.WithBasicAuthenticator("golang-fhir-client", "password567"),
	),
)
