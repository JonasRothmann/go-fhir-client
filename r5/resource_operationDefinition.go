// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/OperationDefinition
type OperationDefinitionParameterReferencedFrom struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Referencing parameter
	Source string `json:"source"`
	// Element id of reference
	SourceId *string `json:"sourceId,omitempty"`
}

type OperationDefinitionOverload struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name of parameter to include in overload
	ParameterName []string `json:"parameterName,omitempty"`
	// Comments to go on overload
	Comment *string `json:"comment,omitempty"`
}

type OperationDefinition struct {
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
	// Canonical identifier for this operation definition, represented as an absolute URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the implementation guide (business identifier)
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the operation definition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this operation definition (computer friendly)
	Name string `json:"name"`
	// Name for this operation definition (human friendly)
	Title *string `json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// operation | query
	Kind custom.Code `json:"kind"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the operation definition
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for operation definition (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this operation definition is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// Whether content is changed by the operation
	AffectsState *bool `json:"affectsState,omitempty"`
	// Recommended name for operation in search url
	Code custom.Code `json:"code"`
	// Additional information about use
	Comment *custom.Markdown `json:"comment,omitempty"`
	// Marks this as a profile of the base
	Base *custom.Canonical[OperationDefinition] `json:"base,omitempty"`
	// Types this operation applies to
	Resource []custom.Code `json:"resource,omitempty"`
	// Invoke at the system level?
	System bool `json:"system"`
	// Invoke at the type level?
	Type bool `json:"type"`
	// Invoke on an instance?
	Instance bool `json:"instance"`
	// Validation information for in parameters
	InputProfile *custom.Canonical[StructureDefinition] `json:"inputProfile,omitempty"`
	// Validation information for out parameters
	OutputProfile *custom.Canonical[StructureDefinition] `json:"outputProfile,omitempty"`
	// Parameters for the operation/query
	Parameter []OperationDefinitionParameter `json:"parameter,omitempty"`
	// Define overloaded variants for when  generating code
	Overload []OperationDefinitionOverload `json:"overload,omitempty"`
}

type OperationDefinitionParameter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name in Parameters.parameter.name or in URL
	Name custom.Code `json:"name"`
	// in | out
	Use custom.Code `json:"use"`
	// instance | type | system
	Scope []custom.Code `json:"scope,omitempty"`
	// Minimum Cardinality
	Min int32 `json:"min"`
	// Maximum Cardinality (a number or *)
	Max string `json:"max"`
	// Description of meaning/use
	Documentation *custom.Markdown `json:"documentation,omitempty"`
	// What type this parameter has
	Type *custom.Code `json:"type,omitempty"`
	// Allowed sub-type this parameter can have (if type is abstract)
	AllowedType []custom.Code `json:"allowedType,omitempty"`
	// If type is Reference | canonical, allowed targets. If type is 'Resource', then this constrains the allowed resource types
	TargetProfile []custom.Canonical[StructureDefinition] `json:"targetProfile,omitempty"`
	// number | date | string | token | reference | composite | quantity | uri | special
	SearchType *custom.Code `json:"searchType,omitempty"`
	// ValueSet details if this is coded
	Binding *OperationDefinitionParameterBinding `json:"binding,omitempty"`
	// References to this parameter
	ReferencedFrom []OperationDefinitionParameterReferencedFrom `json:"referencedFrom,omitempty"`
	// Parts of a nested Parameter
	Part []interface{} `json:"part,omitempty"`
}

type OperationDefinitionParameterBinding struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// required | extensible | preferred | example
	Strength custom.Code `json:"strength"`
	// Source of value set
	ValueSet custom.Canonical[ValueSet] `json:"valueSet"`
}

type OtherOperationDefinition OperationDefinition

func (o OperationDefinition) ResourceType() string {
	return "OperationDefinition"
}

func (o OperationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOperationDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherOperationDefinition: OtherOperationDefinition(o), ResourceType: o.ResourceType()})
}

func UnmarshalOperationDefinition(b []byte) (res OperationDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
