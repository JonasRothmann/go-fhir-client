// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/link-type
*/type LinkType string

const (
	// The patient resource containing this link must no longer be used. The link points forward to another patient resource that must be used in lieu of the patient resource that contains this link.
	LinkTypeReplacedBy LinkType = "replaced-by"
	// The patient resource containing this link is the current active patient record. The link points back to an inactive patient resource that has been merged into this resource, and should be consulted to retrieve additional referenced information.
	LinkTypeReplaces LinkType = "replaces"
	// The patient resource containing this link is in use and valid but not considered the main source of information about a patient. The link points forward to another patient resource that should be consulted to retrieve additional patient information.
	LinkTypeRefer LinkType = "refer"
	// The patient resource containing this link is in use and valid, but points to another Patient or RelatedPerson resource that is known to contain data about the same person. Data in this resource might overlap or contradict information found in the other Patient/RelatedPerson resource. This link does not indicate any relative importance of the resources concerned, and both should be regarded as equally valid.
	LinkTypeSeealso LinkType = "seealso"
)
