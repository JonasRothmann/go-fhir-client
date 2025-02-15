package main

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ElementDefinition
type ElementDefinitionSlicing struct {
	Id            *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	Discriminator []ElementDefinitionSlicingDiscriminator `bson:"discriminator,omitempty" json:"discriminator,omitempty"`
	Description   *string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Ordered       *bool                                   `bson:"ordered,omitempty" json:"ordered,omitempty"`
	Rules         custom.Code                             `bson:"rules" json:"rules"`
}

type ElementDefinitionBase struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Path      string      `bson:"path" json:"path"`
	Min       uint32      `bson:"min" json:"min"`
	Max       string      `bson:"max" json:"max"`
}

type ElementDefinitionExample struct {
	Id        *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	Label     string              `bson:"label" json:"label"`
	Value     custom.Base64binary `bson:"value" json:"value"`
}

type ElementDefinitionBinding struct {
	Id          *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	Strength    custom.Code                          `bson:"strength" json:"strength"`
	Description *custom.Markdown                     `bson:"description,omitempty" json:"description,omitempty"`
	ValueSet    *custom.Canonical[ValueSet]          `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	Additional  []ElementDefinitionBindingAdditional `bson:"additional,omitempty" json:"additional,omitempty"`
}

type ElementDefinition struct {
	Id                  *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Path                string                                  `bson:"path" json:"path"`
	Representation      []custom.Code                           `bson:"representation,omitempty" json:"representation,omitempty"`
	SliceName           *string                                 `bson:"sliceName,omitempty" json:"sliceName,omitempty"`
	SliceIsConstraining *bool                                   `bson:"sliceIsConstraining,omitempty" json:"sliceIsConstraining,omitempty"`
	Label               *string                                 `bson:"label,omitempty" json:"label,omitempty"`
	Code                []Coding                                `bson:"code,omitempty" json:"code,omitempty"`
	Slicing             *ElementDefinitionSlicing               `bson:"slicing,omitempty" json:"slicing,omitempty"`
	Short               *string                                 `bson:"short,omitempty" json:"short,omitempty"`
	Definition          *custom.Markdown                        `bson:"definition,omitempty" json:"definition,omitempty"`
	Comment             *custom.Markdown                        `bson:"comment,omitempty" json:"comment,omitempty"`
	Requirements        *custom.Markdown                        `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Alias               []string                                `bson:"alias,omitempty" json:"alias,omitempty"`
	Min                 *uint32                                 `bson:"min,omitempty" json:"min,omitempty"`
	Max                 *string                                 `bson:"max,omitempty" json:"max,omitempty"`
	Base                *ElementDefinitionBase                  `bson:"base,omitempty" json:"base,omitempty"`
	ContentReference    *custom.URI                             `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
	Type                []ElementDefinitionType                 `bson:"type,omitempty" json:"type,omitempty"`
	DefaultValue        *custom.Base64binary                    `bson:"defaultValue,omitempty" json:"defaultValue,omitempty"`
	MeaningWhenMissing  *custom.Markdown                        `bson:"meaningWhenMissing,omitempty" json:"meaningWhenMissing,omitempty"`
	OrderMeaning        *string                                 `bson:"orderMeaning,omitempty" json:"orderMeaning,omitempty"`
	Fixed               *custom.Base64binary                    `bson:"fixed,omitempty" json:"fixed,omitempty"`
	Pattern             *custom.Base64binary                    `bson:"pattern,omitempty" json:"pattern,omitempty"`
	Example             []ElementDefinitionExample              `bson:"example,omitempty" json:"example,omitempty"`
	MinValue            *custom.Date                            `bson:"minValue,omitempty" json:"minValue,omitempty"`
	MaxValue            *custom.Date                            `bson:"maxValue,omitempty" json:"maxValue,omitempty"`
	MaxLength           *int32                                  `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	Condition           []custom.ID                             `bson:"condition,omitempty" json:"condition,omitempty"`
	Constraint          []ElementDefinitionConstraint           `bson:"constraint,omitempty" json:"constraint,omitempty"`
	MustHaveValue       *bool                                   `bson:"mustHaveValue,omitempty" json:"mustHaveValue,omitempty"`
	ValueAlternatives   []custom.Canonical[StructureDefinition] `bson:"valueAlternatives,omitempty" json:"valueAlternatives,omitempty"`
	MustSupport         *bool                                   `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	IsModifier          *bool                                   `bson:"isModifier,omitempty" json:"isModifier,omitempty"`
	IsModifierReason    *string                                 `bson:"isModifierReason,omitempty" json:"isModifierReason,omitempty"`
	IsSummary           *bool                                   `bson:"isSummary,omitempty" json:"isSummary,omitempty"`
	Binding             *ElementDefinitionBinding               `bson:"binding,omitempty" json:"binding,omitempty"`
	Mapping             []ElementDefinitionMapping              `bson:"mapping,omitempty" json:"mapping,omitempty"`
}

type ElementDefinitionSlicingDiscriminator struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Type      custom.Code `bson:"type" json:"type"`
	Path      string      `bson:"path" json:"path"`
}

type ElementDefinitionType struct {
	Id            *string                                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension                                            `bson:"extension,omitempty" json:"extension,omitempty"`
	Code          custom.URI                                             `bson:"code" json:"code"`
	Profile       []custom.Canonical[ElementDefinitionTypeProfile]       `bson:"profile,omitempty" json:"profile,omitempty"`
	TargetProfile []custom.Canonical[ElementDefinitionTypeTargetProfile] `bson:"targetProfile,omitempty" json:"targetProfile,omitempty"`
	Aggregation   []custom.Code                                          `bson:"aggregation,omitempty" json:"aggregation,omitempty"`
	Versioning    *custom.Code                                           `bson:"versioning,omitempty" json:"versioning,omitempty"`
}

type ElementDefinitionConstraint struct {
	Id           *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension    []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	Key          custom.ID                              `bson:"key" json:"key"`
	Requirements *custom.Markdown                       `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Severity     custom.Code                            `bson:"severity" json:"severity"`
	Suppress     *bool                                  `bson:"suppress,omitempty" json:"suppress,omitempty"`
	Human        string                                 `bson:"human" json:"human"`
	Expression   *string                                `bson:"expression,omitempty" json:"expression,omitempty"`
	Source       *custom.Canonical[StructureDefinition] `bson:"source,omitempty" json:"source,omitempty"`
}

type ElementDefinitionBindingAdditional struct {
	Id            *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	Purpose       custom.Code                `bson:"purpose" json:"purpose"`
	ValueSet      custom.Canonical[ValueSet] `bson:"valueSet" json:"valueSet"`
	Documentation *custom.Markdown           `bson:"documentation,omitempty" json:"documentation,omitempty"`
	ShortDoco     *string                    `bson:"shortDoco,omitempty" json:"shortDoco,omitempty"`
	//Usage         []UsageContext             `bson:"usage,omitempty" json:"usage,omitempty"`
	Any *bool `bson:"any,omitempty" json:"any,omitempty"`
}

type ElementDefinitionMapping struct {
	Id        *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	Identity  custom.ID        `bson:"identity" json:"identity"`
	Language  *custom.Code     `bson:"language,omitempty" json:"language,omitempty"`
	Map       string           `bson:"map" json:"map"`
	Comment   *custom.Markdown `bson:"comment,omitempty" json:"comment,omitempty"`
}

type ElementDefinitionTypeProfile interface {
	gofhirclient.Resource

	Is_ElementDefinitionTypeProfile()
}

func (s StructureDefinition) Is_ElementDefinitionTypeProfile() {}

type ElementDefinitionTypeTargetProfile interface {
	gofhirclient.Resource

	Is_ElementDefinitionTypeTargetProfile()
}

func (s StructureDefinition) Is_ElementDefinitionTypeTargetProfile() {}

func (e ElementDefinition) ResourceType() string {
	return "StructureDefinition"
}
