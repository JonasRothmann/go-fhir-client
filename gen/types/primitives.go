package types

var primitiveTypesToJen = map[string]Jenable{
	"PrimitiveType":                         NewQualType(dataTypesPath, "PrimitiveType"),
	"instant":                               NewQualType(dataTypesPath, "Instant"),
	"time":                                  NewQualType("time", "Time"), // TODO: Change
	"date":                                  NewQualType("time", "Time"),
	"dateTime":                              NewQualType("time", "Time"),
	"decimal":                               NewQualType("encoding/json", "Number"),
	"boolean":                               BuiltInType("bool"),
	"integer":                               BuiltInType("int32"),
	"string":                                BuiltInType("string"),
	"http://hl7.org/fhirpath/System.String": BuiltInType("string"),
	"http://hl7.org/fhirpath/System.DateTime": NewQualType("time", "Time"),
	"http://hl7.org/fhirpath/System.Date":     NewQualType("time", "Time"),
	"uri":                                     NewQualType(dataTypesPath, "URI"),
	"base64Binary":                            NewQualType(dataTypesPath, "Base64Binary"),
	"code":                                    NewQualType(dataTypesPath, "Code"),
	"id":                                      NewQualType(dataTypesPath, "ID"),
	"oid":                                     NewQualType(dataTypesPath, "OID"),
	"unsignedInt":                             BuiltInType("uint32"),
	"positiveInt":                             BuiltInType("uint32"),
	"markdown":                                NewQualType(dataTypesPath, "Markdown"),
	"url":                                     NewQualType(dataTypesPath, "URL"),
	"canonical":                               NewGenericQualType(dataTypesPath, "Canonical"),
	"uuid":                                    NewQualType(dataTypesPath, "UUID"),
	"integer64":                               BuiltInType("int64"),
}
var primitiveTypes = Keys(primitiveTypesToJen)
