// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/operation-outcome
*/type AuditEventOutcomeDetail string

const (
	AuditEventOutcomeDetailErrorMultipleMatchesExistForTheConditionalDelete                      AuditEventOutcomeDetail = "DELETE_MULTIPLE_MATCHES"
	AuditEventOutcomeDetailYouMustAuthenticateBeforeYouCanUseThisService                         AuditEventOutcomeDetail = "MSG_AUTH_REQUIRED"
	AuditEventOutcomeDetailBadSyntaxSMustBeAs                                                    AuditEventOutcomeDetail = "MSG_BAD_FORMAT"
	AuditEventOutcomeDetailBadSyntaxInS                                                          AuditEventOutcomeDetail = "MSG_BAD_SYNTAX"
	AuditEventOutcomeDetailUnableToParseFeedEntryContentTypeS                                    AuditEventOutcomeDetail = "MSG_CANT_PARSE_CONTENT"
	AuditEventOutcomeDetailUnableToParseFeedRootElementNameS                                     AuditEventOutcomeDetail = "MSG_CANT_PARSE_ROOT"
	AuditEventOutcomeDetailNewResourceCreated                                                    AuditEventOutcomeDetail = "MSG_CREATED"
	AuditEventOutcomeDetailTheDateValueSIsNotInTheCorrectFormatXmlDateFormatRequired             AuditEventOutcomeDetail = "MSG_DATE_FORMAT"
	AuditEventOutcomeDetailThisResourceHasBeenDeleted                                            AuditEventOutcomeDetail = "MSG_DELETED"
	AuditEventOutcomeDetailResourceDeleted                                                       AuditEventOutcomeDetail = "MSG_DELETED_DONE"
	AuditEventOutcomeDetailTheResourceSHasBeenDeleted                                            AuditEventOutcomeDetail = "MSG_DELETED_ID"
	AuditEventOutcomeDetailDuplicateIdSForResourceTypeS                                          AuditEventOutcomeDetail = "MSG_DUPLICATE_ID"
	AuditEventOutcomeDetailErrorParsingResourceXmlS                                              AuditEventOutcomeDetail = "MSG_ERROR_PARSING"
	AuditEventOutcomeDetailUnableToResolveExternalReferenceToResourceS                           AuditEventOutcomeDetail = "MSG_EXTERNAL_FAIL"
	AuditEventOutcomeDetailIdSHasAnInvalidCharacterS                                             AuditEventOutcomeDetail = "MSG_ID_INVALID"
	AuditEventOutcomeDetailIdSTooLongLengthLimit36                                               AuditEventOutcomeDetail = "MSG_ID_TOO_LONG"
	AuditEventOutcomeDetailIdNotAccepted                                                         AuditEventOutcomeDetail = "MSG_INVALID_ID"
	AuditEventOutcomeDetailJsonSourceForAResourceShouldStartWithAnObject                         AuditEventOutcomeDetail = "MSG_JSON_OBJECT"
	AuditEventOutcomeDetailUnableToResolveLocalReferenceToResourceS                              AuditEventOutcomeDetail = "MSG_LOCAL_FAIL"
	AuditEventOutcomeDetailResourceIdSDoesNotExist                                               AuditEventOutcomeDetail = "MSG_NO_EXIST"
	AuditEventOutcomeDetailNoResourceFoundMatchingTheQueryS                                      AuditEventOutcomeDetail = "MSG_NO_MATCH"
	AuditEventOutcomeDetailNoModuleCouldBeFoundToHandleTheRequestS                               AuditEventOutcomeDetail = "MSG_NO_MODULE"
	AuditEventOutcomeDetailNoSummaryForThisResource                                              AuditEventOutcomeDetail = "MSG_NO_SUMMARY"
	AuditEventOutcomeDetailOperationSNotAllowedForResourceSDueToLocalConfiguration               AuditEventOutcomeDetail = "MSG_OP_NOT_ALLOWED"
	AuditEventOutcomeDetailUnknownChainedParameterNameS                                          AuditEventOutcomeDetail = "MSG_PARAM_CHAINED"
	AuditEventOutcomeDetailParameterSContentIsInvalid                                            AuditEventOutcomeDetail = "MSG_PARAM_INVALID"
	AuditEventOutcomeDetailParameterSModifierIsInvalid                                           AuditEventOutcomeDetail = "MSG_PARAM_MODIFIER_INVALID"
	AuditEventOutcomeDetailParameterSIsNotAllowedToRepeat                                        AuditEventOutcomeDetail = "MSG_PARAM_NO_REPEAT"
	AuditEventOutcomeDetailParameterSNotUnderstood                                               AuditEventOutcomeDetail = "MSG_PARAM_UNKNOWN"
	AuditEventOutcomeDetailResourcesWithIdentityExampleCannotBeDeletedForTestingTrainingPurposes AuditEventOutcomeDetail = "MSG_RESOURCE_EXAMPLE_PROTECTED"
	AuditEventOutcomeDetailUnableToAllocateResourceId                                            AuditEventOutcomeDetail = "MSG_RESOURCE_ID_FAIL"
	AuditEventOutcomeDetailResourceIdMismatch                                                    AuditEventOutcomeDetail = "MSG_RESOURCE_ID_MISMATCH"
	AuditEventOutcomeDetailResourceIdMissing                                                     AuditEventOutcomeDetail = "MSG_RESOURCE_ID_MISSING"
	AuditEventOutcomeDetailNotAllowedToSubmitAResourceForThisOperation                           AuditEventOutcomeDetail = "MSG_RESOURCE_NOT_ALLOWED"
	AuditEventOutcomeDetailAResourceIsRequired                                                   AuditEventOutcomeDetail = "MSG_RESOURCE_REQUIRED"
	AuditEventOutcomeDetailResourceTypeMismatch                                                  AuditEventOutcomeDetail = "MSG_RESOURCE_TYPE_MISMATCH"
	AuditEventOutcomeDetailUnknownSortParameterNameS                                             AuditEventOutcomeDetail = "MSG_SORT_UNKNOWN"
	AuditEventOutcomeDetailDuplicateIdentifierInTransactionS                                     AuditEventOutcomeDetail = "MSG_TRANSACTION_DUPLICATE_ID"
	AuditEventOutcomeDetailMissingIdentifierInTransactionAnEntryIdMustBeProvided                 AuditEventOutcomeDetail = "MSG_TRANSACTION_MISSING_ID"
	AuditEventOutcomeDetailUnhandledXmlNodeTypeS                                                 AuditEventOutcomeDetail = "MSG_UNHANDLED_NODE_TYPE"
	AuditEventOutcomeDetailUnknownContentSAtS                                                    AuditEventOutcomeDetail = "MSG_UNKNOWN_CONTENT"
	AuditEventOutcomeDetailUnknownFhirHttpOperation                                              AuditEventOutcomeDetail = "MSG_UNKNOWN_OPERATION"
	AuditEventOutcomeDetailResourceTypeSNotRecognised                                            AuditEventOutcomeDetail = "MSG_UNKNOWN_TYPE"
	AuditEventOutcomeDetailExistingResourceUpdated                                               AuditEventOutcomeDetail = "MSG_UPDATED"
	AuditEventOutcomeDetailVersionAwareUpdatesAreRequiredForThisResource                         AuditEventOutcomeDetail = "MSG_VERSION_AWARE"
	AuditEventOutcomeDetailUpdateConflictServerCurrentVersionSClientVersionReferencedS           AuditEventOutcomeDetail = "MSG_VERSION_AWARE_CONFLICT"
	AuditEventOutcomeDetailVersionSpecificUrlNotRecognised                                       AuditEventOutcomeDetail = "MSG_VERSION_AWARE_URL"
	AuditEventOutcomeDetailThisDoesNotAppearToBeAfhirElementOrResourceWrongNamespaceS            AuditEventOutcomeDetail = "MSG_WRONG_NS"
	AuditEventOutcomeDetailErrorMultipleMatchesExistForSSearchParametersS                        AuditEventOutcomeDetail = "SEARCH_MULTIPLE"
	AuditEventOutcomeDetailErrorNoProcessableSearchFoundForSSearchParametersS                    AuditEventOutcomeDetail = "SEARCH_NONE"
	AuditEventOutcomeDetailErrorMultipleMatchesExistForTheConditionalUpdate                      AuditEventOutcomeDetail = "UPDATE_MULTIPLE_MATCHES"
)
