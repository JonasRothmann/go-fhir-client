// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ResearchSubject
type ResearchSubject struct {
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
	// Business Identifier for research subject in a study
	Identifier []Identifier `json:"identifier,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// Subject status
	Progress []ResearchSubjectProgress `json:"progress,omitempty"`
	// Start and end of participation
	Period *Period `json:"period,omitempty"`
	// Study subject is part of
	Study custom.Reference[ResearchStudy] `json:"study"`
	// Who or what is part of study
	Subject custom.Reference[ResearchSubjectSubject] `json:"subject"`
	// What path should be followed
	AssignedComparisonGroup *custom.ID `json:"assignedComparisonGroup,omitempty"`
	// What path was followed
	ActualComparisonGroup *custom.ID `json:"actualComparisonGroup,omitempty"`
	// Agreement to participate in study
	Consent []custom.Reference[Consent] `json:"consent,omitempty"`
}

type ResearchSubjectProgress struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// state | milestone
	Type *CodeableConcept `json:"type,omitempty"`
	// candidate | eligible | follow-up | ineligible | not-registered | off-study | on-study | on-study-intervention | on-study-observation | pending-on-study | potential-candidate | screening | withdrawn
	SubjectState *CodeableConcept `json:"subjectState,omitempty"`
	// SignedUp | Screened | Randomized
	Milestone *CodeableConcept `json:"milestone,omitempty"`
	// State change reason
	Reason *CodeableConcept `json:"reason,omitempty"`
	// State change date
	StartDate *custom.DateTime `json:"startDate,omitempty"`
	// State change date
	EndDate *custom.DateTime `json:"endDate,omitempty"`
}

type OtherResearchSubject ResearchSubject

type ResearchSubjectSubject interface {
	gofhirclient.Resource

	Is_ResearchSubjectSubject()
}

func (p Patient) Is_ResearchSubjectSubject()                    {}
func (g Group) Is_ResearchSubjectSubject()                      {}
func (s Specimen) Is_ResearchSubjectSubject()                   {}
func (d Device) Is_ResearchSubjectSubject()                     {}
func (m Medication) Is_ResearchSubjectSubject()                 {}
func (s Substance) Is_ResearchSubjectSubject()                  {}
func (b BiologicallyDerivedProduct) Is_ResearchSubjectSubject() {}

func (r ResearchSubject) ResourceType() string {
	return "ResearchSubject"
}

func (r ResearchSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchSubject
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherResearchSubject: OtherResearchSubject(r), ResourceType: r.ResourceType()})
}

func UnmarshalResearchSubject(b []byte) (res ResearchSubject, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
