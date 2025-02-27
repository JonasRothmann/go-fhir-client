// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/adverse-event-causality-assess
*/type AdverseEventCausalityAssessment string

const (
	// i) Event or laboratory test abnormality, with plausible time relationship to drug intake; ii) Cannot be explained by disease or other drugs; iii) Response to withdrawal plausible (pharmacologically, pathologically); iv) Event definitive pharmacologically or phenomenologically (i.e. an objective and specific medical disorder or a recognized pharmacological phenomenon); or v) Re-challenge satisfactory, if necessary.
	AdverseEventCausalityAssessmentCertain AdverseEventCausalityAssessment = "certain"
	// i) Event or laboratory test abnormality, with reasonable time relationship to drug intake; ii) Unlikely to be attributed to disease or other drugs; iii) Response to withdrawal clinically reasonable; or iv) Re-challenge not required.
	AdverseEventCausalityAssessmentProbablyLikely AdverseEventCausalityAssessment = "probably-likely"
	// i) Event or laboratory test abnormality, with reasonable time relationship to drug intake; ii) Could also be explained by disease or other drugs; or iii) Information on drug withdrawal may be lacking or unclear.
	AdverseEventCausalityAssessmentPossible AdverseEventCausalityAssessment = "possible"
	// i) Event or laboratory test abnormality, with a time to drug intake that makes a relationship improbable (but not impossible); or ii) Disease or other drugs provide plausible explanations.
	AdverseEventCausalityAssessmentUnlikely AdverseEventCausalityAssessment = "unlikely"
	// i) Event or laboratory test abnormality; ii) More data for proper assessment needed; or iii) Additional data under examination.
	AdverseEventCausalityAssessmentConditionalClassified AdverseEventCausalityAssessment = "conditional-classified"
	// i) Report suggesting an adverse reaction; ii) Cannot be judged because information is insufficient or contradictory; or iii) Data cannot be supplemented or verified.
	AdverseEventCausalityAssessmentUnassessableUnclassifiable AdverseEventCausalityAssessment = "unassessable-unclassifiable"
)
