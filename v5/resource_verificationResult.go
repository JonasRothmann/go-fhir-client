// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/VerificationResult
type VerificationResultPrimarySource struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to the primary source
	Who *custom.Reference[VerificationResultPrimarySourceWho] `bson:"who,omitempty" json:"who,omitempty"`
	// Type of primary source (License Board; Primary Education; Continuing Education; Postal Service; Relationship owner; Registration Authority; legal source; issuing source; authoritative source)
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Method for exchanging information with the primary source
	CommunicationMethod []CodeableConcept `bson:"communicationMethod,omitempty" json:"communicationMethod,omitempty"`
	// successful | failed | unknown
	ValidationStatus *CodeableConcept `bson:"validationStatus,omitempty" json:"validationStatus,omitempty"`
	// When the target was validated against the primary source
	ValidationDate *custom.DateTime `bson:"validationDate,omitempty" json:"validationDate,omitempty"`
	// yes | no | undetermined
	CanPushUpdates *CodeableConcept `bson:"canPushUpdates,omitempty" json:"canPushUpdates,omitempty"`
	// specific | any | source
	PushTypeAvailable []CodeableConcept `bson:"pushTypeAvailable,omitempty" json:"pushTypeAvailable,omitempty"`
}

type VerificationResultAttestation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The individual or organization attesting to information
	Who *custom.Reference[VerificationResultAttestationWho] `bson:"who,omitempty" json:"who,omitempty"`
	// When the who is asserting on behalf of another (organization or individual)
	OnBehalfOf *custom.Reference[VerificationResultAttestationOnBehalfOf] `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	// The method by which attested information was submitted/retrieved
	CommunicationMethod *CodeableConcept `bson:"communicationMethod,omitempty" json:"communicationMethod,omitempty"`
	// The date the information was attested to
	Date *custom.Date `bson:"date,omitempty" json:"date,omitempty"`
	// A digital identity certificate associated with the attestation source
	SourceIdentityCertificate *string `bson:"sourceIdentityCertificate,omitempty" json:"sourceIdentityCertificate,omitempty"`
	// A digital identity certificate associated with the proxy entity submitting attested information on behalf of the attestation source
	ProxyIdentityCertificate *string `bson:"proxyIdentityCertificate,omitempty" json:"proxyIdentityCertificate,omitempty"`
	// Proxy signature (digital or image)
	ProxySignature *Signature `bson:"proxySignature,omitempty" json:"proxySignature,omitempty"`
	// Attester signature (digital or image)
	SourceSignature *Signature `bson:"sourceSignature,omitempty" json:"sourceSignature,omitempty"`
}

type VerificationResultValidator struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to the organization validating information
	Organization custom.Reference[Organization] `bson:"organization" json:"organization"`
	// A digital identity certificate associated with the validator
	IdentityCertificate *string `bson:"identityCertificate,omitempty" json:"identityCertificate,omitempty"`
	// Validator signature (digital or image)
	AttestationSignature *Signature `bson:"attestationSignature,omitempty" json:"attestationSignature,omitempty"`
}

type VerificationResult struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `bson:"contained,omitempty" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A resource that was validated
	Target []custom.Reference[Resource] `bson:"target,omitempty" json:"target,omitempty"`
	// The fhirpath location(s) within the resource that was validated
	TargetLocation []string `bson:"targetLocation,omitempty" json:"targetLocation,omitempty"`
	// none | initial | periodic
	Need *CodeableConcept `bson:"need,omitempty" json:"need,omitempty"`
	// attested | validated | in-process | req-revalid | val-fail | reval-fail | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// When the validation status was updated
	StatusDate *custom.DateTime `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	// nothing | primary | multiple
	ValidationType *CodeableConcept `bson:"validationType,omitempty" json:"validationType,omitempty"`
	// The primary process by which the target is validated (edit check; value set; primary source; multiple sources; standalone; in context)
	ValidationProcess []CodeableConcept `bson:"validationProcess,omitempty" json:"validationProcess,omitempty"`
	// Frequency of revalidation
	Frequency *Timing `bson:"frequency,omitempty" json:"frequency,omitempty"`
	// The date/time validation was last completed (including failed validations)
	LastPerformed *custom.DateTime `bson:"lastPerformed,omitempty" json:"lastPerformed,omitempty"`
	// The date when target is next validated, if appropriate
	NextScheduled *custom.Date `bson:"nextScheduled,omitempty" json:"nextScheduled,omitempty"`
	// fatal | warn | rec-only | none
	FailureAction *CodeableConcept `bson:"failureAction,omitempty" json:"failureAction,omitempty"`
	// Information about the primary source(s) involved in validation
	PrimarySource []VerificationResultPrimarySource `bson:"primarySource,omitempty" json:"primarySource,omitempty"`
	// Information about the entity attesting to information
	Attestation *VerificationResultAttestation `bson:"attestation,omitempty" json:"attestation,omitempty"`
	// Information about the entity validating information
	Validator []VerificationResultValidator `bson:"validator,omitempty" json:"validator,omitempty"`
}

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
	return "StructureDefinition"
}
