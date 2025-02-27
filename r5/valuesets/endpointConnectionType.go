// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/endpoint-connection-type
*/type EndpointConnectionType string

const (
	// DICOMweb RESTful Image Retrieve - http://dicom.nema.org/medical/dicom/current/output/chtml/part18/sect_6.5.html
	EndpointConnectionTypeDicomwadors EndpointConnectionType = "dicom-wado-rs"
	// DICOMweb RESTful Image query - http://dicom.nema.org/medical/dicom/current/output/chtml/part18/sect_6.7.html
	EndpointConnectionTypeDicomqidors EndpointConnectionType = "dicom-qido-rs"
	// DICOMweb RESTful image sending and storage - http://dicom.nema.org/medical/dicom/current/output/chtml/part18/sect_6.6.html
	EndpointConnectionTypeDicomstowrs EndpointConnectionType = "dicom-stow-rs"
	// DICOMweb Image Retrieve - http://dicom.nema.org/dicom/2013/output/chtml/part18/sect_6.3.html
	EndpointConnectionTypeDicomwadouri EndpointConnectionType = "dicom-wado-uri"
	// Interact with the server interface using FHIR's RESTful interface. For details on its version/capabilities you should connect the value in Endpoint.address and retrieve the FHIR CapabilityStatement.
	EndpointConnectionTypeHl7fhir EndpointConnectionType = "hl7-fhir-rest"
	// Use the servers FHIR Messaging interface. Details can be found on the messaging.html page in the FHIR Specification. The FHIR server's base address is specified in the Endpoint.address property.
	EndpointConnectionTypeHl7fhirMessaging EndpointConnectionType = "hl7-fhir-msg"
	// HL7v2 messages over an LLP TCP connection
	EndpointConnectionTypeHl7v2mllp EndpointConnectionType = "hl7v2-mllp"
	// Email delivery using a digital certificate to encrypt the content using the public key, receiver must have the private key to decrypt the content
	EndpointConnectionTypeSecureEmail EndpointConnectionType = "secure-email"
	// Direct Project information - http://wiki.directproject.org/
	EndpointConnectionTypeDirectProject EndpointConnectionType = "direct-project"
	// A CDS Hooks Service connection. The address will be the base URL of the service as described in the CDS specification https://cds-hooks.hl7.org
	EndpointConnectionTypeCdsHooksService EndpointConnectionType = "cds-hooks-service"
	// IHE Cross Community Patient Discovery Profile (XCPD) - http://wiki.ihe.net/index.php/Cross-Community_Patient_Discovery
	EndpointConnectionTypeIhexcpd EndpointConnectionType = "ihe-xcpd"
	// IHE Cross Community Access Profile (XCA) - http://wiki.ihe.net/index.php/Cross-Community_Access
	EndpointConnectionTypeIhexca EndpointConnectionType = "ihe-xca"
	// IHE Cross-Enterprise Document Reliable Exchange (XDR) - http://wiki.ihe.net/index.php/Cross-enterprise_Document_Reliable_Interchange
	EndpointConnectionTypeIhexdr EndpointConnectionType = "ihe-xdr"
	// IHE Cross-Enterprise Document Sharing (XDS) - http://wiki.ihe.net/index.php/Cross-Enterprise_Document_Sharing
	EndpointConnectionTypeIhexds EndpointConnectionType = "ihe-xds"
	// IHE Invoke Image Display (IID) - http://wiki.ihe.net/index.php/Invoke_Image_Display
	EndpointConnectionTypeIheiid EndpointConnectionType = "ihe-iid"
)
