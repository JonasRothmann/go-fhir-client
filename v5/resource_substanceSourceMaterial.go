// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/SubstanceSourceMaterial
type SubstanceSourceMaterialPartDescription struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Entity of anatomical origin of source material within an organism
	Part *CodeableConcept `bson:"part,omitempty" json:"part,omitempty"`
	// The detailed anatomic location when the part can be extracted from different anatomical locations of the organism. Multiple alternative locations may apply
	PartLocation *CodeableConcept `bson:"partLocation,omitempty" json:"partLocation,omitempty"`
}

type SubstanceSourceMaterial struct {
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
	// General high level classification of the source material specific to the origin of the material
	SourceMaterialClass *CodeableConcept `bson:"sourceMaterialClass,omitempty" json:"sourceMaterialClass,omitempty"`
	// The type of the source material shall be specified based on a controlled vocabulary. For vaccines, this subclause refers to the class of infectious agent
	SourceMaterialType *CodeableConcept `bson:"sourceMaterialType,omitempty" json:"sourceMaterialType,omitempty"`
	// The state of the source material when extracted
	SourceMaterialState *CodeableConcept `bson:"sourceMaterialState,omitempty" json:"sourceMaterialState,omitempty"`
	// The unique identifier associated with the source material parent organism shall be specified
	OrganismId *Identifier `bson:"organismId,omitempty" json:"organismId,omitempty"`
	// The organism accepted Scientific name shall be provided based on the organism taxonomy
	OrganismName *string `bson:"organismName,omitempty" json:"organismName,omitempty"`
	// The parent of the herbal drug Ginkgo biloba, Leaf is the substance ID of the substance (fresh) of Ginkgo biloba L. or Ginkgo biloba L. (Whole plant)
	ParentSubstanceId []Identifier `bson:"parentSubstanceId,omitempty" json:"parentSubstanceId,omitempty"`
	// The parent substance of the Herbal Drug, or Herbal preparation
	ParentSubstanceName []string `bson:"parentSubstanceName,omitempty" json:"parentSubstanceName,omitempty"`
	// The country where the plant material is harvested or the countries where the plasma is sourced from as laid down in accordance with the Plasma Master File. For “Plasma-derived substances” the attribute country of origin provides information about the countries used for the manufacturing of the Cryopoor plama or Crioprecipitate
	CountryOfOrigin []CodeableConcept `bson:"countryOfOrigin,omitempty" json:"countryOfOrigin,omitempty"`
	// The place/region where the plant is harvested or the places/regions where the animal source material has its habitat
	GeographicalLocation []string `bson:"geographicalLocation,omitempty" json:"geographicalLocation,omitempty"`
	// Stage of life for animals, plants, insects and microorganisms. This information shall be provided only when the substance is significantly different in these stages (e.g. foetal bovine serum)
	DevelopmentStage *CodeableConcept `bson:"developmentStage,omitempty" json:"developmentStage,omitempty"`
	// Many complex materials are fractions of parts of plants, animals, or minerals. Fraction elements are often necessary to define both Substances and Specified Group 1 Substances. For substances derived from Plants, fraction information will be captured at the Substance information level ( . Oils, Juices and Exudates). Additional information for Extracts, such as extraction solvent composition, will be captured at the Specified Substance Group 1 information level. For plasma-derived products fraction information will be captured at the Substance and the Specified Substance Group 1 levels
	FractionDescription []SubstanceSourceMaterialFractionDescription `bson:"fractionDescription,omitempty" json:"fractionDescription,omitempty"`
	// This subclause describes the organism which the substance is derived from. For vaccines, the parent organism shall be specified based on these subclause elements. As an example, full taxonomy will be described for the Substance Name: ., Leaf
	Organism *SubstanceSourceMaterialOrganism `bson:"organism,omitempty" json:"organism,omitempty"`
	// To do
	PartDescription []SubstanceSourceMaterialPartDescription `bson:"partDescription,omitempty" json:"partDescription,omitempty"`
}

type SubstanceSourceMaterialFractionDescription struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// This element is capturing information about the fraction of a plant part, or human plasma for fractionation
	Fraction *string `bson:"fraction,omitempty" json:"fraction,omitempty"`
	// The specific type of the material constituting the component. For Herbal preparations the particulars of the extracts (liquid/dry) is described in Specified Substance Group 1
	MaterialType *CodeableConcept `bson:"materialType,omitempty" json:"materialType,omitempty"`
}

type SubstanceSourceMaterialOrganism struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The family of an organism shall be specified
	Family *CodeableConcept `bson:"family,omitempty" json:"family,omitempty"`
	// The genus of an organism shall be specified; refers to the Latin epithet of the genus element of the plant/animal scientific name; it is present in names for genera, species and infraspecies
	Genus *CodeableConcept `bson:"genus,omitempty" json:"genus,omitempty"`
	// The species of an organism shall be specified; refers to the Latin epithet of the species of the plant/animal; it is present in names for species and infraspecies
	Species *CodeableConcept `bson:"species,omitempty" json:"species,omitempty"`
	// The Intraspecific type of an organism shall be specified
	IntraspecificType *CodeableConcept `bson:"intraspecificType,omitempty" json:"intraspecificType,omitempty"`
	// The intraspecific description of an organism shall be specified based on a controlled vocabulary. For Influenza Vaccine, the intraspecific description shall contain the syntax of the antigen in line with the WHO convention
	IntraspecificDescription *string `bson:"intraspecificDescription,omitempty" json:"intraspecificDescription,omitempty"`
	// 4.9.13.6.1 Author type (Conditional)
	Author []SubstanceSourceMaterialOrganismAuthor `bson:"author,omitempty" json:"author,omitempty"`
	// 4.9.13.8.1 Hybrid species maternal organism ID (Optional)
	Hybrid *SubstanceSourceMaterialOrganismHybrid `bson:"hybrid,omitempty" json:"hybrid,omitempty"`
	// 4.9.13.7.1 Kingdom (Conditional)
	OrganismGeneral *SubstanceSourceMaterialOrganismOrganismGeneral `bson:"organismGeneral,omitempty" json:"organismGeneral,omitempty"`
}

type SubstanceSourceMaterialOrganismAuthor struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of author of an organism species shall be specified. The parenthetical author of an organism species refers to the first author who published the plant/animal name (of any rank). The primary author of an organism species refers to the first author(s), who validly published the plant/animal name
	AuthorType *CodeableConcept `bson:"authorType,omitempty" json:"authorType,omitempty"`
	// The author of an organism species shall be specified. The author year of an organism shall also be specified when applicable; refers to the year in which the first author(s) published the infraspecific plant/animal name (of any rank)
	AuthorDescription *string `bson:"authorDescription,omitempty" json:"authorDescription,omitempty"`
}

type SubstanceSourceMaterialOrganismHybrid struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The identifier of the maternal species constituting the hybrid organism shall be specified based on a controlled vocabulary. For plants, the parents aren’t always known, and it is unlikely that it will be known which is maternal and which is paternal
	MaternalOrganismId *string `bson:"maternalOrganismId,omitempty" json:"maternalOrganismId,omitempty"`
	// The name of the maternal species constituting the hybrid organism shall be specified. For plants, the parents aren’t always known, and it is unlikely that it will be known which is maternal and which is paternal
	MaternalOrganismName *string `bson:"maternalOrganismName,omitempty" json:"maternalOrganismName,omitempty"`
	// The identifier of the paternal species constituting the hybrid organism shall be specified based on a controlled vocabulary
	PaternalOrganismId *string `bson:"paternalOrganismId,omitempty" json:"paternalOrganismId,omitempty"`
	// The name of the paternal species constituting the hybrid organism shall be specified
	PaternalOrganismName *string `bson:"paternalOrganismName,omitempty" json:"paternalOrganismName,omitempty"`
	// The hybrid type of an organism shall be specified
	HybridType *CodeableConcept `bson:"hybridType,omitempty" json:"hybridType,omitempty"`
}

type SubstanceSourceMaterialOrganismOrganismGeneral struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The kingdom of an organism shall be specified
	Kingdom *CodeableConcept `bson:"kingdom,omitempty" json:"kingdom,omitempty"`
	// The phylum of an organism shall be specified
	Phylum *CodeableConcept `bson:"phylum,omitempty" json:"phylum,omitempty"`
	// The class of an organism shall be specified
	Class *CodeableConcept `bson:"class,omitempty" json:"class,omitempty"`
	// The order of an organism shall be specified,
	Order *CodeableConcept `bson:"order,omitempty" json:"order,omitempty"`
}

func (s SubstanceSourceMaterial) ResourceType() string {
	return "StructureDefinition"
}
