// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ResearchStudy
type ResearchStudy struct {
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
	// Canonical identifier for this study resource
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Business Identifier for study
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// The business version for the study record
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// Name for this study (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Human readable name of the study
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Additional names for the study
	Label []ResearchStudyLabel `bson:"label,omitempty" json:"label,omitempty"`
	// Steps followed in executing study
	Protocol []custom.Reference[PlanDefinition] `bson:"protocol,omitempty" json:"protocol,omitempty"`
	// Part of larger study
	PartOf []custom.Reference[ResearchStudy] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// References, URLs, and attachments
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// Date the resource last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// treatment | prevention | diagnostic | supportive-care | screening | health-services-research | basic-science | device-feasibility
	PrimaryPurposeType *CodeableConcept `bson:"primaryPurposeType,omitempty" json:"primaryPurposeType,omitempty"`
	// n-a | early-phase-1 | phase-1 | phase-1-phase-2 | phase-2 | phase-2-phase-3 | phase-3 | phase-4
	Phase *CodeableConcept `bson:"phase,omitempty" json:"phase,omitempty"`
	// Classifications of the study design characteristics
	StudyDesign []CodeableConcept `bson:"studyDesign,omitempty" json:"studyDesign,omitempty"`
	// Drugs, devices, etc. under study
	Focus []custom.CodeableReference[ResearchStudyFocus] `bson:"focus,omitempty" json:"focus,omitempty"`
	// Condition being studied
	Condition []CodeableConcept `bson:"condition,omitempty" json:"condition,omitempty"`
	// Used to search for the study
	Keyword []CodeableConcept `bson:"keyword,omitempty" json:"keyword,omitempty"`
	// Geographic area for the study
	Region []CodeableConcept `bson:"region,omitempty" json:"region,omitempty"`
	// Brief text explaining the study
	DescriptionSummary *custom.Markdown `bson:"descriptionSummary,omitempty" json:"descriptionSummary,omitempty"`
	// Detailed narrative of the study
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// When the study began and ended
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Facility where study activities are conducted
	Site []custom.Reference[ResearchStudySite] `bson:"site,omitempty" json:"site,omitempty"`
	// Comments made about the study
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Classification for the study
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
	// Sponsors, collaborators, and other parties
	AssociatedParty []ResearchStudyAssociatedParty `bson:"associatedParty,omitempty" json:"associatedParty,omitempty"`
	// Status of study with time for that status
	ProgressStatus []ResearchStudyProgressStatus `bson:"progressStatus,omitempty" json:"progressStatus,omitempty"`
	// accrual-goal-met | closed-due-to-toxicity | closed-due-to-lack-of-study-progress | temporarily-closed-per-study-design
	WhyStopped *CodeableConcept `bson:"whyStopped,omitempty" json:"whyStopped,omitempty"`
	// Target or actual group of participants enrolled in study
	Recruitment *ResearchStudyRecruitment `bson:"recruitment,omitempty" json:"recruitment,omitempty"`
	// Defined path through the study for a subject
	ComparisonGroup []ResearchStudyComparisonGroup `bson:"comparisonGroup,omitempty" json:"comparisonGroup,omitempty"`
	// A goal for the study
	Objective []ResearchStudyObjective `bson:"objective,omitempty" json:"objective,omitempty"`
	// A variable measured during the study
	OutcomeMeasure []ResearchStudyOutcomeMeasure `bson:"outcomeMeasure,omitempty" json:"outcomeMeasure,omitempty"`
	// Link to results generated during the study
	Result []custom.Reference[ResearchStudyResult] `bson:"result,omitempty" json:"result,omitempty"`
}

type ResearchStudyLabel struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// primary | official | scientific | plain-language | subtitle | short-title | acronym | earlier-title | language | auto-translated | human-use | machine-use | duplicate-uid
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The name
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
}

type ResearchStudyAssociatedParty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name of associated party
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// sponsor | lead-sponsor | sponsor-investigator | primary-investigator | collaborator | funding-source | general-contact | recruitment-contact | sub-investigator | study-director | study-chair
	Role CodeableConcept `bson:"role" json:"role"`
	// When active in the role
	Period []Period `bson:"period,omitempty" json:"period,omitempty"`
	// nih | fda | government | nonprofit | academic | industry
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
	// Individual or organization associated with study (use practitionerRole to specify their organisation)
	Party *custom.Reference[ResearchStudyAssociatedPartyParty] `bson:"party,omitempty" json:"party,omitempty"`
}

type ResearchStudyProgressStatus struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Label for status or state (e.g. recruitment status)
	State CodeableConcept `bson:"state" json:"state"`
	// Actual if true else anticipated
	Actual *bool `bson:"actual,omitempty" json:"actual,omitempty"`
	// Date range
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type ResearchStudyRecruitment struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Estimated total number of participants to be enrolled
	TargetNumber *uint32 `bson:"targetNumber,omitempty" json:"targetNumber,omitempty"`
	// Actual total number of participants enrolled in study
	ActualNumber *uint32 `bson:"actualNumber,omitempty" json:"actualNumber,omitempty"`
	// Inclusion and exclusion criteria
	Eligibility *custom.Reference[ResearchStudyRecruitmentEligibility] `bson:"eligibility,omitempty" json:"eligibility,omitempty"`
	// Group of participants who were enrolled in study
	ActualGroup *custom.Reference[Group] `bson:"actualGroup,omitempty" json:"actualGroup,omitempty"`
}

type ResearchStudyComparisonGroup struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Allows the comparisonGroup for the study and the comparisonGroup for the subject to be linked easily
	LinkId *custom.ID `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// Label for study comparisonGroup
	Name string `bson:"name" json:"name"`
	// Categorization of study comparisonGroup
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Short explanation of study path
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Interventions or exposures in this comparisonGroup or cohort
	IntendedExposure []custom.Reference[EvidenceVariable] `bson:"intendedExposure,omitempty" json:"intendedExposure,omitempty"`
	// Group of participants who were enrolled in study comparisonGroup
	ObservedGroup *custom.Reference[Group] `bson:"observedGroup,omitempty" json:"observedGroup,omitempty"`
}

type ResearchStudyObjective struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Label for the objective
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// primary | secondary | exploratory
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Description of the objective
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
}

type ResearchStudyOutcomeMeasure struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Label for the outcome
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// primary | secondary | exploratory
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Description of the outcome
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Structured outcome definition
	Reference *custom.Reference[EvidenceVariable] `bson:"reference,omitempty" json:"reference,omitempty"`
}

type ResearchStudyFocus interface {
	gofhirclient.Resource

	Is_ResearchStudyFocus()
}

func (m Medication) Is_ResearchStudyFocus()                 {}
func (m MedicinalProductDefinition) Is_ResearchStudyFocus() {}
func (s SubstanceDefinition) Is_ResearchStudyFocus()        {}
func (e EvidenceVariable) Is_ResearchStudyFocus()           {}

type ResearchStudySite interface {
	gofhirclient.Resource

	Is_ResearchStudySite()
}

func (l Location) Is_ResearchStudySite()      {}
func (r ResearchStudy) Is_ResearchStudySite() {}
func (o Organization) Is_ResearchStudySite()  {}

type ResearchStudyAssociatedPartyParty interface {
	gofhirclient.Resource

	Is_ResearchStudyAssociatedPartyParty()
}

func (p Practitioner) Is_ResearchStudyAssociatedPartyParty()     {}
func (p PractitionerRole) Is_ResearchStudyAssociatedPartyParty() {}
func (o Organization) Is_ResearchStudyAssociatedPartyParty()     {}

type ResearchStudyRecruitmentEligibility interface {
	gofhirclient.Resource

	Is_ResearchStudyRecruitmentEligibility()
}

func (g Group) Is_ResearchStudyRecruitmentEligibility()            {}
func (e EvidenceVariable) Is_ResearchStudyRecruitmentEligibility() {}

type ResearchStudyResult interface {
	gofhirclient.Resource

	Is_ResearchStudyResult()
}

func (e EvidenceReport) Is_ResearchStudyResult()   {}
func (c Citation) Is_ResearchStudyResult()         {}
func (d DiagnosticReport) Is_ResearchStudyResult() {}

func (r ResearchStudy) ResourceType() string {
	return "StructureDefinition"
}
