// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/reference-handling-policy
*/type ReferenceHandlingPolicy string

const (
	// The server supports and populates Literal references (i.e. using Reference.reference) where they are known (this code does not guarantee that all references are literal; see 'enforced').
	ReferenceHandlingPolicyLiteral ReferenceHandlingPolicy = "literal"
	// The server allows logical references (i.e. using Reference.identifier).
	ReferenceHandlingPolicyLogical ReferenceHandlingPolicy = "logical"
	// The server will attempt to resolve logical references to literal references - i.e. converting Reference.identifier to Reference.reference (if resolution fails, the server may still accept resources; see logical).
	ReferenceHandlingPolicyResolves ReferenceHandlingPolicy = "resolves"
	// The server enforces that references have integrity - e.g. it ensures that references can always be resolved. This is typically the case for clinical record systems, but often not the case for middleware/proxy systems.
	ReferenceHandlingPolicyEnforced ReferenceHandlingPolicy = "enforced"
	// The server does not support references that point to other servers.
	ReferenceHandlingPolicyLocal ReferenceHandlingPolicy = "local"
)
