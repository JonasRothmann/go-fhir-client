// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/request-priority
*/type RequestPriority string

const (
	// The request has normal priority.
	RequestPriorityRoutine RequestPriority = "routine"
	// The request should be actioned promptly - higher priority than routine.
	RequestPriorityUrgent RequestPriority = "urgent"
	// The request should be actioned as soon as possible - higher priority than urgent.
	RequestPriorityAsap RequestPriority = "asap"
	// The request should be actioned immediately - highest possible priority.  E.g. an emergency.
	RequestPriorityStat RequestPriority = "stat"
)
