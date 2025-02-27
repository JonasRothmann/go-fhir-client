// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/VerificationResult
type VerificationResult struct {
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
	// A resource that was validated
	Target []custom.Reference[Resource] `json:"target,omitempty"`
	// The fhirpath location(s) within the resource that was validated
	TargetLocation []string `json:"targetLocation,omitempty"`
	// none | initial | periodic
	Need *CodeableConcept `json:"need,omitempty"`
	// attested | validated | in-process | req-revalid | val-fail | reval-fail | entered-in-error
	Status custom.Code `json:"status"`
	// When the validation status was updated
	StatusDate *custom.DateTime `json:"statusDate,omitempty"`
	// nothing | primary | multiple
	ValidationType *CodeableConcept `json:"validationType,omitempty"`
	// The primary process by which the target is validated (edit check; value set; primary source; multiple sources; standalone; in context)
	ValidationProcess []CodeableConcept `json:"validationProcess,omitempty"`
	// Frequency of revalidation
	Frequency *Timing `json:"frequency,omitempty"`
	// The date/time validation was last completed (including failed validations)
	LastPerformed *custom.DateTime `json:"lastPerformed,omitempty"`
	// The date when target is next validated, if appropriate
	NextScheduled *custom.Date `json:"nextScheduled,omitempty"`
	// fatal | warn | rec-only | none
	FailureAction *CodeableConcept `json:"failureAction,omitempty"`
	// Information about the primary source(s) involved in validation
	PrimarySource []VerificationResultPrimarySource `json:"primarySource,omitempty"`
	// Information about the entity attesting to information
	Attestation *VerificationResultAttestation `json:"attestation,omitempty"`
	// Information about the entity validating information
	Validator []VerificationResultValidator `json:"validator,omitempty"`
}

type VerificationResultPrimarySource struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to the primary source
	Who *custom.Reference[VerificationResultPrimarySourceWho] `json:"who,omitempty"`
	// Type of primary source (License Board; Primary Education; Continuing Education; Postal Service; Relationship owner; Registration Authority; legal source; issuing source; authoritative source)
	Type []CodeableConcept `json:"type,omitempty"`
	// Method for exchanging information with the primary source
	CommunicationMethod []CodeableConcept `json:"communicationMethod,omitempty"`
	// successful | failed | unknown
	ValidationStatus *CodeableConcept `json:"validationStatus,omitempty"`
	// When the target was validated against the primary source
	ValidationDate *custom.DateTime `json:"validationDate,omitempty"`
	// yes | no | undetermined
	CanPushUpdates *CodeableConcept `json:"canPushUpdates,omitempty"`
	// specific | any | source
	PushTypeAvailable []CodeableConcept `json:"pushTypeAvailable,omitempty"`
}

type VerificationResultAttestation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The individual or organization attesting to information
	Who *custom.Reference[VerificationResultAttestationWho] `json:"who,omitempty"`
	// When the who is asserting on behalf of another (organization or individual)
	OnBehalfOf *custom.Reference[VerificationResultAttestationOnBehalfOf] `json:"onBehalfOf,omitempty"`
	// The method by which attested information was submitted/retrieved
	CommunicationMethod *CodeableConcept `json:"communicationMethod,omitempty"`
	// The date the information was attested to
	Date *custom.Date `json:"date,omitempty"`
	// A digital identity certificate associated with the attestation source
	SourceIdentityCertificate *string `json:"sourceIdentityCertificate,omitempty"`
	// A digital identity certificate associated with the proxy entity submitting attested information on behalf of the attestation source
	ProxyIdentityCertificate *string `json:"proxyIdentityCertificate,omitempty"`
	// Proxy signature (digital or image)
	ProxySignature *Signature `json:"proxySignature,omitempty"`
	// Attester signature (digital or image)
	SourceSignature *Signature `json:"sourceSignature,omitempty"`
}

type VerificationResultValidator struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to the organization validating information
	Organization custom.Reference[Organization] `json:"organization"`
	// A digital identity certificate associated with the validator
	IdentityCertificate *string `json:"identityCertificate,omitempty"`
	// Validator signature (digital or image)
	AttestationSignature *Signature `json:"attestationSignature,omitempty"`
}

type OtherVerificationResult VerificationResult

type VerificationResultPrimarySourceWho interface {
	gofhirclient.Resource

	Is_VerificationResultPrimarySourceWho()
}

func (o Organization) Is_VerificationResultPrimarySourceWho()     {}
func (p Practitioner) Is_VerificationResultPrimarySourceWho()     {}
func (p PractitionerRole) Is_VerificationResultPrimarySourceWho() {}

type VerificationResultAttestationWho interface {
	gofhirclient.Resource

	Is_VerificationResultAttestationWho()
}

func (p Practitioner) Is_VerificationResultAttestationWho()     {}
func (p PractitionerRole) Is_VerificationResultAttestationWho() {}
func (o Organization) Is_VerificationResultAttestationWho()     {}

type VerificationResultAttestationOnBehalfOf interface {
	gofhirclient.Resource

	Is_VerificationResultAttestationOnBehalfOf()
}

func (o Organization) Is_VerificationResultAttestationOnBehalfOf()     {}
func (p Practitioner) Is_VerificationResultAttestationOnBehalfOf()     {}
func (p PractitionerRole) Is_VerificationResultAttestationOnBehalfOf() {}

func (v VerificationResult) ResourceType() string {
	return "VerificationResult"
}

func (v VerificationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVerificationResult
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherVerificationResult: OtherVerificationResult(v), ResourceType: v.ResourceType()})
}

func UnmarshalVerificationResult(b []byte) (res VerificationResult, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
