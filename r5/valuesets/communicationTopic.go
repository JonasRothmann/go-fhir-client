// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/communication-topic
*/type CommunicationTopic string

const (
	// The purpose or content of the communication is a prescription refill request.
	CommunicationTopicPrescriptionRefillRequest CommunicationTopic = "prescription-refill-request"
	// The purpose or content of the communication is a progress update.
	CommunicationTopicProgressUpdate CommunicationTopic = "progress-update"
	// The purpose or content of the communication is to report labs.
	CommunicationTopicReportLabs CommunicationTopic = "report-labs"
	// The purpose or content of the communication is an appointment reminder.
	CommunicationTopicAppointmentReminder CommunicationTopic = "appointment-reminder"
	// The purpose or content of the communication is a phone consult.
	CommunicationTopicPhoneConsult CommunicationTopic = "phone-consult"
	// The purpose or content of the communication is a summary report.
	CommunicationTopicSummaryReport CommunicationTopic = "summary-report"
)
