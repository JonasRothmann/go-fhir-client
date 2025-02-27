// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/map-transform
*/type StructureMapTransform string

const (
	// create(type : string) - type is passed through to the application on the standard API, and must be known by it.
	StructureMapTransformCreate StructureMapTransform = "create"
	// copy(source).
	StructureMapTransformCopy StructureMapTransform = "copy"
	// truncate(source, length) - source must be stringy type.
	StructureMapTransformTruncate StructureMapTransform = "truncate"
	// escape(source, fmt1, fmt2) - change source from one kind of escaping to another (plain, java, xml, json). note that this is for when the string itself is escaped.
	StructureMapTransformEscape StructureMapTransform = "escape"
	// cast(source, type?) - cast (convert) source from one type to another. Target type can be left as implicit if there is one and only one target type known. The default namespace for the type is 'FHIR' (see [FHIRPath type specifiers](http://hl7.org/fhirpath/N1/#is-type-specifier))
	StructureMapTransformCast StructureMapTransform = "cast"
	// append(source...) - source is element or string.
	StructureMapTransformAppend StructureMapTransform = "append"
	// translate(source, uri_of_map) - use the translate operation.
	StructureMapTransformTranslate StructureMapTransform = "translate"
	// reference(source : object) - return a string that references the provided tree properly.
	StructureMapTransformReference StructureMapTransform = "reference"
	// Perform a date operation. *Parameters to be documented*.
	StructureMapTransformDateOp StructureMapTransform = "dateOp"
	// Generate a random UUID (in lowercase). No Parameters.
	StructureMapTransformUuid StructureMapTransform = "uuid"
	// Return the appropriate string to put in a reference that refers to the resource provided as a parameter.
	StructureMapTransformPointer StructureMapTransform = "pointer"
	// Execute the supplied FHIRPath expression and use the value returned by that.
	StructureMapTransformEvaluate StructureMapTransform = "evaluate"
	// Create a CodeableConcept. Parameters = (text) or (system. Code[, display]).
	StructureMapTransformCc StructureMapTransform = "cc"
	// Create a Coding. Parameters = (system. Code[, display]).
	StructureMapTransformC StructureMapTransform = "c"
	// Create a quantity. Parameters = (text) or (value, unit, [system, code]) where text is the natural representation e.g. [comparator]value[space]unit.
	StructureMapTransformQty StructureMapTransform = "qty"
	// Create an identifier. Parameters = (system, value[, type]) where type is a code from the identifier type value set.
	StructureMapTransformId StructureMapTransform = "id"
	// Create a contact details. Parameters = (value) or (system, value). If no system is provided, the system should be inferred from the content of the value.
	StructureMapTransformCp StructureMapTransform = "cp"
)
