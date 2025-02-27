// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/location-physical-type
*/type LocationForm string

const (
	// A collection of buildings or other locations such as a site or a campus.
	LocationFormSi LocationForm = "si"
	// Any Building or structure. This may contain rooms, corridors, wings, etc. It might not have walls, or a roof, but is considered a defined/allocated space.
	LocationFormBu LocationForm = "bu"
	// A Wing within a Building, this often contains levels, rooms and corridors.
	LocationFormWi LocationForm = "wi"
	// A Ward is a section of a medical facility that may contain rooms and other types of location.
	LocationFormWa LocationForm = "wa"
	// A Level in a multi-level Building/Structure.
	LocationFormLvl LocationForm = "lvl"
	// Any corridor within a Building, that may connect rooms.
	LocationFormCo LocationForm = "co"
	// A space that is allocated as a room, it may have walls/roof etc., but does not require these.
	LocationFormRo LocationForm = "ro"
	// A space that is allocated for sleeping/laying on. This is not the physical bed/trolley that may be moved about, but the space it may occupy.
	LocationFormBd LocationForm = "bd"
	// A means of transportation.
	LocationFormVe LocationForm = "ve"
	// A residential dwelling. Usually used to reference a location that a person/patient may reside.
	LocationFormHo LocationForm = "ho"
	// A container that can store goods, equipment, medications or other items.
	LocationFormCa LocationForm = "ca"
	// A defined path to travel between 2 points that has a known name.
	LocationFormRd LocationForm = "rd"
	// A defined physical boundary of something, such as a flood risk zone, region, postcode
	LocationFormArea LocationForm = "area"
	// A wide scope that covers a conceptual domain, such as a Nation (Country wide community or Federal Government - e.g. Ministry of Health),  Province or State (community or Government), Business (throughout the enterprise), Nation with a business scope of an agency (e.g. CDC, FDA etc.) or a Business segment (UK Pharmacy), not just an physical boundary
	LocationFormJdn LocationForm = "jdn"
	// A location that is virtual in nature, such as a conference call or virtual meeting space
	LocationFormVi LocationForm = "vi"
)
