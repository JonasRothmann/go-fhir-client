// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/SubstanceNucleicAcid
type SubstanceNucleicAcid struct {
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
	// The type of the sequence shall be specified based on a controlled vocabulary
	SequenceType *CodeableConcept `bson:"sequenceType,omitempty" json:"sequenceType,omitempty"`
	// The number of linear sequences of nucleotides linked through phosphodiester bonds shall be described. Subunits would be strands of nucleic acids that are tightly associated typically through Watson-Crick base pairing. NOTE: If not specified in the reference source, the assumption is that there is 1 subunit
	NumberOfSubunits *int32 `bson:"numberOfSubunits,omitempty" json:"numberOfSubunits,omitempty"`
	// The area of hybridisation shall be described if applicable for double stranded RNA or DNA. The number associated with the subunit followed by the number associated to the residue shall be specified in increasing order. The underscore “” shall be used as separator as follows: “Subunitnumber Residue”
	AreaOfHybridisation *string `bson:"areaOfHybridisation,omitempty" json:"areaOfHybridisation,omitempty"`
	// (TBC)
	OligoNucleotideType *CodeableConcept `bson:"oligoNucleotideType,omitempty" json:"oligoNucleotideType,omitempty"`
	// Subunits are listed in order of decreasing length; sequences of the same length will be ordered by molecular weight; subunits that have identical sequences will be repeated multiple times
	Subunit []SubstanceNucleicAcidSubunit `bson:"subunit,omitempty" json:"subunit,omitempty"`
}

type SubstanceNucleicAcidSubunit struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Index of linear sequences of nucleic acids in order of decreasing length. Sequences of the same length will be ordered by molecular weight. Subunits that have identical sequences will be repeated and have sequential subscripts
	Subunit *int32 `bson:"subunit,omitempty" json:"subunit,omitempty"`
	// Actual nucleotide sequence notation from 5' to 3' end using standard single letter codes. In addition to the base sequence, sugar and type of phosphate or non-phosphate linkage should also be captured
	Sequence *string `bson:"sequence,omitempty" json:"sequence,omitempty"`
	// The length of the sequence shall be captured
	Length *int32 `bson:"length,omitempty" json:"length,omitempty"`
	// (TBC)
	SequenceAttachment *Attachment `bson:"sequenceAttachment,omitempty" json:"sequenceAttachment,omitempty"`
	// The nucleotide present at the 5’ terminal shall be specified based on a controlled vocabulary. Since the sequence is represented from the 5' to the 3' end, the 5’ prime nucleotide is the letter at the first position in the sequence. A separate representation would be redundant
	FivePrime *CodeableConcept `bson:"fivePrime,omitempty" json:"fivePrime,omitempty"`
	// The nucleotide present at the 3’ terminal shall be specified based on a controlled vocabulary. Since the sequence is represented from the 5' to the 3' end, the 5’ prime nucleotide is the letter at the last position in the sequence. A separate representation would be redundant
	ThreePrime *CodeableConcept `bson:"threePrime,omitempty" json:"threePrime,omitempty"`
	// The linkages between sugar residues will also be captured
	Linkage []SubstanceNucleicAcidSubunitLinkage `bson:"linkage,omitempty" json:"linkage,omitempty"`
	// 5.3.6.8.1 Sugar ID (Mandatory)
	Sugar []SubstanceNucleicAcidSubunitSugar `bson:"sugar,omitempty" json:"sugar,omitempty"`
}

type SubstanceNucleicAcidSubunitLinkage struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The entity that links the sugar residues together should also be captured for nearly all naturally occurring nucleic acid the linkage is a phosphate group. For many synthetic oligonucleotides phosphorothioate linkages are often seen. Linkage connectivity is assumed to be 3’-5’. If the linkage is either 3’-3’ or 5’-5’ this should be specified
	Connectivity *string `bson:"connectivity,omitempty" json:"connectivity,omitempty"`
	// Each linkage will be registered as a fragment and have an ID
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Each linkage will be registered as a fragment and have at least one name. A single name shall be assigned to each linkage
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Residues shall be captured as described in 5.3.6.8.3
	ResidueSite *string `bson:"residueSite,omitempty" json:"residueSite,omitempty"`
}

type SubstanceNucleicAcidSubunitSugar struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The Substance ID of the sugar or sugar-like component that make up the nucleotide
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// The name of the sugar or sugar-like component that make up the nucleotide
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// The residues that contain a given sugar will be captured. The order of given residues will be captured in the 5‘-3‘direction consistent with the base sequences listed above
	ResidueSite *string `bson:"residueSite,omitempty" json:"residueSite,omitempty"`
}

func (s SubstanceNucleicAcid) ResourceType() string {
	return "StructureDefinition"
}
