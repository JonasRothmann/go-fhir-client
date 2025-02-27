// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/genomicstudy-type
*/type GenomicStudyType string

const (
	// Identification of multiple different processed mRNA transcripts from the same DNA template
	GenomicStudyTypeAlternativeSplicingDetection GenomicStudyType = "alt-splc"
	// Analysis of the spacial organization of chromatin within a cell
	GenomicStudyTypeChromatin GenomicStudyType = "chromatin"
	// Detection of a change in the number of copies of a defined region of genomic DNA sequence resulting in structural variation when compared to the reference sequence
	GenomicStudyTypeCnv GenomicStudyType = "cnv"
	// Detection of biochemical modifications covalently bound to the N-terminal tail of a histone protein. These modifications may alter chromatin compaction and gene expression
	GenomicStudyTypeEpigeneticAlterationsHistoneModifications GenomicStudyType = "epi-alt-hist"
	// Detection of the presence of an additional methyl group on a DNA nucleobase, which may alter gene transcription
	GenomicStudyTypeEpigeneticAlterationsDnaMethylation GenomicStudyType = "epi-alt-dna"
	// Determining if a variant identified in an individual is present in other family members
	GenomicStudyTypeFamilialVariantSegregation GenomicStudyType = "fam-var-segr"
	// Detection of sequence variants which may alter gene expression or gene product function when compared to the reference sequence
	GenomicStudyTypeFunctionalVariationDetection GenomicStudyType = "func-var"
	// Measurement and characterization of activity from all gene products
	GenomicStudyTypeGeneExpressionProfiling GenomicStudyType = "gene-expression"
	// Detection of biochemical modifications covalently bound to the amino acid monomers of a processed protein
	GenomicStudyTypePostTranslationalModificationIdentification GenomicStudyType = "post-trans-mod"
	// Determination of which nucleotide is base present at a known variable location of the genomic sequence
	GenomicStudyTypeSnp GenomicStudyType = "snp"
	// Quantification of the number of sequential microsatellite units in a repetitive sequence region
	GenomicStudyTypeStr GenomicStudyType = "str"
	// Detection of deletions, insertions, or rearrangements of DNA segments compared to the reference sequence
	GenomicStudyTypeStructuralVariationDetection GenomicStudyType = "struc-var"
)
