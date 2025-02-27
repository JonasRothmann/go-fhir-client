// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/trigger-type
*/type TriggerType string

const (
	// The trigger occurs in response to a specific named event, and no other information about the trigger is specified. Named events are completely pre-coordinated, and the formal semantics of the trigger are not provided.
	TriggerTypeNamedEvent TriggerType = "named-event"
	// The trigger occurs at a specific time or periodically as described by a timing or schedule. A periodic event cannot have any data elements, but may have a name assigned as a shorthand for the event.
	TriggerTypePeriodic TriggerType = "periodic"
	// The trigger occurs whenever data of a particular type is changed in any way, either added, modified, or removed.
	TriggerTypeDataChanged TriggerType = "data-changed"
	// The trigger occurs whenever data of a particular type is accessed.
	TriggerTypeDataAccessed TriggerType = "data-accessed"
	// The trigger occurs whenever access to data of a particular type is completed.
	TriggerTypeDataAccessEnded TriggerType = "data-access-ended"
)
