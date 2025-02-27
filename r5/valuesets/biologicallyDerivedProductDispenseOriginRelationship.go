// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/biologicallyderivedproductdispense-origin-relationship
*/type BiologicallyDerivedProductDispenseOriginRelationship string

const (
	// The product was pre-donated by the recipient
	BiologicallyDerivedProductDispenseOriginRelationshipAutologous BiologicallyDerivedProductDispenseOriginRelationship = "autologous"
	// The product is from a blood relation of the recipient
	BiologicallyDerivedProductDispenseOriginRelationshipRelated BiologicallyDerivedProductDispenseOriginRelationship = "related"
	// The donor has been specifically selected to provide product for the recipient
	BiologicallyDerivedProductDispenseOriginRelationshipDirected BiologicallyDerivedProductDispenseOriginRelationship = "directed"
	// The donor and the recipient are unrelated
	BiologicallyDerivedProductDispenseOriginRelationshipAllogeneic BiologicallyDerivedProductDispenseOriginRelationship = "allogeneic"
	// The product is from a different species to the recipient
	BiologicallyDerivedProductDispenseOriginRelationshipXenogenic BiologicallyDerivedProductDispenseOriginRelationship = "xenogenic"
)
