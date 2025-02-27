// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/metric-calibration-state
*/type DeviceMetricCalibrationState string

const (
	// The metric has not been calibrated.
	DeviceMetricCalibrationStateNotCalibrated DeviceMetricCalibrationState = "not-calibrated"
	// The metric needs to be calibrated.
	DeviceMetricCalibrationStateCalibrationRequired DeviceMetricCalibrationState = "calibration-required"
	// The metric has been calibrated.
	DeviceMetricCalibrationStateCalibrated DeviceMetricCalibrationState = "calibrated"
	// The state of calibration of this metric is unspecified.
	DeviceMetricCalibrationStateUnspecified DeviceMetricCalibrationState = "unspecified"
)
