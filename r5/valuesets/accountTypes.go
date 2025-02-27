// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/v3-ActCode
*/type AccountTypes string

const (
	// An account represents a grouping of financial transactions that are tracked and reported together with a single balance. Examples of account codes (types) are Patient billing accounts (collection of charges), Cost centers; Cash.
	AccountTypesActAccountCode AccountTypes = "_ActAccountCode"
	// An account for collecting charges, reversals, adjustments and payments, including deductibles, copayments, coinsurance (financial transactions) credited or debited to the account receivable account for a patient's encounter.
	AccountTypesAccountReceivable AccountTypes = "ACCTRECEIVABLE"
	AccountTypesCash              AccountTypes = "CASH"
	// Types of advance payment to be made on a plastic card usually issued by a financial institution used of purchasing services and/or products.
	AccountTypesCreditCard AccountTypes = "CC"
	// An account representing charges and credits (financial transactions) for a patient's encounter.
	AccountTypesPatientBillingAccount AccountTypes = "PBILLACCT"
	AccountTypesCreditCard2           AccountTypes = "_CreditCard"
	AccountTypesAmericanExpress       AccountTypes = "AE"
	AccountTypesDinerSClub            AccountTypes = "DN"
	AccountTypesDiscoverCard          AccountTypes = "DV"
	AccountTypesMasterCard            AccountTypes = "MC"
	AccountTypesVisa                  AccountTypes = "V"
)
