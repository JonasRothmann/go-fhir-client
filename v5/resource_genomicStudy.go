// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/GenomicStudy
type GenomicStudyAnalysisOutput struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// File containing output data
	File *custom.Reference[DocumentReference] `bson:"file,omitempty" json:"file,omitempty"`
	// Type of output data (e.g., VCF, MAF, or BAM)
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}

type GenomicStudyAnalysisPerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The organization, healthcare professional, or others who participated in performing this analysis
	Actor *custom.Reference[GenomicStudyAnalysisPerformerActor] `bson:"actor,omitempty" json:"actor,omitempty"`
	// Role of the actor for this analysis
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type GenomicStudyAnalysisDevice struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Device used for the analysis
	Device *custom.Reference[Device] `bson:"device,omitempty" json:"device,omitempty"`
	// Specific function for the device used for the analysis
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
}

type GenomicStudy struct {
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
	// Identifiers for this genomic study
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// registered | available | cancelled | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// The type of the study (e.g., Familial variant segregation, Functional variation detection, or Gene expression profiling)
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The primary subject of the genomic study
	Subject custom.Reference[GenomicStudySubject] `bson:"subject" json:"subject"`
	// The healthcare event with which this genomics study is associated
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When the genomic study was started
	StartDate *custom.DateTime `bson:"startDate,omitempty" json:"startDate,omitempty"`
	// Event resources that the genomic study is based on
	BasedOn []custom.Reference[GenomicStudyBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Healthcare professional who requested or referred the genomic study
	Referrer *custom.Reference[GenomicStudyReferrer] `bson:"referrer,omitempty" json:"referrer,omitempty"`
	// Healthcare professionals who interpreted the genomic study
	Interpreter []custom.Reference[GenomicStudyInterpreter] `bson:"interpreter,omitempty" json:"interpreter,omitempty"`
	// Why the genomic study was performed
	Reason []custom.CodeableReference[GenomicStudyReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// The defined protocol that describes the study
	InstantiatesCanonical *custom.Canonical[PlanDefinition] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// The URL pointing to an externally maintained protocol that describes the study
	InstantiatesUri *custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// Comments related to the genomic study
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Description of the genomic study
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Genomic Analysis Event
	Analysis []GenomicStudyAnalysis `bson:"analysis,omitempty" json:"analysis,omitempty"`
}

type GenomicStudyAnalysis struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Identifiers for the analysis event
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Type of the methods used in the analysis (e.g., FISH, Karyotyping, MSI)
	MethodType []CodeableConcept `bson:"methodType,omitempty" json:"methodType,omitempty"`
	// Type of the genomic changes studied in the analysis (e.g., DNA, RNA, or AA change)
	ChangeType []CodeableConcept `bson:"changeType,omitempty" json:"changeType,omitempty"`
	// Genome build that is used in this analysis
	GenomeBuild *CodeableConcept `bson:"genomeBuild,omitempty" json:"genomeBuild,omitempty"`
	// The defined protocol that describes the analysis
	InstantiatesCanonical *custom.Canonical[GenomicStudyAnalysisInstantiatesCanonical] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// The URL pointing to an externally maintained protocol that describes the analysis
	InstantiatesUri *custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// Name of the analysis event (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// What the genomic analysis is about, when it is not about the subject of record
	Focus []custom.Reference[Resource] `bson:"focus,omitempty" json:"focus,omitempty"`
	// The specimen used in the analysis event
	Specimen []custom.Reference[Specimen] `bson:"specimen,omitempty" json:"specimen,omitempty"`
	// The date of the analysis event
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Any notes capture with the analysis event
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// The protocol that was performed for the analysis event
	ProtocolPerformed *custom.Reference[GenomicStudyAnalysisProtocolPerformed] `bson:"protocolPerformed,omitempty" json:"protocolPerformed,omitempty"`
	// The genomic regions to be studied in the analysis (BED file)
	RegionsStudied []custom.Reference[GenomicStudyAnalysisRegionsStudied] `bson:"regionsStudied,omitempty" json:"regionsStudied,omitempty"`
	// Genomic regions actually called in the analysis event (BED file)
	RegionsCalled []custom.Reference[GenomicStudyAnalysisRegionsCalled] `bson:"regionsCalled,omitempty" json:"regionsCalled,omitempty"`
	// Inputs for the analysis event
	Input []GenomicStudyAnalysisInput `bson:"input,omitempty" json:"input,omitempty"`
	// Outputs for the analysis event
	Output []GenomicStudyAnalysisOutput `bson:"output,omitempty" json:"output,omitempty"`
	// Performer for the analysis event
	Performer []GenomicStudyAnalysisPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// Devices used for the analysis (e.g., instruments, software), with settings and parameters
	Device []GenomicStudyAnalysisDevice `bson:"device,omitempty" json:"device,omitempty"`
}

type GenomicStudyAnalysisInput struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// File containing input data
	File *custom.Reference[DocumentReference] `bson:"file,omitempty" json:"file,omitempty"`
	// Type of input data (e.g., BAM, CRAM, or FASTA)
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The analysis event or other GenomicStudy that generated this input file
	GeneratedBy *Identifier `bson:"generatedBy,omitempty" json:"generatedBy,omitempty"`
}

type GenomicStudySubject interface {
	gofhirclient.Resource

	Is_GenomicStudySubject()
}

func (p Patient) Is_GenomicStudySubject()                    {}
func (g Group) Is_GenomicStudySubject()                      {}
func (s Substance) Is_GenomicStudySubject()                  {}
func (b BiologicallyDerivedProduct) Is_GenomicStudySubject() {}
func (n NutritionProduct) Is_GenomicStudySubject()           {}

type GenomicStudyBasedOn interface {
	gofhirclient.Resource

	Is_GenomicStudyBasedOn()
}

func (s ServiceRequest) Is_GenomicStudyBasedOn() {}
func (t Task) Is_GenomicStudyBasedOn()           {}

type GenomicStudyReferrer interface {
	gofhirclient.Resource

	Is_GenomicStudyReferrer()
}

func (p Practitioner) Is_GenomicStudyReferrer()     {}
func (p PractitionerRole) Is_GenomicStudyReferrer() {}

type GenomicStudyInterpreter interface {
	gofhirclient.Resource

	Is_GenomicStudyInterpreter()
}

func (p Practitioner) Is_GenomicStudyInterpreter()     {}
func (p PractitionerRole) Is_GenomicStudyInterpreter() {}

type GenomicStudyReason interface {
	gofhirclient.Resource

	Is_GenomicStudyReason()
}

func (c Condition) Is_GenomicStudyReason()   {}
func (o Observation) Is_GenomicStudyReason() {}

type GenomicStudyAnalysisInstantiatesCanonical interface {
	gofhirclient.Resource

	Is_GenomicStudyAnalysisInstantiatesCanonical()
}

func (p PlanDefinition) Is_GenomicStudyAnalysisInstantiatesCanonical()     {}
func (a ActivityDefinition) Is_GenomicStudyAnalysisInstantiatesCanonical() {}

type GenomicStudyAnalysisProtocolPerformed interface {
	gofhirclient.Resource

	Is_GenomicStudyAnalysisProtocolPerformed()
}

func (p Procedure) Is_GenomicStudyAnalysisProtocolPerformed() {}
func (t Task) Is_GenomicStudyAnalysisProtocolPerformed()      {}

type GenomicStudyAnalysisRegionsStudied interface {
	gofhirclient.Resource

	Is_GenomicStudyAnalysisRegionsStudied()
}

func (d DocumentReference) Is_GenomicStudyAnalysisRegionsStudied() {}
func (o Observation) Is_GenomicStudyAnalysisRegionsStudied()       {}

type GenomicStudyAnalysisRegionsCalled interface {
	gofhirclient.Resource

	Is_GenomicStudyAnalysisRegionsCalled()
}

func (d DocumentReference) Is_GenomicStudyAnalysisRegionsCalled() {}
func (o Observation) Is_GenomicStudyAnalysisRegionsCalled()       {}

type GenomicStudyAnalysisPerformerActor interface {
	gofhirclient.Resource

	Is_GenomicStudyAnalysisPerformerActor()
}

func (p Practitioner) Is_GenomicStudyAnalysisPerformerActor()     {}
func (p PractitionerRole) Is_GenomicStudyAnalysisPerformerActor() {}
func (o Organization) Is_GenomicStudyAnalysisPerformerActor()     {}
func (d Device) Is_GenomicStudyAnalysisPerformerActor()           {}

func (g GenomicStudy) ResourceType() string {
	return "StructureDefinition"
}
