// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/v3-Confidentiality
  - N
  - R

- http://terminology.hl7.org/CodeSystem/v3-ActCode
  - ETH
  - PSY
  - STD
*/type SecurityLabelDataExamples string

const (
	/*
	   Policy for handling alcohol or drug-abuse information, which will be afforded heightened confidentiality. Information handling protocols based on organizational policies related to alcohol or drug-abuse information that is deemed sensitive.

	   *Usage Note:* If there is a jurisdictional mandate, then use the applicable ActPrivacyLaw code system, and specify the law rather than or in addition to this more generic code.
	*/
	SecurityLabelDataExamplesSubstanceAbuseInformationSensitivity SecurityLabelDataExamples = "ETH"
	/*
	   Policy for handling sexually transmitted disease information, which will be afforded heightened confidentiality. Information handling protocols based on organizational policies related to sexually transmitted disease information that is deemed sensitive.

	   *Usage Note:* If there is a jurisdictional mandate, then use the applicable ActPrivacyLaw code system, and specify the law rather than or in addition to this more generic code.
	*/
	SecurityLabelDataExamplesSexuallyTransmittedDiseaseInformationSensitivity SecurityLabelDataExamples = "STD"
	/*
	   Policy for handling psychiatry psychiatric disorder information, which is afforded heightened confidentiality.

	   *Usage Note:* If there is a jurisdictional mandate, then use the applicable ActPrivacyLaw code system, and specify the law rather than or in addition to this more generic code.
	*/
	SecurityLabelDataExamplesPsychiatryDisorderInformationSensitivity SecurityLabelDataExamples = "PSY"
)
