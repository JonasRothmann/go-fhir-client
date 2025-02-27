// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/knowledge-representation-level
*/type KnowledgeRepresentationLevel string

const (
	// The knowledge is expressed as narrative text, affording broad shareability, but limited computability. The knowledge is typically expressed independent of delivery method (modality), technology platform, and implementation site. Knowledge at this level is typically authored by guideline developers for a broad range of purposes including communication of policy, synthesis of evidence, and dissemination of best-practices.
	KnowledgeRepresentationLevelNarrative KnowledgeRepresentationLevel = "narrative"
	// The knowledge is expressed as organized text, often accompanied by diagrams such as flow charts and decision tables, affording broad shareability and moderate computability. The knowledge is typically expressed independent of delivery method (modality), technology platform, and implementation site. Knowledge at this level is typically authored by clinical domain experts for the purpose of communication the description of recommendations for implementation in various modalities.
	KnowledgeRepresentationLevelSemiStructured KnowledgeRepresentationLevel = "semi-structured"
	// The knowledge is expressed in a structured way that is interpretable by computer, including being coded using standard terminologies and specifications for the representation of structured content, affording broad shareability and high computability. The knowledge is typically expressed independent of delivery method (modality), technology platform, and implementation site. Knowledge at this level is typically authored by knowledge engineersto enable precise communication and validation of the recommendations.
	KnowledgeRepresentationLevelStructured KnowledgeRepresentationLevel = "structured"
	// The knowledge is expressed in a way that is coded and interpretable by CDS systems using a variety of formats, affording direct executability, but potentially limited shareability. The knowledge is typically expressed focusing on a specific delivery method (modality), technology platform, and implementation environment. Knowledge at this level is typically built and developed by implementers working in specific technology platforms, for the purpose of implementation at a specific site, though affordances such as computable mappings and configuration capabilities can broaden the usability of the knowledge artifact.
	KnowledgeRepresentationLevelExecutable KnowledgeRepresentationLevel = "executable"
)
