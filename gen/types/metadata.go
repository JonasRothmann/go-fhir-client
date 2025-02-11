package types

var metadataTypesToJen = map[string]Jenable{
	"ContactDetail":         NewQualType(dataTypesPath, "ContactDetail"),
	"DataRequirement":       NewQualType(dataTypesPath, "DataRequirement"),
	"Expression":            NewQualType(dataTypesPath, "Expression"),
	"ParameterDefinition":   NewQualType(dataTypesPath, "ParameterDefinition"),
	"RelatedArtifact":       NewQualType(dataTypesPath, "RelatedArtifact"),
	"TriggerDefinition":     NewQualType(dataTypesPath, "TriggerDefinition"),
	"UsageContext":          NewQualType(dataTypesPath, "UsageContext"),
	"Availability":          NewQualType(dataTypesPath, "Availability"),
	"ExtendedContactDetail": NewQualType(dataTypesPath, "ExtendedContactDetail"),
}

var metadataTypes = Keys(metadataTypesToJen)
