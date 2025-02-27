// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EnrollmentRequest
type EnrollmentRequest struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Business Identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status *custom.Code `json:"status,omitempty"`
	// Creation date
	Created *custom.DateTime `json:"created,omitempty"`
	// Target
	Insurer *custom.Reference[Organization] `json:"insurer,omitempty"`
	// Responsible practitioner
	Provider *custom.Reference[EnrollmentRequestProvider] `json:"provider,omitempty"`
	// The subject to be enrolled
	Candidate *custom.Reference[Patient] `json:"candidate,omitempty"`
	// Insurance information
	Coverage *custom.Reference[Coverage] `json:"coverage,omitempty"`
}

type OtherEnrollmentRequest EnrollmentRequest

type EnrollmentRequestProvider interface {
	gofhirclient.Resource

	Is_EnrollmentRequestProvider()
}

func (p Practitioner) Is_EnrollmentRequestProvider()     {}
func (p PractitionerRole) Is_EnrollmentRequestProvider() {}
func (o Organization) Is_EnrollmentRequestProvider()     {}

func (e EnrollmentRequest) ResourceType() string {
	return "EnrollmentRequest"
}

func (e EnrollmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentRequest
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherEnrollmentRequest: OtherEnrollmentRequest(e), ResourceType: e.ResourceType()})
}

func UnmarshalEnrollmentRequest(b []byte) (res EnrollmentRequest, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
