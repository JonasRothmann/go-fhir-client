// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/List
type List struct {
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
	// Business identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// current | retired | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// working | snapshot | changes
	Mode custom.Code `bson:"mode" json:"mode"`
	// Descriptive name for the list
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// What the purpose of this list is
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// If all resources have the same subject(s)
	Subject []custom.Reference[Resource] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Context in which list created
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When the list was prepared
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Who and/or what defined the list contents (aka Author)
	Source *custom.Reference[ListSource] `bson:"source,omitempty" json:"source,omitempty"`
	// What order the list has
	OrderedBy *CodeableConcept `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	// Comments about the list
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Entries in the list
	Entry []ListEntry `bson:"entry,omitempty" json:"entry,omitempty"`
	// Why list is empty
	EmptyReason *CodeableConcept `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
}

type ListEntry struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Status/Workflow information about this item
	Flag *CodeableConcept `bson:"flag,omitempty" json:"flag,omitempty"`
	// If this item is actually marked as deleted
	Deleted *bool `bson:"deleted,omitempty" json:"deleted,omitempty"`
	// When item added to list
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Actual entry
	Item custom.Reference[Resource] `bson:"item" json:"item"`
}

type ListSource interface {
	gofhirclient.Resource

	Is_ListSource()
}

func (p Practitioner) Is_ListSource()     {}
func (p PractitionerRole) Is_ListSource() {}
func (p Patient) Is_ListSource()          {}
func (d Device) Is_ListSource()           {}
func (o Organization) Is_ListSource()     {}
func (r RelatedPerson) Is_ListSource()    {}
func (c CareTeam) Is_ListSource()         {}

func (l List) ResourceType() string {
	return "StructureDefinition"
}
