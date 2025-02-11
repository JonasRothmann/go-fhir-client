package datatypesold

/*

import "fmt"

// https://www.hl7.org/fhir/narrative.html
type Narrative struct {
	Element

	Status Code  `json:"modifierExtension"`
	Div    XHTML `json:"div"` // TODO: Validation
}

type XHTML string

// https://build.fhir.org/operationoutcome.html
type OperationOutcome struct {
	DomainResource

	Issue struct {
		Severity    fhir.IssueSeverity `json:"severity"`
		Code        fhir.IssueType     `json:"code"`    // TODO: fix
		Details     CodeableConcept    `json:"details"` // TODO: write
		Diagnostics string             `json:"diagnostics"`
		// Deprecated: TODO
		Location   string `json:"location"`
		Expression string `json:"expression"`
	} `json:"issue"`
}

func (o OperationOutcome) String() string {
	return fmt.Sprintf("[%s] %s - %s", o.Severity(), o.Code().Message(), o.Issue.Diagnostics)
}

func (o OperationOutcome) Severity() fhir.IssueSeverity {
	return o.Issue.Severity
}

func (o OperationOutcome) Code() fhir.IssueType {
	return o.Issue.Code
}
*/
