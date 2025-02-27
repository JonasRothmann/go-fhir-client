// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/VirtualServiceDetail
type VirtualServiceDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Channel Type
	ChannelType *Coding `json:"channelType,omitempty"`
	// Contact address/number
	AddressUrl *custom.URL `json:"addressUrl,omitempty"`
	// Contact address/number
	AddressString *string `json:"addressString,omitempty"`
	// Contact address/number
	AddressContactPoint *ContactPoint `json:"addressContactPoint,omitempty"`
	// Contact address/number
	AddressExtendedContactDetail *ExtendedContactDetail `json:"addressExtendedContactDetail,omitempty"`
	// Address to see alternative connection details
	AdditionalInfo []custom.URL `json:"additionalInfo,omitempty"`
	// Maximum number of participants supported by the virtual service
	MaxParticipants *uint32 `json:"maxParticipants,omitempty"`
	// Session Key required by the virtual service
	SessionKey *string `json:"sessionKey,omitempty"`
}

type OtherVirtualServiceDetail VirtualServiceDetail

func (v VirtualServiceDetail) ResourceType() string {
	return "VirtualServiceDetail"
}

func (v VirtualServiceDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVirtualServiceDetail
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherVirtualServiceDetail: OtherVirtualServiceDetail(v), ResourceType: v.ResourceType()})
}

func UnmarshalVirtualServiceDetail(b []byte) (res VirtualServiceDetail, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
