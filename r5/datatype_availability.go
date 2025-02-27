// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Availability
type Availability struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Times the {item} is available
	AvailableTime []AvailabilityAvailableTime `json:"availableTime,omitempty"`
	// Not available during this time due to provided reason
	NotAvailableTime []AvailabilityNotAvailableTime `json:"notAvailableTime,omitempty"`
}

type AvailabilityAvailableTime struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// mon | tue | wed | thu | fri | sat | sun
	DaysOfWeek []custom.Code `json:"daysOfWeek,omitempty"`
	// Always available? i.e. 24 hour service
	AllDay *bool `json:"allDay,omitempty"`
	// Opening time of day (ignored if allDay = true)
	AvailableStartTime *custom.Time `json:"availableStartTime,omitempty"`
	// Closing time of day (ignored if allDay = true)
	AvailableEndTime *custom.Time `json:"availableEndTime,omitempty"`
}

type AvailabilityNotAvailableTime struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Reason presented to the user explaining why time not available
	Description *string `json:"description,omitempty"`
	// Service not available during this period
	During *Period `json:"during,omitempty"`
}

type OtherAvailability Availability

func (a Availability) ResourceType() string {
	return "Availability"
}

func (a Availability) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAvailability
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherAvailability: OtherAvailability(a), ResourceType: a.ResourceType()})
}

func UnmarshalAvailability(b []byte) (res Availability, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
