// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Timing
type Timing struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// When the event occurs
	Event []custom.DateTime `bson:"event,omitempty" json:"event,omitempty"`
	// When the event is to occur
	Repeat *TimingRepeat `bson:"repeat,omitempty" json:"repeat,omitempty"`
	// C | BID | TID | QID | AM | PM | QD | QOD | +
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}

type TimingRepeat struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Length/Range of lengths, or (Start and/or end) limits
	Bounds *Duration `bson:"bounds,omitempty" json:"bounds,omitempty"`
	// Number of times to repeat
	Count *uint32 `bson:"count,omitempty" json:"count,omitempty"`
	// Maximum number of times to repeat
	CountMax *uint32 `bson:"countMax,omitempty" json:"countMax,omitempty"`
	// How long when it happens
	Duration *json.Number `bson:"duration,omitempty" json:"duration,omitempty"`
	// How long when it happens (Max)
	DurationMax *json.Number `bson:"durationMax,omitempty" json:"durationMax,omitempty"`
	// s | min | h | d | wk | mo | a - unit of time (UCUM)
	DurationUnit *custom.Code `bson:"durationUnit,omitempty" json:"durationUnit,omitempty"`
	// Indicates the number of repetitions that should occur within a period. I.e. Event occurs frequency times per period
	Frequency *uint32 `bson:"frequency,omitempty" json:"frequency,omitempty"`
	// Event occurs up to frequencyMax times per period
	FrequencyMax *uint32 `bson:"frequencyMax,omitempty" json:"frequencyMax,omitempty"`
	// The duration to which the frequency applies. I.e. Event occurs frequency times per period
	Period *json.Number `bson:"period,omitempty" json:"period,omitempty"`
	// Upper limit of period (3-4 hours)
	PeriodMax *json.Number `bson:"periodMax,omitempty" json:"periodMax,omitempty"`
	// s | min | h | d | wk | mo | a - unit of time (UCUM)
	PeriodUnit *custom.Code `bson:"periodUnit,omitempty" json:"periodUnit,omitempty"`
	// mon | tue | wed | thu | fri | sat | sun
	DayOfWeek []custom.Code `bson:"dayOfWeek,omitempty" json:"dayOfWeek,omitempty"`
	// Time of day for action
	TimeOfDay []custom.Time `bson:"timeOfDay,omitempty" json:"timeOfDay,omitempty"`
	// Code for time period of occurrence
	When []custom.Code `bson:"when,omitempty" json:"when,omitempty"`
	// Minutes from event (before or after)
	Offset *uint32 `bson:"offset,omitempty" json:"offset,omitempty"`
}

func (t Timing) ResourceType() string {
	return "StructureDefinition"
}
