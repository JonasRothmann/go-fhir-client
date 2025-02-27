// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/audit-event-severity
*/type AuditEventSeverity string

const (
	// System is unusable. e.g., This level should only be reported by infrastructure and should not be used by applications.
	AuditEventSeverityEmergency AuditEventSeverity = "emergency"
	// Notification should be sent to trigger action be taken. e.g., Loss of the primary network connection needing attention.
	AuditEventSeverityAlert AuditEventSeverity = "alert"
	// Critical conditions. e.g., A failure in the system's primary application that will reset automatically.
	AuditEventSeverityCritical AuditEventSeverity = "critical"
	// Error conditions. e.g., An application has exceeded its file storage limit and attempts to write are failing.
	AuditEventSeverityError AuditEventSeverity = "error"
	// Warning conditions. May indicate that an error will occur if action is not taken. e.g., A non-root file system has only 2GB remaining.
	AuditEventSeverityWarning AuditEventSeverity = "warning"
	// Notice messages. Normal but significant condition. Events that are unusual, but not error conditions.
	AuditEventSeverityNotice AuditEventSeverity = "notice"
	// Normal operational messages that require no action. e.g., An application has started, paused, or ended successfully.
	AuditEventSeverityInformational AuditEventSeverity = "informational"
	// Debug-level messages. Information useful to developers for debugging the application.
	AuditEventSeverityDebug AuditEventSeverity = "debug"
)
