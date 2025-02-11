package types

var otherDatatypesToJen = map[string]Jenable{
	"Extension": NewQualType(dataTypesPath, "Extension"),
	"Narrative": NewQualType(dataTypesPath, "Narrative"),
	"Reference": NewGenericQualType(dataTypesPath, "Reference"),
}
var otherDatatypes = Keys(otherDatatypesToJen)
