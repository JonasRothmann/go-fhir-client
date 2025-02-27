// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DeviceMetric
type DeviceMetric struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Instance identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Identity of metric, for example Heart Rate or PEEP Setting
	Type CodeableConcept `json:"type"`
	// Unit of Measure for the Metric
	Unit *CodeableConcept `json:"unit,omitempty"`
	// Describes the link to the Device
	Device custom.Reference[Device] `json:"device"`
	// on | off | standby | entered-in-error
	OperationalStatus *custom.Code `json:"operationalStatus,omitempty"`
	// Color name (from CSS4) or #RRGGBB code
	Color *custom.Code `json:"color,omitempty"`
	// measurement | setting | calculation | unspecified
	Category custom.Code `json:"category"`
	// Indicates how often the metric is taken or recorded
	MeasurementFrequency *Quantity `json:"measurementFrequency,omitempty"`
	// Describes the calibrations that have been performed or that are required to be performed
	Calibration []DeviceMetricCalibration `json:"calibration,omitempty"`
}

type DeviceMetricCalibration struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// unspecified | offset | gain | two-point
	Type *custom.Code `json:"type,omitempty"`
	// not-calibrated | calibration-required | calibrated | unspecified
	State *custom.Code `json:"state,omitempty"`
	// Describes the time last calibration has been performed
	Time *custom.Instant `json:"time,omitempty"`
}

type OtherDeviceMetric DeviceMetric

func (d DeviceMetric) ResourceType() string {
	return "DeviceMetric"
}

func (d DeviceMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceMetric
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDeviceMetric: OtherDeviceMetric(d), ResourceType: d.ResourceType()})
}

func UnmarshalDeviceMetric(b []byte) (res DeviceMetric, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
