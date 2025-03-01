// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/episode-of-care-status
*/type EpisodeOfCareStatus string

const (
	// This episode of care is planned to start at the date specified in the period.start. During this status, an organization may perform assessments to determine if the patient is eligible to receive services, or be organizing to make resources available to provide care services.
	EpisodeOfCareStatusPlanned EpisodeOfCareStatus = "planned"
	// This episode has been placed on a waitlist, pending the episode being made active (or cancelled).
	EpisodeOfCareStatusWaitlist EpisodeOfCareStatus = "waitlist"
	// This episode of care is current.
	EpisodeOfCareStatusActive EpisodeOfCareStatus = "active"
	// This episode of care is on hold; the organization has limited responsibility for the patient (such as while on respite).
	EpisodeOfCareStatusOnhold EpisodeOfCareStatus = "onhold"
	// This episode of care is finished and the organization is not expecting to be providing further care to the patient. Can also be known as "closed", "completed" or other similar terms.
	EpisodeOfCareStatusFinished EpisodeOfCareStatus = "finished"
	// The episode of care was cancelled, or withdrawn from service, often selected during the planned stage as the patient may have gone elsewhere, or the circumstances have changed and the organization is unable to provide the care. It indicates that services terminated outside the planned/expected workflow.
	EpisodeOfCareStatusCancelled EpisodeOfCareStatus = "cancelled"
	// This instance should not have been part of this patient's medical record.
	EpisodeOfCareStatusEnteredInError EpisodeOfCareStatus = "entered-in-error"
)
