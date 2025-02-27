// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/security-source-type
  - 1
  - 2
  - 3
  - 4
  - 5
  - 6
  - 7
  - 8
*/type AuditEventSourceType string

const (
	// End-user display device, diagnostic device.
	AuditEventSourceTypeUserDevice AuditEventSourceType = "1"
	// Data acquisition device or instrument.
	AuditEventSourceTypeDataInterface AuditEventSourceType = "2"
	// Web Server process or thread.
	AuditEventSourceTypeWebServer AuditEventSourceType = "3"
	// Application Server process or thread.
	AuditEventSourceTypeApplicationServer AuditEventSourceType = "4"
	// Database Server process or thread.
	AuditEventSourceTypeDatabaseServer AuditEventSourceType = "5"
	// Security server, e.g. a domain controller.
	AuditEventSourceTypeSecurityServer AuditEventSourceType = "6"
	// ISO level 1-3 network component.
	AuditEventSourceTypeNetworkDevice AuditEventSourceType = "7"
	// ISO level 4-6 operating software.
	AuditEventSourceTypeNetworkRouter AuditEventSourceType = "8"
)
