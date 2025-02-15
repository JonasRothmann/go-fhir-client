// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Attachment
type Attachment struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Mime type of the content, with charset etc.
	ContentType *custom.Code `bson:"contentType,omitempty" json:"contentType,omitempty"`
	// Human language of the content (BCP-47)
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Data inline, base64ed
	Data *custom.Base64binary `bson:"data,omitempty" json:"data,omitempty"`
	// Uri where the data can be found
	Url *custom.URL `bson:"url,omitempty" json:"url,omitempty"`
	// Number of bytes of content (if url provided)
	Size *int64 `bson:"size,omitempty" json:"size,omitempty"`
	// Hash of the data (sha-1, base64ed)
	Hash *custom.Base64binary `bson:"hash,omitempty" json:"hash,omitempty"`
	// Label to display in place of the data
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Date attachment was first created
	Creation *custom.DateTime `bson:"creation,omitempty" json:"creation,omitempty"`
	// Height of the image in pixels (photo/video)
	Height *uint32 `bson:"height,omitempty" json:"height,omitempty"`
	// Width of the image in pixels (photo/video)
	Width *uint32 `bson:"width,omitempty" json:"width,omitempty"`
	// Number of frames if > 1 (photo)
	Frames *uint32 `bson:"frames,omitempty" json:"frames,omitempty"`
	// Length in seconds (audio / video)
	Duration *json.Number `bson:"duration,omitempty" json:"duration,omitempty"`
	// Number of printed pages
	Pages *uint32 `bson:"pages,omitempty" json:"pages,omitempty"`
}

func (a Attachment) ResourceType() string {
	return "StructureDefinition"
}
