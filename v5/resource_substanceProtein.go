// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/SubstanceProtein
type SubstanceProtein struct {
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
	// The SubstanceProtein descriptive elements will only be used when a complete or partial amino acid sequence is available or derivable from a nucleic acid sequence
	SequenceType *CodeableConcept `bson:"sequenceType,omitempty" json:"sequenceType,omitempty"`
	// Number of linear sequences of amino acids linked through peptide bonds. The number of subunits constituting the SubstanceProtein shall be described. It is possible that the number of subunits can be variable
	NumberOfSubunits *int32 `bson:"numberOfSubunits,omitempty" json:"numberOfSubunits,omitempty"`
	// The disulphide bond between two cysteine residues either on the same subunit or on two different subunits shall be described. The position of the disulfide bonds in the SubstanceProtein shall be listed in increasing order of subunit number and position within subunit followed by the abbreviation of the amino acids involved. The disulfide linkage positions shall actually contain the amino acid Cysteine at the respective positions
	DisulfideLinkage []string `bson:"disulfideLinkage,omitempty" json:"disulfideLinkage,omitempty"`
	// This subclause refers to the description of each subunit constituting the SubstanceProtein. A subunit is a linear sequence of amino acids linked through peptide bonds. The Subunit information shall be provided when the finished SubstanceProtein is a complex of multiple sequences; subunits are not used to delineate domains within a single sequence. Subunits are listed in order of decreasing length; sequences of the same length will be ordered by decreasing molecular weight; subunits that have identical sequences will be repeated multiple times
	Subunit []SubstanceProteinSubunit `bson:"subunit,omitempty" json:"subunit,omitempty"`
}

type SubstanceProteinSubunit struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Index of primary sequences of amino acids linked through peptide bonds in order of decreasing length. Sequences of the same length will be ordered by molecular weight. Subunits that have identical sequences will be repeated and have sequential subscripts
	Subunit *int32 `bson:"subunit,omitempty" json:"subunit,omitempty"`
	// The sequence information shall be provided enumerating the amino acids from N- to C-terminal end using standard single-letter amino acid codes. Uppercase shall be used for L-amino acids and lowercase for D-amino acids. Transcribed SubstanceProteins will always be described using the translated sequence; for synthetic peptide containing amino acids that are not represented with a single letter code an X should be used within the sequence. The modified amino acids will be distinguished by their position in the sequence
	Sequence *string `bson:"sequence,omitempty" json:"sequence,omitempty"`
	// Length of linear sequences of amino acids contained in the subunit
	Length *int32 `bson:"length,omitempty" json:"length,omitempty"`
	// The sequence information shall be provided enumerating the amino acids from N- to C-terminal end using standard single-letter amino acid codes. Uppercase shall be used for L-amino acids and lowercase for D-amino acids. Transcribed SubstanceProteins will always be described using the translated sequence; for synthetic peptide containing amino acids that are not represented with a single letter code an X should be used within the sequence. The modified amino acids will be distinguished by their position in the sequence
	SequenceAttachment *Attachment `bson:"sequenceAttachment,omitempty" json:"sequenceAttachment,omitempty"`
	// Unique identifier for molecular fragment modification based on the ISO 11238 Substance ID
	NTerminalModificationId *Identifier `bson:"nTerminalModificationId,omitempty" json:"nTerminalModificationId,omitempty"`
	// The name of the fragment modified at the N-terminal of the SubstanceProtein shall be specified
	NTerminalModification *string `bson:"nTerminalModification,omitempty" json:"nTerminalModification,omitempty"`
	// Unique identifier for molecular fragment modification based on the ISO 11238 Substance ID
	CTerminalModificationId *Identifier `bson:"cTerminalModificationId,omitempty" json:"cTerminalModificationId,omitempty"`
	// The modification at the C-terminal shall be specified
	CTerminalModification *string `bson:"cTerminalModification,omitempty" json:"cTerminalModification,omitempty"`
}

func (s SubstanceProtein) ResourceType() string {
	return "StructureDefinition"
}
