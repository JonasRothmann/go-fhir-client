// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/versioning-policy
*/type ResourceVersionPolicy string

const (
	// VersionId meta-property is not supported (server) or used (client).
	ResourceVersionPolicyNoVersionIdSupport ResourceVersionPolicy = "no-version"
	// VersionId meta-property is supported (server) or used (client).
	ResourceVersionPolicyVersioned ResourceVersionPolicy = "versioned"
	// Supports version-aware updates (server) or will be specified (If-match header) for updates (client).
	ResourceVersionPolicyVersionIdTrackedFully ResourceVersionPolicy = "versioned-update"
)
