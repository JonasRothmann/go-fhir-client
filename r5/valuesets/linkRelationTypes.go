// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/iana-link-relations
*/type LinkRelationTypes string

const (
	// Refers to a resource that is the subject of the link's context.
	LinkRelationTypesAbout LinkRelationTypes = "about"
	// Asserts that the link target provides an access control description for the link context.
	LinkRelationTypesAcl LinkRelationTypes = "acl"
	// Refers to a substitute for this context
	LinkRelationTypesAlternate LinkRelationTypes = "alternate"
	// Used to reference alternative content that uses the AMP profile of the HTML format.
	LinkRelationTypesAmphtml LinkRelationTypes = "amphtml"
	// Refers to an appendix.
	LinkRelationTypesAppendix LinkRelationTypes = "appendix"
	// Refers to an icon for the context. Synonym for icon.
	LinkRelationTypesRefersToAnIconForTheContextSynonymForIcon LinkRelationTypes = "apple-touch-icon"
	// Refers to a launch screen for the context.
	LinkRelationTypesRefersToALaunchScreenForTheContext LinkRelationTypes = "apple-touch-startup-image"
	/*
	   Refers to a collection of records, documents, or other
	         materials of historical interest.
	*/
	LinkRelationTypesArchives LinkRelationTypes = "archives"
	// Refers to the context's author.
	LinkRelationTypesAuthor LinkRelationTypes = "author"
	/*
	   Identifies the entity that blocks access to a resource
	         following receipt of a legal demand.
	*/
	LinkRelationTypesIdentifiesTheEntityThatBlocksAccessToAResourceFollowingReceiptOfALegalDemand LinkRelationTypes = "blocked-by"
	// Gives a permanent link to use for bookmarking purposes.
	LinkRelationTypesBookmark LinkRelationTypes = "bookmark"
	// Designates the preferred version of a resource (the IRI and its contents).
	LinkRelationTypesCanonical LinkRelationTypes = "canonical"
	// Refers to a chapter in a collection of resources.
	LinkRelationTypesChapter LinkRelationTypes = "chapter"
	// Indicates that the link target is preferred over the link context for the purpose of permanent citation.
	LinkRelationTypesIndicatesThatTheLinkTargetIsPreferredOverTheLinkContextForThePurposeOfPermanentCitation LinkRelationTypes = "cite-as"
	// The target IRI points to a resource which represents the collection resource for the context IRI.
	LinkRelationTypesCollection LinkRelationTypes = "collection"
	// Refers to a table of contents.
	LinkRelationTypesContents LinkRelationTypes = "contents"
	/*
	   The document linked to was later converted to the
	         document that contains this link relation.  For example, an RFC can
	         have a link to the Internet-Draft that became the RFC; in that case,
	         the link relation would be "convertedFrom".
	*/
	LinkRelationTypesConvertedFrom LinkRelationTypes = "convertedFrom"
	/*
	   Refers to a copyright statement that applies to the
	       link's context.
	*/
	LinkRelationTypesCopyright LinkRelationTypes = "copyright"
	// The target IRI points to a resource where a submission form can be obtained.
	LinkRelationTypesTheTargetIriPointsToAResourceWhereASubmissionFormCanBeObtained LinkRelationTypes = "create-form"
	/*
	   Refers to a resource containing the most recent
	         item(s) in a collection of resources.
	*/
	LinkRelationTypesCurrent LinkRelationTypes = "current"
	/*
	   Refers to a resource providing information about the
	         link's context.
	*/
	LinkRelationTypesDescribedby LinkRelationTypes = "describedby"
	/*
	   The relationship A 'describes' B asserts that
	         resource A provides a description of resource B. There are no
	         constraints on the format or representation of either A or B,
	         neither are there any further constraints on either resource.
	*/
	LinkRelationTypesDescribes LinkRelationTypes = "describes"
	/*
	   Refers to a list of patent disclosures made with respect to
	         material for which 'disclosure' relation is specified.
	*/
	LinkRelationTypesDisclosure LinkRelationTypes = "disclosure"
	/*
	   Used to indicate an origin that will be used to fetch required
	         resources for the link context, and that the user agent ought to resolve
	         as early as possible.
	*/
	LinkRelationTypesUsedToIndicateAnOriginThatWillBeUsedToFetchRequiredResourcesForTheLinkContextAndThatTheUserAgentOughtToResolveAsEarlyAsPossible LinkRelationTypes = "dns-prefetch"
	/*
	   Refers to a resource whose available representations
	         are byte-for-byte identical with the corresponding representations of
	         the context IRI.
	*/
	LinkRelationTypesDuplicate LinkRelationTypes = "duplicate"
	/*
	   Refers to a resource that can be used to edit the
	         link's context.
	*/
	LinkRelationTypesEdit LinkRelationTypes = "edit"
	/*
	   The target IRI points to a resource where a submission form for
	         editing associated resource can be obtained.
	*/
	LinkRelationTypesTheTargetIriPointsToAResourceWhereASubmissionFormForEditingAssociatedResourceCanBeObtained LinkRelationTypes = "edit-form"
	/*
	   Refers to a resource that can be used to edit media
	         associated with the link's context.
	*/
	LinkRelationTypesRefersToAResourceThatCanBeUsedToEditMediaAssociatedWithTheLinkSContext LinkRelationTypes = "edit-media"
	/*
	   Identifies a related resource that is potentially
	         large and might require special handling.
	*/
	LinkRelationTypesEnclosure LinkRelationTypes = "enclosure"
	// Refers to a resource that is not part of the same site as the current context.
	LinkRelationTypesExternal LinkRelationTypes = "external"
	/*
	   An IRI that refers to the furthest preceding resource
	       in a series of resources.
	*/
	LinkRelationTypesFirst LinkRelationTypes = "first"
	// Refers to a glossary of terms.
	LinkRelationTypesGlossary LinkRelationTypes = "glossary"
	// Refers to context-sensitive help.
	LinkRelationTypesHelp LinkRelationTypes = "help"
	/*
	   Refers to a resource hosted by the server indicated by
	         the link context.
	*/
	LinkRelationTypesHosts LinkRelationTypes = "hosts"
	/*
	   Refers to a hub that enables registration for
	       notification of updates to the context.
	*/
	LinkRelationTypesHub LinkRelationTypes = "hub"
	// Refers to an icon representing the link's context.
	LinkRelationTypesIcon LinkRelationTypes = "icon"
	// Refers to an index.
	LinkRelationTypesIndex LinkRelationTypes = "index"
	// refers to a resource associated with a time interval that ends before the beginning of the time interval associated with the context resource
	LinkRelationTypesIntervalAfter LinkRelationTypes = "intervalAfter"
	// refers to a resource associated with a time interval that begins after the end of the time interval associated with the context resource
	LinkRelationTypesIntervalBefore LinkRelationTypes = "intervalBefore"
	// refers to a resource associated with a time interval that begins after the beginning of the time interval associated with the context resource, and ends before the end of the time interval associated with the context resource
	LinkRelationTypesIntervalContains LinkRelationTypes = "intervalContains"
	// refers to a resource associated with a time interval that begins after the end of the time interval associated with the context resource, or ends before the beginning of the time interval associated with the context resource
	LinkRelationTypesIntervalDisjoint LinkRelationTypes = "intervalDisjoint"
	// refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource
	LinkRelationTypesIntervalDuring LinkRelationTypes = "intervalDuring"
	// refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource
	LinkRelationTypesIntervalEquals LinkRelationTypes = "intervalEquals"
	// refers to a resource associated with a time interval that begins after the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource
	LinkRelationTypesIntervalFinishedBy LinkRelationTypes = "intervalFinishedBy"
	// refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and whose end coincides with the end of the time interval associated with the context resource
	LinkRelationTypesIntervalFinishes LinkRelationTypes = "intervalFinishes"
	// refers to a resource associated with a time interval that begins before or is coincident with the beginning of the time interval associated with the context resource, and ends after or is coincident with the end of the time interval associated with the context resource
	LinkRelationTypesIntervalIn LinkRelationTypes = "intervalIn"
	// refers to a resource associated with a time interval whose beginning coincides with the end of the time interval associated with the context resource
	LinkRelationTypesIntervalMeets LinkRelationTypes = "intervalMeets"
	// refers to a resource associated with a time interval whose end coincides with the beginning of the time interval associated with the context resource
	LinkRelationTypesIntervalMetBy LinkRelationTypes = "intervalMetBy"
	// refers to a resource associated with a time interval that begins before the beginning of the time interval associated with the context resource, and ends after the beginning of the time interval associated with the context resource
	LinkRelationTypesIntervalOverlappedBy LinkRelationTypes = "intervalOverlappedBy"
	// refers to a resource associated with a time interval that begins before the end of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource
	LinkRelationTypesIntervalOverlaps LinkRelationTypes = "intervalOverlaps"
	// refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and ends before the end of the time interval associated with the context resource
	LinkRelationTypesIntervalStartedBy LinkRelationTypes = "intervalStartedBy"
	// refers to a resource associated with a time interval whose beginning coincides with the beginning of the time interval associated with the context resource, and ends after the end of the time interval associated with the context resource
	LinkRelationTypesIntervalStarts LinkRelationTypes = "intervalStarts"
	// The target IRI points to a resource that is a member of the collection represented by the context IRI.
	LinkRelationTypesItem LinkRelationTypes = "item"
	/*
	   An IRI that refers to the furthest following resource
	         in a series of resources.
	*/
	LinkRelationTypesLast LinkRelationTypes = "last"
	/*
	   Points to a resource containing the latest (e.g.,
	         current) version of the context.
	*/
	LinkRelationTypesPointsToAResourceContainingTheLatestEgCurrentVersionOfTheContext LinkRelationTypes = "latest-version"
	// Refers to a license associated with this context.
	LinkRelationTypesLicense LinkRelationTypes = "license"
	/*
	   The link target of a link with the "linkset" relation
	         type provides a set of links, including links in which the link
	         context of the link participates.

	*/
	LinkRelationTypesLinkset LinkRelationTypes = "linkset"
	/*
	   Refers to further information about the link's context,
	         expressed as a LRDD ("Link-based Resource Descriptor Document")
	         resource.  See  for information about
	         processing this relation type in host-meta documents. When used
	         elsewhere, it refers to additional links and other metadata.
	         Multiple instances indicate additional LRDD resources. LRDD
	         resources MUST have an "application/xrd+xml" representation, and
	         MAY have others.
	*/
	LinkRelationTypesLrdd LinkRelationTypes = "lrdd"
	// Links to a manifest file for the context.
	LinkRelationTypesManifest LinkRelationTypes = "manifest"
	// Refers to a mask that can be applied to the icon for the context.
	LinkRelationTypesRefersToAMaskThatCanBeAppliedToTheIconForTheContext LinkRelationTypes = "mask-icon"
	// Refers to a feed of personalised media recommendations relevant to the link context.
	LinkRelationTypesRefersToAFeedOfPersonalisedMediaRecommendationsRelevantToTheLinkContext LinkRelationTypes = "media-feed"
	// The Target IRI points to a Memento, a fixed resource that will not change state anymore.
	LinkRelationTypesMemento LinkRelationTypes = "memento"
	// Links to the context's Micropub endpoint.
	LinkRelationTypesMicropub LinkRelationTypes = "micropub"
	// Refers to a module that the user agent is to preemptively fetch and store for use in the current context.
	LinkRelationTypesModulepreload LinkRelationTypes = "modulepreload"
	/*
	   Refers to a resource that can be used to monitor changes in an HTTP resource.

	*/
	LinkRelationTypesMonitor LinkRelationTypes = "monitor"
	/*
	   Refers to a resource that can be used to monitor changes in a specified group of HTTP resources.

	*/
	LinkRelationTypesRefersToAResourceThatCanBeUsedToMonitorChangesInASpecifiedGroupOfHttpResources LinkRelationTypes = "monitor-group"
	/*
	   Indicates that the link's context is a part of a series, and
	         that the next in the series is the link target.

	*/
	LinkRelationTypesNext LinkRelationTypes = "next"
	// Refers to the immediately following archive resource.
	LinkRelationTypesRefersToTheImmediatelyFollowingArchiveResource LinkRelationTypes = "next-archive"
	// Indicates that the context’s original author or publisher does not endorse the link target.
	LinkRelationTypesNofollow LinkRelationTypes = "nofollow"
	// Indicates that any newly created top-level browsing context which results from following the link will not be an auxiliary browsing context.
	LinkRelationTypesNoopener LinkRelationTypes = "noopener"
	// Indicates that no referrer information is to be leaked when following the link.
	LinkRelationTypesNoreferrer LinkRelationTypes = "noreferrer"
	// Indicates that any newly created top-level browsing context which results from following the link will be an auxiliary browsing context.
	LinkRelationTypesOpener LinkRelationTypes = "opener"
	// Refers to an OpenID Authentication server on which the context relies for an assertion that the end user controls an Identifier.
	LinkRelationTypesRefersToAnOpenIdAuthenticationServerOnWhichTheContextReliesForAnAssertionThatTheEndUserControlsAnIdentifier LinkRelationTypes = "openid2.local_id"
	// Refers to a resource which accepts OpenID Authentication protocol messages for the context.
	LinkRelationTypesRefersToAResourceWhichAcceptsOpenIdAuthenticationProtocolMessagesForTheContext LinkRelationTypes = "openid2.provider"
	// The Target IRI points to an Original Resource.
	LinkRelationTypesOriginal LinkRelationTypes = "original"
	// Refers to a P3P privacy policy for the context.
	LinkRelationTypesP3pv1 LinkRelationTypes = "P3Pv1"
	// Indicates a resource where payment is accepted.
	LinkRelationTypesPayment LinkRelationTypes = "payment"
	// Gives the address of the pingback resource for the link context.
	LinkRelationTypesPingback LinkRelationTypes = "pingback"
	/*
	   Used to indicate an origin that will be used to fetch required
	         resources for the link context. Initiating an early connection, which
	         includes the DNS lookup, TCP handshake, and optional TLS negotiation,
	         allows the user agent to mask the high latency costs of establishing a
	         connection.
	*/
	LinkRelationTypesPreconnect LinkRelationTypes = "preconnect"
	/*
	   Points to a resource containing the predecessor
	         version in the version history.

	*/
	LinkRelationTypesPointsToAResourceContainingThePredecessorVersionInTheVersionHistory LinkRelationTypes = "predecessor-version"
	/*
	   The prefetch link relation type is used to identify a resource
	         that might be required by the next navigation from the link context, and
	         that the user agent ought to fetch, such that the user agent can deliver a
	         faster response once the resource is requested in the future.
	*/
	LinkRelationTypesPrefetch LinkRelationTypes = "prefetch"
	/*
	   Refers to a resource that should be loaded early in the
	         processing of the link's context, without blocking rendering.
	*/
	LinkRelationTypesPreload LinkRelationTypes = "preload"
	/*
	   Used to identify a resource that might be required by the next
	         navigation from the link context, and that the user agent ought to fetch
	         and execute, such that the user agent can deliver a faster response once
	         the resource is requested in the future.
	*/
	LinkRelationTypesPrerender LinkRelationTypes = "prerender"
	/*
	   Indicates that the link's context is a part of a series, and
	         that the previous in the series is the link target.

	*/
	LinkRelationTypesPrev LinkRelationTypes = "prev"
	// Refers to a resource that provides a preview of the link's context.
	LinkRelationTypesPreview LinkRelationTypes = "preview"
	/*
	   Refers to the previous resource in an ordered series
	         of resources.  Synonym for "prev".
	*/
	LinkRelationTypesPrevious LinkRelationTypes = "previous"
	// Refers to the immediately preceding archive resource.
	LinkRelationTypesRefersToTheImmediatelyPrecedingArchiveResource LinkRelationTypes = "prev-archive"
	// Refers to a privacy policy associated with the link's context.
	LinkRelationTypesRefersToAPrivacyPolicyAssociatedWithTheLinkSContext LinkRelationTypes = "privacy-policy"
	/*
	   Identifying that a resource representation conforms
	   to a certain profile, without affecting the non-profile semantics
	   of the resource representation.
	*/
	LinkRelationTypesProfile LinkRelationTypes = "profile"
	/*
	   Links to a publication manifest. A manifest represents
	         structured information about a publication, such as informative metadata,
	         a list of resources, and a default reading order.
	*/
	LinkRelationTypesPublication LinkRelationTypes = "publication"
	// Identifies a related resource.
	LinkRelationTypesRelated LinkRelationTypes = "related"
	/*
	   Identifies the root of RESTCONF API as configured on this HTTP server.
	         The "restconf" relation defines the root of the API defined in RFC8040.
	         Subsequent revisions of RESTCONF will use alternate relation values to support
	         protocol versioning.
	*/
	LinkRelationTypesRestconf LinkRelationTypes = "restconf"
	/*
	   Identifies a resource that is a reply to the context
	         of the link.

	*/
	LinkRelationTypesReplies LinkRelationTypes = "replies"
	/*
	   The resource identified by the link target provides an input value to an
	       instance of a rule, where the resource which represents the rule instance is
	       identified by the link context.

	*/
	LinkRelationTypesRuleinput LinkRelationTypes = "ruleinput"
	/*
	   Refers to a resource that can be used to search through
	         the link's context and related resources.
	*/
	LinkRelationTypesSearch LinkRelationTypes = "search"
	// Refers to a section in a collection of resources.
	LinkRelationTypesSection LinkRelationTypes = "section"
	/*
	   Conveys an identifier for the link's context.

	*/
	LinkRelationTypesSelf LinkRelationTypes = "self"
	/*
	   Indicates a URI that can be used to retrieve a
	         service document.
	*/
	LinkRelationTypesService LinkRelationTypes = "service"
	/*
	   Identifies service description for the context that
	         is primarily intended for consumption by machines.
	*/
	LinkRelationTypesIdentifiesServiceDescriptionForTheContextThatIsPrimarilyIntendedForConsumptionByMachines LinkRelationTypes = "service-desc"
	/*
	   Identifies service documentation for the context that
	         is primarily intended for human consumption.
	*/
	LinkRelationTypesIdentifiesServiceDocumentationForTheContextThatIsPrimarilyIntendedForHumanConsumption LinkRelationTypes = "service-doc"
	/*
	   Identifies general metadata for the context that is
	         primarily intended for consumption by machines.
	*/
	LinkRelationTypesIdentifiesGeneralMetadataForTheContextThatIsPrimarilyIntendedForConsumptionByMachines LinkRelationTypes = "service-meta"
	/*
	   Refers to a resource that is within a context that is
	   		sponsored (such as advertising or another compensation agreement).
	*/
	LinkRelationTypesSponsored LinkRelationTypes = "sponsored"
	/*
	   Refers to the first resource in a collection of
	         resources.
	*/
	LinkRelationTypesStart LinkRelationTypes = "start"
	/*
	   Identifies a resource that represents the context's
	         status.
	*/
	LinkRelationTypesStatus LinkRelationTypes = "status"
	// Refers to a stylesheet.
	LinkRelationTypesStylesheet LinkRelationTypes = "stylesheet"
	/*
	   Refers to a resource serving as a subsection in a
	         collection of resources.
	*/
	LinkRelationTypesSubsection LinkRelationTypes = "subsection"
	/*
	   Points to a resource containing the successor version
	         in the version history.

	*/
	LinkRelationTypesPointsToAResourceContainingTheSuccessorVersionInTheVersionHistory LinkRelationTypes = "successor-version"
	/*
	   Identifies a resource that provides information about
	         the context's retirement policy.

	*/
	LinkRelationTypesSunset LinkRelationTypes = "sunset"
	/*
	   Gives a tag (identified by the given address) that applies to
	         the current document.

	*/
	LinkRelationTypesTag LinkRelationTypes = "tag"
	// Refers to the terms of service associated with the link's context.
	LinkRelationTypesRefersToTheTermsOfServiceAssociatedWithTheLinkSContext LinkRelationTypes = "terms-of-service"
	// The Target IRI points to a TimeGate for an Original Resource.
	LinkRelationTypesTimegate LinkRelationTypes = "timegate"
	// The Target IRI points to a TimeMap for an Original Resource.
	LinkRelationTypesTimemap LinkRelationTypes = "timemap"
	// Refers to a resource identifying the abstract semantic type of which the link's context is considered to be an instance.
	LinkRelationTypesType LinkRelationTypes = "type"
	/*
	   Refers to a resource that is within a context that is User Generated Content.

	*/
	LinkRelationTypesUgc LinkRelationTypes = "ugc"
	/*
	   Refers to a parent document in a hierarchy of
	         documents.

	*/
	LinkRelationTypesUp LinkRelationTypes = "up"
	/*
	   Points to a resource containing the version history
	         for the context.

	*/
	LinkRelationTypesPointsToAResourceContainingTheVersionHistoryForTheContext LinkRelationTypes = "version-history"
	/*
	   Identifies a resource that is the source of the
	         information in the link's context.

	*/
	LinkRelationTypesVia LinkRelationTypes = "via"
	/*
	   Identifies a target URI that supports the Webmention protocol.
	       This allows clients that mention a resource in some form of publishing process
	       to contact that endpoint and inform it that this resource has been mentioned.
	*/
	LinkRelationTypesWebmention LinkRelationTypes = "webmention"
	// Points to a working copy for this resource.
	LinkRelationTypesPointsToAWorkingCopyForThisResource LinkRelationTypes = "working-copy"
	/*
	   Points to the versioned resource from which this
	         working copy was obtained.

	*/
	LinkRelationTypesPointsToTheVersionedResourceFromWhichThisWorkingCopyWasObtained LinkRelationTypes = "working-copy-of"
)
