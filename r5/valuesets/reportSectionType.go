// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/evidence-report-section
*/type ReportSectionType string

const (
	// Evidence Results.
	ReportSectionTypeEvidence ReportSectionType = "Evidence"
	// Evidence Results for the intervention exposure only.
	ReportSectionTypeEvidenceResultsForTheInterventionExposureOnly ReportSectionType = "Intervention-group-alone-Evidence"
	// Evidence Results for comparison of Intervention and Control.
	ReportSectionTypeEvidenceResultsForComparisonOfInterventionAndControl ReportSectionType = "Intervention-vs-Control-Evidence"
	// Evidence Results for the control exposure only.
	ReportSectionTypeEvidenceResultsForTheControlExposureOnly ReportSectionType = "Control-group-alone-Evidence"
	// Evidence Variables used.
	ReportSectionTypeEvidenceVariable ReportSectionType = "EvidenceVariable"
	// Evidence Variables as observed in the research data.
	ReportSectionTypeEvidenceVariablesActuallyObserved ReportSectionType = "EvidenceVariable-observed"
	// Evidence Variables intended for interpretation.
	ReportSectionTypeEvidenceVariablesIntendedForInterpretation ReportSectionType = "EvidenceVariable-intended"
	// Evidence Variable in variable role Population.
	ReportSectionTypeEvidenceVariableInVariableRolePopulation ReportSectionType = "EvidenceVariable-population"
	// Evidence Variable in variable role Exposure.
	ReportSectionTypeEvidenceVariableInVariableRoleExposure ReportSectionType = "EvidenceVariable-exposure"
	// Evidence Variable in variable role Outcome (MeasuredVariable).
	ReportSectionTypeEvidenceVariableInVariableRoleOutcomeMeasuredVariable ReportSectionType = "EvidenceVariable-outcome"
	// Outcomes related to efficacy or potential benefits of interventions.
	ReportSectionTypeEfficacyOutcomes ReportSectionType = "Efficacy-outcomes"
	// Outcomes related to adverse effects or potential harms of interventions.
	ReportSectionTypeHarmsOutcomes ReportSectionType = "Harms-outcomes"
	// Sample Size.
	ReportSectionTypeSampleSize ReportSectionType = "SampleSize"
	// References.
	ReportSectionTypeReferences ReportSectionType = "References"
	// Assertion.
	ReportSectionTypeAssertion ReportSectionType = "Assertion"
	// Reasons.
	ReportSectionTypeReasons ReportSectionType = "Reasons"
	// Certainty of Evidence.
	ReportSectionTypeCertaintyOfEvidence ReportSectionType = "Certainty-of-Evidence"
	// This section is used for classifiers of the evidence.
	ReportSectionTypeEvidenceClassifierSection ReportSectionType = "Evidence-Classifier"
	// Warnings.
	ReportSectionTypeWarnings ReportSectionType = "Warnings"
	// Denotes a section specifying text summary for a report.
	ReportSectionTypeTextSummary ReportSectionType = "Text-Summary"
	// Summary of Body of Evidence Findings.
	ReportSectionTypeSummaryOfBodyOfEvidenceFindings ReportSectionType = "SummaryOfBodyOfEvidenceFindings"
	// Summary of Individual Study Findings.
	ReportSectionTypeSummaryOfIndividualStudyFindings ReportSectionType = "SummaryOfIndividualStudyFindings"
	// Denotes the header to use for a Text Summary or above a Table.
	ReportSectionTypeHeader ReportSectionType = "Header"
	// Tables.
	ReportSectionTypeTables ReportSectionType = "Tables"
	ReportSectionTypeTable  ReportSectionType = "Table"
	// Denotes a section specifying row headers for a tabular report.
	ReportSectionTypeRowHeaders ReportSectionType = "Row-Headers"
	// Denotes the header to use for the column for a tabular report.
	ReportSectionTypeColumnHeader ReportSectionType = "Column-Header"
	// Denotes a section specifying column headers for a tabular report.
	ReportSectionTypeColumnHeaders ReportSectionType = "Column-Headers"
)
