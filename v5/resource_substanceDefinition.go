// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/SubstanceDefinition
type SubstanceDefinitionName struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The actual name
	Name string `bson:"name" json:"name"`
	// Name type e.g. 'systematic',  'scientific, 'brand'
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The status of the name e.g. 'current', 'proposed'
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	// If this is the preferred name for this substance
	Preferred *bool `bson:"preferred,omitempty" json:"preferred,omitempty"`
	// Human language that the name is written in
	Language []CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	// The use context of this name e.g. as an active ingredient or as a food colour additive
	Domain []CodeableConcept `bson:"domain,omitempty" json:"domain,omitempty"`
	// The jurisdiction where this name applies
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// A synonym of this particular name, by which the substance is also known
	Synonym []interface{} `bson:"synonym,omitempty" json:"synonym,omitempty"`
	// A translation for this name into another human language
	Translation []interface{} `bson:"translation,omitempty" json:"translation,omitempty"`
	// Details of the official nature of this name
	Official []SubstanceDefinitionNameOfficial `bson:"official,omitempty" json:"official,omitempty"`
	// Supporting literature
	Source []custom.Reference[DocumentReference] `bson:"source,omitempty" json:"source,omitempty"`
}

type SubstanceDefinitionRelationship struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A pointer to another substance, as a resource or a representational code
	SubstanceDefinition *custom.Reference[SubstanceDefinition] `bson:"substanceDefinition,omitempty" json:"substanceDefinition,omitempty"`
	// For example "salt to parent", "active moiety"
	Type CodeableConcept `bson:"type" json:"type"`
	// For example where an enzyme strongly bonds with a particular substance, this is a defining relationship for that enzyme, out of several possible relationships
	IsDefining *bool `bson:"isDefining,omitempty" json:"isDefining,omitempty"`
	// A numeric factor for the relationship, e.g. that a substance salt has some percentage of active substance in relation to some other
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	// For use when the numeric has an uncertain range
	RatioHighLimitAmount *Ratio `bson:"ratioHighLimitAmount,omitempty" json:"ratioHighLimitAmount,omitempty"`
	// An operator for the amount, for example "average", "approximately", "less than"
	Comparator *CodeableConcept `bson:"comparator,omitempty" json:"comparator,omitempty"`
	// Supporting literature
	Source []custom.Reference[DocumentReference] `bson:"source,omitempty" json:"source,omitempty"`
}

type SubstanceDefinitionSourceMaterial struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Classification of the origin of the raw material. e.g. cat hair is an Animal source type
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The genus of an organism e.g. the Latin epithet of the plant/animal scientific name
	Genus *CodeableConcept `bson:"genus,omitempty" json:"genus,omitempty"`
	// The species of an organism e.g. the Latin epithet of the species of the plant/animal
	Species *CodeableConcept `bson:"species,omitempty" json:"species,omitempty"`
	// An anatomical origin of the source material within an organism
	Part *CodeableConcept `bson:"part,omitempty" json:"part,omitempty"`
	// The country or countries where the material is harvested
	CountryOfOrigin []CodeableConcept `bson:"countryOfOrigin,omitempty" json:"countryOfOrigin,omitempty"`
}

type SubstanceDefinition struct {
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
	// Identifier by which this substance is known
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// A business level version identifier of the substance
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// Status of substance within the catalogue e.g. active, retired
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	// A categorization, high level e.g. polymer or nucleic acid, or food, chemical, biological, or lower e.g. polymer linear or branch chain, or type of impurity
	Classification []CodeableConcept `bson:"classification,omitempty" json:"classification,omitempty"`
	// If the substance applies to human or veterinary use
	Domain *CodeableConcept `bson:"domain,omitempty" json:"domain,omitempty"`
	// The quality standard, established benchmark, to which substance complies (e.g. USP/NF, BP)
	Grade []CodeableConcept `bson:"grade,omitempty" json:"grade,omitempty"`
	// Textual description of the substance
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Supporting literature
	InformationSource []custom.Reference[Citation] `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	// Textual comment about the substance's catalogue or registry record
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// The entity that creates, makes, produces or fabricates the substance
	Manufacturer []custom.Reference[Organization] `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	// An entity that is the source for the substance. It may be different from the manufacturer
	Supplier []custom.Reference[Organization] `bson:"supplier,omitempty" json:"supplier,omitempty"`
	// Moiety, for structural modifications
	Moiety []SubstanceDefinitionMoiety `bson:"moiety,omitempty" json:"moiety,omitempty"`
	// General specifications for this substance
	Characterization []SubstanceDefinitionCharacterization `bson:"characterization,omitempty" json:"characterization,omitempty"`
	// General specifications for this substance
	Property []SubstanceDefinitionProperty `bson:"property,omitempty" json:"property,omitempty"`
	// General information detailing this substance
	ReferenceInformation *custom.Reference[SubstanceReferenceInformation] `bson:"referenceInformation,omitempty" json:"referenceInformation,omitempty"`
	// The average mass of a molecule of a compound
	MolecularWeight []SubstanceDefinitionMolecularWeight `bson:"molecularWeight,omitempty" json:"molecularWeight,omitempty"`
	// Structural information
	Structure *SubstanceDefinitionStructure `bson:"structure,omitempty" json:"structure,omitempty"`
	// Codes associated with the substance
	Code []SubstanceDefinitionCode `bson:"code,omitempty" json:"code,omitempty"`
	// Names applicable to this substance
	Name []SubstanceDefinitionName `bson:"name,omitempty" json:"name,omitempty"`
	// A link between this substance and another
	Relationship []SubstanceDefinitionRelationship `bson:"relationship,omitempty" json:"relationship,omitempty"`
	// Data items specific to nucleic acids
	NucleicAcid *custom.Reference[SubstanceNucleicAcid] `bson:"nucleicAcid,omitempty" json:"nucleicAcid,omitempty"`
	// Data items specific to polymers
	Polymer *custom.Reference[SubstancePolymer] `bson:"polymer,omitempty" json:"polymer,omitempty"`
	// Data items specific to proteins
	Protein *custom.Reference[SubstanceProtein] `bson:"protein,omitempty" json:"protein,omitempty"`
	// Material or taxonomic/anatomical source
	SourceMaterial *SubstanceDefinitionSourceMaterial `bson:"sourceMaterial,omitempty" json:"sourceMaterial,omitempty"`
}

type SubstanceDefinitionMoiety struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Role that the moiety is playing
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// Identifier by which this moiety substance is known
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Textual name for this moiety substance
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Stereochemistry type
	Stereochemistry *CodeableConcept `bson:"stereochemistry,omitempty" json:"stereochemistry,omitempty"`
	// Optical activity type
	OpticalActivity *CodeableConcept `bson:"opticalActivity,omitempty" json:"opticalActivity,omitempty"`
	// Molecular formula for this moiety (e.g. with the Hill system)
	MolecularFormula *string `bson:"molecularFormula,omitempty" json:"molecularFormula,omitempty"`
	// Quantitative value for this moiety
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	// The measurement type of the quantitative value
	MeasurementType *CodeableConcept `bson:"measurementType,omitempty" json:"measurementType,omitempty"`
}

type SubstanceDefinitionMolecularWeight struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The method by which the weight was determined
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	// Type of molecular weight e.g. exact, average, weight average
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Used to capture quantitative values for a variety of elements
	Amount Quantity `bson:"amount" json:"amount"`
}

type SubstanceDefinitionStructureRepresentation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The kind of structural representation (e.g. full, partial)
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The structural representation as a text string in a standard format
	Representation *string `bson:"representation,omitempty" json:"representation,omitempty"`
	// The format of the representation e.g. InChI, SMILES, MOLFILE (note: not the physical file format)
	Format *CodeableConcept `bson:"format,omitempty" json:"format,omitempty"`
	// An attachment with the structural representation e.g. a structure graphic or AnIML file
	Document *custom.Reference[DocumentReference] `bson:"document,omitempty" json:"document,omitempty"`
}

type SubstanceDefinitionCode struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The specific code
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Status of the code assignment, for example 'provisional', 'approved'
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	// The date at which the code status was changed
	StatusDate *custom.DateTime `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	// Any comment can be provided in this field
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Supporting literature
	Source []custom.Reference[DocumentReference] `bson:"source,omitempty" json:"source,omitempty"`
}

type SubstanceDefinitionCharacterization struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The method used to find the characterization e.g. HPLC
	Technique *CodeableConcept `bson:"technique,omitempty" json:"technique,omitempty"`
	// Describes the nature of the chemical entity and explains, for instance, whether this is a base or a salt form
	Form *CodeableConcept `bson:"form,omitempty" json:"form,omitempty"`
	// The description or justification in support of the interpretation of the data file
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The data produced by the analytical instrument or a pictorial representation of that data. Examples: a JCAMP, JDX, or ADX file, or a chromatogram or spectrum analysis
	File []Attachment `bson:"file,omitempty" json:"file,omitempty"`
}

type SubstanceDefinitionProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A code expressing the type of property
	Type CodeableConcept `bson:"type" json:"type"`
	// A value for the property
	Value *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
}

type SubstanceDefinitionStructure struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Stereochemistry type
	Stereochemistry *CodeableConcept `bson:"stereochemistry,omitempty" json:"stereochemistry,omitempty"`
	// Optical activity type
	OpticalActivity *CodeableConcept `bson:"opticalActivity,omitempty" json:"opticalActivity,omitempty"`
	// An expression which states the number and type of atoms present in a molecule of a substance
	MolecularFormula *string `bson:"molecularFormula,omitempty" json:"molecularFormula,omitempty"`
	// Specified per moiety according to the Hill system
	MolecularFormulaByMoiety *string `bson:"molecularFormulaByMoiety,omitempty" json:"molecularFormulaByMoiety,omitempty"`
	// The molecular weight or weight range
	MolecularWeight *interface{} `bson:"molecularWeight,omitempty" json:"molecularWeight,omitempty"`
	// The method used to find the structure e.g. X-ray, NMR
	Technique []CodeableConcept `bson:"technique,omitempty" json:"technique,omitempty"`
	// Source of information for the structure
	SourceDocument []custom.Reference[DocumentReference] `bson:"sourceDocument,omitempty" json:"sourceDocument,omitempty"`
	// A depiction of the structure of the substance
	Representation []SubstanceDefinitionStructureRepresentation `bson:"representation,omitempty" json:"representation,omitempty"`
}

type SubstanceDefinitionNameOfficial struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Which authority uses this official name
	Authority *CodeableConcept `bson:"authority,omitempty" json:"authority,omitempty"`
	// The status of the official name, for example 'draft', 'active'
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	// Date of official name change
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
}

func (s SubstanceDefinition) ResourceType() string {
	return "StructureDefinition"
}
