// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/property-representation
*/type PropertyRepresentation string

const (
	// In XML, this property is represented as an attribute not an element.
	PropertyRepresentationXmlAttr PropertyRepresentation = "xmlAttr"
	// This element is represented using the XML text attribute (primitives only).
	PropertyRepresentationXmlText PropertyRepresentation = "xmlText"
	// The type of this element is indicated using xsi:type.
	PropertyRepresentationTypeAttr PropertyRepresentation = "typeAttr"
	// Use CDA narrative instead of XHTML.
	PropertyRepresentationCdaText PropertyRepresentation = "cdaText"
	// The property is represented using XHTML.
	PropertyRepresentationXhtml PropertyRepresentation = "xhtml"
)
