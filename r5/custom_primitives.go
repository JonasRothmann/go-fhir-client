package r5

import (
	"encoding/json"
	"fmt"
	"time"
)

var PrimitiveTypes = map[string]any{
	"instant":      Instant{},
	"time":         time.Time{}, // TODO: Change
	"date":         time.Time{},
	"dateTime":     time.Time{},
	"decimal":      json.Number(""),
	"boolean":      bool(true),
	"integer":      int32(0),
	"string":       string(""),
	"uri":          URI(""),
	"base64Binary": Base64binary{},
	"code":         Code(""),
	"id":           ID(""),
	"oid":          OID(""),
	"unsignedInt":  uint32(0),
	"positiveInt":  uint32(0),
	"markdown":     Markdown(""),
	"url":          URL(""),
	"canonical":    Canonical[any](""),
	"uuid":         UUID(""),
	"xhtml":        XHTML(""),
	"integer64":    int64(0),
}

// http://hl7.org/fhir/StructureDefinition/base64Binary
type Base64binary struct {
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	//Extension []string `bson:"extension,omitempty" json:"extension,omitempty"`
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
}

// Canonical represents a canonical URL with optional version
type Canonical[T any] string // regex: \S*

// Code represents a code from controlled string set (no leading/trailing whitespace)
type Code string // regex: [^\s]+( [^\s]+)*

// ID represents an identifier (64 character max)
type ID string // regex: [A-Za-z0-9\-\.]{1,64}

// Instant represents precise instant in time (RFC3339 format)
// Uses standard time.Time for JSON marshaling
type Instant time.Time

// Markdown represents content with markdown formatting
type Markdown string // regex: ^[\s\S]+$

// OID represents object identifier (URI format)
type OID string // regex: urn:oid:[0-2](\.(0|[1-9][0-9]*))+

// URI represents Uniform Resource Identifier
type URI string // regex: \S*

// URL represents Uniform Resource Locator
type URL string // regex: \S*

// UUID represents UUID (URI format)
type UUID string // regex: urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}

type XHTML string

// Instant marshaling (RFC3339 format)
func (i Instant) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(i).Format(time.RFC3339Nano))
}

func (i *Instant) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return fmt.Errorf("invalid Instant value: %v", data)
	}
	*i = Instant(t)
	return nil
}
