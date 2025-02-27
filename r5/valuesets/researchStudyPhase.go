// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/research-study-phase
*/type ResearchStudyPhase string

const (
	// Trials without phases (for example, studies of devices or behavioral interventions).
	ResearchStudyPhaseNa ResearchStudyPhase = "n-a"
	// Designation for optional exploratory trials conducted in accordance with the United States Food and Drug Administration's (FDA) 2006 Guidance on Exploratory Investigational New Drug (IND) Studies. Formerly called Phase 0.
	ResearchStudyPhaseEarlyPhase1 ResearchStudyPhase = "early-phase-1"
	// Includes initial studies to determine the metabolism and pharmacologic actions of drugs in humans, the side effects associated with increasing doses, and to gain early evidence of effectiveness; may include healthy participants and/or patients.
	ResearchStudyPhasePhase1 ResearchStudyPhase = "phase-1"
	// Trials that are a combination of phases 1 and 2.
	ResearchStudyPhasePhase1phase2 ResearchStudyPhase = "phase-1-phase-2"
	// Includes controlled clinical studies conducted to evaluate the effectiveness of the drug for a particular indication or indications in participants with the disease or condition under study and to determine the common short-term side effects and risks.
	ResearchStudyPhasePhase2 ResearchStudyPhase = "phase-2"
	// Trials that are a combination of phases 2 and 3.
	ResearchStudyPhasePhase2phase3 ResearchStudyPhase = "phase-2-phase-3"
	// Includes trials conducted after preliminary evidence suggesting effectiveness of the drug has been obtained, and are intended to gather additional information to evaluate the overall benefit-risk relationship of the drug.
	ResearchStudyPhasePhase3 ResearchStudyPhase = "phase-3"
	// Studies of FDA-approved drugs to delineate additional information including the drug's risks, benefits, and optimal use.
	ResearchStudyPhasePhase4 ResearchStudyPhase = "phase-4"
)
