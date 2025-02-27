// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CapabilityStatement
type CapabilityStatementSoftware struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A name the software is known by
	Name string `json:"name"`
	// Version covered by this statement
	Version *string `json:"version,omitempty"`
	// Date this version was released
	ReleaseDate *custom.DateTime `json:"releaseDate,omitempty"`
}

type CapabilityStatementMessagingEndpoint struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// http | ftp | mllp +
	Protocol Coding `json:"protocol"`
	// Network address or identifier of the end-point
	Address custom.URL `json:"address"`
}

type CapabilityStatementMessagingSupportedMessage struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// sender | receiver
	Mode custom.Code `json:"mode"`
	// Message supported by this system
	Definition custom.Canonical[MessageDefinition] `json:"definition"`
}

type CapabilityStatementRest struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// client | server
	Mode custom.Code `json:"mode"`
	// General description of implementation
	Documentation *custom.Markdown `json:"documentation,omitempty"`
	// Information about security of implementation
	Security *CapabilityStatementRestSecurity `json:"security,omitempty"`
	// Resource served on the REST interface
	Resource []CapabilityStatementRestResource `json:"resource,omitempty"`
	// What operations are supported?
	Interaction []CapabilityStatementRestInteraction `json:"interaction,omitempty"`
	// Search parameters for searching all resources
	SearchParam []interface{} `json:"searchParam,omitempty"`
	// Definition of a system level operation
	Operation []interface{} `json:"operation,omitempty"`
	// Compartments served/used by system
	Compartment []custom.Canonical[CompartmentDefinition] `json:"compartment,omitempty"`
}

type CapabilityStatementRestSecurity struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Adds CORS Headers (http://enable-cors.org/)
	Cors *bool `json:"cors,omitempty"`
	// OAuth | SMART-on-FHIR | NTLM | Basic | Kerberos | Certificates
	Service []CodeableConcept `json:"service,omitempty"`
	// General description of how security works
	Description *custom.Markdown `json:"description,omitempty"`
}

type CapabilityStatementRestResource struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A resource type that is supported
	Type custom.Code `json:"type"`
	// System-wide profile
	Profile *custom.Canonical[StructureDefinition] `json:"profile,omitempty"`
	// Use-case specific profiles
	SupportedProfile []custom.Canonical[StructureDefinition] `json:"supportedProfile,omitempty"`
	// Additional information about the use of the resource type
	Documentation *custom.Markdown `json:"documentation,omitempty"`
	// What operations are supported?
	Interaction []CapabilityStatementRestResourceInteraction `json:"interaction,omitempty"`
	// no-version | versioned | versioned-update
	Versioning *custom.Code `json:"versioning,omitempty"`
	// Whether vRead can return past versions
	ReadHistory *bool `json:"readHistory,omitempty"`
	// If update can commit to a new identity
	UpdateCreate *bool `json:"updateCreate,omitempty"`
	// If allows/uses conditional create
	ConditionalCreate *bool `json:"conditionalCreate,omitempty"`
	// not-supported | modified-since | not-match | full-support
	ConditionalRead *custom.Code `json:"conditionalRead,omitempty"`
	// If allows/uses conditional update
	ConditionalUpdate *bool `json:"conditionalUpdate,omitempty"`
	// If allows/uses conditional patch
	ConditionalPatch *bool `json:"conditionalPatch,omitempty"`
	// not-supported | single | multiple - how conditional delete is supported
	ConditionalDelete *custom.Code `json:"conditionalDelete,omitempty"`
	// literal | logical | resolves | enforced | local
	ReferencePolicy []custom.Code `json:"referencePolicy,omitempty"`
	// _include values supported by the server
	SearchInclude []string `json:"searchInclude,omitempty"`
	// _revinclude values supported by the server
	SearchRevInclude []string `json:"searchRevInclude,omitempty"`
	// Search parameters supported by implementation
	SearchParam []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
	// Definition of a resource operation
	Operation []CapabilityStatementRestResourceOperation `json:"operation,omitempty"`
}

type CapabilityStatementRestResourceSearchParam struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name for parameter in search url
	Name string `json:"name"`
	// Source of definition for parameter
	Definition *custom.Canonical[SearchParameter] `json:"definition,omitempty"`
	// number | date | string | token | reference | composite | quantity | uri | special
	Type custom.Code `json:"type"`
	// Server-specific usage
	Documentation *custom.Markdown `json:"documentation,omitempty"`
}

type CapabilityStatementMessaging struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Where messages should be sent
	Endpoint []CapabilityStatementMessagingEndpoint `json:"endpoint,omitempty"`
	// Reliable Message Cache Length (min)
	ReliableCache *uint32 `json:"reliableCache,omitempty"`
	// Messaging interface behavior details
	Documentation *custom.Markdown `json:"documentation,omitempty"`
	// Messages supported by this system
	SupportedMessage []CapabilityStatementMessagingSupportedMessage `json:"supportedMessage,omitempty"`
}

type CapabilityStatementImplementation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Describes this specific instance
	Description custom.Markdown `json:"description"`
	// Base URL for the installation
	Url *custom.URL `json:"url,omitempty"`
	// Organization that manages the data
	Custodian *custom.Reference[Organization] `json:"custodian,omitempty"`
}

type CapabilityStatement struct {
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
	// Canonical identifier for this capability statement, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the CapabilityStatement (business identifier)
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the capability statement
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this capability statement (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this capability statement (human friendly)
	Title *string `json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Date last changed
	Date custom.DateTime `json:"date"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the capability statement
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for capability statement (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this capability statement is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// instance | capability | requirements
	Kind custom.Code `json:"kind"`
	// Canonical URL of another capability statement this implements
	Instantiates []custom.Canonical[CapabilityStatement] `json:"instantiates,omitempty"`
	// Canonical URL of another capability statement this adds to
	Imports []custom.Canonical[CapabilityStatement] `json:"imports,omitempty"`
	// Software that is covered by this capability statement
	Software *CapabilityStatementSoftware `json:"software,omitempty"`
	// If this describes a specific instance
	Implementation *CapabilityStatementImplementation `json:"implementation,omitempty"`
	// FHIR Version the system supports
	FhirVersion custom.Code `json:"fhirVersion"`
	// formats supported (xml | json | ttl | mime type)
	Format []custom.Code `json:"format"`
	// Patch formats supported
	PatchFormat []custom.Code `json:"patchFormat,omitempty"`
	// Languages supported
	AcceptLanguage []custom.Code `json:"acceptLanguage,omitempty"`
	// Implementation guides supported
	ImplementationGuide []custom.Canonical[ImplementationGuide] `json:"implementationGuide,omitempty"`
	// If the endpoint is a RESTful one
	Rest []CapabilityStatementRest `json:"rest,omitempty"`
	// If messaging is supported
	Messaging []CapabilityStatementMessaging `json:"messaging,omitempty"`
	// Document definition
	Document []CapabilityStatementDocument `json:"document,omitempty"`
}

type CapabilityStatementRestResourceInteraction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// read | vread | update | patch | delete | history-instance | history-type | create | search-type
	Code custom.Code `json:"code"`
	// Anything special about operation behavior
	Documentation *custom.Markdown `json:"documentation,omitempty"`
}

type CapabilityStatementRestResourceOperation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name by which the operation/query is invoked
	Name string `json:"name"`
	// The defined operation/query
	Definition custom.Canonical[OperationDefinition] `json:"definition"`
	// Specific details about operation behavior
	Documentation *custom.Markdown `json:"documentation,omitempty"`
}

type CapabilityStatementRestInteraction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// transaction | batch | search-system | history-system
	Code custom.Code `json:"code"`
	// Anything special about operation behavior
	Documentation *custom.Markdown `json:"documentation,omitempty"`
}

type CapabilityStatementDocument struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// producer | consumer
	Mode custom.Code `json:"mode"`
	// Description of document support
	Documentation *custom.Markdown `json:"documentation,omitempty"`
	// Constraint on the resources used in the document
	Profile custom.Canonical[StructureDefinition] `json:"profile"`
}

type OtherCapabilityStatement CapabilityStatement

func (c CapabilityStatement) ResourceType() string {
	return "CapabilityStatement"
}

func (c CapabilityStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCapabilityStatement
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCapabilityStatement: OtherCapabilityStatement(c), ResourceType: c.ResourceType()})
}

func UnmarshalCapabilityStatement(b []byte) (res CapabilityStatement, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
