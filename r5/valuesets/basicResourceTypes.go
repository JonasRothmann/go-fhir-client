// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/basic-resource-type
*/type BasicResourceTypes string

const (
	// An assertion of permission for an activity or set of activities to occur, possibly subject to particular limitations; e.g. surgical consent, information disclosure consent, etc.
	BasicResourceTypesConsent BasicResourceTypes = "consent"
	// A request that care of a particular type be provided to a patient.  Could involve the transfer of care, a consult, etc.
	BasicResourceTypesReferral BasicResourceTypes = "referral"
	// An undesired reaction caused by exposure to some agent (e.g. a medication, immunization, food, or environmental agent).
	BasicResourceTypesAdvevent BasicResourceTypes = "advevent"
	// A request that a time be scheduled for a type of service for a specified patient, potentially subject to other constraints
	BasicResourceTypesAptmtreq BasicResourceTypes = "aptmtreq"
	// The transition of a patient or set of material from one location to another
	BasicResourceTypesTransfer BasicResourceTypes = "transfer"
	// The specification of a set of food and/or other nutritional material to be delivered to a patient.
	BasicResourceTypesDiet BasicResourceTypes = "diet"
	// An occurrence of a non-care-related event in the healthcare domain, such as approvals, reviews, etc.
	BasicResourceTypesAdminact BasicResourceTypes = "adminact"
	// Record of a situation where a subject was exposed to a substance.  Usually of interest to public health.
	BasicResourceTypesExposure BasicResourceTypes = "exposure"
	// A formalized inquiry into the circumstances surrounding a particular unplanned event or potential event for the purposes of identifying possible causes and contributing factors for the event
	BasicResourceTypesInvestigation BasicResourceTypes = "investigation"
	// A financial instrument used to track costs, charges or other amounts.
	BasicResourceTypesAccount BasicResourceTypes = "account"
	// A request for payment for goods and/or services.  Includes the idea of a healthcare insurance claim.
	BasicResourceTypesInvoice BasicResourceTypes = "invoice"
	// The determination of what will be paid against a particular invoice based on coverage, plan rules, etc.
	BasicResourceTypesAdjudicat BasicResourceTypes = "adjudicat"
	// A request for a pre-determination of the cost that would be paid under an insurance plan for a hypothetical claim for goods or services
	BasicResourceTypesPredetreq BasicResourceTypes = "predetreq"
	// An adjudication of what would be paid under an insurance plan for a hypothetical claim for goods or services
	BasicResourceTypesPredetermine BasicResourceTypes = "predetermine"
	// An investigation to determine information about a particular therapy or product
	BasicResourceTypesStudy BasicResourceTypes = "study"
	// A set of (possibly conditional) steps to be taken to achieve some aim.  Includes study protocols, treatment protocols, emergency protocols, etc.
	BasicResourceTypesProtocol BasicResourceTypes = "protocol"
)
