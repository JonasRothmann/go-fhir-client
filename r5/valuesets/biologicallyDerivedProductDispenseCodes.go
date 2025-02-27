// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/biologicallyderivedproductdispense-status
*/type BiologicallyDerivedProductDispenseCodes string

const (
	// The dispense process has started but not yet completed.
	BiologicallyDerivedProductDispenseCodesPreparation BiologicallyDerivedProductDispenseCodes = "preparation"
	// The dispense process is in progress.
	BiologicallyDerivedProductDispenseCodesInProgress BiologicallyDerivedProductDispenseCodes = "in-progress"
	// The requested product has been allocated and is ready for transport.
	BiologicallyDerivedProductDispenseCodesAllocated BiologicallyDerivedProductDispenseCodes = "allocated"
	// The dispensed product has been picked up.
	BiologicallyDerivedProductDispenseCodesIssued BiologicallyDerivedProductDispenseCodes = "issued"
	// The dispense could not be completed.
	BiologicallyDerivedProductDispenseCodesUnfulfilled BiologicallyDerivedProductDispenseCodes = "unfulfilled"
	// The dispensed product was returned.
	BiologicallyDerivedProductDispenseCodesReturned BiologicallyDerivedProductDispenseCodes = "returned"
	// The dispense was entered in error and therefore nullified.
	BiologicallyDerivedProductDispenseCodesEnteredInError BiologicallyDerivedProductDispenseCodes = "entered-in-error"
	// The authoring system does not know which of the status values applies for this dispense. Note: this concept is not to be used for other - one of the listed statuses is presumed to apply, it's just not known which one.
	BiologicallyDerivedProductDispenseCodesUnknown BiologicallyDerivedProductDispenseCodes = "unknown"
)
