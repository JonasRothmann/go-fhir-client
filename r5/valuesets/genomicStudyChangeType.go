// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/genomicstudy-changetype
*/type GenomicStudyChangeType string

const (
	// Change that involves Deoxyribonucleic acid (DNA) sequences.
	GenomicStudyChangeTypeDnaChange GenomicStudyChangeType = "DNA"
	// Change that involves Ribonucleic Acid (RNA) sequences.
	GenomicStudyChangeTypeRnaChange GenomicStudyChangeType = "RNA"
	// Change that involves Amino Acid (AA) or protein sequences.
	GenomicStudyChangeTypeProteinAminoAcidsChange GenomicStudyChangeType = "AA"
	// Change that involves number or strcture of chromosomes.
	GenomicStudyChangeTypeChromosomalChanges GenomicStudyChangeType = "CHR"
	// Change that involves copy number variations among various genomes.
	GenomicStudyChangeTypeCopyNumberVariations GenomicStudyChangeType = "CNV"
)
