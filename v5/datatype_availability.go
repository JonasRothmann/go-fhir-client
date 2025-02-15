// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Availability
type Availability struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Times the {item} is available
	AvailableTime []AvailabilityAvailableTime `bson:"availableTime,omitempty" json:"availableTime,omitempty"`
	// Not available during this time due to provided reason
	NotAvailableTime []AvailabilityNotAvailableTime `bson:"notAvailableTime,omitempty" json:"notAvailableTime,omitempty"`
}

type AvailabilityAvailableTime struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// mon | tue | wed | thu | fri | sat | sun
	DaysOfWeek []custom.Code `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty"`
	// Always available? i.e. 24 hour service
	AllDay *bool `bson:"allDay,omitempty" json:"allDay,omitempty"`
	// Opening time of day (ignored if allDay = true)
	AvailableStartTime *custom.Time `bson:"availableStartTime,omitempty" json:"availableStartTime,omitempty"`
	// Closing time of day (ignored if allDay = true)
	AvailableEndTime *custom.Time `bson:"availableEndTime,omitempty" json:"availableEndTime,omitempty"`
}

type AvailabilityNotAvailableTime struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Reason presented to the user explaining why time not available
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Service not available during this period
	During *Period `bson:"during,omitempty" json:"during,omitempty"`
}

func (a Availability) ResourceType() string {
	return "StructureDefinition"
}
