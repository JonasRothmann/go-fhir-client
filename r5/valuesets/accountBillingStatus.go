// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/account-billing-status
*/type AccountBillingStatus string

const (
	// The account is open for charging transactions (account.status is active)
	AccountBillingStatusOpen AccountBillingStatus = "open"
	// The account.status is still active and may have charges recorded against it (only for events in the servicePeriod), however the encounters associated are completed. (Also known as Discharged not billed) This BillingStatus is often not used in ongoing accounts. (account.status is active)
	AccountBillingStatusCareCompleteNotBilled AccountBillingStatus = "carecomplete-notbilled"
	// Indicates that all transactions are recorded and the finance system can perform the billing process, including preparing insurance claims, scrubbing charges, invoicing etc. During this time any new charges will not be included in the current billing run/cycle. (account.status is active)
	AccountBillingStatusBilling AccountBillingStatus = "billing"
	// The balance of this debt has not been able to be recovered, and the organization has decided not to persue debt recovery. (account.status is in-active)
	AccountBillingStatusClosedBadDebt AccountBillingStatus = "closed-baddebt"
	// The account was not created in error, however the organization has decided that it will not be charging any transactions associated. (account.status is i n-active)
	AccountBillingStatusClosedVoided AccountBillingStatus = "closed-voided"
	// The account is closed and all charges are processed and accounted for. (account.status is i n-active)
	AccountBillingStatusClosedCompleted AccountBillingStatus = "closed-completed"
	// This account has been merged into another account, all charged have been migrated. This account should no longer be used, and will not be billed. (account.status is i n-active)
	AccountBillingStatusClosedCombined AccountBillingStatus = "closed-combined"
)
