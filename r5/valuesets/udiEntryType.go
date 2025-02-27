// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/udi-entry-type
*/type UDIEntryType string

const (
	// a barcodescanner captured the data from the device label.
	UDIEntryTypeBarcode UDIEntryType = "barcode"
	// An RFID chip reader captured the data from the device label.
	UDIEntryTypeRfid UDIEntryType = "rfid"
	// The data was read from the label by a person and manually entered. (e.g.  via a keyboard).
	UDIEntryTypeManual UDIEntryType = "manual"
	// The data originated from a patient's implant card and was read by an operator.
	UDIEntryTypeCard UDIEntryType = "card"
	// The data originated from a patient source and was not directly scanned or read from a label or card.
	UDIEntryTypeSelfReported UDIEntryType = "self-reported"
	// The UDI information was received electronically from the device through a communication protocol, such as the IEEE 11073 20601 version 4 exchange protocol over Bluetooth or USB.
	UDIEntryTypeElectronicTransmission UDIEntryType = "electronic-transmission"
	// The method of data capture has not been determined.
	UDIEntryTypeUnknown UDIEntryType = "unknown"
)
