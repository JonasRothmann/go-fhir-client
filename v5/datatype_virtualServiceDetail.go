// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/VirtualServiceDetail
type VirtualServiceDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Channel Type
	ChannelType *Coding `bson:"channelType,omitempty" json:"channelType,omitempty"`
	// Contact address/number
	Address *custom.URL `bson:"address,omitempty" json:"address,omitempty"`
	// Address to see alternative connection details
	AdditionalInfo []custom.URL `bson:"additionalInfo,omitempty" json:"additionalInfo,omitempty"`
	// Maximum number of participants supported by the virtual service
	MaxParticipants *uint32 `bson:"maxParticipants,omitempty" json:"maxParticipants,omitempty"`
	// Session Key required by the virtual service
	SessionKey *string `bson:"sessionKey,omitempty" json:"sessionKey,omitempty"`
}

func (v VirtualServiceDetail) ResourceType() string {
	return "StructureDefinition"
}
