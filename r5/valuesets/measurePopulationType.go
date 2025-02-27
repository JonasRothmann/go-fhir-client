// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/measure-population
*/type MeasurePopulationType string

const (
	// The initial population criteria refers to all patients, subjects, or events to be evaluated by a quality measure involving patients or subjects who share a common set of specified characterstics. All patients, subjects, or events counted (for example, as numerator, as denominator) are drawn from the initial population.
	MeasurePopulationTypeInitialPopulation MeasurePopulationType = "initial-population"
	/*
	   The numerator inclusion statement describes the process, condition, event, or outcome that satisfies the measure’s focus or intent. The numerator inclusion statement includes parameters such as:

	    * the Event or events that will satisfy the numerator inclusion criteria
	    * the Performance period or time interval in which the numerator event must occur, if it is different from that used for identifying the denominator.

	    Note that this code is used to identify the starting inclusion criteria for the numerator. The final calculated numerator must be determined by considering all relevant criteria for the scoring type, including numerator inclusion and exclusion criteria.

	    Source: https://mmshub.cms.gov/measure-lifecycle/measure-specification/develop-specification/numerator
	*/
	MeasurePopulationTypeNumerator MeasurePopulationType = "numerator"
	// Numerator exclusion criteria define patients, subjects, or events to be excluded from the numerator. Numerator exclusions are used in proportion and ratio measures to help narrow the numerator (for inverted measures).
	MeasurePopulationTypeNumeratorExclusion MeasurePopulationType = "numerator-exclusion"
	/*
	   Denominator inclusion criteria define the patients, subjects, or events that should be included in the lower portion of a fraction used to calculate a rate, proportion, or ratio.

	    Note that this code is used to identify the starting inclusion criteria for the denominator. The final calculated denominator must be determined by considering all relevant criteria for the scoring type, including denominator inclusion, exclusion, and exception criteria. The final calculated denominator can be the same as the initial population, or a subset of the initial population to further constrain the population for the purpose of the measure.
	*/
	MeasurePopulationTypeDenominator MeasurePopulationType = "denominator"
	// Denominator exclusion criteria define patients, subjects, or events that should be excluded from the denominator. Denominator exclusions are used in proportion and ratio measures to help narrow the denominator. For example, patients with bilateral lower extremity amputations would be listed as a denominator exclusion for a measure requiring foot exams.
	MeasurePopulationTypeDenominatorExclusion MeasurePopulationType = "denominator-exclusion"
	// Denominator exceptions are conditions that should remove a patient, subject, or event from the denominator of a measure only if the numerator criteria are not met. Denominator exception allows for adjustment of the calculated score for those providers with higher risk populations. Denominator exception criteria are only used in proportion measures.
	MeasurePopulationTypeDenominatorException MeasurePopulationType = "denominator-exception"
	/*
	   Measure population inclusion criteria define the patients, subjects, or events for which the measure observation should be taken. Measure populations are used for continuous variable measures rather than numerator and denominator criteria.

	    Note that this code is used to identify the starting inclusion criteria for the measure population. The final calculated measure population must be determined by considering all relevant criteria for the scoring type, including measure population inclusion and exclusion criteria.
	*/
	MeasurePopulationTypeMeasurePopulation MeasurePopulationType = "measure-population"
	// Measure population exclusion criteria define the patients or events that should be excluded from the measure population before determining the outcome of one or more continuous variables defined for the measure observation. Measure population exclusion criteria are used within continuous variable measures to help narrow the measure population.
	MeasurePopulationTypeMeasurePopulationExclusion MeasurePopulationType = "measure-population-exclusion"
	// Measure observation criteria are used to define an individual observation to be performed for each patient, subject, or event in the measure population. Measure observations for each case in the population are aggregated to determine the overall measure score for the population.
	MeasurePopulationTypeMeasureObservation MeasurePopulationType = "measure-observation"
)
