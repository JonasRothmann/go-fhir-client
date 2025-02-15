package custom

import (
	"context"

	fhirclient "github.com/jonasrothmann/go-fhir-client"
	"github.com/jonasrothmann/go-fhir-client/query"
)

//
// [Reference]: https://www.hl7.org/fhir/references.html#Reference
type Reference[T fhirclient.Resource] struct {
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	//Extension  []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Reference  *string     `bson:"reference,omitempty" json:"reference,omitempty"`
	Type       *URI        `bson:"type,omitempty" json:"type,omitempty"`
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Display    *string     `bson:"display,omitempty" json:"display,omitempty"`
}

func (r Reference[T]) Read(ctx context.Context, client *fhirclient.Client, summary *query.SearchSummary) (T, error) {
	return query.Read[T](ctx, client, *r.Id, nil)
}

func (r Reference[T]) ResourceType() string {
	return "Reference"
}

// http://hl7.org/fhir/StructureDefinition/CodeableReference
type CodeableReference[T fhirclient.Resource] struct {
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	//Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// TODO: Reenable
	//Concept   *CodeableConcept `bson:"concept,omitempty" json:"concept,omitempty"`
	Reference *Reference[T] `bson:"reference,omitempty" json:"reference,omitempty"`
}

func (c CodeableReference[T]) ResourceType() string {
	return "CodeableReference"
}
