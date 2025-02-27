// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Signature
type Signature struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Indication of the reason the entity signed the object(s)
	Type []Coding `json:"type,omitempty"`
	// When the signature was created
	When *custom.Instant `json:"when,omitempty"`
	// Who signed
	Who *custom.Reference[SignatureWho] `json:"who,omitempty"`
	// The party represented
	OnBehalfOf *custom.Reference[SignatureOnBehalfOf] `json:"onBehalfOf,omitempty"`
	// The technical format of the signed resources
	TargetFormat *custom.Code `json:"targetFormat,omitempty"`
	// The technical format of the signature
	SigFormat *custom.Code `json:"sigFormat,omitempty"`
	// The actual signature content (XML DigSig. JWS, picture, etc.)
	Data *custom.Base64binary `json:"data,omitempty"`
}

type OtherSignature Signature

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
	return "Signature"
}

func (s Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSignature
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSignature: OtherSignature(s), ResourceType: s.ResourceType()})
}

func UnmarshalSignature(b []byte) (res Signature, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
