// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/permission-rule-combining
*/type PermissionRuleCombining string

const (
	// The deny overrides combining algorithm is intended for those cases where a deny decision should have priority over a permit decision.
	PermissionRuleCombiningDenyOverrides PermissionRuleCombining = "deny-overrides"
	// The permit overrides combining algorithm is intended for those cases where a permit decision should have priority over a deny decision.
	PermissionRuleCombiningPermitOverrides PermissionRuleCombining = "permit-overrides"
	// The behavior of this algorithm is identical to that of the “Deny-overrides” rule-combining algorithm with one exception.  The order in which the collection of rules is evaluated SHALL match the order as listed in the permission.
	PermissionRuleCombiningOrderedDenyOverrides PermissionRuleCombining = "ordered-deny-overrides"
	// The behavior of this algorithm is identical to that of the “Permit-overrides” rule-combining algorithm with one exception.  The order in which the collection of rules is evaluated SHALL match the order as listed in the permission.
	PermissionRuleCombiningOrderedPermitOverrides PermissionRuleCombining = "ordered-permit-overrides"
	// The “Deny-unless-permit” combining algorithm is intended for those cases where a permit decision should have priority over a deny decision, and an “Indeterminate” or “NotApplicable” must never be the result. It is particularly useful at the top level in a policy structure to ensure that a PDP will always return a definite “Permit” or “Deny” result.
	PermissionRuleCombiningDenyUnlessPermit PermissionRuleCombining = "deny-unless-permit"
	// The “Permit-unless-deny” combining algorithm is intended for those cases where a deny decision should have priority over a permit decision, and an “Indeterminate” or “NotApplicable” must never be the result. It is particularly useful at the top level in a policy structure to ensure that a PDP will always return a definite “Permit” or “Deny” result. This algorithm has the following behavior.
	PermissionRuleCombiningPermitUnlessDeny PermissionRuleCombining = "permit-unless-deny"
)
