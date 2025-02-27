// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Timing
type Timing struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// When the event occurs
	Event []custom.DateTime `json:"event,omitempty"`
	// When the event is to occur
	Repeat *TimingRepeat `json:"repeat,omitempty"`
	// C | BID | TID | QID | AM | PM | QD | QOD | +
	Code *CodeableConcept `json:"code,omitempty"`
}

type TimingRepeat struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Length/Range of lengths, or (Start and/or end) limits
	BoundsDuration *Duration `json:"boundsDuration,omitempty"`
	// Length/Range of lengths, or (Start and/or end) limits
	BoundsRange *Range `json:"boundsRange,omitempty"`
	// Length/Range of lengths, or (Start and/or end) limits
	BoundsPeriod *Period `json:"boundsPeriod,omitempty"`
	// Number of times to repeat
	Count *uint32 `json:"count,omitempty"`
	// Maximum number of times to repeat
	CountMax *uint32 `json:"countMax,omitempty"`
	// How long when it happens
	Duration *json.Number `json:"duration,omitempty"`
	// How long when it happens (Max)
	DurationMax *json.Number `json:"durationMax,omitempty"`
	// s | min | h | d | wk | mo | a - unit of time (UCUM)
	DurationUnit *custom.Code `json:"durationUnit,omitempty"`
	// Indicates the number of repetitions that should occur within a period. I.e. Event occurs frequency times per period
	Frequency *uint32 `json:"frequency,omitempty"`
	// Event occurs up to frequencyMax times per period
	FrequencyMax *uint32 `json:"frequencyMax,omitempty"`
	// The duration to which the frequency applies. I.e. Event occurs frequency times per period
	Period *json.Number `json:"period,omitempty"`
	// Upper limit of period (3-4 hours)
	PeriodMax *json.Number `json:"periodMax,omitempty"`
	// s | min | h | d | wk | mo | a - unit of time (UCUM)
	PeriodUnit *custom.Code `json:"periodUnit,omitempty"`
	// mon | tue | wed | thu | fri | sat | sun
	DayOfWeek []custom.Code `json:"dayOfWeek,omitempty"`
	// Time of day for action
	TimeOfDay []custom.Time `json:"timeOfDay,omitempty"`
	// Code for time period of occurrence
	When []custom.Code `json:"when,omitempty"`
	// Minutes from event (before or after)
	Offset *uint32 `json:"offset,omitempty"`
}

type OtherTiming Timing

func (t Timing) ResourceType() string {
	return "Timing"
}

func (t Timing) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTiming
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherTiming: OtherTiming(t), ResourceType: t.ResourceType()})
}

func UnmarshalTiming(b []byte) (res Timing, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
