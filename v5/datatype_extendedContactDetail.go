// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/ExtendedContactDetail
type ExtendedContactDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// The type of contact
	Purpose *CodeableConcept `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Name of an individual to contact
	Name []HumanName `bson:"name,omitempty" json:"name,omitempty"`
	// Contact details (e.g.phone/fax/url)
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
	// Address for the contact
	Address *Address `bson:"address,omitempty" json:"address,omitempty"`
	// This contact detail is handled/monitored by a specific organization
	Organization *custom.Reference[Organization] `bson:"organization,omitempty" json:"organization,omitempty"`
	// Period that this contact was valid for usage
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

func (e ExtendedContactDetail) ResourceType() string {
	return "StructureDefinition"
}
