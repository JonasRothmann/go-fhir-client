// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MolecularSequence
type MolecularSequenceRelativeEdit struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Start position of the edit on the starting sequence
	Start *int32 `bson:"start,omitempty" json:"start,omitempty"`
	// End position of the edit on the starting sequence
	End *int32 `bson:"end,omitempty" json:"end,omitempty"`
	// Allele that was observed
	ReplacementSequence *string `bson:"replacementSequence,omitempty" json:"replacementSequence,omitempty"`
	// Allele in the starting sequence
	ReplacedSequence *string `bson:"replacedSequence,omitempty" json:"replacedSequence,omitempty"`
}

type MolecularSequence struct {
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
	// Unique ID for this particular sequence
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// aa | dna | rna
	Type *custom.Code `bson:"type,omitempty" json:"type,omitempty"`
	// Subject this sequence is associated too
	Subject *custom.Reference[MolecularSequenceSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// What the molecular sequence is about, when it is not about the subject of record
	Focus []custom.Reference[Resource] `bson:"focus,omitempty" json:"focus,omitempty"`
	// Specimen used for sequencing
	Specimen *custom.Reference[Specimen] `bson:"specimen,omitempty" json:"specimen,omitempty"`
	// The method for sequencing
	Device *custom.Reference[Device] `bson:"device,omitempty" json:"device,omitempty"`
	// Who should be responsible for test result
	Performer *custom.Reference[Organization] `bson:"performer,omitempty" json:"performer,omitempty"`
	// Sequence that was observed
	Literal *string `bson:"literal,omitempty" json:"literal,omitempty"`
	// Embedded file or a link (URL) which contains content to represent the sequence
	Formatted []Attachment `bson:"formatted,omitempty" json:"formatted,omitempty"`
	// A sequence defined relative to another sequence
	Relative []MolecularSequenceRelative `bson:"relative,omitempty" json:"relative,omitempty"`
}

type MolecularSequenceRelative struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Ways of identifying nucleotides or amino acids within a sequence
	CoordinateSystem CodeableConcept `bson:"coordinateSystem" json:"coordinateSystem"`
	// Indicates the order in which the sequence should be considered when putting multiple 'relative' elements together
	OrdinalPosition *int32 `bson:"ordinalPosition,omitempty" json:"ordinalPosition,omitempty"`
	// Indicates the nucleotide range in the composed sequence when multiple 'relative' elements are used together
	SequenceRange *Range `bson:"sequenceRange,omitempty" json:"sequenceRange,omitempty"`
	// A sequence used as starting sequence
	StartingSequence *MolecularSequenceRelativeStartingSequence `bson:"startingSequence,omitempty" json:"startingSequence,omitempty"`
	// Changes in sequence from the starting sequence
	Edit []MolecularSequenceRelativeEdit `bson:"edit,omitempty" json:"edit,omitempty"`
}

type MolecularSequenceRelativeStartingSequence struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The genome assembly used for starting sequence, e.g. GRCh38
	GenomeAssembly *CodeableConcept `bson:"genomeAssembly,omitempty" json:"genomeAssembly,omitempty"`
	// Chromosome Identifier
	Chromosome *CodeableConcept `bson:"chromosome,omitempty" json:"chromosome,omitempty"`
	// The reference sequence that represents the starting sequence
	Sequence *CodeableConcept `bson:"sequence,omitempty" json:"sequence,omitempty"`
	// Start position of the window on the starting sequence
	WindowStart *int32 `bson:"windowStart,omitempty" json:"windowStart,omitempty"`
	// End position of the window on the starting sequence
	WindowEnd *int32 `bson:"windowEnd,omitempty" json:"windowEnd,omitempty"`
	// sense | antisense
	Orientation *custom.Code `bson:"orientation,omitempty" json:"orientation,omitempty"`
	// watson | crick
	Strand *custom.Code `bson:"strand,omitempty" json:"strand,omitempty"`
}

type MolecularSequenceSubject interface {
	gofhirclient.Resource

	Is_MolecularSequenceSubject()
}

func (p Patient) Is_MolecularSequenceSubject()                    {}
func (g Group) Is_MolecularSequenceSubject()                      {}
func (s Substance) Is_MolecularSequenceSubject()                  {}
func (b BiologicallyDerivedProduct) Is_MolecularSequenceSubject() {}
func (n NutritionProduct) Is_MolecularSequenceSubject()           {}

func (m MolecularSequence) ResourceType() string {
	return "StructureDefinition"
}
