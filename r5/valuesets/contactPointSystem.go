// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/contact-point-system
*/type ContactPointSystem string

const (
	// The value is a telephone number used for voice calls. Use of full international numbers starting with + is recommended to enable automatic dialing support but not required.
	ContactPointSystemPhone ContactPointSystem = "phone"
	// The value is a fax machine. Use of full international numbers starting with + is recommended to enable automatic dialing support but not required.
	ContactPointSystemFax ContactPointSystem = "fax"
	// The value is an email address.
	ContactPointSystemEmail ContactPointSystem = "email"
	// The value is a pager number. These may be local pager numbers that are only usable on a particular pager system.
	ContactPointSystemPager ContactPointSystem = "pager"
	// A contact that is not a phone, fax, pager or email address and is expressed as a URL.  This is intended for various institutional or personal contacts including web sites, blogs, Skype, Twitter, Facebook, etc. Do not use for email addresses.
	ContactPointSystemUrl ContactPointSystem = "url"
	// A contact that can be used for sending a sms message (e.g. mobile phones, some landlines).
	ContactPointSystemSms ContactPointSystem = "sms"
	// A contact that is not a phone, fax, page or email address and is not expressible as a URL.  E.g. Internal mail address.  This SHOULD NOT be used for contacts that are expressible as a URL (e.g. Skype, Twitter, Facebook, etc.)  Extensions may be used to distinguish "other" contact types.
	ContactPointSystemOther ContactPointSystem = "other"
)
