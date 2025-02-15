package custom

type Identifier struct {
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	//Extension []datatypes.Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	Use *Code `bson:"use,omitempty" json:"use,omitempty"`
	//Type     *CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	System *URI    `bson:"system,omitempty" json:"system,omitempty"`
	Value  *string `bson:"value,omitempty" json:"value,omitempty"`
	//Period   *Period                         `bson:"period,omitempty" json:"period,omitempty"`
	//Assigner *Reference[IdentifierReference] `bson:"assigner,omitempty" json:"assigner,omitempty"`
}
