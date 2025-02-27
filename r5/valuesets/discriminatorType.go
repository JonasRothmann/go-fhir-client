// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/discriminator-type
*/type DiscriminatorType string

const (
	// The slices have different values in the nominated element, as determined by the applicable fixed value, pattern, or required ValueSet binding.
	DiscriminatorTypeValue DiscriminatorType = "value"
	// The slices are differentiated by the presence or absence of the nominated element. There SHALL be no more than two slices. The slices are differentiated by the fact that one must have a max of 0 and the other must have a min of 1 (or more).  The order in which the slices are declared doesn't matter.
	DiscriminatorTypeExists DiscriminatorType = "exists"
	// The slices have different values in the nominated element, as determined by the applicable fixed value, pattern, or required ValueSet binding. This has the same meaning as 'value' and is deprecated.
	DiscriminatorTypePattern DiscriminatorType = "pattern"
	// The slices are differentiated by type of the nominated element.
	DiscriminatorTypeType DiscriminatorType = "type"
	// The slices are differentiated by conformance of the nominated element to a specified profile. Note that if the path specifies .resolve() then the profile is the target profile on the reference. In this case, validation by the possible profiles is required to differentiate the slices.
	DiscriminatorTypeProfile DiscriminatorType = "profile"
	// The slices are differentiated by their index. This is only possible if all but the last slice have min=max cardinality, and the (optional) last slice contains other undifferentiated elements.
	DiscriminatorTypePosition DiscriminatorType = "position"
)
