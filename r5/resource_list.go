// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/List
type List struct {
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
	// Business identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// current | retired | entered-in-error
	Status custom.Code `json:"status"`
	// working | snapshot | changes
	Mode custom.Code `json:"mode"`
	// Descriptive name for the list
	Title *string `json:"title,omitempty"`
	// What the purpose of this list is
	Code *CodeableConcept `json:"code,omitempty"`
	// If all resources have the same subject(s)
	Subject []custom.Reference[Resource] `json:"subject,omitempty"`
	// Context in which list created
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When the list was prepared
	Date *custom.DateTime `json:"date,omitempty"`
	// Who and/or what defined the list contents (aka Author)
	Source *custom.Reference[ListSource] `json:"source,omitempty"`
	// What order the list has
	OrderedBy *CodeableConcept `json:"orderedBy,omitempty"`
	// Comments about the list
	Note []Annotation `json:"note,omitempty"`
	// Entries in the list
	Entry []ListEntry `json:"entry,omitempty"`
	// Why list is empty
	EmptyReason *CodeableConcept `json:"emptyReason,omitempty"`
}

type ListEntry struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Status/Workflow information about this item
	Flag *CodeableConcept `json:"flag,omitempty"`
	// If this item is actually marked as deleted
	Deleted *bool `json:"deleted,omitempty"`
	// When item added to list
	Date *custom.DateTime `json:"date,omitempty"`
	// Actual entry
	Item custom.Reference[Resource] `json:"item"`
}

type OtherList List

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
	return "List"
}

func (l List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherList
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherList: OtherList(l), ResourceType: l.ResourceType()})
}

func UnmarshalList(b []byte) (res List, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
