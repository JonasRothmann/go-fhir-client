// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/subscriber-relationship
*/type SubscriberRelationshipCodes string

const (
	// The Beneficiary is a child of the Subscriber
	SubscriberRelationshipCodesChild SubscriberRelationshipCodes = "child"
	// The Beneficiary is a parent of the Subscriber
	SubscriberRelationshipCodesParent SubscriberRelationshipCodes = "parent"
	// The Beneficiary is a spouse or equivalent of the Subscriber
	SubscriberRelationshipCodesSpouse SubscriberRelationshipCodes = "spouse"
	// The Beneficiary is a common law spouse or equivalent of the Subscriber
	SubscriberRelationshipCodesCommon SubscriberRelationshipCodes = "common"
	// The Beneficiary has some other relationship the Subscriber
	SubscriberRelationshipCodesOther SubscriberRelationshipCodes = "other"
	// The Beneficiary is the Subscriber
	SubscriberRelationshipCodesSelf SubscriberRelationshipCodes = "self"
	// The Beneficiary is covered under insurance of the subscriber due to an injury.
	SubscriberRelationshipCodesInjured SubscriberRelationshipCodes = "injured"
)
