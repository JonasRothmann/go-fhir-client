// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/usage-context-agreement-scope
*/type UsageContextAgreementScope string

const (
	// Defines use-case independent foundational expectations for exchange within a particular country or jurisdiction.  Should be accompanied by a jurisdiction.  Commonly used as a 'base' for more restrictive artifacts.
	UsageContextAgreementScopeRealmBase UsageContextAgreementScope = "realm-base"
	// Defines use-case-specific information or guidance that is relevant to a specific business or health domain but is not mandated for particular use.
	UsageContextAgreementScopeKnowledge UsageContextAgreementScope = "knowledge"
	// Defines use-case-specific requirements for a specific business or health domain.  May vary in jurisdictional scope from international to small region and in business scope from broad to narrow.
	UsageContextAgreementScopeDomain UsageContextAgreementScope = "domain"
	// Sets contractual or business expectations for systems participating in a particular exchange community.
	UsageContextAgreementScopeCommunity UsageContextAgreementScope = "community"
	// Documents the specific capabilities of a single system as 'available' for purchase or use.  May have some variability reflecting options that can be configured.
	UsageContextAgreementScopeSystemDesign UsageContextAgreementScope = "system-design"
	// Documents the specific points of a single production system or endpoint.  This may be time-specific - i.e. reflecting the system 'as it is now' or 'as it was at some point in the past'.
	UsageContextAgreementScopeSystemImplementation UsageContextAgreementScope = "system-implementation"
)
