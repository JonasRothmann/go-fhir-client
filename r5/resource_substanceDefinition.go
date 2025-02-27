// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionMoiety struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Role that the moiety is playing
	Role *CodeableConcept `json:"role,omitempty"`
	// Identifier by which this moiety substance is known
	Identifier *Identifier `json:"identifier,omitempty"`
	// Textual name for this moiety substance
	Name *string `json:"name,omitempty"`
	// Stereochemistry type
	Stereochemistry *CodeableConcept `json:"stereochemistry,omitempty"`
	// Optical activity type
	OpticalActivity *CodeableConcept `json:"opticalActivity,omitempty"`
	// Molecular formula for this moiety (e.g. with the Hill system)
	MolecularFormula *string `json:"molecularFormula,omitempty"`
	// Quantitative value for this moiety
	AmountQuantity *Quantity `json:"amountQuantity,omitempty"`
	// Quantitative value for this moiety
	AmountString *string `json:"amountString,omitempty"`
	// The measurement type of the quantitative value
	MeasurementType *CodeableConcept `json:"measurementType,omitempty"`
}

type SubstanceDefinitionCharacterization struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The method used to find the characterization e.g. HPLC
	Technique *CodeableConcept `json:"technique,omitempty"`
	// Describes the nature of the chemical entity and explains, for instance, whether this is a base or a salt form
	Form *CodeableConcept `json:"form,omitempty"`
	// The description or justification in support of the interpretation of the data file
	Description *custom.Markdown `json:"description,omitempty"`
	// The data produced by the analytical instrument or a pictorial representation of that data. Examples: a JCAMP, JDX, or ADX file, or a chromatogram or spectrum analysis
	File []Attachment `json:"file,omitempty"`
}

type SubstanceDefinitionMolecularWeight struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The method by which the weight was determined
	Method *CodeableConcept `json:"method,omitempty"`
	// Type of molecular weight e.g. exact, average, weight average
	Type *CodeableConcept `json:"type,omitempty"`
	// Used to capture quantitative values for a variety of elements
	Amount Quantity `json:"amount"`
}

type SubstanceDefinitionStructureRepresentation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The kind of structural representation (e.g. full, partial)
	Type *CodeableConcept `json:"type,omitempty"`
	// The structural representation as a text string in a standard format
	Representation *string `json:"representation,omitempty"`
	// The format of the representation e.g. InChI, SMILES, MOLFILE (note: not the physical file format)
	Format *CodeableConcept `json:"format,omitempty"`
	// An attachment with the structural representation e.g. a structure graphic or AnIML file
	Document *custom.Reference[DocumentReference] `json:"document,omitempty"`
}

type SubstanceDefinitionRelationship struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A pointer to another substance, as a resource or a representational code
	SubstanceDefinitionReference *custom.Reference[SubstanceDefinition] `json:"substanceDefinitionReference,omitempty"`
	// A pointer to another substance, as a resource or a representational code
	SubstanceDefinitionCodeableConcept *CodeableConcept `json:"substanceDefinitionCodeableConcept,omitempty"`
	// For example "salt to parent", "active moiety"
	Type CodeableConcept `json:"type"`
	// For example where an enzyme strongly bonds with a particular substance, this is a defining relationship for that enzyme, out of several possible relationships
	IsDefining *bool `json:"isDefining,omitempty"`
	// A numeric factor for the relationship, e.g. that a substance salt has some percentage of active substance in relation to some other
	AmountQuantity *Quantity `json:"amountQuantity,omitempty"`
	// A numeric factor for the relationship, e.g. that a substance salt has some percentage of active substance in relation to some other
	AmountRatio *Ratio `json:"amountRatio,omitempty"`
	// A numeric factor for the relationship, e.g. that a substance salt has some percentage of active substance in relation to some other
	AmountString *string `json:"amountString,omitempty"`
	// For use when the numeric has an uncertain range
	RatioHighLimitAmount *Ratio `json:"ratioHighLimitAmount,omitempty"`
	// An operator for the amount, for example "average", "approximately", "less than"
	Comparator *CodeableConcept `json:"comparator,omitempty"`
	// Supporting literature
	Source []custom.Reference[DocumentReference] `json:"source,omitempty"`
}

type SubstanceDefinitionSourceMaterial struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Classification of the origin of the raw material. e.g. cat hair is an Animal source type
	Type *CodeableConcept `json:"type,omitempty"`
	// The genus of an organism e.g. the Latin epithet of the plant/animal scientific name
	Genus *CodeableConcept `json:"genus,omitempty"`
	// The species of an organism e.g. the Latin epithet of the species of the plant/animal
	Species *CodeableConcept `json:"species,omitempty"`
	// An anatomical origin of the source material within an organism
	Part *CodeableConcept `json:"part,omitempty"`
	// The country or countries where the material is harvested
	CountryOfOrigin []CodeableConcept `json:"countryOfOrigin,omitempty"`
}

type SubstanceDefinition struct {
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
	// Identifier by which this substance is known
	Identifier []Identifier `json:"identifier,omitempty"`
	// A business level version identifier of the substance
	Version *string `json:"version,omitempty"`
	// Status of substance within the catalogue e.g. active, retired
	Status *CodeableConcept `json:"status,omitempty"`
	// A categorization, high level e.g. polymer or nucleic acid, or food, chemical, biological, or lower e.g. polymer linear or branch chain, or type of impurity
	Classification []CodeableConcept `json:"classification,omitempty"`
	// If the substance applies to human or veterinary use
	Domain *CodeableConcept `json:"domain,omitempty"`
	// The quality standard, established benchmark, to which substance complies (e.g. USP/NF, BP)
	Grade []CodeableConcept `json:"grade,omitempty"`
	// Textual description of the substance
	Description *custom.Markdown `json:"description,omitempty"`
	// Supporting literature
	InformationSource []custom.Reference[Citation] `json:"informationSource,omitempty"`
	// Textual comment about the substance's catalogue or registry record
	Note []Annotation `json:"note,omitempty"`
	// The entity that creates, makes, produces or fabricates the substance
	Manufacturer []custom.Reference[Organization] `json:"manufacturer,omitempty"`
	// An entity that is the source for the substance. It may be different from the manufacturer
	Supplier []custom.Reference[Organization] `json:"supplier,omitempty"`
	// Moiety, for structural modifications
	Moiety []SubstanceDefinitionMoiety `json:"moiety,omitempty"`
	// General specifications for this substance
	Characterization []SubstanceDefinitionCharacterization `json:"characterization,omitempty"`
	// General specifications for this substance
	Property []SubstanceDefinitionProperty `json:"property,omitempty"`
	// General information detailing this substance
	ReferenceInformation *custom.Reference[SubstanceReferenceInformation] `json:"referenceInformation,omitempty"`
	// The average mass of a molecule of a compound
	MolecularWeight []SubstanceDefinitionMolecularWeight `json:"molecularWeight,omitempty"`
	// Structural information
	Structure *SubstanceDefinitionStructure `json:"structure,omitempty"`
	// Codes associated with the substance
	Code []SubstanceDefinitionCode `json:"code,omitempty"`
	// Names applicable to this substance
	Name []SubstanceDefinitionName `json:"name,omitempty"`
	// A link between this substance and another
	Relationship []SubstanceDefinitionRelationship `json:"relationship,omitempty"`
	// Data items specific to nucleic acids
	NucleicAcid *custom.Reference[SubstanceNucleicAcid] `json:"nucleicAcid,omitempty"`
	// Data items specific to polymers
	Polymer *custom.Reference[SubstancePolymer] `json:"polymer,omitempty"`
	// Data items specific to proteins
	Protein *custom.Reference[SubstanceProtein] `json:"protein,omitempty"`
	// Material or taxonomic/anatomical source
	SourceMaterial *SubstanceDefinitionSourceMaterial `json:"sourceMaterial,omitempty"`
}

type SubstanceDefinitionProperty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A code expressing the type of property
	Type CodeableConcept `json:"type"`
	// A value for the property
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// A value for the property
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// A value for the property
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// A value for the property
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// A value for the property
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
}

type SubstanceDefinitionStructure struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Stereochemistry type
	Stereochemistry *CodeableConcept `json:"stereochemistry,omitempty"`
	// Optical activity type
	OpticalActivity *CodeableConcept `json:"opticalActivity,omitempty"`
	// An expression which states the number and type of atoms present in a molecule of a substance
	MolecularFormula *string `json:"molecularFormula,omitempty"`
	// Specified per moiety according to the Hill system
	MolecularFormulaByMoiety *string `json:"molecularFormulaByMoiety,omitempty"`
	// The molecular weight or weight range
	MolecularWeight *interface{} `json:"molecularWeight,omitempty"`
	// The method used to find the structure e.g. X-ray, NMR
	Technique []CodeableConcept `json:"technique,omitempty"`
	// Source of information for the structure
	SourceDocument []custom.Reference[DocumentReference] `json:"sourceDocument,omitempty"`
	// A depiction of the structure of the substance
	Representation []SubstanceDefinitionStructureRepresentation `json:"representation,omitempty"`
}

type SubstanceDefinitionCode struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The specific code
	Code *CodeableConcept `json:"code,omitempty"`
	// Status of the code assignment, for example 'provisional', 'approved'
	Status *CodeableConcept `json:"status,omitempty"`
	// The date at which the code status was changed
	StatusDate *custom.DateTime `json:"statusDate,omitempty"`
	// Any comment can be provided in this field
	Note []Annotation `json:"note,omitempty"`
	// Supporting literature
	Source []custom.Reference[DocumentReference] `json:"source,omitempty"`
}

type SubstanceDefinitionName struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The actual name
	Name string `json:"name"`
	// Name type e.g. 'systematic',  'scientific, 'brand'
	Type *CodeableConcept `json:"type,omitempty"`
	// The status of the name e.g. 'current', 'proposed'
	Status *CodeableConcept `json:"status,omitempty"`
	// If this is the preferred name for this substance
	Preferred *bool `json:"preferred,omitempty"`
	// Human language that the name is written in
	Language []CodeableConcept `json:"language,omitempty"`
	// The use context of this name e.g. as an active ingredient or as a food colour additive
	Domain []CodeableConcept `json:"domain,omitempty"`
	// The jurisdiction where this name applies
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// A synonym of this particular name, by which the substance is also known
	Synonym []interface{} `json:"synonym,omitempty"`
	// A translation for this name into another human language
	Translation []interface{} `json:"translation,omitempty"`
	// Details of the official nature of this name
	Official []SubstanceDefinitionNameOfficial `json:"official,omitempty"`
	// Supporting literature
	Source []custom.Reference[DocumentReference] `json:"source,omitempty"`
}

type SubstanceDefinitionNameOfficial struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Which authority uses this official name
	Authority *CodeableConcept `json:"authority,omitempty"`
	// The status of the official name, for example 'draft', 'active'
	Status *CodeableConcept `json:"status,omitempty"`
	// Date of official name change
	Date *custom.DateTime `json:"date,omitempty"`
}

type OtherSubstanceDefinition SubstanceDefinition

func (s SubstanceDefinition) ResourceType() string {
	return "SubstanceDefinition"
}

func (s SubstanceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSubstanceDefinition: OtherSubstanceDefinition(s), ResourceType: s.ResourceType()})
}

func UnmarshalSubstanceDefinition(b []byte) (res SubstanceDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
