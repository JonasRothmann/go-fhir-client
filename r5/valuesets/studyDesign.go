// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/study-design
*/type StudyDesign string

const (
	// A study design in which an independent variable (an exposure or intervention) is prospectively assigned or modified by the investigator to evaluate a response in the dependent variable (an effect or outcome).
	StudyDesignInterventionalResearch StudyDesign = "SEVCO:01001"
	// A study design in which the independent variables (exposures or interventions) are not prospectively assigned or modified by the investigator.
	StudyDesignObservationalResearch StudyDesign = "SEVCO:01002"
	// A study design in which two or more groups are compared.
	StudyDesignComparativeStudyDesign StudyDesign = "SEVCO:01010"
	// A study design with no comparisons between groups with different exposures and no comparisons between groups with different outcomes.
	StudyDesignNonComparativeStudyDesign StudyDesign = "SEVCO:01023"
	// A study design in which the unit of observation is a population or community.
	StudyDesignPopulationBasedDesign StudyDesign = "SEVCO:01022"
	// A study design process in which data is collected at a single point in time.
	StudyDesignCrossSectionalDataCollection StudyDesign = "SEVCO:01027"
	// A study design process in which data is collected at two or more points in time.
	StudyDesignLongitudinalDataCollection StudyDesign = "SEVCO:01028"
	// A study design process in which the data are recorded and collected during the study for the purpose of the same study.
	StudyDesignPrimaryDataCollection StudyDesign = "SEVCO:01045"
	// A study design process in which the study data are obtained from a source of data collected during a routine process in the natural environment rather than using a process designed or controlled by the researcher.
	StudyDesignRealWorldDataCollection StudyDesign = "SEVCO:01026"
	// A study design process in which the data are collected from data obtained during a different study than the current study.
	StudyDesignSecondaryDataCollectionFromPriorResearch StudyDesign = "SEVCO:01049"
	// A study design process in which the data are collected from a system organized to obtain and maintain uniform data for discovery and analysis, and this system is organized prior to the current study.
	StudyDesignSecondaryDataCollectionFromARegistry StudyDesign = "SEVCO:01042"
	// A study design process in which data are collected from two or more geographic locations.
	StudyDesignMultisiteDataCollection StudyDesign = "SEVCO:01051"
	// A study design process in which data are analyzed with mathematical or statistical methods and formulas.
	StudyDesignQuantitativeAnalysis StudyDesign = "SEVCO:01086"
	// A study design process in which data are analyzed, without primary reliance on mathematical or statistical techniques, by coding and organizing data to provide interpretation or understanding of experiences or hypotheses.
	StudyDesignQualitativeAnalysis StudyDesign = "SEVCO:01087"
	// A study design process in which study participants are not informed of their intervention assignment.
	StudyDesignBlindingOfStudyParticipants StudyDesign = "SEVCO:01060"
	// A study design process in which the people administering the intervention are not informed of the intervention assignment.
	StudyDesignBlindingOfInterventionProviders StudyDesign = "SEVCO:01061"
	// A study design process in which the people determining the outcome are not informed of the intervention assignment.
	StudyDesignBlindingOfOutcomeAssessors StudyDesign = "SEVCO:01062"
	// A study design process in which the people managing or processing the data and statistical analysis are not informed of the intervention assignment.
	StudyDesignBlindingOfDataAnalysts StudyDesign = "SEVCO:01063"
	// A study design process in which all parties influencing study enrollment and allocation to study groups are unaware of the group assignment for the study participant at the time of enrollment and allocation.
	StudyDesignAllocationConcealment StudyDesign = "SEVCO:01064"
	// A study design feature in which two or more institutions are responsible for the conduct of the study.
	StudyDesignMulticentric StudyDesign = "SEVCO:01043"
	// A study design feature in which one or more outcomes are reported directly from the patient without interpretation by a clinician or researcher.
	StudyDesignIncludesPatientReportedOutcome StudyDesign = "SEVCO:01052"
	// A study design feature in which one or more measures are outcomes that patients directly care about, i.e. outcomes that are directly related to patients' experience of their life.
	StudyDesignIncludesPatientCenteredOutcome StudyDesign = "SEVCO:01053"
	// A study design feature in which one or more measures are outcomes that relate to a health or illness condition but are not outcomes which patients directly care about.
	StudyDesignIncludesDiseaseOrientedOutcome StudyDesign = "SEVCO:01054"
	// A study design feature in which one or more outcomes are actions or behaviors of a healthcare professional or care team.
	StudyDesignIncludesProcessMeasure StudyDesign = "SEVCO:01085"
	// A study design feature specifying the intent of the study.
	StudyDesignStudyGoal StudyDesign = "SEVCO:01089"
)
