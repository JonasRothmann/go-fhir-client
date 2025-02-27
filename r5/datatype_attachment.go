// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Attachment
type Attachment struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Mime type of the content, with charset etc.
	ContentType *custom.Code `json:"contentType,omitempty"`
	// Human language of the content (BCP-47)
	Language *custom.Code `json:"language,omitempty"`
	// Data inline, base64ed
	Data *custom.Base64binary `json:"data,omitempty"`
	// Uri where the data can be found
	Url *custom.URL `json:"url,omitempty"`
	// Number of bytes of content (if url provided)
	Size *int64 `json:"size,omitempty"`
	// Hash of the data (sha-1, base64ed)
	Hash *custom.Base64binary `json:"hash,omitempty"`
	// Label to display in place of the data
	Title *string `json:"title,omitempty"`
	// Date attachment was first created
	Creation *custom.DateTime `json:"creation,omitempty"`
	// Height of the image in pixels (photo/video)
	Height *uint32 `json:"height,omitempty"`
	// Width of the image in pixels (photo/video)
	Width *uint32 `json:"width,omitempty"`
	// Number of frames if > 1 (photo)
	Frames *uint32 `json:"frames,omitempty"`
	// Length in seconds (audio / video)
	Duration *json.Number `json:"duration,omitempty"`
	// Number of printed pages
	Pages *uint32 `json:"pages,omitempty"`
}

type OtherAttachment Attachment

func (a Attachment) ResourceType() string {
	return "Attachment"
}

func (a Attachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAttachment
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherAttachment: OtherAttachment(a), ResourceType: a.ResourceType()})
}

func UnmarshalAttachment(b []byte) (res Attachment, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
