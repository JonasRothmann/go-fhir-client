// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeatRepeatUnitStructuralRepresentation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of structure (e.g. Full, Partial, Representative)
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The structural representation as text string in a standard format e.g. InChI, SMILES, MOLFILE, CDX, SDF, PDB, mmCIF
	Representation *string `bson:"representation,omitempty" json:"representation,omitempty"`
	// The format of the representation e.g. InChI, SMILES, MOLFILE, CDX, SDF, PDB, mmCIF
	Format *CodeableConcept `bson:"format,omitempty" json:"format,omitempty"`
	// An attached file with the structural representation
	Attachment *Attachment `bson:"attachment,omitempty" json:"attachment,omitempty"`
}

type SubstancePolymer struct {
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
	// A business idenfier for this polymer, but typically this is handled by a SubstanceDefinition identifier
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Overall type of the polymer
	Class *CodeableConcept `bson:"class,omitempty" json:"class,omitempty"`
	// Polymer geometry, e.g. linear, branched, cross-linked, network or dendritic
	Geometry *CodeableConcept `bson:"geometry,omitempty" json:"geometry,omitempty"`
	// Descrtibes the copolymer sequence type (polymer connectivity)
	CopolymerConnectivity []CodeableConcept `bson:"copolymerConnectivity,omitempty" json:"copolymerConnectivity,omitempty"`
	// Todo - this is intended to connect to a repeating full modification structure, also used by Protein and Nucleic Acid . String is just a placeholder
	Modification *string `bson:"modification,omitempty" json:"modification,omitempty"`
	// Todo
	MonomerSet []SubstancePolymerMonomerSet `bson:"monomerSet,omitempty" json:"monomerSet,omitempty"`
	// Specifies and quantifies the repeated units and their configuration
	Repeat []SubstancePolymerRepeat `bson:"repeat,omitempty" json:"repeat,omitempty"`
}

type SubstancePolymerMonomerSet struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Captures the type of ratio to the entire polymer, e.g. Monomer/Polymer ratio, SRU/Polymer Ratio
	RatioType *CodeableConcept `bson:"ratioType,omitempty" json:"ratioType,omitempty"`
	// The starting materials - monomer(s) used in the synthesis of the polymer
	StartingMaterial []SubstancePolymerMonomerSetStartingMaterial `bson:"startingMaterial,omitempty" json:"startingMaterial,omitempty"`
}

type SubstancePolymerMonomerSetStartingMaterial struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of substance for this starting material
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Substance high level category, e.g. chemical substance
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Used to specify whether the attribute described is a defining element for the unique identification of the polymer
	IsDefining *bool `bson:"isDefining,omitempty" json:"isDefining,omitempty"`
	// A percentage
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
}

type SubstancePolymerRepeat struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A representation of an (average) molecular formula from a polymer
	AverageMolecularFormula *string `bson:"averageMolecularFormula,omitempty" json:"averageMolecularFormula,omitempty"`
	// How the quantitative amount of Structural Repeat Units is captured (e.g. Exact, Numeric, Average)
	RepeatUnitAmountType *CodeableConcept `bson:"repeatUnitAmountType,omitempty" json:"repeatUnitAmountType,omitempty"`
	// An SRU - Structural Repeat Unit
	RepeatUnit []SubstancePolymerRepeatRepeatUnit `bson:"repeatUnit,omitempty" json:"repeatUnit,omitempty"`
}

type SubstancePolymerRepeatRepeatUnit struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Structural repeat units are essential elements for defining polymers
	Unit *string `bson:"unit,omitempty" json:"unit,omitempty"`
	// The orientation of the polymerisation, e.g. head-tail, head-head, random
	Orientation *CodeableConcept `bson:"orientation,omitempty" json:"orientation,omitempty"`
	// Number of repeats of this unit
	Amount *int32 `bson:"amount,omitempty" json:"amount,omitempty"`
	// Applies to homopolymer and block co-polymers where the degree of polymerisation within a block can be described
	DegreeOfPolymerisation []SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation `bson:"degreeOfPolymerisation,omitempty" json:"degreeOfPolymerisation,omitempty"`
	// A graphical structure for this SRU
	StructuralRepresentation []SubstancePolymerRepeatRepeatUnitStructuralRepresentation `bson:"structuralRepresentation,omitempty" json:"structuralRepresentation,omitempty"`
}

type SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of the degree of polymerisation shall be described, e.g. SRU/Polymer Ratio
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// An average amount of polymerisation
	Average *int32 `bson:"average,omitempty" json:"average,omitempty"`
	// A low expected limit of the amount
	Low *int32 `bson:"low,omitempty" json:"low,omitempty"`
	// A high expected limit of the amount
	High *int32 `bson:"high,omitempty" json:"high,omitempty"`
}

func (s SubstancePolymer) ResourceType() string {
	return "StructureDefinition"
}
