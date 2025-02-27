// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/research-subject-state
*/type ResearchSubjectState string

const (
	// An identified person that can be considered for inclusion in a study.
	ResearchSubjectStateCandidate ResearchSubjectState = "candidate"
	// A person that has met the eligibility criteria for inclusion in a study.
	ResearchSubjectStateEligible ResearchSubjectState = "eligible"
	// A person is no longer receiving study intervention and/or being evaluated with tests and procedures according to the protocol, but they are being monitored on a protocol-prescribed schedule.
	ResearchSubjectStateFollowUp ResearchSubjectState = "follow-up"
	/*
	   A person who did not meet one or more criteria required for participation in a study is considered to have failed screening or
	   is ineligible for the study.
	*/
	ResearchSubjectStateIneligible ResearchSubjectState = "ineligible"
	// A person for whom registration was not completed.
	ResearchSubjectStateNotRegistered ResearchSubjectState = "not-registered"
	/*
	   A person that has ended their participation on a study either because their treatment/observation is complete or through not
	   responding, withdrawal, non-compliance and/or adverse event.
	*/
	ResearchSubjectStateOffStudy ResearchSubjectState = "off-study"
	// A person that is enrolled or registered on a study.
	ResearchSubjectStateOnStudy ResearchSubjectState = "on-study"
	// The person is receiving the treatment or participating in an activity (e.g. yoga, diet, etc.) that the study is evaluating.
	ResearchSubjectStateOnStudyIntervention ResearchSubjectState = "on-study-intervention"
	// The subject is being evaluated via tests and assessments according to the study calendar, but is not receiving any intervention. Note that this state is study-dependent and might not exist in all studies.  A synonym for this is "short-term follow-up".
	ResearchSubjectStateOnStudyObservation ResearchSubjectState = "on-study-observation"
	// A person is pre-registered for a study.
	ResearchSubjectStatePendingOnStudy ResearchSubjectState = "pending-on-study"
	// A person that is potentially eligible for participation in the study.
	ResearchSubjectStatePotentialCandidate ResearchSubjectState = "potential-candidate"
	// A person who is being evaluated for eligibility for a study.
	ResearchSubjectStateScreening ResearchSubjectState = "screening"
	// The person has withdrawn their participation in the study before registration.
	ResearchSubjectStateWithdrawn ResearchSubjectState = "withdrawn"
)
