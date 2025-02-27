// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Contributor
type Contributor struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// author | editor | reviewer | endorser
	Type custom.Code `json:"type"`
	// Who contributed the content
	Name string `json:"name"`
	// Contact details of the contributor
	Contact []ContactDetail `json:"contact,omitempty"`
}

type OtherContributor Contributor

func (c Contributor) ResourceType() string {
	return "Contributor"
}

func (c Contributor) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherContributor
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherContributor: OtherContributor(c), ResourceType: c.ResourceType()})
}

func UnmarshalContributor(b []byte) (res Contributor, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
