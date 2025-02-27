// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/characteristic-combination
*/type CharacteristicCombination string

const (
	// Combine characteristics with AND.
	CharacteristicCombinationAllOf CharacteristicCombination = "all-of"
	// Combine characteristics with OR.
	CharacteristicCombinationAnyOf CharacteristicCombination = "any-of"
	// Meet at least the threshold number of characteristics for definition.
	CharacteristicCombinationAtLeast CharacteristicCombination = "at-least"
	// Meet at most the threshold number of characteristics for definition.
	CharacteristicCombinationAtMost CharacteristicCombination = "at-most"
	// Combine characteristics statistically. Use method to specify the statistical method.
	CharacteristicCombinationStatistical CharacteristicCombination = "statistical"
	// Combine characteristics by addition of benefits and subtraction of harms.
	CharacteristicCombinationNetEffect CharacteristicCombination = "net-effect"
	// Combine characteristics as a collection used as the dataset.
	CharacteristicCombinationDataset CharacteristicCombination = "dataset"
)
