package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jonasrothmann/go-fhir-client/custom"
	"github.com/rs/zerolog/log"
)

type StructureDefinitionKind string

const (
	StructureDefinitionKindComplex   StructureDefinitionKind = "complex-type"
	StructureDefinitionKindPrimitive StructureDefinitionKind = "primitive-type"
	StructureDefinitionKindResource  StructureDefinitionKind = "resource"
)

func (s StructureDefinitionKind) Path() string {
	return fmt.Sprintf("github.com/jonasrothmann/go-fhir-client/r5")
}

func (s StructureDefinitionKind) Name() string {
	switch s {
	case StructureDefinitionKindComplex, StructureDefinitionKindPrimitive:
		return "datatypes"
	case StructureDefinitionKindResource:
		return "resources"
	default:
		log.Fatal().Msg("unknown StructureDefinitionKind")
	}

	return ""
}

type StructureDefinition struct {
	Id           *string                 `json:"id,omitempty"`
	ResourceType string                  `json:"resourceType"`
	Name         *string                 `json:"name,omitempty"`
	Url          *string                 `json:"url,omitempty"`
	Kind         StructureDefinitionKind `json:"kind,omitempty"`
	Snapshot     SnapshotDefinition      `json:"snapshot"`
}

type SnapshotDefinition struct {
	Element []ElementDefinition `json:"element"`
}

type TypeRef struct {
	Code          string   `json:"code"`
	TargetProfile []string `json:"targetProfile"`
	// Additional fields (Profile, TargetProfile, etc.) can be added if required.
}

type ResourceCollections struct {
	StructureDefinitions map[string]StructureDefinition
	// You could add ValueSets and CodeSystems here as needed.
}

type Bundle struct {
	Id            *string `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *string `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *string `bson:"language,omitempty" json:"language,omitempty"`
	//Identifier    *datatypes.Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type      string        `bson:"type" json:"type"`
	Timestamp *string       `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Total     *uint32       `bson:"total,omitempty" json:"total,omitempty"`
	Link      []BundleLink  `bson:"link,omitempty" json:"link,omitempty"`
	Entry     []BundleEntry `bson:"entry,omitempty" json:"entry,omitempty"`
	//Signature     *Signature  `bson:"signature,omitempty" json:"signature,omitempty"`
	Issues *json.RawMessage `bson:"issues,omitempty" json:"issues,omitempty"`
}
type BundleLink struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relation          string      `bson:"relation" json:"relation"`
	Url               string      `bson:"url" json:"url"`
}
type BundleEntry struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Link              []interface{}        `bson:"link,omitempty" json:"link,omitempty"`
	FullUrl           *string              `bson:"fullUrl,omitempty" json:"fullUrl,omitempty"`
	Resource          *json.RawMessage     `bson:"resource,omitempty" json:"resource,omitempty"`
	Search            *BundleEntrySearch   `bson:"search,omitempty" json:"search,omitempty"`
	Request           *BundleEntryRequest  `bson:"request,omitempty" json:"request,omitempty"`
	Response          *BundleEntryResponse `bson:"response,omitempty" json:"response,omitempty"`
}
type BundleEntrySearch struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Mode              *string      `bson:"mode,omitempty" json:"mode,omitempty"`
	Score             *json.Number `bson:"score,omitempty" json:"score,omitempty"`
}
type BundleEntryRequest struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Method            string      `bson:"method" json:"method"`
	Url               string      `bson:"url" json:"url"`
	IfNoneMatch       *string     `bson:"ifNoneMatch,omitempty" json:"ifNoneMatch,omitempty"`
	IfModifiedSince   *time.Time  `bson:"ifModifiedSince,omitempty" json:"ifModifiedSince,omitempty"`
	IfMatch           *string     `bson:"ifMatch,omitempty" json:"ifMatch,omitempty"`
	IfNoneExist       *string     `bson:"ifNoneExist,omitempty" json:"ifNoneExist,omitempty"`
}
type BundleEntryResponse struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            string           `bson:"status" json:"status"`
	Location          *string          `bson:"location,omitempty" json:"location,omitempty"`
	Etag              *string          `bson:"etag,omitempty" json:"etag,omitempty"`
	LastModified      *time.Time       `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	Outcome           *json.RawMessage `bson:"outcome,omitempty" json:"outcome,omitempty"`
}

type Coding struct {
	Id           *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension    []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	System       *custom.URI  `bson:"system,omitempty" json:"system,omitempty"`
	Version      *string      `bson:"version,omitempty" json:"version,omitempty"`
	Code         *custom.Code `bson:"code,omitempty" json:"code,omitempty"`
	Display      *string      `bson:"display,omitempty" json:"display,omitempty"`
	UserSelected *bool        `bson:"userSelected,omitempty" json:"userSelected,omitempty"`
}

type Meta struct {
	Id          *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	VersionId   *custom.ID                              `bson:"versionId,omitempty" json:"versionId,omitempty"`
	LastUpdated *custom.Instant                         `bson:"lastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Source      *custom.URI                             `bson:"source,omitempty" json:"source,omitempty"`
	Profile     []custom.Canonical[StructureDefinition] `bson:"profile,omitempty" json:"profile,omitempty"`
	Security    []Coding                                `bson:"security,omitempty" json:"security,omitempty"`
	Tag         []Coding                                `bson:"tag,omitempty" json:"tag,omitempty"`
}

// http://hl7.org/fhir/StructureDefinition/Extension
type Extension struct {
	Id        *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	Url       string               `bson:"url" json:"url"`
	Value     *custom.Base64binary `bson:"value,omitempty" json:"value,omitempty"`
}

// http://hl7.org/fhir/StructureDefinition/ValueSet
type ValueSetCompose struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LockedDate        *custom.Date             `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	Inactive          *bool                    `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Include           []ValueSetComposeInclude `bson:"include" json:"include"`
	Exclude           []interface{}            `bson:"exclude,omitempty" json:"exclude,omitempty"`
	Property          []string                 `bson:"property,omitempty" json:"property,omitempty"`
}

type ValueSetComposeIncludeConcept struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              custom.Code                                `bson:"code" json:"code"`
	Display           *string                                    `bson:"display,omitempty" json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
}

type ValueSetComposeIncludeConceptDesignation struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	Use               *Coding      `bson:"use,omitempty" json:"use,omitempty"`
	AdditionalUse     []Coding     `bson:"additionalUse,omitempty" json:"additionalUse,omitempty"`
	Value             string       `bson:"value" json:"value"`
}

type ValueSetComposeIncludeFilter struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Property          custom.Code `bson:"property" json:"property"`
	Op                custom.Code `bson:"op" json:"op"`
	Value             string      `bson:"value" json:"value"`
}

type ValueSetExpansionParameter struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Value             *string     `bson:"value,omitempty" json:"value,omitempty"`
}

type ValueSetExpansionProperty struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              custom.Code `bson:"code" json:"code"`
	Uri               *custom.URI `bson:"uri,omitempty" json:"uri,omitempty"`
}

type ValueSetExpansionContainsPropertySubProperty struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              custom.Code `bson:"code" json:"code"`
	Value             custom.Code `bson:"value" json:"value"`
}

type ValueSet struct {
	Id            *string      `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *custom.URI  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	//Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	//Contained         []Resource         `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	//Identifier        []Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version          *string          `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithm *string          `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	Name             *string          `bson:"name,omitempty" json:"name,omitempty"`
	Title            *string          `bson:"title,omitempty" json:"title,omitempty"`
	Status           custom.Code      `bson:"status" json:"status"`
	Experimental     *bool            `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date             *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	Publisher        *string          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	//Contact           []ContactDetail    `bson:"contact,omitempty" json:"contact,omitempty"`
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	//UseContext        []UsageContext     `bson:"useContext,omitempty" json:"useContext,omitempty"`
	//Jurisdiction      []CodeableConcept  `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Immutable      *bool            `bson:"immutable,omitempty" json:"immutable,omitempty"`
	Purpose        *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright      *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel *string          `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	ApprovalDate   *custom.Date     `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate *custom.Date     `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	//EffectivePeriod   *Period            `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	//Topic             []CodeableConcept  `bson:"topic,omitempty" json:"topic,omitempty"`
	//Author            []ContactDetail    `bson:"author,omitempty" json:"author,omitempty"`
	//Editor            []ContactDetail    `bson:"editor,omitempty" json:"editor,omitempty"`
	//Reviewer          []ContactDetail    `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	//Endorser          []ContactDetail    `bson:"endorser,omitempty" json:"endorser,omitempty"`
	//RelatedArtifact   []RelatedArtifact  `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Compose   *ValueSetCompose   `bson:"compose,omitempty" json:"compose,omitempty"`
	Expansion *ValueSetExpansion `bson:"expansion,omitempty" json:"expansion,omitempty"`
	Scope     *ValueSetScope     `bson:"scope,omitempty" json:"scope,omitempty"`
}

type ValueSetComposeInclude struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            *custom.URI                     `bson:"system,omitempty" json:"system,omitempty"`
	Version           *string                         `bson:"version,omitempty" json:"version,omitempty"`
	Concept           []ValueSetComposeIncludeConcept `bson:"concept,omitempty" json:"concept,omitempty"`
	Filter            []ValueSetComposeIncludeFilter  `bson:"filter,omitempty" json:"filter,omitempty"`
	ValueSet          []custom.Canonical[ValueSet]    `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	Copyright         *string                         `bson:"copyright,omitempty" json:"copyright,omitempty"`
}

type ValueSetExpansion struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *custom.URI                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Next              *custom.URI                  `bson:"next,omitempty" json:"next,omitempty"`
	Timestamp         custom.DateTime              `bson:"timestamp" json:"timestamp"`
	Total             *int32                       `bson:"total,omitempty" json:"total,omitempty"`
	Offset            *int32                       `bson:"offset,omitempty" json:"offset,omitempty"`
	Parameter         []ValueSetExpansionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Property          []ValueSetExpansionProperty  `bson:"property,omitempty" json:"property,omitempty"`
	Contains          []ValueSetExpansionContains  `bson:"contains,omitempty" json:"contains,omitempty"`
}

type ValueSetExpansionContains struct {
	Id                *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            *custom.URI                         `bson:"system,omitempty" json:"system,omitempty"`
	Abstract          *bool                               `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Inactive          *bool                               `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Version           *string                             `bson:"version,omitempty" json:"version,omitempty"`
	Code              *custom.Code                        `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                             `bson:"display,omitempty" json:"display,omitempty"`
	Designation       []interface{}                       `bson:"designation,omitempty" json:"designation,omitempty"`
	Property          []ValueSetExpansionContainsProperty `bson:"property,omitempty" json:"property,omitempty"`
	Contains          []interface{}                       `bson:"contains,omitempty" json:"contains,omitempty"`
}

type ValueSetExpansionContainsProperty struct {
	Id                *string                                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              custom.Code                                    `bson:"code" json:"code"`
	Value             custom.Code                                    `bson:"value" json:"value"`
	SubProperty       []ValueSetExpansionContainsPropertySubProperty `bson:"subProperty,omitempty" json:"subProperty,omitempty"`
}

type ValueSetScope struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	InclusionCriteria *string     `bson:"inclusionCriteria,omitempty" json:"inclusionCriteria,omitempty"`
	ExclusionCriteria *string     `bson:"exclusionCriteria,omitempty" json:"exclusionCriteria,omitempty"`
}

// http://hl7.org/fhir/StructureDefinition/CodeSystem
type CodeSystem struct {
	Id            *string      `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *custom.URI  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	//Text              *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	//Contained         []Resource                    `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	//Identifier        []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version          *string          `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithm *string          `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	Name             *string          `bson:"name,omitempty" json:"name,omitempty"`
	Title            *string          `bson:"title,omitempty" json:"title,omitempty"`
	Status           custom.Code      `bson:"status" json:"status"`
	Experimental     *bool            `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date             *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	Publisher        *string          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	//Contact           []ContactDetail               `bson:"contact,omitempty" json:"contact,omitempty"`
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	//UseContext        []UsageContext                `bson:"useContext,omitempty" json:"useContext,omitempty"`
	//Jurisdiction      []CodeableConcept             `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose        *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright      *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel *string          `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	ApprovalDate   *custom.Date     `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate *custom.Date     `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	/*EffectivePeriod   *Period                       `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Topic             []CodeableConcept             `bson:"topic,omitempty" json:"topic,omitempty"`
	Author            []ContactDetail               `bson:"author,omitempty" json:"author,omitempty"`
	Editor            []ContactDetail               `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer          []ContactDetail               `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser          []ContactDetail               `bson:"endorser,omitempty" json:"endorser,omitempty"`
	RelatedArtifact   []RelatedArtifact             `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	*/CaseSensitive  *bool                         `bson:"caseSensitive,omitempty" json:"caseSensitive,omitempty"`
	ValueSet         *custom.Canonical[ValueSet]   `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	HierarchyMeaning *custom.Code                  `bson:"hierarchyMeaning,omitempty" json:"hierarchyMeaning,omitempty"`
	Compositional    *bool                         `bson:"compositional,omitempty" json:"compositional,omitempty"`
	VersionNeeded    *bool                         `bson:"versionNeeded,omitempty" json:"versionNeeded,omitempty"`
	Content          custom.Code                   `bson:"content" json:"content"`
	Supplements      *custom.Canonical[CodeSystem] `bson:"supplements,omitempty" json:"supplements,omitempty"`
	Count            *uint32                       `bson:"count,omitempty" json:"count,omitempty"`
	Filter           []CodeSystemFilter            `bson:"filter,omitempty" json:"filter,omitempty"`
	Property         []CodeSystemProperty          `bson:"property,omitempty" json:"property,omitempty"`
	Concept          []CodeSystemConcept           `bson:"concept,omitempty" json:"concept,omitempty"`
}

type CodeSystemFilter struct {
	Id                *string       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              custom.Code   `bson:"code" json:"code"`
	Description       *string       `bson:"description,omitempty" json:"description,omitempty"`
	Operator          []custom.Code `bson:"operator" json:"operator"`
	Value             string        `bson:"value" json:"value"`
}

type CodeSystemProperty struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              custom.Code `bson:"code" json:"code"`
	Uri               *custom.URI `bson:"uri,omitempty" json:"uri,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Type              custom.Code `bson:"type" json:"type"`
}

type CodeSystemConcept struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              custom.Code                    `bson:"code" json:"code"`
	Display           *string                        `bson:"display,omitempty" json:"display,omitempty"`
	Definition        *string                        `bson:"definition,omitempty" json:"definition,omitempty"`
	Designation       []CodeSystemConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
	Property          []CodeSystemConceptProperty    `bson:"property,omitempty" json:"property,omitempty"`
	Concept           []interface{}                  `bson:"concept,omitempty" json:"concept,omitempty"`
}

func (c CodeSystemConcept) GetProperty(name string) (parent CodeSystemConceptProperty, ok bool) {
	for _, property := range c.Property {
		if string(property.Code) == name {
			return property, true
		}
	}

	return parent, false
}

type CodeSystemConceptDesignation struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	Use               *Coding      `bson:"use,omitempty" json:"use,omitempty"`
	AdditionalUse     []Coding     `bson:"additionalUse,omitempty" json:"additionalUse,omitempty"`
	Value             string       `bson:"value" json:"value"`
}

type CodeSystemConceptProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to CodeSystem.property.code
	Code custom.Code `bson:"code" json:"code"`
	// Value of the property for this concept
	ValueCode *custom.Code `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	// Value of the property for this concept
	ValueCoding *Coding `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	// Value of the property for this concept
	ValueString *string `bson:"valueString,omitempty" json:"valueString,omitempty"`
	// Value of the property for this concept
	ValueInteger *int32 `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	// Value of the property for this concept
	ValueBoolean *bool `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	// Value of the property for this concept
	ValueDateTime *custom.DateTime `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	// Value of the property for this concept
	ValueDecimal *json.Number `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
}

func (c CodeSystem) ResourceType() string {
	return "StructureDefinition"
}
