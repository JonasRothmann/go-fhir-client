// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformation struct {
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
	// Todo
	Comment *string `json:"comment,omitempty"`
	// Todo
	Gene []SubstanceReferenceInformationGene `json:"gene,omitempty"`
	// Todo
	GeneElement []SubstanceReferenceInformationGeneElement `json:"geneElement,omitempty"`
	// Todo
	Target []SubstanceReferenceInformationTarget `json:"target,omitempty"`
}

type SubstanceReferenceInformationGene struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Todo
	GeneSequenceOrigin *CodeableConcept `json:"geneSequenceOrigin,omitempty"`
	// Todo
	Gene *CodeableConcept `json:"gene,omitempty"`
	// Todo
	Source []custom.Reference[DocumentReference] `json:"source,omitempty"`
}

type SubstanceReferenceInformationGeneElement struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Todo
	Type *CodeableConcept `json:"type,omitempty"`
	// Todo
	Element *Identifier `json:"element,omitempty"`
	// Todo
	Source []custom.Reference[DocumentReference] `json:"source,omitempty"`
}

type SubstanceReferenceInformationTarget struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Todo
	Target *Identifier `json:"target,omitempty"`
	// Todo
	Type *CodeableConcept `json:"type,omitempty"`
	// Todo
	Interaction *CodeableConcept `json:"interaction,omitempty"`
	// Todo
	Organism *CodeableConcept `json:"organism,omitempty"`
	// Todo
	OrganismType *CodeableConcept `json:"organismType,omitempty"`
	// Todo
	AmountQuantity *Quantity `json:"amountQuantity,omitempty"`
	// Todo
	AmountRange *Range `json:"amountRange,omitempty"`
	// Todo
	AmountString *string `json:"amountString,omitempty"`
	// Todo
	AmountType *CodeableConcept `json:"amountType,omitempty"`
	// Todo
	Source []custom.Reference[DocumentReference] `json:"source,omitempty"`
}

type OtherSubstanceReferenceInformation SubstanceReferenceInformation

func (s SubstanceReferenceInformation) ResourceType() string {
	return "SubstanceReferenceInformation"
}

func (s SubstanceReferenceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceReferenceInformation
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSubstanceReferenceInformation: OtherSubstanceReferenceInformation(s), ResourceType: s.ResourceType()})
}

func UnmarshalSubstanceReferenceInformation(b []byte) (res SubstanceReferenceInformation, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
