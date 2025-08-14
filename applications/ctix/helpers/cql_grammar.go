package helpers

const CQL_grammar_rule = `This tool returns all the grammar rules required to build a CQL (Cyware Query Language) query.  
In the backend, these CQL queries are used to perform searches.  
CQL must be only used when you want to search the list of threat data object based on the condition.
Example query: type = "indicator" AND ioc_type = "ipv4-addr"  

❗❗❗ Caution : Don't use parentheses when generating sub-queries. CQL should follow natural left-to-right precedence without brackets.

This is invalid query : (type = \"malware\" AND relationship_type = \"indicates\" AND related_object = \"indicator\") OR (type = \"indicator\" AND relationship_type = \"related-to\" AND related_object = \"report\")
Valid query := type = \"malware\" AND relationship_type = \"indicates\" OR type = \"indicator\" AND related_object = \"report\" AND relationship_type = \"related-to\"

Below are the grammar rules used to construct a valid CQL:

1. type  
Represents the object type in the query.  
Supported values, valaues must be from one of these:  
indicator, malware, threat-actor, vulnerability, attack-pattern, campaign, course-of-action, identity, infrastructure, intrusion-set, location, malware-analysis, observed-data, opinion, tool, report, custom-object, observable, incident, note, grouping

Supported operators: =, !=, IN, NOT  
Examples:  
type = "indicator"  
type != "indicator"  
type IN ("indicator", "malware")  
type NOT ("indicator", "threat-actor")

2. ioc_type  
Represents the IOC (Indicator of Compromise) type for the SDO.  
Usually used when type = "indicator".  
Supported values:  
artifact, asn, directory, domain-name, email-addr, email-message, file, ipv4-addr, ipv6-addr, mac-addr, MD5, mutex, network-traffic, process, SHA-1, SHA-224, SHA-256, SHA-384, SHA-512, software, SSDEEP, url, user-account, window-registry-key, x509-certificate, yara

Supported operators: =, !=, IN, NOT  
Examples:  
ioc_type = "ipv4-addr"  
ioc_type != "ipv4-addr"  
ioc_type IN ("email-addr", "file")  
ioc_type NOT ("email-addr", "mutex")

3. source  
Represents the source from which threat intelligence is received.  
Values must be strings.  
Supported operators: =, !=, IN, NOT  
Examples:  
source = "import"  
source != "crowdstrike"  
source IN ("import", "mandiant")  
source NOT ("rss", "mail")

4. value  
Represents a value (e.g., title or IOC value) used for free-text search.  
Example: If you want to search for domain "abc.com", use value = "abc.com"  
Values must be strings.  

Supported operators: =, !=, CONTAINS, IN, NOT, BEGINS_WITH, ENDS_WITH  
Examples:  
value = "domain.com"  
value != "domain.com"  
value CONTAINS "dom"  
value IN ("domain.com", "abc.com")  
value NOT ("domain.com", "abc.com")  
value BEGINS_WITH "sd"  
value ENDS_WITH "sd"

5. is_deprecated
Represents the deprecation status of the indicator, if they are deprecated or not.
Example: If you want to search all the indicator which are deprecated, use type="indicator" AND is_deprecated="true"
If you want to search all the indicator which are not deprecated, use type="indicator" AND is_deprecated="false"
Supported values are "true", "false"

Examples:
is_deprecated="true"
is_deprecated="false"


6. ctix_modified
Represent the exact date, timestamp when the threat is updatd in CTIX application. ctix_modified is applicable to all type SDO(stix domain objects) which exists in CTIX.
Example: If you want to get the data which is modified later than 1746901800, use ctix_modified >= "1746901800"
The value must be in string and which is epoch equivalent of the date.
Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
ctix_modified >= "1746901800"
ctix_modified = "1746901800"
ctix_modified != "1746901800"
ctix_modified <= "1746901800"
ctix_modified < "1746901800"
ctix_modified RANGE ("1746297000","1746988199")

7. tag
Represents the tag attached to the threat intel(stix domain objects)
Example: If you want search all the objects having tag malicious then use tag = "malicious"
The value must be in string.
Supported operators: =, !=, IN, NOT

Examples:  
tag = "malicious"  
tag != "non-malicious"  
tag IN ("non-malicious", "malicious")  
tag NOT ("non-malicious", "raided")

8. tlp
Represent the TLP of the SDO(stix domain objects), It can be RED, AMBER_STRICT, AMBER, GREEN, CLEAR, NONE
Example: If you want search all the objects having tlp red then use tlp = "RED"
The value must be in string.
Supported operators: =, !=, IN, NOT

Examples: 
tlp = "RED"  
tlp != "GREEN"  
tlp IN ("RED", "AMBER_STRICT")  
tlp NOT ("GREEN", "NONE")


9. published_collection
Represent the STIX collection to which the data is published.
Example: If you want to search all the objects which are published to Collection1, then use published_collection = "Collection1"
The values must be in string.
Supported operators: =, !=, IN, NOT

Examples: 
published_collection = "Collection1"  
published_collection != "Collection1" 
published_collection IN ("Collection1", "Collection2")  
published_collection NOT ("Collection1", "Collection2")

10. rule

Represent the rules which has run on the threat data object. It will all the actioned object by rule.
You can directly hit the CQL with "rule" = "Rule Name", no need to get the details of the rule until and unless specified by the user.
Example: If you want to search all the threat data objects which are passed by the rules. For example: If you want to search all the threat data objects which are passed by rule PublishToColl then you will do something : "rules" = "PublishToColl"
The values must be in string.
Supported operators: =, !=, IN, NOT

Examples: 
rule = "rule1"  
rule != "rule1" 
rule IN ("rule1", "rule2")  
rule NOT ("rule3", "rule4")


11. enrichment_tool

Represent the enrichment_tool which has enriched the threat data objects.
Example: If you want to search all the threat data objects which are enriched by the enrichment tool AbuseIPDB, then you will do something : "enrichment_tool" = "AbuseIPDB"
The values must be in string.
Supported operators: =, !=, IN, NOT

Examples: 
enrichment_tool = "AbuseIPDB"  
enrichment_tool != "AbuseIPDB" 
enrichment_tool IN ("AbuseIPDB", "Alien Vault")  
enrichment_tool NOT ("AbuseIPDB", "Alien Vault") 

If you want to search whether ip 1.1.1.1 enriched by tool AbuseIPDB then use "value" = "1.1.1.1" AND "enrichment_tool" = "AbuseIPDB"

12. enriched_status

Represent the enrichment status of the threat data object, which means if it tells the object which are enriched.
Example: If you want to search all the threat data objects which are enriched, then you will do something : "enriched_status" = "1"
The values must be in string and they are fixed. Use 1 for ENRICHED, 2 for Tried and Failed and 3 for Quota Completed
Supported operators: =, !=, IN, NOT

Examples: 
enriched_status = "1"  
enriched_status != "2" 
enriched_status IN ("1", "2")  
enriched_status NOT ("3", "2") 

13. relationship_type
This will give you all the object which are having this relation_type in any of its relations. For example user searches for  
"relationship_type = "related-to" it means give all the objects which are having realtion type as related-to with the related object.
"type"= "indicator" AND "relationship_type = "related-to" --> It means give all the indicator having relation type related-to with the related object.

The values must be in string, and name should be in the available relation types list.

Supported operators: =, !=, IN, NOT

Examples: 
relationship_type = "related-to"
relationship_type != "related-to"
relationship_type IN ("related-to", "associated_actor",  "authored-by")
relationship_type NOT ("related-to", "associated_actor",  "authored-by")


14. related_object
This will give you all the object which are related to the specified object type. For example user searches for  
"related_object = "malware" it means give all the objects which are having relation with malware object type.
"type"= "indicator" AND "related_object = "malware" --> It means give all the indicator having relation with malware.

The values must be in string, and should be a valid object type. you can reference "type" to see the values.

Supported operators: =, !=, IN, NOT

Examples: 
related_object = "indicator"
related_object != "indicator"
related_object IN ("malware", "indicator",  "report")
related_object NOT ("malware", "indicator",  "report")

15. related_object_value
This represents the value for the related object of a object. Please note As its a related object property, so it must be followed by related_object field.
The value of this must be a string value
Incorrect query ❌:  related_object_value = "domain.com"
Correct query ✅ : related_object = "indicator" and related_object_value = "domain.com"

Supported operators: =, !=, CONTAINS, IN, NOT, BEGINS_WITH, ENDS_WITH  
Examples:  
related_object_value = "domain.com"  
valrelated_object_valueue != "domain.com"  
related_object_value CONTAINS "dom"  
related_object_value IN ("domain.com", "abc.com")  
related_object_value NOT ("domain.com", "abc.com")  
related_object_value BEGINS_WITH "sd"  
related_object_value ENDS_WITH "sd"

16. related_object_property
This represents the fields/properties for the related object, so it can search/filter data based on the related object as well.
Please note As its a related object property, so it must be followed by related_object field.
There are selected fields only which can be fetched from related object. List of fields are 'type', 'value', 'source'. These fields are already defined in the CQL, so supported operators and usages are same.
Syntax to fetch the fiels is related_object_property.fieldName.

Supported operators are based on fields type which is used.

Example queries -  
related_object_property.source = "alien vault"
related_object_property.type = "indicator"
related_object_property.value CONTAINS "has"



17. has_relations

This will give you all the objects which has atleast one relation with other object, if has_relations = "true".
has_relations = "false" --> Gives the list of objects which doesn't have any relation with any object.

The values must be in string.

Supported operators: =, !=

Examples: 
has_relations = "true"
has_relations != "false"

18. enrichment_verdict

This indicates the enrichment verdict of the threat data objects. It must values from 'Malicious' or 'Non malicious' depending on the query. Dont' start fetching the indicator enrichment details until unless explicitly asked.
The values must be in string.

Supported operators: =, !=, IN, NOT

Examples: 
enrichment_verdict = "Malicious"  
enriched_status != "Malicious" 
enrichment_verdict IN ("Malicious", "Non malicious")  
enrichment_verdict NOT ("Malicious", "Non malicious")  

19. ctix_created
Represent the exact date, timestamp when the threat is created in CTIX application. ctix_created is applicable to all type SDO(stix domain objects) which exists in CTIX.
Example: If you want to get the data which is created later than 1746901800, use ctix_created >= "1746901800"
The value must be in string and which is epoch equivalent of the date.
Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
ctix_created > "1746901800"
ctix_created >= "1746901800"
ctix_created = "1746901800"
ctix_created != "1746901800"
ctix_created <= "1746901800"
ctix_created < "1746901800"
ctix_created RANGE ("1746297000","1746988199")

20. confidence_score
Represet the confidence score(also knowns as risk score) of the threat data object.
Example: If you want to get the data having risk score greater than 75, use confidence_score > "75"
The value must be in string and must be between 0 to 100 inclusive.
Supported operators: =, !=, >, >=, <, <=, RANGE

Examples:
confidence_score > "75"
confidence_score >= "75"
confidence_score = "75"
confidence_score != "75"
confidence_score <= "75"
confidence_score < "75"
confidence_score RANGE ("75","90")
---

📌 Additional Notes for Query Construction:

- Use IN and NOT when querying multiple values for a single field.  
  ✅ Correct: type IN ("indicator", "malware")  
  ❌ Incorrect: type = "indicator" OR type = "malware"

- Don't use parentheses. The query must be written in flat form using AND/OR only.  
  Correct: type = "indicator" AND ioc_type = "ipv4-addr" OR source = "import"  
  This will be interpreted as: (type = "indicator" AND ioc_type = "ipv4-addr") OR source = "import"

- All values (except field names and operators) should be **strings**, i.e., wrapped in double quotes.

- !! parentheses must only be used only if there is IN or NOT operator to enclose the list of values, not anywhere else.

---

📋 Field Summary:

| Field      | Type   | Operators                                       | Multi-value Support |
|------------|--------|--------------------------------------------------|---------------------|
| type       | Enum   | =, !=, IN, NOT                                   | Yes                 |
| ioc_type   | Enum   | =, !=, IN, NOT                                   | Yes                 |
| source     | String | =, !=, IN, NOT                                   | Yes                 |
| value      | String | =, !=, CONTAINS, BEGINS_WITH, ENDS_WITH, IN, NOT | Yes                 |

Use these grammar rules to generate valid CQL queries for searching threat data within the CTIX platform.
`
