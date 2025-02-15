// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ResearchSubject
type ResearchSubjectProgress struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// state | milestone
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// candidate | eligible | follow-up | ineligible | not-registered | off-study | on-study | on-study-intervention | on-study-observation | pending-on-study | potential-candidate | screening | withdrawn
	SubjectState *CodeableConcept `bson:"subjectState,omitempty" json:"subjectState,omitempty"`
	// SignedUp | Screened | Randomized
	Milestone *CodeableConcept `bson:"milestone,omitempty" json:"milestone,omitempty"`
	// State change reason
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	// State change date
	StartDate *custom.DateTime `bson:"startDate,omitempty" json:"startDate,omitempty"`
	// State change date
	EndDate *custom.DateTime `bson:"endDate,omitempty" json:"endDate,omitempty"`
}

type ResearchSubject struct {
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
	// Business Identifier for research subject in a study
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Subject status
	Progress []ResearchSubjectProgress `bson:"progress,omitempty" json:"progress,omitempty"`
	// Start and end of participation
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Study subject is part of
	Study custom.Reference[ResearchStudy] `bson:"study" json:"study"`
	// Who or what is part of study
	Subject custom.Reference[ResearchSubjectSubject] `bson:"subject" json:"subject"`
	// What path should be followed
	AssignedComparisonGroup *custom.ID `bson:"assignedComparisonGroup,omitempty" json:"assignedComparisonGroup,omitempty"`
	// What path was followed
	ActualComparisonGroup *custom.ID `bson:"actualComparisonGroup,omitempty" json:"actualComparisonGroup,omitempty"`
	// Agreement to participate in study
	Consent []custom.Reference[Consent] `bson:"consent,omitempty" json:"consent,omitempty"`
}

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
	return "StructureDefinition"
}
