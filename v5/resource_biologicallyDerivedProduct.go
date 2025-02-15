// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductCollection struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Individual performing collection
	Collector *custom.Reference[BiologicallyDerivedProductCollectionCollector] `bson:"collector,omitempty" json:"collector,omitempty"`
	// The patient who underwent the medical procedure to collect the product or the organization that facilitated the collection
	Source *custom.Reference[BiologicallyDerivedProductCollectionSource] `bson:"source,omitempty" json:"source,omitempty"`
	// Time of product collection
	Collected *custom.DateTime `bson:"collected,omitempty" json:"collected,omitempty"`
}

type BiologicallyDerivedProductProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code that specifies the property
	Type CodeableConcept `bson:"type" json:"type"`
	// Property values
	Value bool `bson:"value" json:"value"`
}

type BiologicallyDerivedProduct struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `bson:"contained,omitempty" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// organ | tissue | fluid | cells | biologicalAgent
	ProductCategory *Coding `bson:"productCategory,omitempty" json:"productCategory,omitempty"`
	// A code that identifies the kind of this biologically derived product
	ProductCode *CodeableConcept `bson:"productCode,omitempty" json:"productCode,omitempty"`
	// The parent biologically-derived product
	Parent []custom.Reference[BiologicallyDerivedProduct] `bson:"parent,omitempty" json:"parent,omitempty"`
	// Request to obtain and/or infuse this product
	Request []custom.Reference[ServiceRequest] `bson:"request,omitempty" json:"request,omitempty"`
	// Instance identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// An identifier that supports traceability to the event during which material in this product from one or more biological entities was obtained or pooled
	BiologicalSourceEvent *Identifier `bson:"biologicalSourceEvent,omitempty" json:"biologicalSourceEvent,omitempty"`
	// Processing facilities responsible for the labeling and distribution of this biologically derived product
	ProcessingFacility []custom.Reference[Organization] `bson:"processingFacility,omitempty" json:"processingFacility,omitempty"`
	// A unique identifier for an aliquot of a product
	Division *string `bson:"division,omitempty" json:"division,omitempty"`
	// available | unavailable
	ProductStatus *Coding `bson:"productStatus,omitempty" json:"productStatus,omitempty"`
	// Date, and where relevant time, of expiration
	ExpirationDate *custom.DateTime `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	// How this product was collected
	Collection *BiologicallyDerivedProductCollection `bson:"collection,omitempty" json:"collection,omitempty"`
	// Product storage temperature requirements
	StorageTempRequirements *Range `bson:"storageTempRequirements,omitempty" json:"storageTempRequirements,omitempty"`
	// A property that is specific to this BiologicallyDerviedProduct instance
	Property []BiologicallyDerivedProductProperty `bson:"property,omitempty" json:"property,omitempty"`
}

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
	return "StructureDefinition"
}
