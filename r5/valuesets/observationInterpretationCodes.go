// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/v3-ObservationInterpretation
*/type ObservationInterpretationCodes string

const (
	// Codes that specify interpretation of genetic analysis, such as "positive", "negative", "carrier", "responsive", etc.
	ObservationInterpretationCodesGeneticObservationInterpretation ObservationInterpretationCodes = "_GeneticObservationInterpretation"
	// Interpretations of change of quantity and/or severity. At most one of B or W and one of U or D allowed.
	ObservationInterpretationCodesObservationInterpretationChange ObservationInterpretationCodes = "_ObservationInterpretationChange"
	// Technical exceptions resulting in the inability to provide an interpretation. At most one allowed. Does not imply normality or severity.
	ObservationInterpretationCodesObservationInterpretationExceptions ObservationInterpretationCodes = "_ObservationInterpretationExceptions"
	// Interpretation of normality or degree of abnormality (including critical or "alert" level). Concepts in this category are mutually exclusive, i.e., at most one is allowed.
	ObservationInterpretationCodesObservationInterpretationNormality ObservationInterpretationCodes = "_ObservationInterpretationNormality"
	// Interpretations of anti-microbial susceptibility testing results (microbiology). At most one allowed.
	ObservationInterpretationCodesObservationInterpretationSusceptibility ObservationInterpretationCodes = "_ObservationInterpretationSusceptibility"
	/*
	   The observation/test result is interpreted as being outside the inclusion range for a particular protocol within which the result is being reported.

	   **Example:** A positive result on a Hepatitis screening test.

	   *Open Issue:* EX, HX, LX: These three concepts do not seem to meet a clear need in the vocabulary, and their use in observation interpretation appears likely to be covered by other existing concepts (e.g., A, H, L). The only apparent significant difference is their reference to use in protocols for exclusion of study subjects. These concepts/codes were proposed by RCRIM for use in the CTLaboratory message. They were submitted and approved in the November 2005 Harmonization cycle in proposal "030103C\_VOCAB\_RCRIM\_l\_quade\_RCRIM Obs Interp\_20051028154455". However, this proposal was not fully implemented in the vocabulary. The proposal recommended creation of the x\_ClinicalResearchExclusion domain in ObservationInterpretation with a value set including those three concepts/codes, but there is no subdomain of that name or equivalent with a binding to either of the value sets that contain these concepts/codes. Members of the OO WG have recently attempted to contact members of RCRIM regarding these concepts, both by email and at the recent WGM in Atlanta, without response. It is felt by OO that the best course of action to take at this time is to add this comprehensive Open Issue rather than deprecate these three concepts at this time, until further discussion is held.
	*/
	ObservationInterpretationCodesOutsideThreshold ObservationInterpretationCodes = "EX"
	/*
	   Hold for Medical Review

	   *Usage Note:* This code is not intended for use in V3 artifacts. It is included in the code system to maintain alignment with the V2 Table 0078 "Interpretation Codes."
	*/
	ObservationInterpretationCodesHoldForMedicalReview ObservationInterpretationCodes = "HM"
	// Interpretations of the presence or absence of a component / analyte or organism in a test or of a sign in a clinical observation. In keeping with laboratory data processing practice, these concepts provide a categorical interpretation of the "meaning" of the quantitative value for the same observation.
	ObservationInterpretationCodesObservationInterpretationDetection ObservationInterpretationCodes = "ObservationInterpretationDetection"
	// Interpretation of the observed result taking into account additional information (contraindicators) about the patient's situation. Concepts in this category are mutually exclusive, i.e., at most one is allowed.
	ObservationInterpretationCodesObservationInterpretationExpectation ObservationInterpretationCodes = "ObservationInterpretationExpectation"
	/*
	   Interpretation qualifiers in separate OBX segments

	   *Usage Note:* This code is not intended for use in V3 artifacts. It is included in the code system to maintain alignment with the V2 Table 0078 "Interpretation Codes."
	*/
	ObservationInterpretationCodesInterpretationQualifiersInSeparateObxSegments ObservationInterpretationCodes = "OBX"
	// Interpretations of the presence and level of reactivity of the specified component / analyte with the reagent in the performed laboratory test.
	ObservationInterpretationCodesReactivityObservationInterpretation ObservationInterpretationCodes = "ReactivityObservationInterpretation"
	// The patient is considered as carrier based on the testing results. A carrier is an individual who carries an altered form of a gene which can lead to having a child or offspring in future generations with a genetic disorder.
	ObservationInterpretationCodesCarrier ObservationInterpretationCodes = "CAR"
	/*
	   The patient is considered as carrier based on the testing results. A carrier is an individual who carries an altered form of a gene which can lead to having a child or offspring in future generations with a genetic disorder.

	   ***Deprecation Comment:*** This code is currently the same string as the print name for this concept and is inconsistent with the conventions being used for the other codes in the coding system, as it is a full word with initial capitalization, rather than an all upper case mnemonic. The recommendation from OO is to deprecate the code "Carrier" and to add "CAR" as the new active code representation for this concept.
	*/
	ObservationInterpretationCodesCarrier2 ObservationInterpretationCodes = "Carrier"
	/*
	   The current result or observation value has improved compared to the previous result or observation value (the change is significant as defined in the respective test procedure).

	   \[Note: This can be applied to quantitative or qualitative observations.\]
	*/
	ObservationInterpretationCodesBetter ObservationInterpretationCodes = "B"
	// The current result has decreased from the previous result for a quantitative observation (the change is significant as defined in the respective test procedure).
	ObservationInterpretationCodesSignificantChangeDown ObservationInterpretationCodes = "D"
	// The current result has increased from the previous result for a quantitative observation (the change is significant as defined in the respective test procedure).
	ObservationInterpretationCodesSignificantChangeUp ObservationInterpretationCodes = "U"
	/*
	   The current result or observation value has degraded compared to the previous result or observation value (the change is significant as defined in the respective test procedure).

	   \[Note: This can be applied to quantitative or qualitative observations.\]
	*/
	ObservationInterpretationCodesWorse ObservationInterpretationCodes = "W"
	/*
	   The result is below the minimum detection limit (the test procedure or equipment is the limiting factor).

	   Synonyms: Below analytical limit, low off scale.
	*/
	ObservationInterpretationCodesOffScaleLow ObservationInterpretationCodes = "<"
	/*
	   The result is above the maximum quantifiable limit (the test procedure or equipment is the limiting factor).

	   Synonyms: Above analytical limit, high off scale.
	*/
	ObservationInterpretationCodesOffScaleHigh ObservationInterpretationCodes = ">"
	/*
	   A valid result cannot be obtained for the specified component / analyte due to the presence of anti-complementary substances in the sample.

	   ***Deprecation Comment:*** This code is being deprecated to match the status in V2 Table 0078 "Interpretation Codes.
	*/
	ObservationInterpretationCodesAntiComplementarySubstancesPresent ObservationInterpretationCodes = "AC"
	/*
	   There is insufficient evidence that the species in question is a good target for therapy with the drug. A categorical interpretation is not possible.

	   \[Note: A MIC with "IE" and/or a comment may be reported (without an accompanying S, I or R-categorization).\]
	*/
	ObservationInterpretationCodesInsufficientEvidence ObservationInterpretationCodes = "IE"
	/*
	   A result cannot be considered valid for the specified component / analyte or organism due to failure in the quality control testing component.

	   ***Deprecation Comment:*** This code is being deprecated to match the status in V2 Table 0078 "Interpretation Codes.
	*/
	ObservationInterpretationCodesQualityControlFailure ObservationInterpretationCodes = "QCF"
	/*
	   A valid result cannot be obtained for the specified organism or cell line due to the presence of cytotoxic substances in the sample or culture.

	   ***Deprecation Comment:*** This code is being deprecated to match the status in V2 Table 0078 "Interpretation Codes.
	*/
	ObservationInterpretationCodesCytotoxicSubstancePresent ObservationInterpretationCodes = "TOX"
	/*
	   The result or observation value is outside the reference range or expected norm (as defined for the respective test procedure).

	   \[Note: Typically applies to non-numeric results.\]
	*/
	ObservationInterpretationCodesAbnormal ObservationInterpretationCodes = "A"
	/*
	   The result or observation value is within the reference range or expected norm (as defined for the respective test procedure).

	   \[Note: Applies to numeric or non-numeric results.\]
	*/
	ObservationInterpretationCodesNormal ObservationInterpretationCodes = "N"
	/*
	   Bacterial strain inhibited in vitro by a concentration of an antimicrobial agent that is associated with uncertain therapeutic effect. Reference: CLSI (http://www.clsi.org/Content/NavigationMenu/Resources/HarmonizedTerminologyDatabase/Harmonized\_Terminolo.htm) Projects: ISO 20776-1, ISO 20776-2

	   \[Note 1: Bacterial strains are categorized as intermediate by applying the appropriate breakpoints in a defined phenotypic test system.\]

	   \[Note 2: This class of susceptibility implies that an infection due to the isolate can be appropriately treated in body sites where the drugs are physiologically concentrated or when a high dosage of drug can be used.\]

	   \[Note 3: This class also indicates a "buffer zone," to prevent small, uncontrolled, technical factors from causing major discrepancies in interpretations.\]

	   \[Note 4: These breakpoints can be altered due to changes in circumstances (e.g., changes in commonly used drug dosages, emergence of new resistance mechanisms).\]
	*/
	ObservationInterpretationCodesIntermediate ObservationInterpretationCodes = "I"
	/*
	   The patient is considered as carrier based on the testing results. A carrier is an individual who carries an altered form of a gene which can lead to having a child or offspring in future generations with a genetic disorder.

	   ***Deprecation Comment:*** This antimicrobial susceptibility test interpretation concept is recommended by OO to be deprecated as it is no longer recommended for use in susceptibility testing by CLSI (reference CLSI document M100-S22; Vol. 32 No.3; CLSI Performance Standards for Antimicrobial Susceptibility Testing; Twenty-Second Informational Supplement. Jan 2012).
	*/
	ObservationInterpretationCodesModeratelySusceptible ObservationInterpretationCodes = "MS"
	// Use when not enough clinical trial data published by the Clinical and Laboratory Standards Institutes (CLSI) is available to establish the breakpoints for susceptible / intermediate and resistant.
	ObservationInterpretationCodesNoClsiDefinedBreakpoint ObservationInterpretationCodes = "NCL"
	/*
	   A category used for isolates for which only a susceptible interpretive criterion has been designated because of the absence or rare occurrence of resistant strains. Isolates that have MICs above or zone diameters below the value indicated for the susceptible breakpoint should be reported as non-susceptible.

	   NOTE 1: An isolate that is interpreted as non-susceptible does not necessarily mean that the isolate has a resistance mechanism. It is possible that isolates with MICs above the susceptible breakpoint that lack resistance mechanisms may be encountered within the wild-type distribution subsequent to the time the susceptible-only breakpoint is set.

	   NOTE 2: For strains yielding results in the "nonsusceptible" category, organism identification and antimicrobial susceptibility test results should be confirmed.

	   Synonym: decreased susceptibility.
	*/
	ObservationInterpretationCodesNonSusceptible ObservationInterpretationCodes = "NS"
	/*
	   Bacterial strain inhibited in vitro by a concentration of an antimicrobial agent that is associated with a high likelihood of therapeutic failure. Reference: CLSI (http://www.clsi.org/Content/NavigationMenu/Resources/HarmonizedTerminologyDatabase/Harmonized\_Terminolo.htm) Projects: ISO 20776-1, ISO 20776-2

	   \[Note 1: Bacterial strains are categorized as resistant by applying the appropriate breakpoints in a defined phenotypic test system.\]

	   \[Note 2: This breakpoint can be altered due to changes in circumstances (e.g., changes in commonly used drug dosages, emergence of new resistance mechanisms).\]
	*/
	ObservationInterpretationCodesResistant ObservationInterpretationCodes = "R"
	/*
	   Bacterial strain inhibited by in vitro concentration of an antimicrobial agent that is associated with a high likelihood of therapeutic success. Reference: CLSI (http://www.clsi.org/Content/NavigationMenu/Resources/HarmonizedTerminologyDatabase/Harmonized\_Terminolo.htm) Synonym (earlier term): Sensitive Projects: ISO 20776-1, ISO 20776-2

	   \[Note 1: Bacterial strains are categorized as susceptible by applying the appropriate breakpoints in a defined phenotypic system.\]

	   \[Note 2: This breakpoint can be altered due to changes in circumstances (e.g., changes in commonly used drug dosages, emergence of new resistance mechanisms).\]
	*/
	ObservationInterpretationCodesSusceptible ObservationInterpretationCodes = "S"
	/*
	   The patient is considered as carrier based on the testing results. A carrier is an individual who carries an altered form of a gene which can lead to having a child or offspring in future generations with a genetic disorder.

	   ***Deprecation Comment:*** This antimicrobial susceptibility test interpretation concept is recommended by OO to be deprecated as it is no longer recommended for use in susceptibility testing by CLSI (reference CLSI document M100-S22; Vol. 32 No.3; CLSI Performance Standards for Antimicrobial Susceptibility Testing; Twenty-Second Informational Supplement. Jan 2012).
	*/
	ObservationInterpretationCodesVerySusceptible ObservationInterpretationCodes = "VS"
	/*
	   The result or observation value is outside a reference range or expected norm at a level at which immediate action should be considered for patient safety (as defined for the respective test procedure).

	   \[Note: Typically applies to non-numeric results. Analogous to critical/panic limits for numeric results.\]
	*/
	ObservationInterpretationCodesCriticalAbnormal ObservationInterpretationCodes = "AA"
	/*
	   The result for a quantitative observation is above the upper limit of the reference range (as defined for the respective test procedure).

	   Synonym: Above high normal
	*/
	ObservationInterpretationCodesHigh ObservationInterpretationCodes = "H"
	/*
	   The result for a quantitative observation is below the lower limit of the reference range (as defined for the respective test procedure).

	   Synonym: Below low normal
	*/
	ObservationInterpretationCodesLow ObservationInterpretationCodes = "L"
	/*
	   The result for a quantitative observation is above a reference level at which immediate action should be considered for patient safety (as defined for the respective test procedure).

	   Synonym: Above upper panic limits.
	*/
	ObservationInterpretationCodesCriticalHigh ObservationInterpretationCodes = "HH"
	/*
	   The result for a quantitative observation is below a reference level at which immediate action should be considered for patient safety (as defined for the respective test procedure).

	   Synonym: Below lower panic limits.
	*/
	ObservationInterpretationCodesCriticalLow ObservationInterpretationCodes = "LL"
	/*
	   The observation/test result is interpreted as being outside the inclusion range for a particular protocol within which the result is being reported.

	   **Example:** A positive result on a Hepatitis screening test.

	   *Open Issue:* EX, HX, LX: These three concepts do not seem to meet a clear need in the vocabulary, and their use in observation interpretation appears likely to be covered by other existing concepts (e.g., A, H, L). The only apparent significant difference is their reference to use in protocols for exclusion of study subjects. These concepts/codes were proposed by RCRIM for use in the CTLaboratory message. They were submitted and approved in the November 2005 Harmonization cycle in proposal "030103C\_VOCAB\_RCRIM\_l\_quade\_RCRIM Obs Interp\_20051028154455". However, this proposal was not fully implemented in the vocabulary. The proposal recommended creation of the x\_ClinicalResearchExclusion domain in ObservationInterpretation with a value set including those three concepts/codes, but there is no subdomain of that name or equivalent with a binding to either of the value sets that contain these concepts/codes. Members of the OO WG have recently attempted to contact members of RCRIM regarding these concepts, both by email and at the recent WGM in Atlanta, without response. It is felt by OO that the best course of action to take at this time is to add this comprehensive Open Issue rather than deprecate these three concepts at this time, until further discussion is held.
	*/
	ObservationInterpretationCodesAboveHighThreshold ObservationInterpretationCodes = "HX"
	/*
	   The numeric observation/test result is interpreted as being below the low threshold value for a particular protocol within which the result is being reported.

	   **Example:** A Total White Blood Cell Count falling below a protocol-defined threshold value of 3000/mm^3

	   *Open Issue:* EX, HX, LX: These three concepts do not seem to meet a clear need in the vocabulary, and their use in observation interpretation appears likely to be covered by other existing concepts (e.g., A, H, L). The only apparent significant difference is their reference to use in protocols for exclusion of study subjects. These concepts/codes were proposed by RCRIM for use in the CTLaboratory message. They were submitted and approved in the November 2005 Harmonization cycle in proposal "030103C\_VOCAB\_RCRIM\_l\_quade\_RCRIM Obs Interp\_20051028154455". However, this proposal was not fully implemented in the vocabulary. The proposal recommended creation of the x\_ClinicalResearchExclusion domain in ObservationInterpretation with a value set including those three concepts/codes, but there is no subdomain of that name or equivalent with a binding to either of the value sets that contain these concepts/codes. Members of the OO WG have recently attempted to contact members of RCRIM regarding these concepts, both by email and at the recent WGM in Atlanta, without response. It is felt by OO that the best course of action to take at this time is to add this comprehensive Open Issue rather than deprecate these three concepts at this time, until further discussion is held.
	*/
	ObservationInterpretationCodesBelowLowThreshold ObservationInterpretationCodes = "LX"
	/*
	   A test result that is significantly higher than the reference (normal) or therapeutic interval, but has not reached the critically high value and might need special attention, as defined by the laboratory or the clinician.

	   \[Note: This level is situated between 'H' and 'HH'.\]

	   *Deprecation Comment:* The code 'H>' is being deprecated in order to align with the use of the code 'HU' for "Very high" in V2 Table 0078 "Interpretation Codes".

	   \[Note: The use of code 'H>' is non-preferred, as this code is deprecated and on track to be retired; use code 'HU' instead.
	*/
	ObservationInterpretationCodesSignificantlyHigh ObservationInterpretationCodes = "H>"
	// A test result that is significantly higher than the reference (normal) or therapeutic interval, but has not reached the critically high value and might need special attention, as defined by the laboratory or the clinician.
	ObservationInterpretationCodesSignificantlyHigh2 ObservationInterpretationCodes = "HU"
	// The test or procedure was successfully performed, but the results are borderline and can neither be declared positive / negative nor detected / not detected according to the current established criteria.
	ObservationInterpretationCodesEquivocal ObservationInterpretationCodes = "E"
	/*
	   A test result that is significantly lower than the reference (normal) or therapeutic interval, but has not reached the critically low value and might need special attention, as defined by the laboratory or the clinician.

	   \[Note: This level is situated between 'L' and 'LL'.\]

	   *Deprecation Comment:* The code 'L<' is being deprecated in order to align with the use of the code 'LU' for "Very low" in V2 Table 0078 "Interpretation Codes".

	   \[Note: The use of code 'L<' is non-preferred, as this code is deprecated and on track to be retired; use code 'LU' instead.
	*/
	ObservationInterpretationCodesSignificantlyLow ObservationInterpretationCodes = "L<"
	// A test result that is significantly lower than the reference (normal) or therapeutic interval, but has not reached the critically low value and might need special attention, as defined by the laboratory or the clinician.
	ObservationInterpretationCodesSignificantlyLow2 ObservationInterpretationCodes = "LU"
	// The presence of the specified component / analyte, organism or clinical sign could not be determined within the limit of detection of the performed test or procedure.
	ObservationInterpretationCodesNotDetected ObservationInterpretationCodes = "ND"
	/*
	   The specified component / analyte, organism or clinical sign could neither be declared positive / negative nor detected / not detected by the performed test or procedure.

	   *Usage Note:* For example, if the specimen was degraded, poorly processed, or was missing the required anatomic structures, then "indeterminate" (i.e. "cannot be determined") is the appropriate response, not "equivocal".
	*/
	ObservationInterpretationCodesIndeterminate ObservationInterpretationCodes = "IND"
	/*
	   An absence finding of the specified component / analyte, organism or clinical sign based on the established threshold of the performed test or procedure.

	   \[Note: Negative does not necessarily imply the complete absence of the specified item.\]
	*/
	ObservationInterpretationCodesNegative ObservationInterpretationCodes = "NEG"
	// A presence finding of the specified component / analyte, organism or clinical sign based on the established threshold of the performed test or procedure.
	ObservationInterpretationCodesPositive ObservationInterpretationCodes = "POS"
	// This result has been evaluated in light of known contraindicators. Once those contraindicators have been taken into account the result is determined to be "Expected" (e.g., presence of drugs in a patient that is taking prescription medication for pain management).
	ObservationInterpretationCodesExpected ObservationInterpretationCodes = "EXP"
	// This result has been evaluated in light of known contraindicators. Once those contraindicators have been taken into account the result is determined to be "Unexpected" (e.g., presence of non-prescribed drugs in a patient that is taking prescription medication for pain management).
	ObservationInterpretationCodesUnexpected ObservationInterpretationCodes = "UNE"
	// The measurement of the specified component / analyte, organism or clinical sign above the limit of detection of the performed test or procedure.
	ObservationInterpretationCodesDetected ObservationInterpretationCodes = "DET"
	/*
	   A category for isolates where the bacteria (e.g. enterococci) are not susceptible in vitro to a combination therapy (e.g., high-level aminoglycoside and cell wall active agent). This is predictive that this combination therapy will not be effective.

	   *Usage Note:* Since the use of penicillin or ampicillin alone often results in treatment failure of serious enterococcal or other bacterial infections, combination therapy is usually indicated to enhance bactericidal activity. The synergy between a cell wall active agent (such as penicillin, ampicillin, or vancomycin) and an aminoglycoside (such as gentamicin, kanamycin or streptomycin) is best predicted by screening for high-level bacterial resistance to the aminoglycoside.

	   *Open Issue:* The print name of the code is very general and the description is very specific to a pair of classes of agents, which may lead to confusion of these concepts in the future should other synergies be found.
	*/
	ObservationInterpretationCodesSynergyResistant ObservationInterpretationCodes = "SYN-R"
	// An absence finding used to indicate that the specified component / analyte did not react measurably with the reagent.
	ObservationInterpretationCodesNonReactive ObservationInterpretationCodes = "NR"
	// A presence finding used to indicate that the specified component / analyte reacted with the reagent above the reliably measurable limit of the performed test.
	ObservationInterpretationCodesReactive ObservationInterpretationCodes = "RR"
	// A weighted presence finding used to indicate that the specified component / analyte reacted with the reagent, but below the reliably measurable limit of the performed test.
	ObservationInterpretationCodesWeaklyReactive ObservationInterpretationCodes = "WR"
	/*
	   A category that includes isolates with antimicrobial agent minimum inhibitory concentrations (MICs) that approach usually attainable blood and tissue levels and for which response rates may be lower than for susceptible isolates.

	   Reference: CLSI document M44-A2 2009 "Method for antifungal disk diffusion susceptibility testing of yeasts; approved guideline - second edition" - page 2.
	*/
	ObservationInterpretationCodesSusceptibleDoseDependent ObservationInterpretationCodes = "SDD"
	/*
	   A category for isolates where the bacteria (e.g. enterococci) are susceptible in vitro to a combination therapy (e.g., high-level aminoglycoside and cell wall active agent). This is predictive that this combination therapy will be effective.

	   *Usage Note:* Since the use of penicillin or ampicillin alone often results in treatment failure of serious enterococcal or other bacterial infections, combination therapy is usually indicated to enhance bactericidal activity. The synergy between a cell wall active agent (such as penicillin, ampicillin, or vancomycin) and an aminoglycoside (such as gentamicin, kanamycin or streptomycin) is best predicted by screening for high-level bacterial resistance to the aminoglycoside.

	   *Open Issue:* The print name of the code is very general and the description is very specific to a pair of classes of agents, which may lead to confusion of these concepts in the future should other synergies be found.
	*/
	ObservationInterpretationCodesSynergySusceptible ObservationInterpretationCodes = "SYN-S"
)
