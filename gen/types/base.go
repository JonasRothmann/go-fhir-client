package types

var baseDataTypesToJen = map[string]Jenable{
	"Element":           NewQualType(dataTypesPath, "Element"),
	"BackboneElement":   NewQualType(dataTypesPath, "BackboneElement"),
	"DataType":          NewQualType(dataTypesPath, "DataType"),
	"Resource":          NewQualType("encoding/json", "RawMessage"),
	"DomainResource":    NewQualType(dataTypesPath, "DomainResource"),
	"PrimitiveType":     NewQualType(dataTypesPath, "PrimitiveType"),
	"BackboneType":      NewQualType(dataTypesPath, "BackboneType"),
	"CanonicalResource": NewGenericQualType(dataTypesPath, "CanonicalResource"),
	"MetadataResource":  NewGenericQualType(dataTypesPath, "MetadataResource"),
}

var baseDataTypes = Keys(baseDataTypesToJen)
