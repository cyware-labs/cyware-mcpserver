package helpers

const CQL_grammar_rule = `This tool returns all the grammar rules required to build a CQL (Cyware Query Language) query.
In the backend, these CQL queries are used to perform searches.
CQL must be only used when you want to search the list of threat data object based on the condition.
Example query: 'type' = "indicator" AND 'ioc_type' = "ipv4-addr"

Note: Parentheses are optional but can improve query clarity and control evaluation order.

Query without parentheses (relies on default AND precedence):
'type' = "malware" AND 'relationship_type' = "indicates" OR 'type' = "indicator" AND 'related_object' = "report"

Same query with parentheses (explicit grouping):
('type' = "malware" AND 'relationship_type' = "indicates") OR ('type' = "indicator" AND 'related_object' = "report")

Both formats are valid. Parentheses are required for IN, NOT, and RANGE operators to enclose value lists.



LIKE Operator (for fuzzy matching)
The LIKE operator performs case-insensitive fuzzy/partial matching and should be used for source and source_collection fields.

When to use LIKE:
- Always use for 'source' field: source LIKE "Virus Total" will match "VirusTotal", "virus_total", etc.
- Always use for 'source_collection' field: source_collection LIKE "Free Text"

LIKE is case-insensitive and matches partial strings, making it more flexible than exact matching with =.

Examples:
'source' LIKE "Virus"  # Matches "VirusTotal", "virus_feed", "VIRUS", etc.
'source_collection' LIKE "Free Text"  # Matches variations of "Free Text"

CRITICAL: Each field has specific supported operators. NEVER assume a field supports CONTAINS or any other operator without verifying it in that field's "Supported operators" list. Always check the field's documentation below before using any operator.

Below are the grammar rules used to construct a valid CQL:



1. type
Represents the object type in the query.
Supported values, values must be from one of these:
indicator, malware, threat-actor, vulnerability, attack-pattern, campaign, course-of-action, identity, infrastructure, intrusion-set, location, malware-analysis, observed-data, opinion, tool, report, custom-object, observable, incident, note, grouping

Supported operators: =, !=, IN, NOT
Examples:
'type' = "indicator"
'type' != "indicator"
'type' IN ("indicator", "malware")
'type' NOT ("indicator", "threat-actor")

2. ioc_type
Represents the IOC (Indicator of Compromise) type for the SDO.
Usually used when type = "indicator".
Supported values:
artifact, autonomous-system, directory, domain-name, email-addr, email-message, file, ipv4-addr, ipv6-addr, mac-addr, MD5, mutex, network-traffic, process, SHA-1, SHA-224, SHA-256, SHA-384, SHA-512, software, SSDEEP, url, user-account, window-registry-key, x509-certificate, yara

Supported operators: =, !=, IN, NOT
Examples:
'ioc_type' = "ipv4-addr"
'ioc_type' != "ipv4-addr"
'ioc_type' IN ("email-addr", "file")
'ioc_type' NOT ("email-addr", "mutex")

3. source
Represents the source from which threat intelligence is received.
Values must be strings.
Supported operators: =, !=, IN, NOT , ~, LIKE
Examples:
'source' = "import"
'source' != "crowdstrike"
'source' IN ("import", "mandiant")
'source' NOT ("rss", "mail")
'source' LIKE "VirusTotal"
'source' LIKE "Import"

4. value
Represents a value (e.g., title or IOC value) used for free-text search.
Example: If you want to search for domain "abc.com", use value = "abc.com"
Values must be strings.

Supported operators: =, !=, MATCHES, IN, NOT, BEGINS_WITH, ENDS_WITH
Examples:
'value' = "domain.com"
'value' != "domain.com"
'value' MATCHES "dom"
'value' IN ("domain.com", "abc.com")
'value' NOT ("domain.com", "abc.com")
'value' BEGINS_WITH "sd"
'value' ENDS_WITH "sd"

5. is_deprecated
Represents the deprecation status of the indicator, if they are deprecated or not.
Example: If you want to search all the indicator which are deprecated, use type="indicator" AND is_deprecated="true"
If you want to search all the indicator which are not deprecated, use type="indicator" AND is_deprecated="false"
Supported values are "true", "false"

Examples:
'is_deprecated' = "true"
'is_deprecated' = "false"


6. ctix_modified
Represent the exact date, timestamp when the threat is updatd in CTIX application. ctix_modified is applicable to all type SDO(stix domain objects) which exists in CTIX.
Example: If you want to get the data which is modified later than 1746901800000, use ctix_modified >= "1746901800000"
The value must be in string and which is epoch equivalent of the date.
Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'ctix_modified' >= "1746901800000"
'ctix_modified' = "1746901800000"
'ctix_modified' != "1746901800000"
'ctix_modified' <= "1746901800000"
'ctix_modified' < "1746901800000"
'ctix_modified' RANGE ("1746297000000","1746988199000")

7. tag
Represents the tag attached to the threat intel(stix domain objects)
Example: If you want search all the objects having tag malicious then use tag = "malicious"
The value must be in string.
Supported operators: =, !=, IN, NOT

Examples:
'tag' = "malicious"
'tag' != "non-malicious"
'tag' IN ("non-malicious", "malicious")
'tag' NOT ("non-malicious", "raided")

8. tlp
Represent the TLP of the SDO(stix domain objects), It can be RED, AMBER_STRICT, AMBER, GREEN, CLEAR, NONE
Example: If you want search all the objects having tlp red then use tlp = "RED"
The value must be in string.
Supported operators: =, !=, IN, NOT

Examples:
'tlp' = "RED"
'tlp' != "GREEN"
'tlp' IN ("RED", "AMBER_STRICT")
'tlp' NOT ("GREEN", "NONE")


9. published_collection
Represent the STIX collection to which the data is published.
Example: If you want to search all the objects which are published to Collection1, then use published_collection = "Collection1"
The values must be in string.
Supported operators: =, !=, IN, NOT

Examples:
'published_collection' = "Collection1"
'published_collection' != "Collection1"
'published_collection' IN ("Collection1", "Collection2")
'published_collection' NOT ("Collection1", "Collection2")

10. rule

Represent the rules which has run on the threat data object. It will all the actioned object by rule.
You can directly hit the CQL with "rule" = "Rule Name", no need to get the details of the rule until and unless specified by the user.
Example: If you want to search all the threat data objects which are passed by the rules. For example: If you want to search all the threat data objects which are passed by rule PublishToColl then you will do something : "rule" = "PublishToColl"
The values must be in string.
Supported operators: =, !=, IN, NOT

Special value:
To get all objects on which ANY rule has been processed, use 'rule' = "ALL"

Examples:
'rule' = "rule1"
'rule' != "rule1"
'rule' IN ("rule1", "rule2")
'rule' NOT ("rule3", "rule4")
'rule' = "ALL"  # Get all objects processed by any rule


11. enrichment_tool

Represent the enrichment_tool which has enriched the threat data objects.
Example: If you want to search all the threat data objects which are enriched by the enrichment tool AbuseIPDB, then you will do something : "enrichment_tool" = "AbuseIPDB"
The values must be in string.
Supported operators: =, !=, IN, NOT

Examples:
'enrichment_tool' = "AbuseIPDB"
'enrichment_tool' != "AbuseIPDB"
'enrichment_tool' IN ("AbuseIPDB", "Alien Vault")
'enrichment_tool' NOT ("AbuseIPDB", "Alien Vault")

If you want to search whether ip 1.1.1.1 enriched by tool AbuseIPDB then use "value" = "1.1.1.1" AND "enrichment_tool" = "AbuseIPDB"

12. enriched_status

Represent the enrichment status of the threat data object, which means if it tells the object which are enriched.
Example: If you want to search all the threat data objects which are enriched, then you will do something : "enriched_status" = "1"
The values must be in string and they are fixed. Use 1 for ENRICHED, 2 for Tried and Failed and 3 for Quota Completed
Supported operators: =, !=, IN, NOT

Examples:
'enriched_status' = "1"
'enriched_status' != "2"
'enriched_status' IN ("1", "2")
'enriched_status' NOT ("3", "2")

13. relationship_type
This will give you all the object which are having this relation_type in any of its relations. For example user searches for
"relationship_type = "related-to" it means give all the objects which are having realtion type as related-to with the related object.
"type"= "indicator" AND "relationship_type = "related-to" --> It means give all the indicator having relation type related-to with the related object.

The values must be in string, and name should be in the available relation types list.

Supported operators: =, !=, IN, NOT

Examples:
'relationship_type' = "related-to"
'relationship_type' != "related-to"
'relationship_type' IN ("related-to", "associated_actor",  "authored-by")
'relationship_type' NOT ("related-to", "associated_actor",  "authored-by")


14. related_object
This will give you all the object which are related to the specified object type. For example user searches for
"related_object = "malware" it means give all the objects which are having relation with malware object type.
"type"= "indicator" AND "related_object = "malware" --> It means give all the indicator having relation with malware.

The values must be in string, and should be a valid object type. you can reference "type" to see the values.

Supported operators: =, !=, IN, NOT

Examples:
'related_object' = "indicator"
'related_object' != "indicator"
'related_object' IN ("malware", "indicator",  "report")
'related_object' NOT ("malware", "indicator",  "report")

15. related_object_value
This represents the value for the related object of a object. Please note As its a related object property, so it must be followed by related_object field.
The value of this must be a string value
Incorrect query ❌:  related_object_value = "domain.com"
Correct query ✅ : related_object = "indicator" and related_object_value = "domain.com"

Supported operators: =, !=, CONTAINS, IN, NOT, BEGINS_WITH, ENDS_WITH, MATCHES
Examples:
'related_object_value' = "domain.com"
'related_object_value' != "domain.com"
'related_object_value' CONTAINS "dom"
'related_object_value' IN ("domain.com", "abc.com")
'related_object_value' NOT ("domain.com", "abc.com")
'related_object_value' BEGINS_WITH "sd"
'related_object_value' ENDS_WITH "sd"

16. related_object_property
This represents the fields/properties for the related object, so it can search/filter data based on the related object as well.
Please note As its a related object property, so it must be followed by related_object field.
There are selected fields only which can be fetched from related object. List of fields are 'type', 'value', 'source'. These fields are already defined in the CQL, so supported operators and usages are same.
Syntax to fetch the fiels is related_object_property.fieldName.

Supported operators are based on fields type which is used.

Example queries -
'related_object_property.source' = "alien vault"
'related_object_property.type' = "indicator"
'related_object_property.value' CONTAINS "has"



17. has_relations

This will give you all the objects which has atleast one relation with other object, if has_relations = "true".
has_relations = "false" --> Gives the list of objects which doesn't have any relation with any object.

The values must be in string.

Supported operators: =, !=

Examples:
'has_relations' = "true"
'has_relations' != "false"

18. enrichment_verdict

This indicates the enrichment verdict of the threat data objects. It must values from 'Malicious' or 'Non malicious' depending on the query. Dont' start fetching the indicator enrichment details until unless explicitly asked.
The values must be in string.

Supported operators: =, !=, IN, NOT

Examples:
'enrichment_verdict' = "Malicious"
'enrichment_verdict' != "Malicious"
'enrichment_verdict' IN ("Malicious", "Non malicious")
'enrichment_verdict' NOT ("Malicious", "Non malicious")

19. ctix_created
Represent the exact date, timestamp when the threat is created in CTIX application. ctix_created is applicable to all type SDO(stix domain objects) which exists in CTIX.
Example: If you want to get the data which is created later than 1746901800000, use ctix_created >= "1746901800000"
The value must be in string and which is epoch equivalent of the date.
Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'ctix_created' > "1746901800000"
'ctix_created' >= "1746901800000"
'ctix_created' = "1746901800000"
'ctix_created' != "1746901800000"
'ctix_created' <= "1746901800000"
'ctix_created' < "1746901800000"
'ctix_created' RANGE ("1746297000000","1746988199000")

20. confidence_score
Represet the confidence score(also knowns as risk score) of the threat data object.
Example: If you want to get the data having risk score greater than 75, use confidence_score > "75"
The value must be in string and must be between 0 to 100 inclusive.
Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'confidence_score' > "75"
'confidence_score' >= "75"
'confidence_score' = "75"
'confidence_score' != "75"
'confidence_score' <= "75"
'confidence_score' < "75"
'confidence_score' RANGE ("75","90")
—


21.Source Type:
Represents the source type from which threat intelligence is received.
Values must be strings.


Supported operators: =, !=, IN, NOT

Supported Values : CUSTOM_STIX_SOURCES, API_FEEDS, EMAIL_ACCOUNTS, RSS_FEED, TWITTER_FEEDS, WEB_SCRAPPER, MALWARE_SANDBOX, MISCELLANEOUS








Example :

If you want to search for all indicators received from APIs, use 'type' = "indicator" AND 'source_type' = "API_FEEDS"
If you want to search for all indicators received from STIX, use 'type' = "indicator" AND 'source_type' = "CUSTOM_STIX_SOURCES"
If you want to search for all indicators received from Email, use 'type' = "indicator" AND 'source_type' = "EMAIL_ACCOUNTS"
If you want to search for all indicators received from X (Twitter), use 'type' = "indicator" AND 'source_type' = "TWITTER_FEEDS"
If you want to search for all indicators received from Webscraper, use 'type' = "indicator" AND 'source_type' = "WEB_SCRAPPER"
If you want to search for all indicators received from Sandbox, use 'type' = "indicator" AND 'source_type' = "MALWARE_SANDBOX"
If you want to search for all indicators received from Miscellaneous, use 'type' = "indicator" AND 'source_type' = "MISCELLANEOUS"



Examples :
'source_type' = "API_FEEDS"
'source_type' != "RSS_FEED"
'source_type' IN ("API_FEEDS", "RSS_FEED")
'source_type' NOT ("EMAIL_ACCOUNTS", "TWITTER_FEEDS")

22. Custom Object Type:


Represents the custom object type for custom objects in the system. Custom objects are user-defined STIX object types beyond the standard STIX types.

Supported operators: =, !=, IN, NOT

Example: If you want to search for custom objects of type "my-custom-threat", use custom_object_type = "x-my-custom-threat"


Examples :
'custom_object_type' = "x-my-custom-threat"
'custom_object_type' != "x-my-custom-threat"
'custom_object_type' IN ("x-custom-threat-1", "x-custom-threat-2")
'custom_object_type' NOT ("x-custom-threat-1", "x-custom-threat-2")





23. Source Collections :


Represents the source collection ID from which the threat data object was ingested. Collections are sub-groups within a source.
Example: If you want to search for all threat data from a specific collection, use source_collection = "default"

Supported operators: =, !=, IN, NOT, ~, LIKE

Example :
source_collection = "Amazon GuardDuty"
source_collection = "api_testing_source_darkreading_edited"



24.Source Confidence:
Represents the confidence level assigned by the source. This is a categorical representation of the source's confidence score.

Supported Values : "HIGH","LOW", "MEDIUM", "NONE"

Example: If you want to search for all indicators with HIGH source confidence, use 'type' = "indicator" AND 'source_confidence' = "HIGH"


'source_confidence' = "HIGH"
'source_confidence' != "NONE"
'source_confidence' IN ("HIGH", "MEDIUM")
'source_confidence' NOT ("LOW", "NONE")







25.Source Confidence Value :

Represents the numeric confidence score value assigned by the source (0-100).
Example: If you want to search for threat data with source confidence value greater than 75, use source_confidence_value > "75"
The value must be in string format and must be between 0 to 100 inclusive.
Supported operators: =, !=, >, >=, <, <=, RANGE



Examples :
'source_confidence_value' > "75"
'source_confidence_value' >= "50"
'source_confidence_value' = "100"
'source_confidence_value' != "0"
'source_confidence_value' <= "30"
'source_confidence_value' < "25"
'source_confidence_value' RANGE ("50", "80")


26.Published Date

Represents the exact date and timestamp when the threat data object was published to a STIX collection. This field tracks when an object was shared/published from CTIX to external TAXII servers or internal collections.

This field is applicable to all types of threat data objects (SDOs and indicators) that have been published to at least one collection. Objects that have never been published will not have this field.

Example Use Cases:
- If you want to find all threat data published after a specific date, use published_on >= "1746901800000"
- If you want to find threat data published within a specific time range, use published_on RANGE ("1746297000000","1746988199000")
- If you want to find threat data published on an exact date, use published_on = "1746901800000"

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) equivalent of the publication date.

Note: This field is different from "published_collection" which identifies WHICH collection the data was published to, while "published_on" identifies WHEN it was published.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'published_on' > "1746901800000"
'published_on' >= "1746901800000"
'published_on' = "1746901800000"
'published_on' != "1746901800000"
'published_on' <= "1746901800000"
'published_on' < "1746901800000"
'published_on' RANGE ("1746297000000","1746988199000")



27.Imported File
Represents the name of the file from which IOCs (Indicators of Compromise) were imported through the IOC Lookup feature. This field allows you to track and filter threat data based on the source file used during bulk IOC imports.

This field is specifically applicable to indicators that were imported via the IOC Lookup/Bulk Import functionality. It stores the original filename (e.g., "malicious_ips.csv", "threat_indicators.txt") used during the import process.

Example Use Cases:
- If you want to find all IOCs imported from a specific file, use imported_file = "malicious_domains.csv"
- If you want to find IOCs from multiple import files, use imported_file IN ("file1.csv", "file2.txt")
- If you want to exclude IOCs from a particular import file, use imported_file != "test_data.csv"

The value must be in string format and represents the exact filename as it was uploaded/imported into CTIX.

Note: This field is typically used in combination with "imported_file_on" to identify both the file and when it was imported. Objects that were not imported via the IOC Lookup feature will not have this field.

Supported operators: =, !=, CONTAINS, BEGINS_WITH, ENDS_WITH, IN, NOT, MATCHES

Examples:
'imported_file' = "iocs_2024.csv"
'imported_file' != "test_file.txt"
'imported_file' IN ("threat_feed_jan.csv", "threat_feed_feb.csv")
'imported_file' NOT ("old_data.csv", "deprecated_iocs.txt")
'imported_file' CONTAINS "2024"
'imported_file' BEGINS_WITH "threat_feed"



28.Imported File Date


Represents the exact date and timestamp when the IOC file was imported into CTIX through the IOC Lookup/Bulk Import feature. This field tracks the import operation timestamp, not when the IOCs were created or discovered.

This field is applicable only to indicators that were imported via the IOC Lookup functionality. It records when the bulk import process was executed, allowing you to track and audit import operations over time.

Example Use Cases:
- If you want to find all IOCs imported after a specific date, use imported_file_on >= "1746901800000"
- If you want to find IOCs imported within a specific time window, use imported_file_on RANGE ("1746297000000","1746988199000")
- If you want to find IOCs imported on a particular day, use imported_file_on = "1746901800000"

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) of when the import operation occurred.

Note: This field should typically be used together with "imported_file" to get complete import context. The timestamp represents the server time when the file was processed, not the file's creation or modification time.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'imported_file_on' > "1746901800000"
'imported_file_on' >= "1746901800000"
'imported_file_on' = "1746901800000"
'imported_file_on' != "1746901800000"
'imported_file_on' <= "1746901800000"
'imported_file_on' < "1746901800000"
'imported_file_on' RANGE ("1746297000000","1746988199000")

29. Valid From

Represents the start date and timestamp from which an indicator is considered valid or active. This is a standard STIX property that defines the beginning of an indicator's validity period, indicating when the indicator first became relevant or observable.

This field is primarily applicable to indicators and represents the timestamp when the indicator should be considered valid for detection and analysis purposes. It's part of the indicator lifecycle management in STIX.

Example Use Cases:
- If you want to find indicators that became valid after a certain date, use valid_from >= "1746901800000"
- If you want to find indicators valid within a specific time period, use valid_from RANGE ("1746297000000","1746988199000")
- If you want to find recently activated indicators, use valid_from > "1746901800000"

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) of when the indicator's validity period begins.

Note: This field is different from "ctix_created" which represents when the indicator was ingested into CTIX. An indicator can be created in CTIX on one date but have a "valid_from" date in the past or future. The "valid_from" field is often used with "valid_until" to define a complete validity window for temporal indicator analysis.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'valid_from' > "1746901800000"
'valid_from' >= "1746901800000"
'valid_from' = "1746901800000"
'valid_from' != "1746901800000"
'valid_from' <= "1746901800000"
'valid_from' < "1746901800000"
'valid_from' RANGE ("1746297000000","1746988199000")

30. Valid Until

Represents the end date and timestamp until which an indicator is considered valid or active. This is a standard STIX property that defines the expiration of an indicator's validity period, indicating when the indicator should no longer be considered relevant for detection.

This field is primarily applicable to indicators and represents the timestamp when the indicator's validity expires. After this date, the indicator may be considered stale or outdated for active threat detection, though it may still have historical value.

Example Use Cases:
- If you want to find indicators that will expire soon, use valid_until <= "1746901800000"
- If you want to find currently valid indicators (not yet expired), use valid_until > "current_timestamp"
- If you want to find indicators with a specific expiration range, use valid_until RANGE ("1746297000000","1746988199000")

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) of when the indicator's validity period ends.

Note: This field is different from "ctix_modified" which tracks when the indicator was last updated in CTIX. The "valid_until" represents the semantic validity of the indicator itself. Indicators without a "valid_until" date may be considered valid indefinitely. This field is commonly used with "valid_from" to filter indicators based on their active validity window for temporal threat analysis.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'valid_until' > "1746901800000"
'valid_until' >= "1746901800000"
'valid_until' = "1746901800000"
'valid_until' != "1746901800000"
'valid_until' <= "1746901800000"
'valid_until' < "1746901800000"
'valid_until' RANGE ("1746901800000","1746901810000")


31.Tag Category

Represents the category or type of tag applied to threat data objects. Tags in CTIX are classified into different types based on who created them and their access control level. This field allows you to filter objects based on the tag category.

Tag types are used to organize and control access to tags within the system. Different tag types have different permission levels and visual representations in the UI.

Supported Values (case-sensitive):
- "user" - Tags created by end users
- "source" - Tags automatically applied based on the ingestion source
- "system" - System-generated tags (prefixed with "__")
- "privileged" - Tags with restricted access based on user groups
- "group" - Tag groups that contain multiple associated tags

Example Use Cases:
- If you want to find all objects tagged with user-created tags, use tag_type = "USER"
- If you want to find objects with source-assigned tags, use tag_type = "SOURCE"
- If you want to exclude system tags from results, use tag_type != "SYSTEM"

The value must be in string format and must match one of the supported tag type values exactly.

Note: This field works in conjunction with the "tag" field. You typically use both together to filter by specific tags of a certain type. System tags are automatically prefixed with "__" in their names.

Supported operators: =

Examples:
'tag_type' = "user"



32.Analyst Score

Represents the risk score manually assigned by a security analyst to a threat data object. This is a subjective assessment that allows analysts to override or supplement the automatically calculated confidence score based on their expertise and contextual knowledge.

The analyst score is an integer value that represents the analyst's assessment of how dangerous or critical a particular threat is. It's independent of the system-calculated confidence_score and allows for human judgment in threat assessment.

Example Use Cases:
- If you want to find high-priority threats marked by analysts, use analyst_score > "75"
- If you want to find objects that analysts have scored, use analyst_score != "0"
- If you want to find analyst-reviewed threats in a specific score range, use analyst_score RANGE ("50", "100")

The value must be in string format and typically ranges from 0 to 100, where higher values indicate higher risk or criticality as assessed by the analyst.

Note: This field is different from "confidence_score" (system-calculated risk score). The analyst_score provides human expertise input, while confidence_score is algorithmically determined. Objects without analyst review may not have this field set or may have a value of 0.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'analyst_score' > "75"
'analyst_score' >= "50"
'analyst_score' = "100"
'analyst_score' != "0"
'analyst_score' <= "30"
'analyst_score' < "25"
'analyst_score' RANGE ("50", "80")

Combined query example:
'type' = "indicator" AND 'analyst_score' >= "80" AND 'confidence_score' >= "70"
This will fetch all indicators with both high analyst assessment and high system-calculated confidence.

33.Analyst CVSS Score

Represents the CVSS (Common Vulnerability Scoring System) score manually assigned by a security analyst to a threat data object, particularly vulnerabilities. This allows analysts to provide or override CVSS scores based on their assessment of the vulnerability's severity in the organization's specific context.

The analyst CVSS score follows the CVSS v3.x scoring standard, ranging from 0.0 to 10.0, where higher values indicate more severe vulnerabilities. This is particularly useful for vulnerability objects where analysts may adjust the score based on organizational impact.

CVSS Score Ranges (for reference):
- 0.0: None
- 0.1 - 3.9: Low severity
- 4.0 - 6.9: Medium severity
- 7.0 - 8.9: High severity
- 9.0 - 10.0: Critical severity

Example Use Cases:
- If you want to find critical vulnerabilities assessed by analysts, use analyst_cvss_score >= "9.0"
- If you want to find analyst-reviewed vulnerabilities, use 'type' = "vulnerability" AND 'analyst_cvss_score' > "0"
- If you want to find high to critical severity issues, use analyst_cvss_score RANGE ("7.0", "10.0")

The value must be in string format and represents a decimal number between 0.0 and 10.0.

Note: This field is primarily applicable to vulnerability objects but may be used for other object types where CVSS scoring is relevant. It differs from automated CVSS scores that may come from sources, as this represents the analyst's expert assessment.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'analyst_cvss_score' > "7.0"
'analyst_cvss_score' >= "9.0"
'analyst_cvss_score' = "10.0"
'analyst_cvss_score' != "0"
'analyst_cvss_score' <= "4.0"
'analyst_cvss_score' < "7.0"
'analyst_cvss_score' RANGE ("4.0", "6.9")

Combined query example:
'type' = "vulnerability" AND 'analyst_cvss_score' >= "7.0" AND 'tag' = "production_systems"
This will fetch all vulnerabilities with high/critical CVSS scores assigned by analysts that are tagged as affecting production systems.

34.Countries
Represents the country or countries associated with a threat data object. This field tracks geographic information related to threats, such as the country of origin for threat actors, the location of infrastructure, or the target countries for campaigns.

The countries field stores ISO country codes or country names associated with the threat intelligence. This geographic attribution helps in understanding the geopolitical context of threats and filtering threats relevant to specific regions.

Example Use Cases:
- If you want to find threats originating from or associated with a specific country, use countries = "United Arab Emirates"
- If you want to find threats from multiple countries, use countries IN ("Andorra", "United Arab Emirates", "Albania")
- If you want to exclude threats from certain countries, use countries NOT ("Albania", "Andorra")

The value must be in string format and typically uses full country names.

Note: This field is particularly relevant for threat-actor, infrastructure, and campaign objects where geographic attribution is important. Multiple countries can be associated with a single object if the threat has presence or activities in multiple locations. The field may also be used in risk scoring calculations based on geographic risk profiles.

Supported operators: =, !=, IN, NOT

Examples:
'countries' = "United Arab Emirates"
'countries' != "Andorra"
'countries' IN ("Andorra", "United Arab Emirates", "Albania")
'countries' NOT ("Albania", "Andorra")

Combined query example:
'type' = "threat-actor" AND 'countries' IN ("CN", "RU") AND 'first_seen' >= "1704067200000"
This will fetch all threat actors associated with China or Russia that were first observed after January 1, 2024.

35.First Seen

Represents the exact date and timestamp when a threat was first observed or detected. This field is primarily applicable to malware, threat-actors, and campaigns, providing temporal context about when the threat first appeared in the wild or was first identified.

The first_seen timestamp helps track the emergence and evolution of threats over time. It's particularly valuable for understanding threat timelines, identifying new versus established threats, and analyzing threat trends.

Example Use Cases:
- If you want to find recently emerged threats, use first_seen >= "1704067200000" (threats first seen after Jan 1, 2024)
- If you want to find threats first observed within a specific time window, use first_seen RANGE ("1704067200000", "1735689600000")
- If you want to find established threats (observed long ago), use first_seen < "1640995200000" (before 2022)

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) of when the threat was first observed.

Note: This field is different from "ctix_created" which represents when the object was ingested into CTIX. An object can be first seen in the wild months or years before being ingested into your CTIX instance. This field is typically populated from source data and is particularly relevant for malware, threat-actor, and campaign object types. It's often used with "last_seen" to define the active period of a threat.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'first_seen' > "1704067200000"
'first_seen' >= "1704067200000"
'first_seen' = "1704067200000"
'first_seen' != "1704067200000"
'first_seen' <= "1640995200000"
'first_seen' < "1640995200000"
'first_seen' RANGE ("1704067200000", "1735689600000")


36.Last Seen
Represents the exact date and timestamp when a threat was last observed or detected in the wild. This field is primarily applicable to malware, threat-actors, and campaigns, providing temporal context about the most recent activity or sighting of the threat.

The last_seen timestamp helps track threat activity timelines and identify whether threats are currently active or have become dormant. This is crucial for prioritizing response efforts and understanding threat persistence.

Example Use Cases:
- If you want to find currently active threats, use last_seen >= "1735689600000" (active in 2025)
- If you want to find threats that haven't been seen recently (potentially dormant), use last_seen < "1704067200000"
- If you want to find threats with activity in a specific time window, use last_seen RANGE ("1704067200000", "1735689600000")

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) of when the threat was most recently observed.

Note: This field is different from "ctix_modified" which tracks when the object was last updated in CTIX. The last_seen represents actual threat activity in the wild. This field is typically populated from source data and is particularly relevant for malware, threat-actor, and campaign object types. It's often used with "first_seen" to understand the threat's active period and assess if it's a current or historical threat.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'last_seen' > "1735689600000"
'last_seen' >= "1735689600000"
'last_seen' = "1735689600000"
'last_seen' != "1704067200000"
'last_seen' <= "1704067200000"
'last_seen' < "1704067200000"
'last_seen' RANGE ("1704067200000", "1735689600000")

Combined query example:
'type' = "malware" AND 'first_seen' <= "1704067200000" AND 'last_seen' >= "1735689600000"
This will fetch all malware that was first seen before 2024 but is still active in 2025, indicating persistent threats.


37.False Positive Status
Represents whether a threat data object (typically an indicator) has been marked as a false positive by security analysts. This boolean flag helps filter out indicators that have been determined to be benign or incorrectly classified as malicious.

False positive marking is crucial for maintaining the quality of threat intelligence and preventing unnecessary blocking or alerting on legitimate entities. Analysts mark indicators as false positives when investigation reveals they are not actually threats.

Supported Values:
- "true" - Object is marked as a false positive
- "false" - Object is not marked as a false positive (default/normal threat intel)

Example Use Cases:
- If you want to exclude false positives from your search, use is_false_positive = "false"
- If you want to review objects marked as false positives, use is_false_positive = "true"
- If you want to find all valid indicators (not false positives), use 'type' = "indicator" AND 'is_false_positive' = "false"

The value must be in string format ("true" or "false").

Note: This field is primarily used with indicator objects. When an indicator is marked as a false positive, it's typically excluded from active detection and blocking rules. This helps reduce alert fatigue and improves the signal-to-noise ratio of security operations. The action is reversible - analysts can unmark false positives if they determine an indicator is actually malicious.

Supported operators: =, !=

Examples:
'is_false_positive' = "true"
'is_false_positive' = "false"
'is_false_positive' != "true"
'is_false_positive' != "false"


38.Review Status
Represents whether a threat data object has been reviewed and validated by a security analyst. This boolean flag indicates that an analyst has examined the object and confirmed its relevance or accuracy.

The review status helps distinguish between automated threat intelligence (unreviewed) and analyst-validated intelligence (reviewed). Reviewed objects typically have higher trustworthiness as they've undergone human verification.

Supported Values:
- "true" - Object has been reviewed by an analyst
- "false" - Object has not been reviewed (default state)

Example Use Cases:
- If you want to find only analyst-reviewed threats, use is_reviewed = "true"
- If you want to find objects pending review, use is_reviewed = "false"
- If you want to prioritize reviewed high-confidence threats, use is_reviewed = "true" AND confidence_score >= "80"

The value must be in string format ("true" or "false").

Note: The review process involves an analyst examining the threat data, validating its accuracy, and potentially adding additional context or scores. Objects marked as reviewed are considered more reliable and actionable. This field is distinct from "is_under_review" which indicates objects currently being reviewed.

Supported operators: =, !=

Examples:
'is_reviewed' = "true"
'is_reviewed' = "false"
'is_reviewed' != "true"
'is_reviewed' != "false"


39.Revoke Status
Represents whether a threat data object has been revoked or invalidated. Revoked objects are those that were previously considered valid threat intelligence but have since been determined to be outdated, incorrect, or no longer applicable.

Revocation is used to mark threat intelligence that should no longer be used for detection or decision-making, without completely deleting it from the system. This maintains an audit trail while preventing the use of outdated intelligence.

Supported Values:
- "true" - Object has been revoked/invalidated
- "false" - Object is currently valid and active (default state)

Example Use Cases:
- If you want to find only active, valid threats, use is_revoked = "false"
- If you want to review revoked objects, use is_revoked = "true"
- If you want to exclude revoked intelligence from your analysis, use is_revoked != "true"

The value must be in string format ("true" or "false").

Note: Revoked objects are typically excluded from active threat detection and sharing. This is particularly important for indicators that were valid at one time but are no longer accurate (e.g., a malicious IP that is now used for legitimate purposes). The STIX standard includes revocation as a way to update threat intelligence without losing historical context.

Supported operators: =, !=

Examples:
'is_revoked' = "true"
'is_revoked' = "false"
'is_revoked' != "true"
'is_revoked' != "false"



40.Manual Review
Represents whether a threat data object is currently under manual review by a security analyst. This boolean flag indicates that the object has been flagged for analyst examination but the review process has not been completed yet.

Objects under review are in a pending state where analysts are actively investigating or validating the threat intelligence. This helps teams track their review queue and identify objects requiring attention.

Supported Values:
- "true" - Object is currently under manual review
- "false" - Object is not under review (default state)

Example Use Cases:
- If you want to find your review queue, use is_under_review = "true"
- If you want to find objects that need review assignment, use is_reviewed = "false" AND is_under_review = "false"
- If you want to prioritize high-score threats pending review, use is_under_review = "true" AND confidence_score >= "80"

The value must be in string format ("true" or "false").

Note: This field represents an intermediate state in the review workflow. Objects can be marked for manual review either automatically (by rules) or manually (by analysts). Once review is completed, the object transitions to "is_reviewed = true" and "is_under_review" is set back to false. This workflow helps teams manage their analyst workload and track review progress.

Supported operators: =, !=

Examples:
'is_under_review' = "true"
'is_under_review' = "false"
'is_under_review' != "true"
'is_under_review' != "false"


41.Included in Allowed Indicators
Represents whether an indicator has been added to the allowed/whitelist. Whitelisted indicators are those explicitly marked as safe or trusted, typically to prevent false positive alerts or to exclude known-good entities from detection rules.

Whitelisting (also called "allow-listing") is crucial for managing false positives and ensuring that legitimate business activities or trusted entities are not flagged as threats. This improves the accuracy of threat detection.

Supported Values:
- "true" - Indicator is whitelisted/allowed
- "false" - Indicator is not whitelisted (default state)

Example Use Cases:
- If you want to find whitelisted indicators to review your allow list, use is_whitelisted = "true"
- If you want to ensure you're only processing non-whitelisted threats, use is_whitelisted = "false"
- If you want to find potentially over-whitelisted items, use is_whitelisted = "true" AND confidence_score >= "75"

The value must be in string format ("true" or "false").

Note: This field is primarily applicable to indicator objects. Whitelisted indicators are typically excluded from blocking actions and detection rules, though they remain in the system for audit and review purposes. Organizations often whitelist their own infrastructure, partner domains, or known false positives. When combined with "whitelist_reason" (not in CQL but in metadata), it provides context for why an indicator was allowed.

Supported operators: =, !=

Examples:
'is_whitelisted' = "true"
'is_whitelisted' = "false"
'is_whitelisted' != "true"
'is_whitelisted' != "false"


42.Actioned By
Represents the user or system entity that performed an action on a threat data object. This field tracks who made changes, executed rules, or performed operations on the threat intelligence, providing accountability and audit trail information.

The actioned_by field stores the username, user ID, or system identifier (e.g., "admin", "analyst@example.com", "automation_system") that initiated the action. This is essential for compliance, auditing, and understanding the provenance of threat intelligence modifications.

Example Use Cases:
- If you want to find all actions performed by a specific analyst, use actioned_by = "john.doe@example.com"
- If you want to audit automation actions, use actioned_by = "automation_system"
- If you want to find objects modified by multiple specific users, use actioned_by IN ("analyst1", "analyst2")

The value must be in string format and typically contains the username or user identifier.

Note: This field is part of the action tracking system and is stored in the CQL Actions index. It works in conjunction with "action_type" (manual vs automatic), "actioned_on" (timestamp), and "action_name" to provide complete audit trail information. System-generated actions may have system identifiers like "SYSTEM" or "automation_engine" as the actioned_by value.

Supported operators: =, !=, IN, NOT

Examples:
'actioned_by' = "analyst@example.com"
'actioned_by' != "system"
'actioned_by' IN ("admin", "security_lead")
'actioned_by' NOT ("test_user", "deprecated_user")



43.Actioned Date
Represents the exact date and timestamp when an action was performed on a threat data object. This field provides temporal context for all operations, helping track when changes were made, rules were executed, or analysts interacted with the threat intelligence.

The actioned_on timestamp is automatically recorded whenever any action (manual or automatic) is performed on an object. This creates an audit trail and helps understand the timeline of threat intelligence handling.

Example Use Cases:
- If you want to find recent actions, use actioned_on >= "1735689600000" (actions in 2025)
- If you want to audit actions within a specific timeframe, use actioned_on RANGE ("1704067200000", "1735689600000")
- If you want to find objects recently modified by analysts, use action_type = "manual" AND actioned_on >= "1735689600000"

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) when the action was performed.

Note: This field is part of the action tracking system in the CQL Actions index. It differs from "ctix_modified" which tracks when the object itself was last modified - actioned_on tracks when any action (including non-modifying actions like running rules or adding to watchlists) was performed. Multiple actions can be performed on a single object, each with its own actioned_on timestamp.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'actioned_on' > "1735689600000"
'actioned_on' >= "1735689600000"
'actioned_on' = "1735689600000"
'actioned_on' != "1704067200000"
'actioned_on' <= "1704067200000"
'actioned_on' < "1704067200000"
'actioned_on' RANGE ("1704067200000", "1735689600000")


44.Action Medium
Represents the type or category of action that was performed on a threat data object. This field distinguishes between manual actions (performed by human analysts) and automatic actions (performed by automation rules or system processes).

Understanding action type is crucial for auditing workflows, evaluating automation effectiveness, and distinguishing human decisions from automated processing in your threat intelligence operations.

Supported Values:
- "manual" - Action was performed manually by a human analyst
- "automatic" - Action was performed automatically by rules, workflows, or system processes

Example Use Cases:
- If you want to find analyst-initiated actions, use action_type = "manual"
- If you want to audit automated rule executions, use action_type = "automatic"
- If you want to find manual reviews of high-confidence threats, use action_type = "manual" AND confidence_score >= "80"

The value must be in string format and must be either "manual" or "automatic".

Note: This field is part of the action tracking system in the CQL Actions index. Manual actions might include analyst reviews, tag additions, score updates, or whitelist additions. Automatic actions typically come from rule engine executions, automated enrichment workflows, or integration responses. Combined with "actioned_by" and "actioned_on", this provides complete context about how and when threat intelligence was processed.

Supported operators: =, !=, IN, NOT

Examples:
'action_type' = "manual"
'action_type' = "automatic"
'action_type' != "manual"
'action_type' IN ("manual", "automatic")
'action_type' NOT ("automatic")

45.Actioned App Type

Represents the type or category of application that performed an action on a threat data object. This field distinguishes between actions performed by the CTIX platform itself and actions performed by third-party integrations or external tools.

Understanding the application type helps track which systems are interacting with your threat intelligence, evaluate integration effectiveness, and audit external tool actions versus internal CTIX operations.

Supported Values:
- "ctix" - Action was performed by the CTIX platform (internal operations)
- "third_party" - Action was performed by external integrations or third-party tools

Example Use Cases:
- If you want to find CTIX internal actions, use app_type = "ctix"
- If you want to audit third-party integration actions, use app_type = "third_party"
- If you want to find external tool actions on high-priority threats, use app_type = "third_party" AND confidence_score >= "80"

The value must be in string format and must be either "ctix" or "third_party".

Note: This field is part of the action tracking system in the CQL Actions index and works with "app_name" which provides the specific application name (e.g., "CTIX", "QRadar", "Splunk", "PaloAlto"). The app_type categorizes actions at a high level, while app_name provides granular details. This helps organizations understand how different tools in their security stack are interacting with threat intelligence.

Supported operators: =, !=, IN, NOT

Examples:
'app_type' = "ctix"
'app_type' = "third_party"
'app_type' != "ctix"
'app_type' IN ("ctix", "third_party")
'app_type' NOT ("third_party")


46.Actioned App

Represents the specific name of the application or tool that performed an action on a threat data object. This field provides granular identification of which system executed the action, whether it's CTIX itself or a third-party integration.

The app_name field stores the actual application identifier (e.g., "CTIX", "QRadar", "Splunk", "PaloAlto Firewall", "AbuseIPDB", "FortiGate") that interacted with the threat intelligence. This enables detailed tracking and auditing of which specific tools in your security ecosystem are processing threat data.

Example Use Cases:
- If you want to find all actions performed by CTIX platform, use app_name = "CTIX"
- If you want to audit specific SIEM integration actions, use app_name = "QRadar"
- If you want to find actions from multiple specific tools, use app_name IN ("Splunk", "ArcSight")

The value must be in string format and represents the application name as configured in the system.

Note: This field works in conjunction with "app_type" which categorizes apps as "ctix" or "third_party". While app_type provides high-level categorization, app_name gives specific application details. This is particularly useful for organizations with multiple integrations to understand which tools are actively processing their threat intelligence.

Supported operators: =, !=, IN, NOT

Examples:
'app_name' = "CTIX"
'app_name' != "test_app"
'app_name' IN ("QRadar", "Splunk", "FortiSIEM")
'app_name' NOT ("deprecated_tool", "old_integration")

Combined query example:
'type' = "indicator" AND 'app_name' = "QRadar" AND 'action_type' = "automatic" AND 'actioned_on' >= "1735689600000"
This will fetch all indicators where QRadar performed automatic actions in 2025.

47.Relation Created Date
Represents the exact date and timestamp when a relationship between two threat data objects was created in CTIX. This field tracks when connections between objects (like indicator-to-malware, threat-actor-to-campaign) were first established.

The relation_created timestamp is crucial for understanding the evolution of threat intelligence and how different threat entities become connected over time. It helps track when analysts or automated systems discovered or documented relationships between threat objects.

Example Use Cases:
- If you want to find recently discovered relationships, use relation_created >= "1735689600000"
- If you want to audit relationships created within a specific timeframe, use relation_created RANGE ("1704067200000", "1735689600000")
- If you want to find old relationships that may need review, use relation_created < "1640995200000"

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) when the relationship was created.

Note: This field is specific to relationship objects (STIX SROs - STIX Relationship Objects). It differs from "ctix_created" which applies to all objects - relation_created specifically tracks when the relationship itself was documented. This is useful for temporal analysis of threat intelligence connections and understanding when threat campaigns became linked to specific actors or techniques.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'relation_created' > "1735689600000"
'relation_created' >= "1735689600000"
'relation_created' = "1735689600000"
'relation_created' != "1704067200000"
'relation_created' <= "1704067200000"
'relation_created' < "1704067200000"
'relation_created' RANGE ("1704067200000", "1735689600000")



48.Relation Modified Date
Represents the exact date and timestamp when a relationship between two threat data objects was last modified in CTIX. This field tracks when existing connections between objects were updated or edited.

The relation_modified timestamp helps track changes to relationship properties, such as updates to relationship confidence, description changes, or metadata modifications. This is important for maintaining an audit trail of how threat intelligence relationships evolve.

Example Use Cases:
- If you want to find recently updated relationships, use relation_modified >= "1735689600000"
- If you want to audit relationship changes within a timeframe, use relation_modified RANGE ("1704067200000", "1735689600000")
- If you want to find stale relationships (not updated recently), use relation_modified < "1704067200000"

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) when the relationship was last modified.

Note: This field applies specifically to relationship objects. It differs from "ctix_modified" on related objects - relation_modified tracks changes to the relationship itself, not the source or target objects. Relationships can be modified when analysts update confidence scores, add contextual information, or adjust relationship types. This field helps track the maintenance and refinement of threat intelligence connections over time.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'relation_modified' > "1735689600000"
'relation_modified' >= "1735689600000"
'relation_modified' = "1735689600000"
'relation_modified' != "1704067200000"
'relation_modified' <= "1704067200000"
'relation_modified' < "1704067200000"
'relation_modified' RANGE ("1704067200000", "1735689600000")

49-51. Custom Scores (Composite Field)

Custom Scores allow organizations to define and track custom scoring metrics for threat data objects. These fields MUST be used together in a specific format as they form a composite query structure.

**Required Query Structure:**
custom_score_name = "score_name" AND custom_score_type = "type" AND custom_score_value [operator] "value"

**Field Components:**

**custom_score_name**: The name/identifier of the custom score attribute
- Supported operators: =, IN
- Example values: "severity_level", "business_impact", "remediation_priority"

**custom_score_type**: The data type of the custom score value
- Must use operator: =
- Supported values: "string", "integer"

**When custom_score_type = "integer":**
- **Comparison operators**: =, !=, >, >=, <, <= (most common for numeric scores)
- **Range operator**: RANGE (for finding scores within a numeric range)
- **Multi-value operators**: IN, NOT (for matching or excluding specific integer values)

**When custom_score_type = "string":**
- **Equality operators**: =, != (for exact text matches)
- **Text operators**: CONTAINS, BEGINS_WITH, ENDS_WITH (for pattern matching within text)
- **Multi-value operators**: IN, NOT (for matching or excluding multiple string values)
- **Comparison operators**: >, >=, <, <=, MATCHES (lexicographic comparison)
- **Range operator**: RANGE (lexicographic range, rarely used)

**Important Notes:**
1. ❗ All three fields MUST be used together in the specified order
2. ❗ They must be connected with AND operators
3. ❗ The query order matters: name → type → value
4. ❗ Do NOT use these fields individually
5. Custom scores are user-defined attributes that extend CTIX's native scoring capabilities

Example Use Cases:
- Find high business impact threats: custom_score_name = "business_impact" AND custom_score_type = "integer" AND custom_score_value > "7"
- Find specific severity levels: custom_score_name = "severity" AND custom_score_type = "string" AND custom_score_value = "critical"

Examples:
'custom_score_name' = "business_impact" AND 'custom_score_type' = "integer" AND 'custom_score_value' >= "8"
'custom_score_name' = "severity_level" AND 'custom_score_type' = "string" AND 'custom_score_value' IN ("high", "critical")
'custom_score_name' IN ("impact", "urgency") AND 'custom_score_type' = "integer" AND 'custom_score_value' >= "5"

Combined query example:
'type' = "vulnerability" AND 'custom_score_name' = "exploitability" AND 'custom_score_type' = "integer" AND 'custom_score_value' >= "8" AND 'confidence_score' >= "75"
This will fetch all vulnerabilities with high custom exploitability scores and high confidence.



52.CIDR Lookup

Represents the ability to search for IP address indicators that fall within a specific CIDR (Classless Inter-Domain Routing) notation range. This field enables network-based searches to find all IP indicators within a subnet or IP range.

CIDR notation expresses IP address ranges using a base IP address followed by a slash and the number of network bits (e.g., "192.168.1.0/24" represents all IPs from 192.168.1.0 to 192.168.1.255). This is particularly useful for identifying threats within specific network segments.

Example Use Cases:
- If you want to find all malicious IPs within your corporate network range, use 'type' = "indicator" AND 'ioc_type' = "ipv4-addr" AND 'ip_match' = "10.0.0.0/8"
- If you want to find threats in a specific subnet, use ip_match = "192.168.1.0/24"
- If you want to find all IPs in a small range, use ip_match = "203.0.113.0/28"

The value must be in string format and must follow valid CIDR notation (IPv4 or IPv6 address followed by /prefix_length).

Note: This field is primarily used with indicator objects where ioc_type is "ipv4-addr" or "ipv6-addr". The CIDR lookup performs intelligent subnet matching, finding all IP indicators that fall within the specified range. For example, searching ip_match = "192.168.1.0/24" will match indicators like 192.168.1.1, 192.168.1.50, 192.168.1.255, etc. This is more efficient than using pattern matching with BEGINS_WITH or CONTAINS on the value field for network-based searches.

Supported operators: =

Examples:
'ip_match' = "192.168.1.0/24"
'ip_match' = "10.0.0.0/8"
'ip_match' = "2001:db8::/32"


53-55. Custom Attributes (Composite Field)

Represents the name or identifier of a custom attribute field defined in your organization. Custom attributes are user-defined metadata fields that extend the standard STIX properties, allowing organizations to track additional context specific to their needs.

**This field can be used in TWO ways:**

**1. Standalone Query (Simple Mode):**
Query 'custom_attribute_name' alone to find all objects that have a custom attribute with that name, regardless of its type or value.

**2. Composite Query (Advanced Mode):**
Combine with 'custom_attribute_type' and 'custom_attribute_value' to filter by specific attribute values.

**Standalone Usage:**

When used alone, this field finds all threat data objects that have the specified custom attribute defined, regardless of what value it contains. This is useful for:
- Discovering which objects have been tagged with specific metadata
- Finding objects with particular custom fields populated
- Auditing custom attribute usage across your threat intelligence

Supported operators (standalone): =, !=, IN, NOT

**Standalone Examples:**
'custom_attribute_name' = "owner"
'custom_attribute_name' = "department"
'custom_attribute_name' IN ("criticality", "priority", "owner")
'custom_attribute_name' NOT ("deprecated_field", "old_attribute")

**Standalone Query Use Cases:**
- Find all objects with an owner assigned: custom_attribute_name = "owner"
- Find objects with any of several attributes: custom_attribute_name IN ("department", "business_unit", "team")
- Audit attribute usage: type = "indicator" AND custom_attribute_name = "asset_criticality"

**Composite Query Format:**
custom_attribute_name [operator] "name" AND custom_attribute_type = "type" AND custom_attribute_value [operator] "value"


**custom_attribute_name**: The name/identifier of the custom attribute
- Supported operators: =, !=, IN, NOT
- Example values: "department", "asset_criticality", "owner", "location"

**custom_attribute_type**: The data type of the custom attribute value
- Must use operator: =
- Supported values: "string", "integer", "float", "boolean", "date"

**custom_attribute_value**: The actual attribute value being queried
**All Supported Operators:**
=, !=, >, >=, <, <=, CONTAINS, BEGINS_WITH, ENDS_WITH, RANGE, IN, NOT

**Operator Usage by Data Type:**

**When custom_attribute_type = "string":**
- **Equality operators**: =, != (most common for exact text matches)
- **Text operators**: CONTAINS, BEGINS_WITH, ENDS_WITH, MATCHES (for pattern matching within text)
- **Multi-value operators**: IN, NOT (for matching or excluding multiple string values)
- **Comparison operators**: >, >=, <, <= (lexicographic comparison, rarely used)

**When custom_attribute_type = "integer":**
- **Comparison operators**: =, !=, >, >=, <, <= (most common for numeric comparisons)


**When custom_attribute_type = "boolean":**
- **Equality operators**: =, (primary operators for true/false values)


**When custom_attribute_type = "date":**
- **Comparison operators**: =, !=, >, >=, <, <= (for date comparisons using epoch timestamps)
- **Range operator**: RANGE (for finding dates within a time period)



**Important Notes:**
1. ❗ All three fields MUST be used together in the specified order
2. ❗ They must be connected with AND operators
3. ❗ The query order matters: name → type → value
4. ❗ Do NOT use these fields individually
5. Custom attributes are organization-specific metadata that extend STIX standard properties
6. Different object types can have different custom attributes configured

Example Use Cases:
- Find threats affecting specific departments: custom_attribute_name = "department" AND custom_attribute_type = "string" AND custom_attribute_value = "Finance"
- Find high-criticality assets: custom_attribute_name = "asset_criticality" AND custom_attribute_type = "integer" AND custom_attribute_value >= "8"
- Find date-based attributes: custom_attribute_name = "last_review_date" AND custom_attribute_type = "date" AND custom_attribute_value >= "1735689600000"



When used in composite format, all three fields must be present to query for specific attribute values.

**Composite Examples:**
'custom_attribute_name' = "owner" AND 'custom_attribute_type' = "string" AND 'custom_attribute_value' = "security_team"
'custom_attribute_name' = "department" AND 'custom_attribute_type' = "string" AND 'custom_attribute_value' = "Finance"
'custom_attribute_name' = "criticality" AND 'custom_attribute_type' = "integer" AND 'custom_attribute_value' >= "8"

**Complete Examples by Usage Mode:**

Standalone Mode:
'type' = "indicator" AND 'custom_attribute_name' = "owner"
'type' = "vulnerability" AND 'custom_attribute_name' IN ("owner", "department", "location")

Composite Mode (with type and value):
'type' = "indicator" AND 'custom_attribute_name' = "owner" AND 'custom_attribute_type' = "string" AND 'custom_attribute_value' = "SecOps"
'type' = "vulnerability" AND 'custom_attribute_name' = "criticality" AND 'custom_attribute_type' = "integer" AND 'custom_attribute_value' >= "7"

The value must be in string format and should match the name of a custom attribute configured in your CTIX instance.


56. Enriched Date
Represents the exact date and timestamp when a threat data object was enriched by an external enrichment tool or service. This field tracks when additional context, verdicts, or intelligence was added to an object through enrichment integrations.

The enriched_on timestamp records when enrichment services (like AbuseIPDB, VirusTotal, AlienVault, etc.) processed and enhanced the threat data with additional information. This helps track the freshness of enrichment data and identify recently enriched objects.

Example Use Cases:
- If you want to find recently enriched objects, use enriched_on >= "1735689600000"
- If you want to find objects enriched within a specific timeframe, use enriched_on RANGE ("1704067200000", "1735689600000")
- If you want to find stale enrichments that may need updating, use enriched_on < "1704067200000"
- If you want to audit enrichment operations from last week, use enriched_on >= "1735084800000"

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) when the enrichment occurred.

Note: This field is primarily applicable to indicator objects that have been processed by enrichment tools. It differs from "ctix_modified" which tracks any modification to the object - enriched_on specifically tracks enrichment operations. Multiple enrichment operations can occur on the same object over time, with enriched_on tracking the most recent enrichment. This is useful for identifying objects that need re-enrichment based on your enrichment refresh policies.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'enriched_on' > "1735689600000"
'enriched_on' >= "1735689600000"
'enriched_on' = "1735689600000"
'enriched_on' != "1704067200000"
'enriched_on' <= "1704067200000"
'enriched_on' < "1704067200000"
'enriched_on' RANGE ("1704067200000", "1735689600000")

57.Sighting ID
Represents the unique identifier for a sighting record associated with a threat data object. Sightings in STIX represent observations of threat intelligence in the real world, and each sighting has a unique ID for tracking and reference purposes.

The sighting_id field stores the STIX identifier (typically in UUID format) of the sighting object. This allows you to track specific observations of threats, correlate multiple sightings, and reference particular threat observations. While typically used for exact matching, pattern matching operators allow for partial ID searches when needed.

Example Use Cases:
- If you want to find a specific sighting by its ID, use sighting_id = "sighting--550e8400-e29b-41d4-a716-446655440000"
- If you want to find objects with multiple specific sightings, use sighting_id IN ("sighting--id1", "sighting--id2")
- If you want to find sightings from a particular batch (with common prefix), use sighting_id BEGINS_WITH "sighting--batch-2024"
- If you want to exclude specific test sightings, use sighting_id NOT ("sighting--test-id1", "sighting--test-id2")

The value must be in string format and typically follows the STIX identifier pattern (e.g., "sighting--<uuid>").

Note: Sightings represent observations of threat intelligence in real environments. Each sighting records who saw the threat, where it was seen, when it was seen, and how many times. The sighting_id uniquely identifies each observation record. This field is particularly useful when working with threat sharing communities where multiple organizations report sightings of the same threat indicator.

Supported operators: =, !=, CONTAINS, BEGINS_WITH, ENDS_WITH, IN, NOT, MATCHES

Examples:
'sighting_id' = "sighting--550e8400-e29b-41d4-a716-446655440000"
'sighting_id' != "sighting--test-123"
'sighting_id' CONTAINS "550e8400"
'sighting_id' BEGINS_WITH "sighting--batch-2024"
'sighting_id' ENDS_WITH "446655440000"
'sighting_id' IN ("sighting--id1", "sighting--id2", "sighting--id3")
'sighting_id' NOT ("sighting--false-positive-1", "sighting--test-data")
'sighting_id' MATCHES "sighting--.*-test-.*"


58.Sighting Count
Represents the total number of times a threat data object (typically an indicator) has been sighted or observed. This field tracks how frequently a threat has been seen in real-world environments, providing a measure of threat prevalence.

The sighting_count is a numeric value that accumulates as organizations report observations of the same threat. Higher sighting counts generally indicate more widespread or active threats that are being observed by multiple sources.

Example Use Cases:
- If you want to find highly observed threats, use sighting_count >= "10"
- If you want to find rarely seen threats, use sighting_count <= "2"
- If you want to find threats with moderate observation levels, use sighting_count RANGE ("5", "20")
- If you want to prioritize widely seen threats, use sighting_count >= "50"

The value must be in string format and represents an integer count of sightings.

Note: This field is particularly valuable for threat prioritization. Indicators with higher sighting counts are being observed more frequently in the wild, which may indicate active campaigns or widespread threats. Combined with other fields like confidence_score or enrichment_verdict, sighting_count helps identify the most critical and actively exploited threats requiring immediate attention.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'sighting_count' > "10"
'sighting_count' >= "5"
'sighting_count' = "1"
'sighting_count' != "0"
'sighting_count' <= "3"
'sighting_count' < "2"
'sighting_count' RANGE ("5", "50")

Combined query examples:
'type' = "indicator" AND 'sighting_count' >= "20" AND 'confidence_score' >= "75"
This will fetch all high-confidence indicators with 20 or more sightings, representing active widespread threats.

'type' = "indicator" AND 'ioc_type' = "ipv4-addr" AND 'sighting_count' > "10" AND 'enrichment_verdict' = "Malicious"
This will fetch all malicious IPv4 addresses with more than 10 sightings, indicating actively malicious IPs.

59.Has Sighting

Represents whether a threat data object has any associated sighting records. This boolean flag indicates if the threat has been observed in real-world environments by any source.

Sightings provide real-world validation of threat intelligence. Objects with sightings (has_sighting = "true") have been confirmed as observed in actual environments, making them more actionable than purely theoretical or unconfirmed threats.

Supported Values:
- "true" - Object has one or more sighting records
- "false" - Object has no sighting records (default state)

Example Use Cases:
- If you want to find only observed threats, use has_sighting = "true"
- If you want to find unconfirmed threats without sightings, use has_sighting = "false"
- If you want to prioritize confirmed threats, use has_sighting = "true" AND confidence_score >= "70"

The value must be in string format ("true" or "false").

Note: This field is a quick filter for determining if threat intelligence has been validated through real-world observations. Objects with has_sighting = "true" are generally considered higher priority because they represent confirmed active threats rather than just potential threats. This is particularly valuable when combined with other fields to identify actionable, confirmed threats.

Supported operators: =, !=

Examples:
'has_sighting' = "true"
'has_sighting' = "false"
'has_sighting' != "false"
'has_sighting' != "true"


60.Sighting Located


Represents whether a sighting record has associated location information. This boolean flag indicates if geographic/location data was captured when the threat was observed, helping identify sightings with geographic context.

Sightings with location information (sighting_located = "true") provide valuable geographic intelligence about where threats are being observed, while sightings without location data may have been reported anonymously or without geographic attribution.

Supported Values:
- "true" - Sighting has location/geographic information
- "false" - Sighting has no location information (default state)

Example Use Cases:
- If you want to find sightings with location data, use sighting_located = "true"
- If you want to find sightings without geographic attribution, use sighting_located = "false"
- If you want to prioritize geo-located threats for regional analysis, use has_sighting = "true" AND sighting_located = "true"

The value must be in string format ("true" or "false").

Note: This field is useful for filtering sightings based on whether they include geographic context. Sightings with sighting_located = "true" can be analyzed for regional threat patterns, targeted attack analysis, and geographic risk assessment. This differs from other location fields that might store the actual location - sighting_located simply indicates presence or absence of location data.

Supported operators: =, !=

Examples:
'sighting_located' = "true"
'sighting_located' = "false"
'sighting_located' != "true"
'sighting_located' != "false"

61.Sighting Observed
Represents whether a sighting has been observed or confirmed. This boolean flag indicates if the sighting record has observation data attached to it, helping distinguish between different types of sighting reports.

Sightings with sighting_observed = "true" indicate confirmed observations of the threat, while false may indicate reported but unconfirmed or pending verification sightings.

Supported Values:
- "true" - Sighting has been observed/confirmed
- "false" - Sighting has not been observed/confirmed (default state)

Example Use Cases:
- If you want to find confirmed observed sightings, use sighting_observed = "true"
- If you want to find unconfirmed sightings, use sighting_observed = "false"
- If you want to prioritize confirmed sightings, use has_sighting = "true" AND sighting_observed = "true"

The value must be in string format ("true" or "false").

Note: This field helps filter sightings based on their observation/confirmation status. Sightings with sighting_observed = "true" are generally more reliable and actionable as they represent confirmed observations rather than unverified reports. This is useful for threat validation and prioritization workflows.

Supported operators: =, !=

Examples:
'sighting_observed' = "true"
'sighting_observed' = "false"
'sighting_observed' != "true"
'sighting_observed' != "false"

62.Sighting First Seen
Represents the timestamp of the first sighting/observation of a threat data object. When an indicator or threat has multiple sighting records, this field captures the earliest observation, helping identify when a threat first appeared in real-world environments.

The sighting_first_seen timestamp provides the initial detection time across all sighting records for an object. This is valuable for understanding threat emergence and distinguishing newly observed threats from long-standing ones.

Example Use Cases:
- If you want to find threats first observed recently, use sighting_first_seen >= "1735689600000"
- If you want to find emerging threats from last month, use sighting_first_seen >= "1733097600000"
- If you want to find established threats (first seen long ago), use sighting_first_seen < "1704067200000"
- If you want to find threats first observed in a specific period, use sighting_first_seen RANGE ("1704067200000", "1735689600000")

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) of the earliest sighting.

Note: This field differs from "first_seen" (malware/threat-actor field) and "valid_from" (indicator validity). The sighting_first_seen specifically tracks the first real-world observation across all sighting records. An indicator might have a valid_from date in 2020, but sighting_first_seen in 2024 if it was only recently observed in the wild. This helps identify when theoretical threats became active, confirmed threats.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'sighting_first_seen' > "1735689600000"
'sighting_first_seen' >= "1735689600000"
'sighting_first_seen' = "1735689600000"
'sighting_first_seen' != "1704067200000"
'sighting_first_seen' <= "1704067200000"
'sighting_first_seen' < "1704067200000"
'sighting_first_seen' RANGE ("1704067200000", "1735689600000")


63.Sighting Last Seen
Represents the timestamp of the most recent sighting/observation of a threat data object. When an indicator or threat has multiple sighting records, this field captures the latest observation, helping identify currently active versus dormant threats.

The sighting_last_seen timestamp provides the most recent detection time across all sighting records for an object. This is critical for determining if threats are currently active and require immediate attention or have become inactive.

Example Use Cases:
- If you want to find currently active threats, use sighting_last_seen >= "1735689600000"
- If you want to find dormant threats (not recently seen), use sighting_last_seen < "1704067200000"
- If you want to find threats active within a specific period, use sighting_last_seen RANGE ("1704067200000", "1735689600000")
- If you want to prioritize fresh sightings from last 7 days, use sighting_last_seen >= "1735084800000"

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) of the most recent sighting.

Note: This field is crucial for threat prioritization. Threats with recent sighting_last_seen values are actively being observed and pose current risks. Combined with sighting_first_seen, you can calculate threat lifespan and activity patterns. A threat with sighting_last_seen from years ago may be considered dormant, while one with sighting_last_seen from today is an active, immediate concern.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'sighting_last_seen' > "1735689600000"
'sighting_last_seen' >= "1735689600000"
'sighting_last_seen' = "1735689600000"
'sighting_last_seen' != "1704067200000"
'sighting_last_seen' <= "1704067200000"
'sighting_last_seen' < "1704067200000"
'sighting_last_seen' RANGE ("1704067200000", "1735689600000")

64.Sighting Source Created
Represents the timestamp when the sighting record was created at the original source before being ingested into CTIX. This field tracks when the source organization or system that reported the sighting originally created the sighting record.

The sighting_source_created timestamp provides provenance information about the sighting data, indicating when the reporting source documented their observation. This differs from when CTIX received or processed the sighting.

Example Use Cases:
- If you want to find sightings created by sources recently, use sighting_source_created >= "1735689600000"
- If you want to find sightings with delayed reporting, compare sighting_source_created with ctix_created
- If you want to audit sighting sources by creation time, use sighting_source_created RANGE ("1704067200000", "1735689600000")

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) when the source created the sighting record.

Note: This field is part of sighting metadata and helps understand the timeline of threat intelligence sharing. The sighting_source_created might be earlier than when CTIX received it, indicating the lag between observation and sharing. This is important for assessing the timeliness of threat intelligence feeds and understanding how quickly partners share sighting information.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'sighting_source_created' > "1735689600000"
'sighting_source_created' >= "1735689600000"
'sighting_source_created' = "1735689600000"
'sighting_source_created' != "1704067200000"
'sighting_source_created' <= "1704067200000"
'sighting_source_created' < "1704067200000"
'sighting_source_created' RANGE ("1704067200000", "1735689600000")

65.Sighting Source Modified
Represents the timestamp when the sighting record was last modified at the original source before being ingested into CTIX. This field tracks when the source organization or system that reported the sighting last updated the sighting information.

The sighting_source_modified timestamp provides update provenance for sighting data, indicating when the reporting source made changes to their sighting record. This helps track the evolution and refinement of sighting information at the source level.

Example Use Cases:
- If you want to find recently updated sightings from sources, use sighting_source_modified >= "1735689600000"
- If you want to identify sightings that sources have refined, compare sighting_source_created with sighting_source_modified
- If you want to audit sighting updates by sources, use sighting_source_modified RANGE ("1704067200000", "1735689600000")

The value must be in string format and represents the epoch timestamp (Unix timestamp in milliseconds, 13-digit) when the source last modified the sighting record.

Note: This field differs from sighting_source_created (initial creation) and ctix_modified (CTIX updates). The sighting_source_modified tracks when the reporting source updated their sighting information, which might include adding context, updating sighting counts, or correcting observation details. Differences between sighting_source_created and sighting_source_modified indicate the source refined their sighting data after initial reporting.

Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
'sighting_source_modified' > "1735689600000"
'sighting_source_modified' >= "1735689600000"
'sighting_source_modified' = "1735689600000"
'sighting_source_modified' != "1704067200000"
'sighting_source_modified' <= "1704067200000"
'sighting_source_modified' < "1704067200000"
'sighting_source_modified' RANGE ("1704067200000", "1735689600000")




66. Aliases

Represents the aliases, alternative names, or "also-known-as" (AKA) names for a threat data object. This is commonly used for Threat Actors (e.g., "Fancy Bear" for "APT28") and Malware (e.g., "Zeus" for "Zbot").

Example Use Cases:

If you want to find a threat actor by any of its known names, use aliases CONTAINS "APT28"

If you want to find all objects associated with either "Fancy Bear" or "APT28", use aliases IN ("Fancy Bear", "APT28")

Value Format: The value must be a string. This field typically stores a list of names, so CONTAINS or IN are the most common operators.

Note: This field is primarily applicable to Threat Actor and Malware objects.

Supported operators: =, !=, CONTAINS, IN, NOT, BEGINS_WITH, ENDS_WITH, MATCHES

Examples:
'aliases' CONTAINS "APT28"
'aliases' != "Cozy Bear"
'aliases' IN ("Fancy Bear", "APT28")
'aliases' NOT ("Dragonfly", "Sandworm")
'aliases' BEGINS_WITH "APT"
'aliases' MATCHES "APT[0-9]+"

Combined query example:
'type' = "threat-actor" AND 'aliases' CONTAINS "Fancy Bear"

67. External References
Represents external references associated with a threat data object. This field stores external identifiers, such as CVE IDs, MITRE ATT&CK technique IDs, URLs to analysis reports, or internal tracking IDs.

Example Use Cases:

If you want to find all objects related to a specific vulnerability, use external_references = "CVE-2021-44228"

If you want to find all objects mapped to a MITRE ATT&CK technique, use external_references = "T1566"

If you want to find all vulnerability data, use external_references BEGINS_WITH "CVE-"

Value Format: The value must be in string format and represents the external ID, name, or URL.

Note: This field is crucial for correlating CTIX data with external intelligence sources, vulnerability databases (like NVD), or frameworks (like ATT&CK).

Supported operators: =, !=, IN, NOT, BEGINS_WITH

Examples:
'external_references' = "CVE-2021-44228"
'external_references' != "MITRE ATT&CK"
'external_references' IN ("CVE-2021-44228", "CVE-2021-45046")
'external_references' NOT ("Internal-ID-123")
'external_references' BEGINS_WITH "CVE-"


68. Kill Chain phases
Represents the kill chain phases associated with the object (e.g., from Lockheed Martin kill chain).
The value must be a string.

Example Use Cases:

If you want to find all TTPs or indicators related to the "Exploitation" phase, use kill_chain_phases CONTAINS "exploitation"

If you want to find all malware used for "Installation", use 'type' = "malware" AND 'kill_chain_phases' CONTAINS "installation"

Supported operators: =, !=, CONTAINS, IN, NOT, BEGINS_WITH, ENDS_WITH, MATCHES

Examples:
'kill_chain_phases' CONTAINS "exploitation"
'kill_chain_phases' != "reconnaissance"
'kill_chain_phases' IN ("exploitation", "installation")
'kill_chain_phases' NOT ("delivery", "C2")
'kill_chain_phases' BEGINS_WITH "initial"

69. is_defanged
Represents whether the object (e.g., an indicator value) has been "defanged" (made non-malicious, like 'hxxp://' or '[.]').
Supported values are "true", "false".

Example Use Cases:

If you want to find all indicators that were imported in a "safe" (defanged) format, use is_defanged = "true"

If you want to find indicators that are still in their original, potentially harmful format, use is_defanged = "false"

Supported operators: =, !=

Examples:
'is_defanged' = "true"
'is_defanged' != "true"

70. is_published
Represents whether the object has been published to any collection.
Supported values are "true", "false".

Example Use Cases:

If you want to find all objects that have been shared/published, use is_published = "true"

If you want to find all objects that are still in a "draft" or "local-only" state, use is_published = "false"

Supported operators: =, !=

Examples:
'is_published' = "true"
'is_published' = "false"

71. has_alias
Represents whether the object has any aliases defined.
Supported values are "true", "false".

Example Use Cases:

If you want to find all threat actors that have known aliases, use 'type' = "threat-actor" AND 'has_alias' = "true"

If you want to find malware objects that have no other known names, use 'type' = "malware" AND 'has_alias' = "false"

Supported operators: =, !=

Examples:
'has_alias' = "true"
'has_alias' = "false"

72. has_analyst_notes
Represents whether an object has analyst notes attached.
Supported values are "true", "false".

Example Use Cases:

If you want to find all objects that have been reviewed and annotated by analysts, use has_analyst_notes = "true"

If you want to find objects that have no human-added context, use has_analyst_notes = "false"

Supported operators: =, !=

Examples:
'has_analyst_notes' = "true"
'has_analyst_notes' = "false"

73. has_tasks
Represents whether an object has any tasks associated with it.
Supported values are "true", "false".

Example Use Cases:

If you want to find all objects that have open or closed tasks, use has_tasks = "true"

If you want to find objects that have no associated workflow or tasks, use has_tasks = "false"

Supported operators: =, !=

Examples:
'has_tasks' = "true"
'has_tasks' = "false"

74. goals
Represents the goals or objectives of a Threat Actor, Intrusion Set, or Campaign.
The value must be a string.

Example Use Cases:

If you want to find all threat actors focused on espionage, use 'type' = "threat-actor" AND 'goals' CONTAINS "espionage"

If you want to find campaigns not related to financial gain, use 'type' = "campaign" AND 'goals' != "financial gain"

Supported operators: =, !=, CONTAINS, IN, NOT, BEGINS_WITH, ENDS_WITH, MATCHES

Examples:
'goals' = "espionage"
'goals' != "financial gain"
'goals' IN ("espionage", "data theft")
'goals' NOT ("destruction")
'goals' BEGINS_WITH "cyber"

75. infrastructure_type
Represents the type of infrastructure (e.g., "command-and-control").
The value must be a string.

Example Use Cases:
If you want to find all C2 (command-and-control) infrastructure, use 'type' = "infrastructure" AND 'infrastructure_type' = "command-and-control"

If you want to find all infrastructure except C2, use 'type' = "infrastructure" AND 'infrastructure_type' != "command-and-control"

Supported operators: =, !=, IN, NOT

Examples:
'infrastructure_type' = "command-and-control"
'infrastructure_type' != "malware-hosting"
'infrastructure_type' IN ("command-and-control", "staging")
'infrastructure_type' NOT ("exfiltration")

76. threat_actor_type
Represents the type of threat actor (e.g., "state-sponsored").
The value must be a string.

Example Use Cases:

If you want to find all state-sponsored actors, use 'type' = "threat-actor" AND 'threat_actor_type' = "state-sponsored"

If you want to find all financially motivated criminal groups, use 'type' = "threat-actor" AND 'threat_actor_type' IN ("crime-syndicate", "criminal")

Supported operators: =, !=, IN, NOT

Examples:
'threat_actor_type' = "state-sponsored"
'threat_actor_type' != "hacktivist"
'threat_actor_type' IN ("state-sponsored", "crime-syndicate")
'threat_actor_type' NOT ("insider-threat")

77. threat_actor_role
Represents the role of a threat actor.
The value must be a string.

Example Use Cases:

If you want to find actors identified as "directors" of operations, use 'type' = "threat-actor" AND 'threat_actor_role' = "director"

Supported operators: =, !=, IN, NOT

Examples:
'threat_actor_role' = "director"
'threat_actor_role' != "agent"
'threat_actor_role' IN ("director", "infiltrator")
'threat_actor_role' NOT ("sponsor")
'threat_actor_role' BEGINS_WITH "direct"

78. sophistication
Represents the sophistication level of a threat actor (e.g., "high").
The value must be a string.

Example Use Cases:

If you want to find all highly sophisticated actors, use 'type' = "threat-actor" AND 'sophistication' = "high"

If you want to filter for advanced or moderate threats, use 'type' = "threat-actor" AND 'sophistication' IN ("high", "medium")

Supported operators: =, !=, IN, NOT

Examples:
'sophistication' = "high"
'sophistication' != "low"
'sophistication' IN ("high", "medium")
'sophistication' NOT ("none")

79. resource_level
Represents the resource level of a threat actor (e.g., "organization").
The value must be a string.

Example Use Cases:

If you want to find actors with government-level backing, use 'type' = "threat-actor" AND 'resource_level' = "government"

If you want to find individual or small-team actors, use 'type' = "threat-actor" AND 'resource_level' IN ("individual", "team")

Supported operators: =, !=, IN, NOT

Examples:
'resource_level' = "organization"
'resource_level' != "individual"
'resource_level' IN ("organization", "government")
'resource_level' NOT ("team")

80. primary_motivation
Represents the primary motivation of a threat actor (e.g., "espionage").
The value must be a string.

Example Use Cases:

If you want to find all actors driven by espionage, use 'type' = "threat-actor" AND 'primary_motivation' = "espionage"

If you want to find actors whose main goal is not financial, use 'type' = "threat-actor" AND 'primary_motivation' != "financial-gain"

Supported operators: =, !=, IN, NOT

Examples:
'primary_motivation' = "espionage"
'primary_motivation' != "financial-gain"
'primary_motivation' IN ("espionage", "ideological")
'primary_motivation' NOT ("accidental", "unknown")

81. secondary_motivation
Represents the secondary motivation of a threat actor.
The value must be a string.

Example Use Cases:

If you want to find actors who have a secondary motivation of financial gain (e.g., state actors moonlighting), use 'type' = "threat-actor" AND 'secondary_motivation' = "financial-gain"

Supported operators: =, !=, IN, NOT

Examples:
'secondary_motivation' = "financial-gain"
'secondary_motivation' != "espionage"
'secondary_motivation' IN ("financial-gain", "destruction")
'secondary_motivation' NOT ("ideological")

82. personal_motivations
Represents the personal motivations of a threat actor (e.g., "greed").
The value must be a string.

Example Use Cases:

If you are investigating insider threats, you might search for type = "threat-actor" AND personal_motivations = "revenge"


Supported operators: =, !=, IN, NOT

Examples:
'personal_motivations' = "greed"
'personal_motivations' != "revenge"
'personal_motivations' IN ("greed", "ideology")
'personal_motivations' NOT ("boredom")

83. version_string
Represents the version of a tool.
The value must be a string.

Example Use Cases:

If you want to find a specific version of a tool, use 'type' = "tool" AND 'version_string' = "1.2.3"

If you want to find all objects related to version 1.2, use 'type' = "tool" AND 'version_string' BEGINS_WITH "1.2"

Supported operators: =, !=, IN, NOT, BEGINS_WITH, ENDS_WITH, MATCHES, CONTAINS

Examples:
'version_string' = "1.2.3"
'version_string' != "1.0"
'version_string' IN ("1.2.3", "1.2.4")
'version_string' NOT ("2.0")
'version_string' BEGINS_WITH "1.2"
'version_string' CONTAINS "beta"
'version_string' ENDS_WITH "release"

84. tool_type
Represents the type of a tool (e.g., "RAT").
The value must be a string.

Example Use Cases:

If you want to find all known Remote Access Trojans (RATs), use 'type' = "tool" AND 'tool_type' = "RAT"

If you want to find all scanning or downloading tools, use 'type' = "tool" AND 'tool_type' IN ("scanner", "downloader")

Supported operators: =, !=, IN, NOT

Examples:
'tool_type' = "RAT"
'tool_type' != "downloader"
'tool_type' IN ("RAT", "keylogger")
'tool_type' NOT ("scanner")

85. malware_type
Represents the type of malware (e.g., "trojan").
The value must be a string.

Example Use Cases:

If you want to find all ransomware objects, use 'type' = "malware" AND 'malware_type' = "ransomware"

If you want to find various types of spyware, use 'type' = "malware" AND 'malware_type' IN ("spyware", "keylogger")

Supported operators: =, !=, IN, NOT

Examples:
'malware_type' = "trojan"
'malware_type' != "ransomware"
'malware_type' IN ("trojan", "RAT", "spyware")
'malware_type' NOT ("worm", "virus")

86. is_family
Represents whether the malware object is a "family" (a group of related malware) or a specific instance.
Supported values are "true" (it is a family), "false" (it is an instance).

Example Use Cases:

If you want to find high-level malware families (like "Zeus" or "Emotet"), use 'type' = "malware" AND 'is_family' = "true"

If you want to find specific malware samples or instances, use 'type' = "malware" AND 'is_family' = "false"

Supported operators: =, !=

Examples:
'is_family' = "true"
'is_family' = "false"

87. architecture_execution_envs
Represents the list of execution environments (e.g., "x86-64").
The value must be a string.

Example Use Cases:

If you want to find all malware that targets 64-bit systems, use 'type' = "malware" AND 'architecture_execution_envs' = "x86-64"

If you want to find malware targeting ARM (e.g., IoT or mobile), use 'type' = "malware" AND 'architecture_execution_envs' = "arm"

Supported operators: =, !=, IN, NOT

Examples:
'architecture_execution_envs' = "x86-64"
'architecture_execution_envs' != "arm"
'architecture_execution_envs' IN ("x86-64", "x86")
'architecture_execution_envs' NOT ("ia-64", "mips")

88. implementation_languages
Represents the list of programming languages used to implement the malware.
The value must be a string.

Example Use Cases:

If you want to find all malware written in Python, use 'type' = "malware" AND 'implementation_languages' = "python"

If you want to find malware written in either Go or Python, use 'type' = "malware" AND 'implementation_languages' IN ("python", "go")

Supported operators: =, !=, IN, NOT

Examples:
'implementation_languages' = "python"
'implementation_languages' != "c++"
'implementation_languages' IN ("python", "go")
'implementation_languages' NOT ("powershell", "c#")

89. malware_capabilities
Represents the list of capabilities for the malware (e.g., "data-theft").
The value must be a string.

Example Use Cases:

If you want to find all malware capable of stealing data, use 'type' = "malware" AND 'malware_capabilities' = "data-theft"

If you want to find malware that tries to evade analysis, use 'type' = "malware" AND 'malware_capabilities' = "anti-vm"

Supported operators: =, !=, IN, NOT

Examples:
'malware_capabilities' = "data-theft"
'malware_capabilities' != "anti-vm"
'malware_capabilities' IN ("data-theft", "spreads-laterally")
'malware_capabilities' NOT ("privilege-escalation", "rootkit")

90. identity_roles
Represents the list of roles for an identity (e.g., "victim").
The value must be a string.

Example Use Cases:

If you want to find all identities that have been identified as victims, use 'type' = "identity" AND 'identity_roles' CONTAINS "victim"

Supported operators: =, !=, BEGINS_WITH, ENDS_WITH, MATCHES, CONTAINS, IN, NOT

Examples:
'identity_roles' = "victim"
'identity_roles' != "source"
'identity_roles' IN ("victim", "target")
'identity_roles' NOT ("attacker")
'identity_roles' CONTAINS "target"

91. identity_class
Represents the class of an identity (e.g., "organization").
The value must be a string.

Example Use Cases:

If you want to find all identities that are organizations, use 'type' = "identity" AND 'identity_class' = "organization"

If you want to find all identities that are people, use 'type' = "identity" AND 'identity_class' = "individual"

Supported operators: =, !=, IN, NOT

Examples:
'identity_class' = "organization"
'identity_class' != "individual"
'identity_class' IN ("organization", "group")
'identity_class' NOT ("unknown")

92. identity_sectors
Represents the list of industry sectors for an identity.
The value must be a string.

Example Use Cases:

If you want to find all identities in the finance sector, use type = "identity" AND identity_sectors CONTAINS "finance"

If you want to find all targeted identities in the technology or government sectors, use type = "identity" AND identity_roles CONTAINS "target" AND identity_sectors IN ("technology", "government")

Supported operators: =, !=, IN, NOT

Examples:
'identity_sectors' = "technology"
'identity_sectors' != "finance"
'identity_sectors' IN ("technology", "government")
'identity_sectors' NOT ("retail")

93. identity_email
Represents the email address for an identity.
The value must be a string.

Example Use Cases:

If you want to find the identity for a specific person, use type = "identity" AND identity_email = "victim@example.com"

If you want to find all identities starting with "victim", use type = "identity" AND identity_email BEGINS_WITH "victim@"

Supported operators: =, !=, IN, NOT, BEGINS_WITH

Examples:
'identity_email' = "victim@example.com"
'identity_email' != "attacker@evil.com"
'identity_email' IN ("victim@example.com", "contact@org.com")
'identity_email' NOT ("test@test.com")
'identity_email' BEGINS_WITH "victim@"

94. identity_phone_number
Represents the phone number for an identity.
The value must be a string.

Example Use Cases:

If you want to find the identity associated with a specific phone number, use type = "identity" AND identity_phone_number = "+1-555-123-4567"

Supported operators: =, !=, IN, NOT, BEGINS_WITH

Examples:
'identity_phone_number' = "+1-555-123-4567"
'identity_phone_number' != "+1-800-000-0000"
'identity_phone_number' IN ("+1-555-123-4567", "+44-20-7946-0123")
'identity_phone_number' NOT ("N/A")
'identity_phone_number' BEGINS_WITH "+1-555"

95. location_region
Represents the region of a location (e.g., "North America").
The value must be a string.

IMPORTANT: Unlike other location fields (location_country, location_city, location_administrative_area), location_region does NOT support text matching operators like CONTAINS, BEGINS_WITH, ENDS_WITH, or MATCHES. Only use exact matching operators.

Example Use Cases:

If you want to find all locations within North America, use type = "location" AND location_region = "North America"

If you want to find locations in either Europe or Asia, use type = "location" AND location_region IN ("Europe", "Asia")

Supported operators: =, !=, IN, NOT

Examples:
'location_region' = "North America"
'location_region' != "Europe"
'location_region' IN ("North America", "Asia")
'location_region' NOT ("Antarctica")

96. location_country
Represents the country of a location (e.g., "US").
The value must be a string.

Example Use Cases:

If you want to find all locations in the United States, use type = "location" AND location_country = "US"

If you want to find locations associated with specific high-risk countries, use type = "location" AND location_country IN ("RU", "CN")

Supported operators: =, !=, IN, NOT, CONTAINS, BEGINS_WITH, ENDS_WITH, MATCHES

Examples:
'location_country' = "US"
'location_country' != "CA"
'location_country' IN ("US", "GB")
'location_country' NOT ("RU", "CN")

97. location_city
Represents the city of a location.
The value must be a string.

Example Use Cases:

If you want to find all locations in New York, use type = "location" AND location_city = "New York"
Supported operators: =, !=, IN, NOT, CONTAINS, BEGINS_WITH, ENDS_WITH, MATCHES

Examples:
'location_city' = "New York"
'location_city' != "London"
'location_city' IN ("New York", "Washington D.C.")
'location_city' NOT ("Moscow")

98. location_administrative_area
Represents the administrative area (e.g., state or province) of a location.
The value must be a string.

Example Use Cases:

If you want to find all locations in the state of California, use type = "location" AND location_administrative_area = "California"

Supported operators: =, !=, IN, NOT, CONTAINS, BEGINS_WITH, ENDS_WITH, MATCHES

Examples:
'location_administrative_area' = "California"
'location_administrative_area' != "New York"
'location_administrative_area' IN ("California", "Virginia")
'location_administrative_area' NOT ("Texas")



99. description
Represents a unified search across all description fields of a threat data object. When used, it searches across three fields simultaneously using OR logic:
- source_description: The description provided by the original intelligence source
- analyst_description: Notes and descriptions added by analysts
- ai_summary: AI-generated summaries of the threat object

This parameter abstracts multi-field description searching into a single convenient field. It allows analysts to discover threat objects based on contextual narratives stored in descriptions.

IMPORTANT: The description field MUST ALWAYS use the CONTAINS operator. This is because description searches are inherently partial/fuzzy text matches against unstructured narrative content. Do NOT use =, !=, BEGINS_WITH, or ENDS_WITH with description.

Example Use Cases:
- If you want to find all threat objects mentioning a ransomware campaign in any description, use 'description' CONTAINS "ransomware campaign"
- If you want to find indicators related to phishing in their descriptions, use 'type' = "indicator" AND 'description' CONTAINS "phishing"
- If you want to find threat actors described as state-sponsored, use 'type' = "threat-actor" AND 'description' CONTAINS "state-sponsored"

The value must be in string format.

Supported operators: CONTAINS (always use CONTAINS for description)

Examples:
'description' CONTAINS "ransomware"
'description' CONTAINS "APT29"
'description' CONTAINS "lateral movement"
'description' CONTAINS "remote code execution"
'type' = "indicator" AND 'description' CONTAINS "phishing"
'type' = "malware" AND 'description' CONTAINS "lateral movement"
'confidence_score' >= "75" AND 'description' CONTAINS "critical vulnerability"
'type' = "threat-actor" AND 'description' CONTAINS "state-sponsored"


---

📋 Field Summary:

| Field       | Type   | Operators                                       | Multi-value Support |
|-------------|--------|--------------------------------------------------|---------------------|
| type        | Enum   | =, !=, IN, NOT                                   | Yes                 |
| ioc_type    | Enum   | =, !=, IN, NOT                                   | Yes                 |
| source      | String | =, !=, IN, NOT                                   | Yes                 |
| value       | String | =, !=, MATCHES, BEGINS_WITH, ENDS_WITH, IN, NOT  | Yes                 |
| description | String | CONTAINS                                          | No                  |

Use these grammar rules to generate valid CQL queries for searching threat data within the CTIX platform.`
