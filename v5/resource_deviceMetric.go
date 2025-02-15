// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/DeviceMetric
type DeviceMetric struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `bson:"contained,omitempty" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Instance identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Identity of metric, for example Heart Rate or PEEP Setting
	Type CodeableConcept `bson:"type" json:"type"`
	// Unit of Measure for the Metric
	Unit *CodeableConcept `bson:"unit,omitempty" json:"unit,omitempty"`
	// Describes the link to the Device
	Device custom.Reference[Device] `bson:"device" json:"device"`
	// on | off | standby | entered-in-error
	OperationalStatus *custom.Code `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	// Color name (from CSS4) or #RRGGBB code
	Color *custom.Code `bson:"color,omitempty" json:"color,omitempty"`
	// measurement | setting | calculation | unspecified
	Category custom.Code `bson:"category" json:"category"`
	// Indicates how often the metric is taken or recorded
	MeasurementFrequency *Quantity `bson:"measurementFrequency,omitempty" json:"measurementFrequency,omitempty"`
	// Describes the calibrations that have been performed or that are required to be performed
	Calibration []DeviceMetricCalibration `bson:"calibration,omitempty" json:"calibration,omitempty"`
}

type DeviceMetricCalibration struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// unspecified | offset | gain | two-point
	Type *custom.Code `bson:"type,omitempty" json:"type,omitempty"`
	// not-calibrated | calibration-required | calibrated | unspecified
	State *custom.Code `bson:"state,omitempty" json:"state,omitempty"`
	// Describes the time last calibration has been performed
	Time *custom.Instant `bson:"time,omitempty" json:"time,omitempty"`
}

func (d DeviceMetric) ResourceType() string {
	return "StructureDefinition"
}
