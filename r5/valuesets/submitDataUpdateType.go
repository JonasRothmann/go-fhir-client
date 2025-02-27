// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/submit-data-update-type
*/type SubmitDataUpdateType string

const (
	// In contrast to the Snapshot Update, the FHIR Parameters resource used in a Submit Data or the Collect Data scenario contains only the new and updated DEQM and QI Core Profiles since the last transaction. If the Consumer supports incremental updates, the contents of the updated payload updates the previous payload data.
	SubmitDataUpdateTypeIncremental SubmitDataUpdateType = "incremental"
	// In contrast to the Incremental Update, the FHIR Parameters resource used in a Submit Data or the Collect Data scenario contains all the DEQM and QI Core Profiles for each transaction.  If the Consumer supports snapshot updates, the contents of the updated payload entirely replaces the previous payload
	SubmitDataUpdateTypeSnapshot SubmitDataUpdateType = "snapshot"
)
