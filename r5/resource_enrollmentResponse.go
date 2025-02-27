// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EnrollmentResponse
type EnrollmentResponse struct {
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
	// Claim reference
	Request *custom.Reference[EnrollmentRequest] `json:"request,omitempty"`
	// queued | complete | error | partial
	Outcome *custom.Code `json:"outcome,omitempty"`
	// Disposition Message
	Disposition *string `json:"disposition,omitempty"`
	// Creation date
	Created *custom.DateTime `json:"created,omitempty"`
	// Insurer
	Organization *custom.Reference[Organization] `json:"organization,omitempty"`
	// Responsible practitioner
	RequestProvider *custom.Reference[EnrollmentResponseRequestProvider] `json:"requestProvider,omitempty"`
}

type OtherEnrollmentResponse EnrollmentResponse

type EnrollmentResponseRequestProvider interface {
	gofhirclient.Resource

	Is_EnrollmentResponseRequestProvider()
}

func (p Practitioner) Is_EnrollmentResponseRequestProvider()     {}
func (p PractitionerRole) Is_EnrollmentResponseRequestProvider() {}
func (o Organization) Is_EnrollmentResponseRequestProvider()     {}

func (e EnrollmentResponse) ResourceType() string {
	return "EnrollmentResponse"
}

func (e EnrollmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentResponse
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherEnrollmentResponse: OtherEnrollmentResponse(e), ResourceType: e.ResourceType()})
}

func UnmarshalEnrollmentResponse(b []byte) (res EnrollmentResponse, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
