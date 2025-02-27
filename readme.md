
# go-fhir-client

**⚠️ Work in Progress: This library is currently under active development and is not yet functional for production use. APIs and functionality are subject to change without notice.**

This is a Go client for interacting with FHIR (Fast Healthcare Interoperability Resources) servers.

## Features (Planned/In Progress)

### Todo:
- [ ] **Versioning:** Implement semantic versioning and tagging for releases.
- [ ] **HTTP Client:**
    - [ ] **Search:** Add comprehensive search parameter support, including composite parameters.
    - [ ] **Constraints:** Enforce FHIR constraints defined in StructureDefinitions.
    - [ ] **Validation:** Implement resource validation against FHIR schemas and profiles.
- [ ] **Authentication/Authorization:** Add support for OAuth2, SMART on FHIR, and custom authentication providers.
- [ ] **Testing:** Expand unit and integration tests, including mocking and performance testing.
- [ ] **Documentation:** Provide detailed API documentation, usage guides, and error case documentation.
- [ ] **Code Generation Rewrite:** Improve code generation for better maintainability and flexibility.

### Future Features:
- **Batch/Transaction Operations:** Implement `$batch` and `$transaction` operations for efficient handling of multiple requests.
- **FHIRPath Evaluation:** Implement FHIRPath evaluation for data extraction and filtering.
- **Terminology Services:** Implement operations like `$validate-code`, `$lookup`, and `$expand`.
- **Data Streaming/WebSockets:** Implement FHIR subscriptions using WebSockets for real-time data updates.
- **Smart On FHIR:** Implement support for SMART on FHIR authorization flows and launch contexts.
- **Message Definition Support:** Implement support for FHIR messaging and MessageDefinition resources.
- **Questionnaire Resource Support:** Implement support for Questionnaire and QuestionnaireResponse resources.

## Installation

```bash
go get github.com/jonasrothmann/go-fhir-client
```

## Usage
```go
import (
	"log"
	"github.com/jonasrothmann/go-fhir-client/r5"
	"github.com/jonasrothmann/go-fhir-client/r5/valuesets"
)

func main() {
	client, err := fhirclient.NewClient(
		"https://hapi.fhir.org/baseR5",
		fhirclient.WithBasicAuthenticator("username", "password"),
	)
	if err != nil {
		log.Fatal(err)
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
	if err != nil {
		log.Fatalf("failed to create patient: %s", err)
	}

	readPatient, err := query.Read[r5.Patient](context.Background(), client, *patient.Id, nil)
	if err != nil {
		log.Fatalf("failed to read patient: %s", err)
	}

	fmt.Printf("created patient with name %s", *readPatient.Name[0].Text)
}
```

## Contributing
Contributions are welcome, but expect that the whole api might change at any moment.
