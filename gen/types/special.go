package types

var specialDatatypesToJen = map[string]Jenable{
	"Meta":   NewQualType(dataTypesPath, "Meta"),
	"Dosage": NewQualType(dataTypesPath, "Dosage"),
}
var specialDatatypes = Keys(specialDatatypesToJen)
