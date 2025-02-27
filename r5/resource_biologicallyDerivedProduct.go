// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProduct struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// organ | tissue | fluid | cells | biologicalAgent
	ProductCategory *Coding `json:"productCategory,omitempty"`
	// A code that identifies the kind of this biologically derived product
	ProductCode *CodeableConcept `json:"productCode,omitempty"`
	// The parent biologically-derived product
	Parent []custom.Reference[BiologicallyDerivedProduct] `json:"parent,omitempty"`
	// Request to obtain and/or infuse this product
	Request []custom.Reference[ServiceRequest] `json:"request,omitempty"`
	// Instance identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// An identifier that supports traceability to the event during which material in this product from one or more biological entities was obtained or pooled
	BiologicalSourceEvent *Identifier `json:"biologicalSourceEvent,omitempty"`
	// Processing facilities responsible for the labeling and distribution of this biologically derived product
	ProcessingFacility []custom.Reference[Organization] `json:"processingFacility,omitempty"`
	// A unique identifier for an aliquot of a product
	Division *string `json:"division,omitempty"`
	// available | unavailable
	ProductStatus *Coding `json:"productStatus,omitempty"`
	// Date, and where relevant time, of expiration
	ExpirationDate *custom.DateTime `json:"expirationDate,omitempty"`
	// How this product was collected
	Collection *BiologicallyDerivedProductCollection `json:"collection,omitempty"`
	// Product storage temperature requirements
	StorageTempRequirements *Range `json:"storageTempRequirements,omitempty"`
	// A property that is specific to this BiologicallyDerviedProduct instance
	Property []BiologicallyDerivedProductProperty `json:"property,omitempty"`
}

type BiologicallyDerivedProductCollection struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Individual performing collection
	Collector *custom.Reference[BiologicallyDerivedProductCollectionCollector] `json:"collector,omitempty"`
	// The patient who underwent the medical procedure to collect the product or the organization that facilitated the collection
	Source *custom.Reference[BiologicallyDerivedProductCollectionSource] `json:"source,omitempty"`
	// Time of product collection
	CollectedDateTime *custom.DateTime `json:"collectedDateTime,omitempty"`
	// Time of product collection
	CollectedPeriod *Period `json:"collectedPeriod,omitempty"`
}

type BiologicallyDerivedProductProperty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code that specifies the property
	Type CodeableConcept `json:"type"`
	// Property values
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Property values
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Property values
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Property values
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// Property values
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Property values
	ValueRange *Range `json:"valueRange,omitempty"`
	// Property values
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// Property values
	ValueString *string `json:"valueString,omitempty"`
	// Property values
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
}

type OtherBiologicallyDerivedProduct BiologicallyDerivedProduct

type BiologicallyDerivedProductCollectionCollector interface {
	gofhirclient.Resource

	Is_BiologicallyDerivedProductCollectionCollector()
}

func (p Practitioner) Is_BiologicallyDerivedProductCollectionCollector()     {}
func (p PractitionerRole) Is_BiologicallyDerivedProductCollectionCollector() {}

type BiologicallyDerivedProductCollectionSource interface {
	gofhirclient.Resource

	Is_BiologicallyDerivedProductCollectionSource()
}

func (p Patient) Is_BiologicallyDerivedProductCollectionSource()      {}
func (o Organization) Is_BiologicallyDerivedProductCollectionSource() {}

func (b BiologicallyDerivedProduct) ResourceType() string {
	return "BiologicallyDerivedProduct"
}

func (b BiologicallyDerivedProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBiologicallyDerivedProduct
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherBiologicallyDerivedProduct: OtherBiologicallyDerivedProduct(b), ResourceType: b.ResourceType()})
}

func UnmarshalBiologicallyDerivedProduct(b []byte) (res BiologicallyDerivedProduct, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
