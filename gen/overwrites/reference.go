package overwrites

/*
//
// [Reference]: https://www.hl7.org/fhir/references.html#Reference
type Reference[T fhirclient.Resource] struct {
	Id         *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension  []datatypes.Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Reference  *string               `bson:"reference,omitempty" json:"reference,omitempty"`
	Type       *datatypes.Uri        `bson:"type,omitempty" json:"type,omitempty"`
	Identifier *datatypes.Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Display    *string               `bson:"display,omitempty" json:"display,omitempty"`
}

func (r Reference[T]) Read(ctx context.Context, client *fhirclient.Client, summary *query.SearchSummary) (T, error) {
	return query.Read[T](ctx, client, *r.Id, nil)
}
*/
