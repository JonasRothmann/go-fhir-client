// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/CapabilityStatement
type CapabilityStatementRestResourceOperation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name by which the operation/query is invoked
	Name string `bson:"name" json:"name"`
	// The defined operation/query
	Definition custom.Canonical[OperationDefinition] `bson:"definition" json:"definition"`
	// Specific details about operation behavior
	Documentation *custom.Markdown `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type CapabilityStatementMessaging struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Where messages should be sent
	Endpoint []CapabilityStatementMessagingEndpoint `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	// Reliable Message Cache Length (min)
	ReliableCache *uint32 `bson:"reliableCache,omitempty" json:"reliableCache,omitempty"`
	// Messaging interface behavior details
	Documentation *custom.Markdown `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// Messages supported by this system
	SupportedMessage []CapabilityStatementMessagingSupportedMessage `bson:"supportedMessage,omitempty" json:"supportedMessage,omitempty"`
}

type CapabilityStatementMessagingSupportedMessage struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// sender | receiver
	Mode custom.Code `bson:"mode" json:"mode"`
	// Message supported by this system
	Definition custom.Canonical[MessageDefinition] `bson:"definition" json:"definition"`
}

type CapabilityStatementSoftware struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A name the software is known by
	Name string `bson:"name" json:"name"`
	// Version covered by this statement
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// Date this version was released
	ReleaseDate *custom.DateTime `bson:"releaseDate,omitempty" json:"releaseDate,omitempty"`
}

type CapabilityStatementRest struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// client | server
	Mode custom.Code `bson:"mode" json:"mode"`
	// General description of implementation
	Documentation *custom.Markdown `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// Information about security of implementation
	Security *CapabilityStatementRestSecurity `bson:"security,omitempty" json:"security,omitempty"`
	// Resource served on the REST interface
	Resource []CapabilityStatementRestResource `bson:"resource,omitempty" json:"resource,omitempty"`
	// What operations are supported?
	Interaction []CapabilityStatementRestInteraction `bson:"interaction,omitempty" json:"interaction,omitempty"`
	// Search parameters for searching all resources
	SearchParam []interface{} `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	// Definition of a system level operation
	Operation []interface{} `bson:"operation,omitempty" json:"operation,omitempty"`
	// Compartments served/used by system
	Compartment []custom.Canonical[CompartmentDefinition] `bson:"compartment,omitempty" json:"compartment,omitempty"`
}

type CapabilityStatementRestResourceInteraction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// read | vread | update | patch | delete | history-instance | history-type | create | search-type
	Code custom.Code `bson:"code" json:"code"`
	// Anything special about operation behavior
	Documentation *custom.Markdown `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type CapabilityStatementRestInteraction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// transaction | batch | search-system | history-system
	Code custom.Code `bson:"code" json:"code"`
	// Anything special about operation behavior
	Documentation *custom.Markdown `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type CapabilityStatementDocument struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// producer | consumer
	Mode custom.Code `bson:"mode" json:"mode"`
	// Description of document support
	Documentation *custom.Markdown `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// Constraint on the resources used in the document
	Profile custom.Canonical[StructureDefinition] `bson:"profile" json:"profile"`
}

type CapabilityStatement struct {
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
	// Canonical identifier for this capability statement, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the CapabilityStatement (business identifier)
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the capability statement
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this capability statement (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this capability statement (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date custom.DateTime `bson:"date" json:"date"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the capability statement
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for capability statement (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this capability statement is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// instance | capability | requirements
	Kind custom.Code `bson:"kind" json:"kind"`
	// Canonical URL of another capability statement this implements
	Instantiates []custom.Canonical[CapabilityStatement] `bson:"instantiates,omitempty" json:"instantiates,omitempty"`
	// Canonical URL of another capability statement this adds to
	Imports []custom.Canonical[CapabilityStatement] `bson:"imports,omitempty" json:"imports,omitempty"`
	// Software that is covered by this capability statement
	Software *CapabilityStatementSoftware `bson:"software,omitempty" json:"software,omitempty"`
	// If this describes a specific instance
	Implementation *CapabilityStatementImplementation `bson:"implementation,omitempty" json:"implementation,omitempty"`
	// FHIR Version the system supports
	FhirVersion custom.Code `bson:"fhirVersion" json:"fhirVersion"`
	// formats supported (xml | json | ttl | mime type)
	Format []custom.Code `bson:"format" json:"format"`
	// Patch formats supported
	PatchFormat []custom.Code `bson:"patchFormat,omitempty" json:"patchFormat,omitempty"`
	// Languages supported
	AcceptLanguage []custom.Code `bson:"acceptLanguage,omitempty" json:"acceptLanguage,omitempty"`
	// Implementation guides supported
	ImplementationGuide []custom.Canonical[ImplementationGuide] `bson:"implementationGuide,omitempty" json:"implementationGuide,omitempty"`
	// If the endpoint is a RESTful one
	Rest []CapabilityStatementRest `bson:"rest,omitempty" json:"rest,omitempty"`
	// If messaging is supported
	Messaging []CapabilityStatementMessaging `bson:"messaging,omitempty" json:"messaging,omitempty"`
	// Document definition
	Document []CapabilityStatementDocument `bson:"document,omitempty" json:"document,omitempty"`
}

type CapabilityStatementRestResource struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A resource type that is supported
	Type custom.Code `bson:"type" json:"type"`
	// System-wide profile
	Profile *custom.Canonical[StructureDefinition] `bson:"profile,omitempty" json:"profile,omitempty"`
	// Use-case specific profiles
	SupportedProfile []custom.Canonical[StructureDefinition] `bson:"supportedProfile,omitempty" json:"supportedProfile,omitempty"`
	// Additional information about the use of the resource type
	Documentation *custom.Markdown `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// What operations are supported?
	Interaction []CapabilityStatementRestResourceInteraction `bson:"interaction,omitempty" json:"interaction,omitempty"`
	// no-version | versioned | versioned-update
	Versioning *custom.Code `bson:"versioning,omitempty" json:"versioning,omitempty"`
	// Whether vRead can return past versions
	ReadHistory *bool `bson:"readHistory,omitempty" json:"readHistory,omitempty"`
	// If update can commit to a new identity
	UpdateCreate *bool `bson:"updateCreate,omitempty" json:"updateCreate,omitempty"`
	// If allows/uses conditional create
	ConditionalCreate *bool `bson:"conditionalCreate,omitempty" json:"conditionalCreate,omitempty"`
	// not-supported | modified-since | not-match | full-support
	ConditionalRead *custom.Code `bson:"conditionalRead,omitempty" json:"conditionalRead,omitempty"`
	// If allows/uses conditional update
	ConditionalUpdate *bool `bson:"conditionalUpdate,omitempty" json:"conditionalUpdate,omitempty"`
	// If allows/uses conditional patch
	ConditionalPatch *bool `bson:"conditionalPatch,omitempty" json:"conditionalPatch,omitempty"`
	// not-supported | single | multiple - how conditional delete is supported
	ConditionalDelete *custom.Code `bson:"conditionalDelete,omitempty" json:"conditionalDelete,omitempty"`
	// literal | logical | resolves | enforced | local
	ReferencePolicy []custom.Code `bson:"referencePolicy,omitempty" json:"referencePolicy,omitempty"`
	// _include values supported by the server
	SearchInclude []string `bson:"searchInclude,omitempty" json:"searchInclude,omitempty"`
	// _revinclude values supported by the server
	SearchRevInclude []string `bson:"searchRevInclude,omitempty" json:"searchRevInclude,omitempty"`
	// Search parameters supported by implementation
	SearchParam []CapabilityStatementRestResourceSearchParam `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	// Definition of a resource operation
	Operation []CapabilityStatementRestResourceOperation `bson:"operation,omitempty" json:"operation,omitempty"`
}

type CapabilityStatementRestSecurity struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Adds CORS Headers (http://enable-cors.org/)
	Cors *bool `bson:"cors,omitempty" json:"cors,omitempty"`
	// OAuth | SMART-on-FHIR | NTLM | Basic | Kerberos | Certificates
	Service []CodeableConcept `bson:"service,omitempty" json:"service,omitempty"`
	// General description of how security works
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
}

type CapabilityStatementMessagingEndpoint struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// http | ftp | mllp +
	Protocol Coding `bson:"protocol" json:"protocol"`
	// Network address or identifier of the end-point
	Address custom.URL `bson:"address" json:"address"`
}

type CapabilityStatementImplementation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Describes this specific instance
	Description custom.Markdown `bson:"description" json:"description"`
	// Base URL for the installation
	Url *custom.URL `bson:"url,omitempty" json:"url,omitempty"`
	// Organization that manages the data
	Custodian *custom.Reference[Organization] `bson:"custodian,omitempty" json:"custodian,omitempty"`
}

type CapabilityStatementRestResourceSearchParam struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name for parameter in search url
	Name string `bson:"name" json:"name"`
	// Source of definition for parameter
	Definition *custom.Canonical[SearchParameter] `bson:"definition,omitempty" json:"definition,omitempty"`
	// number | date | string | token | reference | composite | quantity | uri | special
	Type custom.Code `bson:"type" json:"type"`
	// Server-specific usage
	Documentation *custom.Markdown `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

func (c CapabilityStatement) ResourceType() string {
	return "StructureDefinition"
}
