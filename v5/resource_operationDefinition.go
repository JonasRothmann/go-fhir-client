// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/OperationDefinition
type OperationDefinitionParameterReferencedFrom struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Referencing parameter
	Source string `bson:"source" json:"source"`
	// Element id of reference
	SourceId *string `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
}

type OperationDefinitionOverload struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name of parameter to include in overload
	ParameterName []string `bson:"parameterName,omitempty" json:"parameterName,omitempty"`
	// Comments to go on overload
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty"`
}

type OperationDefinition struct {
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
	// Canonical identifier for this operation definition, represented as an absolute URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the implementation guide (business identifier)
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the operation definition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this operation definition (computer friendly)
	Name string `bson:"name" json:"name"`
	// Name for this operation definition (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// operation | query
	Kind custom.Code `bson:"kind" json:"kind"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the operation definition
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for operation definition (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this operation definition is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// Whether content is changed by the operation
	AffectsState *bool `bson:"affectsState,omitempty" json:"affectsState,omitempty"`
	// Recommended name for operation in search url
	Code custom.Code `bson:"code" json:"code"`
	// Additional information about use
	Comment *custom.Markdown `bson:"comment,omitempty" json:"comment,omitempty"`
	// Marks this as a profile of the base
	Base *custom.Canonical[OperationDefinition] `bson:"base,omitempty" json:"base,omitempty"`
	// Types this operation applies to
	Resource []custom.Code `bson:"resource,omitempty" json:"resource,omitempty"`
	// Invoke at the system level?
	System bool `bson:"system" json:"system"`
	// Invoke at the type level?
	Type bool `bson:"type" json:"type"`
	// Invoke on an instance?
	Instance bool `bson:"instance" json:"instance"`
	// Validation information for in parameters
	InputProfile *custom.Canonical[StructureDefinition] `bson:"inputProfile,omitempty" json:"inputProfile,omitempty"`
	// Validation information for out parameters
	OutputProfile *custom.Canonical[StructureDefinition] `bson:"outputProfile,omitempty" json:"outputProfile,omitempty"`
	// Parameters for the operation/query
	Parameter []OperationDefinitionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	// Define overloaded variants for when  generating code
	Overload []OperationDefinitionOverload `bson:"overload,omitempty" json:"overload,omitempty"`
}

type OperationDefinitionParameter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name in Parameters.parameter.name or in URL
	Name custom.Code `bson:"name" json:"name"`
	// in | out
	Use custom.Code `bson:"use" json:"use"`
	// instance | type | system
	Scope []custom.Code `bson:"scope,omitempty" json:"scope,omitempty"`
	// Minimum Cardinality
	Min int32 `bson:"min" json:"min"`
	// Maximum Cardinality (a number or *)
	Max string `bson:"max" json:"max"`
	// Description of meaning/use
	Documentation *custom.Markdown `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// What type this parameter has
	Type *custom.Code `bson:"type,omitempty" json:"type,omitempty"`
	// Allowed sub-type this parameter can have (if type is abstract)
	AllowedType []custom.Code `bson:"allowedType,omitempty" json:"allowedType,omitempty"`
	// If type is Reference | canonical, allowed targets. If type is 'Resource', then this constrains the allowed resource types
	TargetProfile []custom.Canonical[StructureDefinition] `bson:"targetProfile,omitempty" json:"targetProfile,omitempty"`
	// number | date | string | token | reference | composite | quantity | uri | special
	SearchType *custom.Code `bson:"searchType,omitempty" json:"searchType,omitempty"`
	// ValueSet details if this is coded
	Binding *OperationDefinitionParameterBinding `bson:"binding,omitempty" json:"binding,omitempty"`
	// References to this parameter
	ReferencedFrom []OperationDefinitionParameterReferencedFrom `bson:"referencedFrom,omitempty" json:"referencedFrom,omitempty"`
	// Parts of a nested Parameter
	Part []interface{} `bson:"part,omitempty" json:"part,omitempty"`
}

type OperationDefinitionParameterBinding struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// required | extensible | preferred | example
	Strength custom.Code `bson:"strength" json:"strength"`
	// Source of value set
	ValueSet custom.Canonical[ValueSet] `bson:"valueSet" json:"valueSet"`
}

func (o OperationDefinition) ResourceType() string {
	return "StructureDefinition"
}
