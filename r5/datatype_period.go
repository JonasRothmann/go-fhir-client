// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Period
type Period struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Starting time with inclusive boundary
	Start *custom.DateTime `json:"start,omitempty"`
	// End time with inclusive boundary, if not ongoing
	End *custom.DateTime `json:"end,omitempty"`
}

type OtherPeriod Period

func (p Period) ResourceType() string {
	return "Period"
}

func (p Period) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPeriod
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPeriod: OtherPeriod(p), ResourceType: p.ResourceType()})
}

func UnmarshalPeriod(b []byte) (res Period, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
