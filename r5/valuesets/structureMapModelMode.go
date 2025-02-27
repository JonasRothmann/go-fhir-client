// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/map-model-mode
*/type StructureMapModelMode string

const (
	// This structure describes an instance passed to the mapping engine that is used a source of data.
	StructureMapModelModeSource StructureMapModelMode = "source"
	// This structure describes an instance that the mapping engine may ask for that is used a source of data.
	StructureMapModelModeQueried StructureMapModelMode = "queried"
	// This structure describes an instance passed to the mapping engine that is used a target of data.
	StructureMapModelModeTarget StructureMapModelMode = "target"
	// This structure describes an instance that the mapping engine may ask to create that is used a target of data.
	StructureMapModelModeProduced StructureMapModelMode = "produced"
)
