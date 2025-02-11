//go:generate go-enum --marshal
package datatypesold

/*
// https://www.hl7.org/fhir/types.html#Element
type Element struct {
	ID        string    `json:"id"`
	Extension Extension `json:"extension"`
}

// https://www.hl7.org/fhir/types.html#DataType
type DataType Element

// https://www.hl7.org/fhir/types.html#BackboneType
type BackboneType struct {
	DataType

	ModifierExtension Extension `json:"modifierExtension"`
}

// https://www.hl7.org/fhir/types.html#BackboneElement
type BackboneElement struct {
	DataType

	ModifierExtension Extension `json:"modifierExtension"`
}

// https://www.hl7.org/fhir/resource.html#Meta
// TODO: Remove
type Meta struct {
	VersionID   ID      `json:"id"`
	LastUpdated Instant `json:"lastUpdated"`
	Source      URL     `json:"source"`
	//Profile     Canonical[StructureDefinition] `json:"profile"`
	Security Coding `json:"security"`
	Tag      Coding `json:"tag"`
}

// https://www.hl7.org/fhir/resource.html#Resource
type Resource struct {
	ID            ID                     `json:"id"`
	Meta          Meta                   `json:"meta"`
	ImplicitRules URI                    `json:"implicitRules"`
	Language      valuesets.AllLanguages `json:"language"`
}

// https://www.hl7.org/fhir/types.html#DomainResource
type DomainResource struct {
	Resource

	//Text              Narrative `json:"text"`
	Contained         Resource  `json:"contained"`
	Extension         Extension `json:"extension"`
	ModifierExtension Extension `json:"modifierExtension"`
}
*/
