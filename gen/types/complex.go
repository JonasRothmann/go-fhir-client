package types

var complexDatatypesToJen = map[string]Jenable{
	"DataType":        NewQualType(dataTypesPath, "DataType"),
	"Quantity":        NewQualType(dataTypesPath, "Quantity"),
	"Age":             NewQualType(dataTypesPath, "Age"),
	"Distance":        NewQualType(dataTypesPath, "Distance"),
	"Duration":        NewQualType(dataTypesPath, "Duration"),
	"Count":           NewQualType(dataTypesPath, "Count"),
	"MoneyQuantity":   NewQualType(dataTypesPath, "MoneyQuantity"),
	"SimpleQuantity":  NewQualType(dataTypesPath, "SimpleQuantity"),
	"Attachment":      NewQualType(dataTypesPath, "Attachment"),
	"Range":           NewQualType(dataTypesPath, "Range"),
	"Period":          NewQualType(dataTypesPath, "Period"),
	"Ratio":           NewQualType(dataTypesPath, "Ratio"),
	"RatioRange":      NewQualType(dataTypesPath, "RatioRange"),
	"CodeableConcept": NewQualType(dataTypesPath, "CodeableConcept"),
	"Coding":          NewQualType(dataTypesPath, "Coding"),
	"Annotation":      NewQualType(dataTypesPath, "Annotation"),
	"Money":           NewQualType(dataTypesPath, "Money"),
	"Identifier":      NewQualType(dataTypesPath, "Identifier"),
	"HumanName":       NewQualType(dataTypesPath, "HumanName"),
	"Address":         NewQualType(dataTypesPath, "Address"),
	"ContactPoint":    NewQualType(dataTypesPath, "ContactPoint"),
	"Timing":          NewQualType(dataTypesPath, "Timing"),
	"Repeat":          NewQualType(dataTypesPath, "Repeat"), // Nested within Timing
	"Signature":       NewQualType(dataTypesPath, "Signature"),
	"SampledData":     NewQualType(dataTypesPath, "SampledData"),
	"BackboneType":    NewQualType(dataTypesPath, "BackboneType"),
}
var complexDatatypes = Keys(complexDatatypesToJen)
