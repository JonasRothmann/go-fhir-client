// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/service-referral-method
*/type ReferralMethod string

const (
	// Referrals may be accepted by fax.
	ReferralMethodFax ReferralMethod = "fax"
	// Referrals may be accepted over the phone.
	ReferralMethodPhone ReferralMethod = "phone"
	// Referrals may be accepted via a secure messaging system. To determine the types of secure messaging systems supported, refer to the identifiers collection. Callers will need to understand the specific identifier system used to know that they are able to transmit messages.
	ReferralMethodElec ReferralMethod = "elec"
	// Referrals may be accepted via a secure email. To send please encrypt with the services public key.
	ReferralMethodSemail ReferralMethod = "semail"
	// Referrals may be accepted via regular postage (or hand delivered).
	ReferralMethodMail ReferralMethod = "mail"
	// Referrals may be accepted via self-referral.
	ReferralMethodSelf ReferralMethod = "self"
)
