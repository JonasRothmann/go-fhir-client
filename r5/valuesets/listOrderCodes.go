// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/list-order
*/type ListOrderCodes string

const (
	// The list was sorted by a user. The criteria the user used are not specified.
	ListOrderCodesUser ListOrderCodes = "user"
	// The list was sorted by the system. The criteria the user used are not specified; define additional codes to specify a particular order (or use other defined codes).
	ListOrderCodesSystem ListOrderCodes = "system"
	// The list is sorted by the data of the event. This can be used when the list has items which are dates with past or future events.
	ListOrderCodesSortedByEventDate ListOrderCodes = "event-date"
	// The list is sorted by the date the item was added to the list. Note that the date added to the list is not explicit in the list itself.
	ListOrderCodesSortedByItemDate ListOrderCodes = "entry-date"
	// The list is sorted by priority. The exact method in which priority has been determined is not specified.
	ListOrderCodesPriority ListOrderCodes = "priority"
	// The list is sorted alphabetically by an unspecified property of the items in the list.
	ListOrderCodesAlphabetic ListOrderCodes = "alphabetic"
	// The list is sorted categorically by an unspecified property of the items in the list.
	ListOrderCodesCategory ListOrderCodes = "category"
	// The list is sorted by patient, with items for each patient grouped together.
	ListOrderCodesPatient ListOrderCodes = "patient"
)
