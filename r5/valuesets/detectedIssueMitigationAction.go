// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/v3-ActCode
*/type DetectedIssueMitigationAction string

const (
	// Codes dealing with the management of Detected Issue observations
	DetectedIssueMitigationActionActDetectedIssueManagementCode   DetectedIssueMitigationAction = "_ActDetectedIssueManagementCode"
	DetectedIssueMitigationActionAuthorizationIssueManagementCode DetectedIssueMitigationAction = "_AuthorizationIssueManagementCode"
	// Codes dealing with the management of Detected Issue observations for the administrative and patient administrative acts domains.
	DetectedIssueMitigationActionActAdministrativeDetectedIssueManagementCode DetectedIssueMitigationAction = "_ActAdministrativeDetectedIssueManagementCode"
	// Confirmed drug therapy appropriate
	DetectedIssueMitigationActionTherapyAppropriate DetectedIssueMitigationAction = "1"
	// Confirmed supply action appropriate
	DetectedIssueMitigationActionSupplyAppropriate DetectedIssueMitigationAction = "14"
	// Order is performed as issued, but other action taken to mitigate potential adverse effects
	DetectedIssueMitigationActionOtherActionTaken DetectedIssueMitigationAction = "8"
	// Codes dealing with the management of Detected Issue observations for the financial acts domain.
	DetectedIssueMitigationActionActFinancialDetectedIssueManagementCode DetectedIssueMitigationAction = "_ActFinancialDetectedIssueManagementCode"
	// Used to temporarily override normal authorization rules to gain access to data in a case of emergency. Use of this override code will typically be monitored, and a procedure to verify its proper use may be triggered when used.
	DetectedIssueMitigationActionEmergencyAuthorizationOverride DetectedIssueMitigationAction = "EMAUTH"
	// Consulted other supplier/pharmacy, therapy confirmed
	DetectedIssueMitigationActionConsultedSupplier DetectedIssueMitigationAction = "19"
	// Assessed patient, therapy is appropriate
	DetectedIssueMitigationActionAssessedPatient DetectedIssueMitigationAction = "2"
	// The patient has the appropriate indication or diagnosis for the action to be taken.
	DetectedIssueMitigationActionAppropriateIndicationOrDiagnosis DetectedIssueMitigationAction = "22"
	// It has been confirmed that the appropriate pre-requisite therapy has been tried.
	DetectedIssueMitigationActionPriorTherapyDocumented DetectedIssueMitigationAction = "23"
	// Patient gave adequate explanation
	DetectedIssueMitigationActionPatientExplanation DetectedIssueMitigationAction = "3"
	// Consulted other supply source, therapy still appropriate
	DetectedIssueMitigationActionConsultedOtherSource DetectedIssueMitigationAction = "4"
	// Consulted prescriber, therapy confirmed
	DetectedIssueMitigationActionConsultedPrescriber DetectedIssueMitigationAction = "5"
	// Concurrent therapy triggering alert is no longer on-going or planned
	DetectedIssueMitigationActionInteractingTherapyNoLongerActivePlanned DetectedIssueMitigationAction = "7"
	// Patient's existing supply was lost/wasted
	DetectedIssueMitigationActionReplacement DetectedIssueMitigationAction = "15"
	// Supply date is due to patient vacation
	DetectedIssueMitigationActionVacationSupply DetectedIssueMitigationAction = "16"
	// Supply date is intended to carry patient over weekend
	DetectedIssueMitigationActionWeekendSupply DetectedIssueMitigationAction = "17"
	// Supply is intended for use during a leave of absence from an institution.
	DetectedIssueMitigationActionLeaveOfAbsence DetectedIssueMitigationAction = "18"
	// Supply is different than expected as an additional quantity has been supplied in a separate dispense.
	DetectedIssueMitigationActionAdditionalQuantityOnSeparateDispense DetectedIssueMitigationAction = "20"
	// Consulted prescriber and recommended change, prescriber declined
	DetectedIssueMitigationActionPrescriberDeclinedChange DetectedIssueMitigationAction = "6"
	// Provided education or training to the patient on appropriate therapy use
	DetectedIssueMitigationActionProvidedPatientEducation DetectedIssueMitigationAction = "10"
	// Instituted an additional therapy to mitigate potential negative effects
	DetectedIssueMitigationActionAddedConcurrentTherapy DetectedIssueMitigationAction = "11"
	// Suspended existing therapy that triggered interaction for the duration of this therapy
	DetectedIssueMitigationActionTemporarilySuspendedConcurrentTherapy DetectedIssueMitigationAction = "12"
	// Aborted existing therapy that triggered interaction.
	DetectedIssueMitigationActionStoppedConcurrentTherapy DetectedIssueMitigationAction = "13"
	// Arranged to monitor patient for adverse effects
	DetectedIssueMitigationActionInstitutedOngoingMonitoringProgram DetectedIssueMitigationAction = "9"
	// Indicates that the permissions have been externally verified and the request should be processed.
	DetectedIssueMitigationActionAuthorizationConfirmed DetectedIssueMitigationAction = "21"
)
