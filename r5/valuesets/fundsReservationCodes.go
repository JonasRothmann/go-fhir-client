// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/fundsreserve
*/type FundsReservationCodes string

const (
	// The payor is requested to reserve funds for the provision of the named services by any provider for settlement of future claims related to this request.
	FundsReservationCodesPatient FundsReservationCodes = "patient"
	// The payor is requested to reserve funds solely for the named provider for settlement of future claims related to this request.
	FundsReservationCodesProvider FundsReservationCodes = "provider"
	// The payor is not being requested to reserve any funds for the settlement of future claims.
	FundsReservationCodesNone FundsReservationCodes = "none"
)
