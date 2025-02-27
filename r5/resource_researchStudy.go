// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ResearchStudy
type ResearchStudyLabel struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// primary | official | scientific | plain-language | subtitle | short-title | acronym | earlier-title | language | auto-translated | human-use | machine-use | duplicate-uid
	Type *CodeableConcept `json:"type,omitempty"`
	// The name
	Value *string `json:"value,omitempty"`
}

type ResearchStudyAssociatedParty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name of associated party
	Name *string `json:"name,omitempty"`
	// sponsor | lead-sponsor | sponsor-investigator | primary-investigator | collaborator | funding-source | general-contact | recruitment-contact | sub-investigator | study-director | study-chair
	Role CodeableConcept `json:"role"`
	// When active in the role
	Period []Period `json:"period,omitempty"`
	// nih | fda | government | nonprofit | academic | industry
	Classifier []CodeableConcept `json:"classifier,omitempty"`
	// Individual or organization associated with study (use practitionerRole to specify their organisation)
	Party *custom.Reference[ResearchStudyAssociatedPartyParty] `json:"party,omitempty"`
}

type ResearchStudyProgressStatus struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Label for status or state (e.g. recruitment status)
	State CodeableConcept `json:"state"`
	// Actual if true else anticipated
	Actual *bool `json:"actual,omitempty"`
	// Date range
	Period *Period `json:"period,omitempty"`
}

type ResearchStudyRecruitment struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Estimated total number of participants to be enrolled
	TargetNumber *uint32 `json:"targetNumber,omitempty"`
	// Actual total number of participants enrolled in study
	ActualNumber *uint32 `json:"actualNumber,omitempty"`
	// Inclusion and exclusion criteria
	Eligibility *custom.Reference[ResearchStudyRecruitmentEligibility] `json:"eligibility,omitempty"`
	// Group of participants who were enrolled in study
	ActualGroup *custom.Reference[Group] `json:"actualGroup,omitempty"`
}

type ResearchStudyComparisonGroup struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Allows the comparisonGroup for the study and the comparisonGroup for the subject to be linked easily
	LinkId *custom.ID `json:"linkId,omitempty"`
	// Label for study comparisonGroup
	Name string `json:"name"`
	// Categorization of study comparisonGroup
	Type *CodeableConcept `json:"type,omitempty"`
	// Short explanation of study path
	Description *custom.Markdown `json:"description,omitempty"`
	// Interventions or exposures in this comparisonGroup or cohort
	IntendedExposure []custom.Reference[EvidenceVariable] `json:"intendedExposure,omitempty"`
	// Group of participants who were enrolled in study comparisonGroup
	ObservedGroup *custom.Reference[Group] `json:"observedGroup,omitempty"`
}

type ResearchStudyObjective struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Label for the objective
	Name *string `json:"name,omitempty"`
	// primary | secondary | exploratory
	Type *CodeableConcept `json:"type,omitempty"`
	// Description of the objective
	Description *custom.Markdown `json:"description,omitempty"`
}

type ResearchStudyOutcomeMeasure struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Label for the outcome
	Name *string `json:"name,omitempty"`
	// primary | secondary | exploratory
	Type []CodeableConcept `json:"type,omitempty"`
	// Description of the outcome
	Description *custom.Markdown `json:"description,omitempty"`
	// Structured outcome definition
	Reference *custom.Reference[EvidenceVariable] `json:"reference,omitempty"`
}

type ResearchStudy struct {
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
	// Canonical identifier for this study resource
	Url *custom.URI `json:"url,omitempty"`
	// Business Identifier for study
	Identifier []Identifier `json:"identifier,omitempty"`
	// The business version for the study record
	Version *string `json:"version,omitempty"`
	// Name for this study (computer friendly)
	Name *string `json:"name,omitempty"`
	// Human readable name of the study
	Title *string `json:"title,omitempty"`
	// Additional names for the study
	Label []ResearchStudyLabel `json:"label,omitempty"`
	// Steps followed in executing study
	Protocol []custom.Reference[PlanDefinition] `json:"protocol,omitempty"`
	// Part of larger study
	PartOf []custom.Reference[ResearchStudy] `json:"partOf,omitempty"`
	// References, URLs, and attachments
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// Date the resource last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// treatment | prevention | diagnostic | supportive-care | screening | health-services-research | basic-science | device-feasibility
	PrimaryPurposeType *CodeableConcept `json:"primaryPurposeType,omitempty"`
	// n-a | early-phase-1 | phase-1 | phase-1-phase-2 | phase-2 | phase-2-phase-3 | phase-3 | phase-4
	Phase *CodeableConcept `json:"phase,omitempty"`
	// Classifications of the study design characteristics
	StudyDesign []CodeableConcept `json:"studyDesign,omitempty"`
	// Drugs, devices, etc. under study
	Focus []custom.CodeableReference[ResearchStudyFocus] `json:"focus,omitempty"`
	// Condition being studied
	Condition []CodeableConcept `json:"condition,omitempty"`
	// Used to search for the study
	Keyword []CodeableConcept `json:"keyword,omitempty"`
	// Geographic area for the study
	Region []CodeableConcept `json:"region,omitempty"`
	// Brief text explaining the study
	DescriptionSummary *custom.Markdown `json:"descriptionSummary,omitempty"`
	// Detailed narrative of the study
	Description *custom.Markdown `json:"description,omitempty"`
	// When the study began and ended
	Period *Period `json:"period,omitempty"`
	// Facility where study activities are conducted
	Site []custom.Reference[ResearchStudySite] `json:"site,omitempty"`
	// Comments made about the study
	Note []Annotation `json:"note,omitempty"`
	// Classification for the study
	Classifier []CodeableConcept `json:"classifier,omitempty"`
	// Sponsors, collaborators, and other parties
	AssociatedParty []ResearchStudyAssociatedParty `json:"associatedParty,omitempty"`
	// Status of study with time for that status
	ProgressStatus []ResearchStudyProgressStatus `json:"progressStatus,omitempty"`
	// accrual-goal-met | closed-due-to-toxicity | closed-due-to-lack-of-study-progress | temporarily-closed-per-study-design
	WhyStopped *CodeableConcept `json:"whyStopped,omitempty"`
	// Target or actual group of participants enrolled in study
	Recruitment *ResearchStudyRecruitment `json:"recruitment,omitempty"`
	// Defined path through the study for a subject
	ComparisonGroup []ResearchStudyComparisonGroup `json:"comparisonGroup,omitempty"`
	// A goal for the study
	Objective []ResearchStudyObjective `json:"objective,omitempty"`
	// A variable measured during the study
	OutcomeMeasure []ResearchStudyOutcomeMeasure `json:"outcomeMeasure,omitempty"`
	// Link to results generated during the study
	Result []custom.Reference[ResearchStudyResult] `json:"result,omitempty"`
}

type OtherResearchStudy ResearchStudy

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
	return "ResearchStudy"
}

func (r ResearchStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchStudy
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherResearchStudy: OtherResearchStudy(r), ResourceType: r.ResourceType()})
}

func UnmarshalResearchStudy(b []byte) (res ResearchStudy, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
