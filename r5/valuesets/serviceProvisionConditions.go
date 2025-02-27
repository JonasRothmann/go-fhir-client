// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/service-provision-conditions
*/type ServiceProvisionConditions string

const (
	// This service is available for no patient cost.
	ServiceProvisionConditionsFree ServiceProvisionConditions = "free"
	// There are discounts available on this service for qualifying patients.
	ServiceProvisionConditionsDisc ServiceProvisionConditions = "disc"
	// Fees apply for this service.
	ServiceProvisionConditionsCost ServiceProvisionConditions = "cost"
)
