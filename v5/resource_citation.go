// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Citation
type CitationCitedArtifactRelatesTo struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// documentation | justification | citation | predecessor | successor | derived-from | depends-on | composed-of | part-of | amends | amended-with | appends | appended-with | cites | cited-by | comments-on | comment-in | contains | contained-in | corrects | correction-in | replaces | replaced-with | retracts | retracted-by | signs | similar-to | supports | supported-with | transforms | transformed-into | transformed-with | documents | specification-of | created-with | cite-as | reprint | reprint-of
	Type custom.Code `bson:"type" json:"type"`
	// Additional classifiers
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
	// Short label
	Label *string `bson:"label,omitempty" json:"label,omitempty"`
	// Brief description of the related artifact
	Display *string `bson:"display,omitempty" json:"display,omitempty"`
	// Bibliographic citation for the artifact
	Citation *custom.Markdown `bson:"citation,omitempty" json:"citation,omitempty"`
	// What document is being referenced
	Document *Attachment `bson:"document,omitempty" json:"document,omitempty"`
	// What artifact is being referenced
	Resource *custom.Canonical `bson:"resource,omitempty" json:"resource,omitempty"`
	// What artifact, if not a conformance resource
	ResourceReference *custom.Reference `bson:"resourceReference,omitempty" json:"resourceReference,omitempty"`
}

type CitationCitedArtifactPublicationFormPublishedIn struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Kind of container (e.g. Periodical, database, or book)
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Journal identifiers include ISSN, ISO Abbreviation and NLMuniqueID; Book identifiers include ISBN
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Name of the database or title of the book or journal
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Name of or resource describing the publisher
	Publisher *custom.Reference[Organization] `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Geographic location of the publisher
	PublisherLocation *string `bson:"publisherLocation,omitempty" json:"publisherLocation,omitempty"`
}

type CitationCitedArtifactClassification struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The kind of classifier (e.g. publication type, keyword)
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The specific classification value
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
	// Complex or externally created classification
	ArtifactAssessment []custom.Reference[ArtifactAssessment] `bson:"artifactAssessment,omitempty" json:"artifactAssessment,omitempty"`
}

type CitationCitedArtifactContributorshipSummary struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Such as author list, contributorship statement, funding statement, acknowledgements statement, or conflicts of interest statement
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The format for the display string
	Style *CodeableConcept `bson:"style,omitempty" json:"style,omitempty"`
	// Used to code the producer or rule for creating the display string
	Source *CodeableConcept `bson:"source,omitempty" json:"source,omitempty"`
	// The display string for the author list, contributor list, or contributorship statement
	Value custom.Markdown `bson:"value" json:"value"`
}

type CitationClassification struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The kind of classifier (e.g. publication type, keyword)
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The specific classification value
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
}

type CitationStatusDate struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Classification of the status
	Activity CodeableConcept `bson:"activity" json:"activity"`
	// Either occurred or expected
	Actual *bool `bson:"actual,omitempty" json:"actual,omitempty"`
	// When the status started and/or ended
	Period Period `bson:"period" json:"period"`
}

type CitationCitedArtifactVersion struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The version number or other version identifier
	Value string `bson:"value" json:"value"`
	// Citation for the main version of the cited artifact
	BaseCitation *custom.Reference[Citation] `bson:"baseCitation,omitempty" json:"baseCitation,omitempty"`
}

type CitationCitedArtifactAbstract struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The kind of abstract
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Used to express the specific language
	Language *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	// Abstract content
	Text custom.Markdown `bson:"text" json:"text"`
	// Copyright notice for the abstract
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
}

type CitationCitedArtifactContributorshipEntry struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The identity of the individual contributor
	Contributor custom.Reference[CitationCitedArtifactContributorshipEntryContributor] `bson:"contributor" json:"contributor"`
	// For citation styles that use initials
	ForenameInitials *string `bson:"forenameInitials,omitempty" json:"forenameInitials,omitempty"`
	// Organizational affiliation
	Affiliation []custom.Reference[CitationCitedArtifactContributorshipEntryAffiliation] `bson:"affiliation,omitempty" json:"affiliation,omitempty"`
	// The specific contribution
	ContributionType []CodeableConcept `bson:"contributionType,omitempty" json:"contributionType,omitempty"`
	// The role of the contributor (e.g. author, editor, reviewer, funder)
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// Contributions with accounting for time or number
	ContributionInstance []CitationCitedArtifactContributorshipEntryContributionInstance `bson:"contributionInstance,omitempty" json:"contributionInstance,omitempty"`
	// Whether the contributor is the corresponding contributor for the role
	CorrespondingContact *bool `bson:"correspondingContact,omitempty" json:"correspondingContact,omitempty"`
	// Ranked order of contribution
	RankingOrder *uint32 `bson:"rankingOrder,omitempty" json:"rankingOrder,omitempty"`
}

type CitationSummary struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Format for display of the citation summary
	Style *CodeableConcept `bson:"style,omitempty" json:"style,omitempty"`
	// The human-readable display of the citation summary
	Text custom.Markdown `bson:"text" json:"text"`
}

type CitationCitedArtifactTitle struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The kind of title
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Used to express the specific language
	Language *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	// The title of the article or artifact
	Text custom.Markdown `bson:"text" json:"text"`
}

type CitationCitedArtifactPublicationForm struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The collection the cited article or artifact is published in
	PublishedIn *CitationCitedArtifactPublicationFormPublishedIn `bson:"publishedIn,omitempty" json:"publishedIn,omitempty"`
	// Internet or Print
	CitedMedium *CodeableConcept `bson:"citedMedium,omitempty" json:"citedMedium,omitempty"`
	// Volume number of journal or other collection in which the article is published
	Volume *string `bson:"volume,omitempty" json:"volume,omitempty"`
	// Issue, part or supplement of journal or other collection in which the article is published
	Issue *string `bson:"issue,omitempty" json:"issue,omitempty"`
	// The date the article was added to the database, or the date the article was released
	ArticleDate *custom.DateTime `bson:"articleDate,omitempty" json:"articleDate,omitempty"`
	// Text representation of the date on which the issue of the cited artifact was published
	PublicationDateText *string `bson:"publicationDateText,omitempty" json:"publicationDateText,omitempty"`
	// Season in which the cited artifact was published
	PublicationDateSeason *string `bson:"publicationDateSeason,omitempty" json:"publicationDateSeason,omitempty"`
	// The date the article was last revised or updated in the database
	LastRevisionDate *custom.DateTime `bson:"lastRevisionDate,omitempty" json:"lastRevisionDate,omitempty"`
	// Language(s) in which this form of the article is published
	Language []CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	// Entry number or identifier for inclusion in a database
	AccessionNumber *string `bson:"accessionNumber,omitempty" json:"accessionNumber,omitempty"`
	// Used for full display of pagination
	PageString *string `bson:"pageString,omitempty" json:"pageString,omitempty"`
	// Used for isolated representation of first page
	FirstPage *string `bson:"firstPage,omitempty" json:"firstPage,omitempty"`
	// Used for isolated representation of last page
	LastPage *string `bson:"lastPage,omitempty" json:"lastPage,omitempty"`
	// Number of pages or screens
	PageCount *string `bson:"pageCount,omitempty" json:"pageCount,omitempty"`
	// Copyright notice for the full article or artifact
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
}

type CitationCitedArtifactPart struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The kind of component
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The specification of the component
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
	// The citation for the full article or artifact
	BaseCitation *custom.Reference[Citation] `bson:"baseCitation,omitempty" json:"baseCitation,omitempty"`
}

type CitationCitedArtifactWebLocation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code the reason for different URLs, e.g. abstract and full-text
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
	// The specific URL
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
}

type CitationCitedArtifactContributorship struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Indicates if the list includes all authors and/or contributors
	Complete *bool `bson:"complete,omitempty" json:"complete,omitempty"`
	// An individual entity named as a contributor
	Entry []CitationCitedArtifactContributorshipEntry `bson:"entry,omitempty" json:"entry,omitempty"`
	// Used to record a display of the author/contributor list without separate data element for each list member
	Summary []CitationCitedArtifactContributorshipSummary `bson:"summary,omitempty" json:"summary,omitempty"`
}

type CitationCitedArtifactContributorshipEntryContributionInstance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The specific contribution
	Type CodeableConcept `bson:"type" json:"type"`
	// The time that the contribution was made
	Time *custom.DateTime `bson:"time,omitempty" json:"time,omitempty"`
}

type Citation struct {
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
	// Canonical identifier for this citation record, represented as a globally unique URI
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Identifier for the citation record itself
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the citation record
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this citation record (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this citation record (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// The publisher of the citation record, not the publisher of the article or artifact being cited
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher of the citation record
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the citation
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the citation record content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for citation record (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this citation is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions for the citation record, not for the cited artifact
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s) for the ciation record, not for the cited artifact
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the citation record was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the citation record was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// When the citation record is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// Who authored the citation record
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the citation record
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the citation record
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the citation record
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// A human-readable display of key concepts to represent the citation
	Summary []CitationSummary `bson:"summary,omitempty" json:"summary,omitempty"`
	// The assignment to an organizing scheme
	Classification []CitationClassification `bson:"classification,omitempty" json:"classification,omitempty"`
	// Used for general notes and annotations not coded elsewhere
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// The status of the citation record
	CurrentState []CodeableConcept `bson:"currentState,omitempty" json:"currentState,omitempty"`
	// An effective date or period for a status of the citation record
	StatusDate []CitationStatusDate `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	// Artifact related to the citation record
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// The article or artifact being described
	CitedArtifact *CitationCitedArtifact `bson:"citedArtifact,omitempty" json:"citedArtifact,omitempty"`
}

type CitationCitedArtifact struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Unique identifier. May include DOI, PMID, PMCID, etc
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Identifier not unique to the cited artifact. May include trial registry identifiers
	RelatedIdentifier []Identifier `bson:"relatedIdentifier,omitempty" json:"relatedIdentifier,omitempty"`
	// When the cited artifact was accessed
	DateAccessed *custom.DateTime `bson:"dateAccessed,omitempty" json:"dateAccessed,omitempty"`
	// The defined version of the cited artifact
	Version *CitationCitedArtifactVersion `bson:"version,omitempty" json:"version,omitempty"`
	// The status of the cited artifact
	CurrentState []CodeableConcept `bson:"currentState,omitempty" json:"currentState,omitempty"`
	// An effective date or period for a status of the cited artifact
	StatusDate []CitationCitedArtifactStatusDate `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	// The title details of the article or artifact
	Title []CitationCitedArtifactTitle `bson:"title,omitempty" json:"title,omitempty"`
	// Summary of the article or artifact
	Abstract []CitationCitedArtifactAbstract `bson:"abstract,omitempty" json:"abstract,omitempty"`
	// The component of the article or artifact
	Part *CitationCitedArtifactPart `bson:"part,omitempty" json:"part,omitempty"`
	// The artifact related to the cited artifact
	RelatesTo []CitationCitedArtifactRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	// If multiple, used to represent alternative forms of the article that are not separate citations
	PublicationForm []CitationCitedArtifactPublicationForm `bson:"publicationForm,omitempty" json:"publicationForm,omitempty"`
	// Used for any URL for the article or artifact cited
	WebLocation []CitationCitedArtifactWebLocation `bson:"webLocation,omitempty" json:"webLocation,omitempty"`
	// The assignment to an organizing scheme
	Classification []CitationCitedArtifactClassification `bson:"classification,omitempty" json:"classification,omitempty"`
	// Attribution of authors and other contributors
	Contributorship *CitationCitedArtifactContributorship `bson:"contributorship,omitempty" json:"contributorship,omitempty"`
	// Any additional information or content for the article or artifact
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type CitationCitedArtifactStatusDate struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Classification of the status
	Activity CodeableConcept `bson:"activity" json:"activity"`
	// Either occurred or expected
	Actual *bool `bson:"actual,omitempty" json:"actual,omitempty"`
	// When the status started and/or ended
	Period Period `bson:"period" json:"period"`
}

type CitationCitedArtifactContributorshipEntryContributor interface {
	gofhirclient.Resource

	Is_CitationCitedArtifactContributorshipEntryContributor()
}

func (p Practitioner) Is_CitationCitedArtifactContributorshipEntryContributor() {}
func (o Organization) Is_CitationCitedArtifactContributorshipEntryContributor() {}

type CitationCitedArtifactContributorshipEntryAffiliation interface {
	gofhirclient.Resource

	Is_CitationCitedArtifactContributorshipEntryAffiliation()
}

func (o Organization) Is_CitationCitedArtifactContributorshipEntryAffiliation()     {}
func (p PractitionerRole) Is_CitationCitedArtifactContributorshipEntryAffiliation() {}

func (c Citation) ResourceType() string {
	return "StructureDefinition"
}
