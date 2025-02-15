// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Signature
type Signature struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Indication of the reason the entity signed the object(s)
	Type []Coding `bson:"type,omitempty" json:"type,omitempty"`
	// When the signature was created
	When *custom.Instant `bson:"when,omitempty" json:"when,omitempty"`
	// Who signed
	Who *custom.Reference[SignatureWho] `bson:"who,omitempty" json:"who,omitempty"`
	// The party represented
	OnBehalfOf *custom.Reference[SignatureOnBehalfOf] `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	// The technical format of the signed resources
	TargetFormat *custom.Code `bson:"targetFormat,omitempty" json:"targetFormat,omitempty"`
	// The technical format of the signature
	SigFormat *custom.Code `bson:"sigFormat,omitempty" json:"sigFormat,omitempty"`
	// The actual signature content (XML DigSig. JWS, picture, etc.)
	Data *custom.Base64binary `bson:"data,omitempty" json:"data,omitempty"`
}

type SignatureWho interface {
	gofhirclient.Resource

	Is_SignatureWho()
}

func (p Practitioner) Is_SignatureWho()     {}
func (p PractitionerRole) Is_SignatureWho() {}
func (r RelatedPerson) Is_SignatureWho()    {}
func (p Patient) Is_SignatureWho()          {}
func (d Device) Is_SignatureWho()           {}
func (o Organization) Is_SignatureWho()     {}

type SignatureOnBehalfOf interface {
	gofhirclient.Resource

	Is_SignatureOnBehalfOf()
}

func (p Practitioner) Is_SignatureOnBehalfOf()     {}
func (p PractitionerRole) Is_SignatureOnBehalfOf() {}
func (r RelatedPerson) Is_SignatureOnBehalfOf()    {}
func (p Patient) Is_SignatureOnBehalfOf()          {}
func (d Device) Is_SignatureOnBehalfOf()           {}
func (o Organization) Is_SignatureOnBehalfOf()     {}

func (s Signature) ResourceType() string {
	return "StructureDefinition"
}
