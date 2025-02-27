// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MolecularSequence
type MolecularSequence struct {
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
	// Unique ID for this particular sequence
	Identifier []Identifier `json:"identifier,omitempty"`
	// aa | dna | rna
	Type *custom.Code `json:"type,omitempty"`
	// Subject this sequence is associated too
	Subject *custom.Reference[MolecularSequenceSubject] `json:"subject,omitempty"`
	// What the molecular sequence is about, when it is not about the subject of record
	Focus []custom.Reference[Resource] `json:"focus,omitempty"`
	// Specimen used for sequencing
	Specimen *custom.Reference[Specimen] `json:"specimen,omitempty"`
	// The method for sequencing
	Device *custom.Reference[Device] `json:"device,omitempty"`
	// Who should be responsible for test result
	Performer *custom.Reference[Organization] `json:"performer,omitempty"`
	// Sequence that was observed
	Literal *string `json:"literal,omitempty"`
	// Embedded file or a link (URL) which contains content to represent the sequence
	Formatted []Attachment `json:"formatted,omitempty"`
	// A sequence defined relative to another sequence
	Relative []MolecularSequenceRelative `json:"relative,omitempty"`
}

type MolecularSequenceRelative struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Ways of identifying nucleotides or amino acids within a sequence
	CoordinateSystem CodeableConcept `json:"coordinateSystem"`
	// Indicates the order in which the sequence should be considered when putting multiple 'relative' elements together
	OrdinalPosition *int32 `json:"ordinalPosition,omitempty"`
	// Indicates the nucleotide range in the composed sequence when multiple 'relative' elements are used together
	SequenceRange *Range `json:"sequenceRange,omitempty"`
	// A sequence used as starting sequence
	StartingSequence *MolecularSequenceRelativeStartingSequence `json:"startingSequence,omitempty"`
	// Changes in sequence from the starting sequence
	Edit []MolecularSequenceRelativeEdit `json:"edit,omitempty"`
}

type MolecularSequenceRelativeStartingSequence struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The genome assembly used for starting sequence, e.g. GRCh38
	GenomeAssembly *CodeableConcept `json:"genomeAssembly,omitempty"`
	// Chromosome Identifier
	Chromosome *CodeableConcept `json:"chromosome,omitempty"`
	// The reference sequence that represents the starting sequence
	SequenceCodeableConcept *CodeableConcept `json:"sequenceCodeableConcept,omitempty"`
	// The reference sequence that represents the starting sequence
	SequenceString *string `json:"sequenceString,omitempty"`
	// The reference sequence that represents the starting sequence
	SequenceReference *custom.Reference[MolecularSequence] `json:"sequenceReference,omitempty"`
	// Start position of the window on the starting sequence
	WindowStart *int32 `json:"windowStart,omitempty"`
	// End position of the window on the starting sequence
	WindowEnd *int32 `json:"windowEnd,omitempty"`
	// sense | antisense
	Orientation *custom.Code `json:"orientation,omitempty"`
	// watson | crick
	Strand *custom.Code `json:"strand,omitempty"`
}

type MolecularSequenceRelativeEdit struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Start position of the edit on the starting sequence
	Start *int32 `json:"start,omitempty"`
	// End position of the edit on the starting sequence
	End *int32 `json:"end,omitempty"`
	// Allele that was observed
	ReplacementSequence *string `json:"replacementSequence,omitempty"`
	// Allele in the starting sequence
	ReplacedSequence *string `json:"replacedSequence,omitempty"`
}

type OtherMolecularSequence MolecularSequence

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
	return "MolecularSequence"
}

func (m MolecularSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMolecularSequence
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMolecularSequence: OtherMolecularSequence(m), ResourceType: m.ResourceType()})
}

func UnmarshalMolecularSequence(b []byte) (res MolecularSequence, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
