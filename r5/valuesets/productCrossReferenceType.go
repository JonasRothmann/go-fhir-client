// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/medicinal-product-cross-reference-type
*/type ProductCrossReferenceType string

const (
	// Link to Investigational (Development) Product
	ProductCrossReferenceTypeInvestigationalProduct ProductCrossReferenceType = "InvestigationalProduct"
	// Link Actual to Virtual Product
	ProductCrossReferenceTypeVirtualProduct ProductCrossReferenceType = "VirtualProduct"
	// Link Virtual to Actual Product
	ProductCrossReferenceTypeActualProduct ProductCrossReferenceType = "ActualProduct"
	// Link Generic to Branded Product
	ProductCrossReferenceTypeBrandedProduct ProductCrossReferenceType = "BrandedProduct"
	// Link Branded to Generic Product
	ProductCrossReferenceTypeGenericProduct ProductCrossReferenceType = "GenericProduct"
	// Link to Parallel Import Product
	ProductCrossReferenceTypeParallel ProductCrossReferenceType = "Parallel"
)
