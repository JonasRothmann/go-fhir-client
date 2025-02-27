// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/evidence-classifier-code
*/type EvidenceClassifier string

const (
	// About COVID-19.
	EvidenceClassifierCovid19specific EvidenceClassifier = "COVID19Specific"
	// Not about COVID-19 but relevant to COVID-19 management or understanding.
	EvidenceClassifierCovid19relevant EvidenceClassifier = "COVID19Relevant"
	// contains human COVID-19 disease in the research data as any variable (population, exposure or outcome).
	EvidenceClassifierCovid19humanResearch EvidenceClassifier = "COVID19HumanResearch"
	// such as randomized trial, observational study.
	EvidenceClassifierOriginalResearch EvidenceClassifier = "OriginalResearch"
	// such as systematic review, meta-analysis, rapid review.
	EvidenceClassifierResearchSynthesis EvidenceClassifier = "ResearchSynthesis"
	// for clinical practice guidelines.
	EvidenceClassifierGuideline EvidenceClassifier = "Guideline"
	// for research protocols.
	EvidenceClassifierResearchProtocol EvidenceClassifier = "ResearchProtocol"
	// for things that are not research reports, research protocols or guidelines.
	EvidenceClassifierNotResearchNotGuideline EvidenceClassifier = "NotResearchNotGuideline"
	// about therapeutic interventions.
	EvidenceClassifierTreatment EvidenceClassifier = "Treatment"
	// about preventive care and interventions.
	EvidenceClassifierPreventionAndControl EvidenceClassifier = "PreventionAndControl"
	// about methods to distinguish having or not having a condition.
	EvidenceClassifierDiagnosis EvidenceClassifier = "Diagnosis"
	// about predicting risk for something or risk factors for it.
	EvidenceClassifierPrognosisPrediction EvidenceClassifier = "PrognosisPrediction"
	EvidenceClassifierRatedAsYes          EvidenceClassifier = "RatedAsYes"
	// Rated as no, negative, absent, or exclude.
	EvidenceClassifierRatedAsNo EvidenceClassifier = "RatedAsNo"
	// Neither rated as yes nor rated as no.
	EvidenceClassifierNotAssessed EvidenceClassifier = "NotAssessed"
	// classified as randomized controlled trial.
	EvidenceClassifierRatedAsRct EvidenceClassifier = "RatedAsRCT"
	// classified as nonrandomized controlled trial (experimental).
	EvidenceClassifierRatedAsControlledTrial EvidenceClassifier = "RatedAsControlledTrial"
	// classified as comparative cohort study (observational).
	EvidenceClassifierRatedAsComparativeCohort EvidenceClassifier = "RatedAsComparativeCohort"
	// classified as case-control study.
	EvidenceClassifierRatedAsCaseControl EvidenceClassifier = "RatedAsCaseControl"
	// classified as uncontrolled cohort (case series).
	EvidenceClassifierRatedAsUncontrolledSeries EvidenceClassifier = "RatedAsUncontrolledSeries"
	// classified as mixed-methods study.
	EvidenceClassifierRatedAsMixedMethods EvidenceClassifier = "RatedAsMixedMethods"
	// classified as other concept (not elsewhere classified).
	EvidenceClassifierRatedAsOther EvidenceClassifier = "RatedAsOther"
	// Risk of bias assessment.
	EvidenceClassifierRiskOfBias EvidenceClassifier = "RiskOfBias"
	// No blinding.
	EvidenceClassifierNoBlinding EvidenceClassifier = "NoBlinding"
	// Allocation concealment not stated.
	EvidenceClassifierAllocConcealNotStated EvidenceClassifier = "AllocConcealNotStated"
	// Early trial termination.
	EvidenceClassifierEarlyTrialTermination EvidenceClassifier = "EarlyTrialTermination"
	// No intention-to-treat analysis.
	EvidenceClassifierNoItt EvidenceClassifier = "NoITT"
	// Results presented in preprint (pre-final publication) form.
	EvidenceClassifierPreprint EvidenceClassifier = "Preprint"
	// Preliminary analysis.
	EvidenceClassifierPreliminaryAnalysis EvidenceClassifier = "PreliminaryAnalysis"
	// Differences between groups at start of trial may confound or bias the findings.
	EvidenceClassifierBaselineImbalance EvidenceClassifier = "BaselineImbalance"
	// Subgroup analysis.
	EvidenceClassifierSubgroupAnalysis EvidenceClassifier = "SubgroupAnalysis"
)
