package query

import (
	"context"
	"fmt"
	"net/url"

	fhirclient "github.com/jonasrothmann/go-fhir-client"
)

// https://www.hl7.org/fhir/http.html#read
func Read[T fhirclient.Resource](ctx context.Context, client *fhirclient.Client, id string, summary *SearchSummary) (result T, err error) {
	url := url.URL{
		Path: fmt.Sprintf("/%s/%s", result.ResourceType(), id),
	}

	q := url.Query()
	if summary != nil {
		q.Set("_summary", string(*summary))
	}
	url.RawQuery = q.Encode()

	_, err = client.HttpClient.GetWithContext(ctx, url, &result)

	return result, err
}

/*
https://www.hl7.org/fhir/http.html#create
id, meta.versionId and meta.lastUpdated are ignored
*/
func Create[T fhirclient.Resource](ctx context.Context, client *fhirclient.Client, input T) (result T, err error) {
	url := url.URL{
		Path: fmt.Sprintf("/%s", result.ResourceType()),
	}

	q := url.Query()
	url.RawQuery = q.Encode()

	_, err = client.HttpClient.PostWithContext(ctx, url, input, &result)

	return result, err
}
