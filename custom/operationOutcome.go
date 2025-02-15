package custom

import (
	"fmt"

	valuesets "github.com/jonasrothmann/go-fhir-client/value-sets"
)

// https://build.fhir.org/operationoutcome.html
type OperationOutcome struct {
	Issue []struct {
		Severity    valuesets.IssueSeverity `json:"severity"`
		Code        valuesets.IssueType     `json:"code"`
		Diagnostics string                  `json:"diagnostics"`
	} `json:"issue"`
}

func (o OperationOutcome) String() string {
	return fmt.Sprintf("[%s] %s - %s", o.Severity(), o.Code().Message(), o.Issue[0].Diagnostics)
}

func (o OperationOutcome) Severity() valuesets.IssueSeverity {
	return o.Issue[0].Severity
}

func (o OperationOutcome) Code() valuesets.IssueType {
	return o.Issue[0].Code
}
