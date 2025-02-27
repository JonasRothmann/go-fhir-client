// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/resource-status
*/type CanonicalStatusCodesForFHIRResources string

const (
	// The resource was created in error, and should not be treated as valid (note: in many cases, for various data integrity related reasons, the information cannot be removed from the record)
	CanonicalStatusCodesForFHIRResourcesError CanonicalStatusCodesForFHIRResources = "error"
	// The resource describes an action or plan that is proposed, and not yet approved by the participants
	CanonicalStatusCodesForFHIRResourcesProposed CanonicalStatusCodesForFHIRResources = "proposed"
	// The resource describes a course of action that is planned and agreed/approved, but at the time of recording was still future
	CanonicalStatusCodesForFHIRResourcesPlanned CanonicalStatusCodesForFHIRResources = "planned"
	// The information in the resource is still being prepared and edited
	CanonicalStatusCodesForFHIRResourcesDraft CanonicalStatusCodesForFHIRResources = "draft"
	// A fulfiller has been asked to perform this action, but it has not yet occurred
	CanonicalStatusCodesForFHIRResourcesRequested CanonicalStatusCodesForFHIRResources = "requested"
	// The fulfiller has received the request, but not yet agreed to carry out the action
	CanonicalStatusCodesForFHIRResourcesReceived CanonicalStatusCodesForFHIRResources = "received"
	// The fulfiller chose not to perform the action
	CanonicalStatusCodesForFHIRResourcesDeclined CanonicalStatusCodesForFHIRResources = "declined"
	// The fulfiller has decided to perform the action, and plans are in train to do this in the future
	CanonicalStatusCodesForFHIRResourcesAccepted CanonicalStatusCodesForFHIRResources = "accepted"
	// The pre-conditions for the action are all fulfilled, and it is imminent
	CanonicalStatusCodesForFHIRResourcesArrived CanonicalStatusCodesForFHIRResources = "arrived"
	// The resource describes information that is currently valid or a process that is presently occuring
	CanonicalStatusCodesForFHIRResourcesActive CanonicalStatusCodesForFHIRResources = "active"
	// The process described/requested in this resource has been halted for some reason
	CanonicalStatusCodesForFHIRResourcesSuspended CanonicalStatusCodesForFHIRResources = "suspended"
	// The process described/requested in the resource could not be completed, and no further action is planned
	CanonicalStatusCodesForFHIRResourcesFailed CanonicalStatusCodesForFHIRResources = "failed"
	// The information in this resource has been replaced by information in another resource
	CanonicalStatusCodesForFHIRResourcesReplaced CanonicalStatusCodesForFHIRResources = "replaced"
	// The process described/requested in the resource has been completed, and no further action is planned
	CanonicalStatusCodesForFHIRResourcesComplete CanonicalStatusCodesForFHIRResources = "complete"
	// The resource describes information that is no longer valid or a process that is stopped occurring
	CanonicalStatusCodesForFHIRResourcesInactive CanonicalStatusCodesForFHIRResources = "inactive"
	// The process described/requested in the resource did not complete - usually due to some workflow error, and no further action is planned
	CanonicalStatusCodesForFHIRResourcesAbandoned CanonicalStatusCodesForFHIRResources = "abandoned"
	// Authoring system does not know the status
	CanonicalStatusCodesForFHIRResourcesUnknown CanonicalStatusCodesForFHIRResources = "unknown"
	// The information in this resource is not yet approved
	CanonicalStatusCodesForFHIRResourcesUnconfirmed CanonicalStatusCodesForFHIRResources = "unconfirmed"
	// The information in this resource is approved
	CanonicalStatusCodesForFHIRResourcesConfirmed CanonicalStatusCodesForFHIRResources = "confirmed"
	// The issue identified by this resource is no longer of concern
	CanonicalStatusCodesForFHIRResourcesResolved CanonicalStatusCodesForFHIRResources = "resolved"
	// This information has been ruled out by testing and evaluation
	CanonicalStatusCodesForFHIRResourcesRefuted CanonicalStatusCodesForFHIRResources = "refuted"
	// Potentially true?
	CanonicalStatusCodesForFHIRResourcesDifferential CanonicalStatusCodesForFHIRResources = "differential"
	// This information is still being assembled
	CanonicalStatusCodesForFHIRResourcesPartial CanonicalStatusCodesForFHIRResources = "partial"
	// not available at this time/location
	CanonicalStatusCodesForFHIRResourcesBusyUnavailable CanonicalStatusCodesForFHIRResources = "busy-unavailable"
	// Free for scheduling
	CanonicalStatusCodesForFHIRResourcesFree CanonicalStatusCodesForFHIRResources = "free"
	// Ready to act
	CanonicalStatusCodesForFHIRResourcesOnTarget CanonicalStatusCodesForFHIRResources = "on-target"
	// Ahead of the planned timelines
	CanonicalStatusCodesForFHIRResourcesAheadOfTarget CanonicalStatusCodesForFHIRResources = "ahead-of-target"
	// Behind the planned timelines
	CanonicalStatusCodesForFHIRResourcesBehindTarget CanonicalStatusCodesForFHIRResources = "behind-target"
	// Not ready to act
	CanonicalStatusCodesForFHIRResourcesNotReady CanonicalStatusCodesForFHIRResources = "not-ready"
	// The device transducer is disconnected
	CanonicalStatusCodesForFHIRResourcesTransducDiscon CanonicalStatusCodesForFHIRResources = "transduc-discon"
	// The hardware is disconnected
	CanonicalStatusCodesForFHIRResourcesHwDiscon CanonicalStatusCodesForFHIRResources = "hw-discon"
)
