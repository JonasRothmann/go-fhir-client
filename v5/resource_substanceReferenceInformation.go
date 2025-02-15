// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformationGene struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Todo
	GeneSequenceOrigin *CodeableConcept `bson:"geneSequenceOrigin,omitempty" json:"geneSequenceOrigin,omitempty"`
	// Todo
	Gene *CodeableConcept `bson:"gene,omitempty" json:"gene,omitempty"`
	// Todo
	Source []custom.Reference[DocumentReference] `bson:"source,omitempty" json:"source,omitempty"`
}

type SubstanceReferenceInformationGeneElement struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Todo
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Todo
	Element *Identifier `bson:"element,omitempty" json:"element,omitempty"`
	// Todo
	Source []custom.Reference[DocumentReference] `bson:"source,omitempty" json:"source,omitempty"`
}

type SubstanceReferenceInformationTarget struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Todo
	Target *Identifier `bson:"target,omitempty" json:"target,omitempty"`
	// Todo
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Todo
	Interaction *CodeableConcept `bson:"interaction,omitempty" json:"interaction,omitempty"`
	// Todo
	Organism *CodeableConcept `bson:"organism,omitempty" json:"organism,omitempty"`
	// Todo
	OrganismType *CodeableConcept `bson:"organismType,omitempty" json:"organismType,omitempty"`
	// Todo
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	// Todo
	AmountType *CodeableConcept `bson:"amountType,omitempty" json:"amountType,omitempty"`
	// Todo
	Source []custom.Reference[DocumentReference] `bson:"source,omitempty" json:"source,omitempty"`
}

type SubstanceReferenceInformation struct {
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
	// Todo
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty"`
	// Todo
	Gene []SubstanceReferenceInformationGene `bson:"gene,omitempty" json:"gene,omitempty"`
	// Todo
	GeneElement []SubstanceReferenceInformationGeneElement `bson:"geneElement,omitempty" json:"geneElement,omitempty"`
	// Todo
	Target []SubstanceReferenceInformationTarget `bson:"target,omitempty" json:"target,omitempty"`
}

func (s SubstanceReferenceInformation) ResourceType() string {
	return "StructureDefinition"
}
