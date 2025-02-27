// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/restful-security-service
*/type RestfulSecurityService string

const (
	// OAuth (unspecified version see oauth.net).
	RestfulSecurityServiceOAuth RestfulSecurityService = "OAuth"
	// OAuth2 using SMART-on-FHIR profile (see http://docs.smarthealthit.org/).
	RestfulSecurityServiceSmartOnFhir RestfulSecurityService = "SMART-on-FHIR"
	// Microsoft NTLM Authentication.
	RestfulSecurityServiceNtlm RestfulSecurityService = "NTLM"
	// Basic authentication defined in HTTP specification.
	RestfulSecurityServiceBasic RestfulSecurityService = "Basic"
	// see http://www.ietf.org/rfc/rfc4120.txt.
	RestfulSecurityServiceKerberos RestfulSecurityService = "Kerberos"
	// SSL where client must have a certificate registered with the server.
	RestfulSecurityServiceCertificates RestfulSecurityService = "Certificates"
)
