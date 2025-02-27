// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/location-mode
*/type LocationMode string

const (
	// The Location resource represents a specific instance of a location (e.g. Operating Theatre 1A).
	LocationModeInstance LocationMode = "instance"
	// The Location represents a class of locations (e.g. Any Operating Theatre) although this class of locations could be constrained within a specific boundary (such as organization, or parent location, address etc.).
	LocationModeKind LocationMode = "kind"
)
