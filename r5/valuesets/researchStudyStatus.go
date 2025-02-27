// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/research-study-status
*/type ResearchStudyStatus string

const (
	// Used for documenting the start and end of the overall study, distinct from progress states.
	ResearchStudyStatusOverallStudy ResearchStudyStatus = "overall-study"
	// Study is opened for accrual.
	ResearchStudyStatusActive ResearchStudyStatus = "active"
	// The study is ongoing, and participants are receiving an intervention or being examined, but potential participants are not currently being recruited or enrolled.
	ResearchStudyStatusActiveNotRecruiting ResearchStudyStatus = "active-but-not-recruiting"
	// Study is completed prematurely and will not resume; patients are no longer examined nor treated.
	ResearchStudyStatusAdministrativelyCompleted ResearchStudyStatus = "administratively-completed"
	// Protocol is approved by the review board.
	ResearchStudyStatusApproved ResearchStudyStatus = "approved"
	// Study is closed for accrual; patients can be examined and treated.
	ResearchStudyStatusClosedToAccrual ResearchStudyStatus = "closed-to-accrual"
	// Study is closed to accrual and intervention, i.e. the study is closed to enrollment, all study subjects have completed treatment or intervention but are still being followed according to the primary objective of the study.
	ResearchStudyStatusClosedToAccrualAndIntervention ResearchStudyStatus = "closed-to-accrual-and-intervention"
	// The study closed according to the study plan. There will be no further treatments, interventions or data collection.
	ResearchStudyStatusCompleted ResearchStudyStatus = "completed"
	// Protocol was disapproved by the review board.
	ResearchStudyStatusDisapproved ResearchStudyStatus = "disapproved"
	// The study is selecting its participants from a population, or group of people, decided on by the researchers in advance. These studies are not open to everyone who meets the eligibility criteria but only to people in that particular population, who are specifically invited to participate.
	ResearchStudyStatusEnrollingByInvitation ResearchStudyStatus = "enrolling-by-invitation"
	// Protocol is submitted to the review board for approval.
	ResearchStudyStatusInReview ResearchStudyStatus = "in-review"
	// The study has not started recruiting participants.
	ResearchStudyStatusNotYetRecruiting ResearchStudyStatus = "not-yet-recruiting"
	// The study is currently recruiting participants.
	ResearchStudyStatusRecruiting ResearchStudyStatus = "recruiting"
	// Study is temporarily closed for accrual; can be potentially resumed in the future; patients can be examined and treated.
	ResearchStudyStatusTemporarilyClosedToAccrual ResearchStudyStatus = "temporarily-closed-to-accrual"
	// Study is temporarily closed for accrual and intervention and potentially can be resumed in the future.
	ResearchStudyStatusTemporarilyClosedToAccrualAndIntervention ResearchStudyStatus = "temporarily-closed-to-accrual-and-intervention"
	// The study has stopped early and will not start again. Participants are no longer being examined or treated.
	ResearchStudyStatusTerminated ResearchStudyStatus = "terminated"
	// Protocol was withdrawn by the lead organization.
	ResearchStudyStatusWithdrawn ResearchStudyStatus = "withdrawn"
)
