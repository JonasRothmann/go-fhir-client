package query

import (
	"context"
	"fmt"
	"time"

	fhirclient "github.com/jonasrothmann/go-fhir-client"
)

type searchQuery[T fhirclient.Resource] struct {
	ID          *fmt.Stringer `params:"_id"`
	Content     *string       `params:"_content"`
	Filter      string        `params:"_filter"`
	Has         *string       `params:"_has"`
	In          *string       `params:"_in"`
	Language    *string       `params:"_language"`
	LastUpdated *time.Time    `params:"_lastUpdated"`
	List        *string       `params:"_list"`
	Query       *string       `params:"_query"`
	Security    *string       `params:"_security"`
	Source      *string       `params:"_source"`
	Tag         *string       `params:"_tag"`
	Text        *string       `params:"_text"`
	Type        *string       `params:"_type"`
}

func Search[T fhirclient.Resource](ctx context.Context, client *fhirclient.Client) { //(result datatypes.Bundle[T], err error) {
	panic("not implemented")
}

// https://www.hl7.org/fhir/search.html#_summary
type SearchSummary string

const (
	/*Return a limited subset of elements from the resource.
	This subset SHOULD consist solely of all supported elements that are marked as "summary" in the base definition of the resource(s) (see ElementDefinition.isSummary)*/
	SearchSummaryTrue SearchSummary = "true"

	/*[TRIAL USE]
	Return only the text, id, meta, and top-level mandatory elements (these mandatory elements are included to ensure that the payload is valid FHIR; servers MAY omit elements within these sub-trees as long as they ensure that the payload is valid). Servers MAY return extensions, but clients SHOULD NOT rely on extensions being present and SHOULD use another search mode if data contained in extensions is required.*/
	SearchSummaryText SearchSummary = "text"

	// Remove the text element
	SearchSummaryData SearchSummary = "data"

	// Search only: just return a count of the matching resources, without returning the actual matches
	SearchSummaryCount SearchSummary = "count"

	// Return all parts of the resource(s)
	SearchSummaryFalse SearchSummary = "false"
)
