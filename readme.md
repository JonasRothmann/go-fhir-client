##

### Todo:
[ ] Versioning
[ ] Search


## HTTP support:

This document outlines the compatibility of the implemented FHIR operations with our HTTP REST API. Currently, only the **Read** operation is supported via the FHIR client. All other operations are not yet implemented.


| Implemented | Interaction    | HTTP Method(s)                                   | FHIR Client Function Signature                                                                                                                        |
|-------------|----------------|--------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| [x]         | **Read**       | GET `[base]/[type]/[id]{?_format,...}`           | `query.Read[T resources.Resource](ctx context.Context, client *fhirclient.Client, id datatypes.ID, summary *SearchSummary)`                              |
| [ ]         | **vRead**      | GET `[base]/[type]/[id]/_history/[vid]{?_format,...}` | `query.VRead[T resources.Resource](ctx context.Context, client *fhirclient.Client, id datatypes.ID, version datatypes.ID)`                                |
| [ ]         | **Update**     | PUT `[base]/[type]/[id]{?_format,...}`           | `query.Update[T resources.Resource](ctx context.Context, client *fhirclient.Client, id datatypes.ID, resource T)`                                         |
| [ ]         | **Patch**      | PATCH `[base]/[type]/[id]{?_format,...}`         | `query.Patch[T resources.Resource](ctx context.Context, client *fhirclient.Client, id datatypes.ID, patch PatchData)`                                      |
| [ ]         | **Delete**     | DELETE `[base]/[type]/[id]`                        | `query.Delete[T resources.Resource](ctx context.Context, client *fhirclient.Client, id datatypes.ID)`                                                    |
| [ ]         | **Create**     | POST `[base]/[type]{?_format,...}`               | `query.Create[T resources.Resource](ctx context.Context, client *fhirclient.Client, resource T)`                                                         |
| [ ]         | **Search**     | GET or POST `[base]/[type]?…` or `[base]/[type]/_search` | `query.Search[T resources.Resource](ctx context.Context, client *fhirclient.Client, params SearchParams)`                                                |
| [ ]         | **History**    | GET `[base]/[type]/[id]/_history{?_format,...}` or `[base]/_history` | `query.History[T resources.Resource](ctx context.Context, client *fhirclient.Client, id datatypes.ID, params HistoryParams)`                              |
| [ ]         | **Capabilities** | GET `[base]/metadata{?mode,...}`             | `query.Capabilities(ctx context.Context, client *fhirclient.Client)`                                                                                   |
| [ ]         | **Transaction**  | POST `[base]{?_format,...}`                    | `query.Transaction(ctx context.Context, client *fhirclient.Client, bundle Bundle)`                                                                       |
| [ ]         | **Batch**        | POST `[base]{?_format,...}`                    | `query.Batch(ctx context.Context, client *fhirclient.Client, bundle Bundle)`                                                                             |
