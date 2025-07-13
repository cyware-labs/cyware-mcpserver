package helpers

const (
	bulk_update_add_tag_endpoint                  = "ingestion/threat-data/bulk-action/add_tag/"
	bulk_update_mark_indicator_allowed_endpoint   = "ingestion/threat-data/bulk-action/whitelist/"
	bulk_update_unmark_indicator_allowed_endpoint = "ingestion/threat-data/bulk-action/un_whitelist/"
	bulk_add_task_endpoint                        = "ingestion/tasks/bulk-action/add_task/"
	bulk_add_manual_review                        = "ingestion/threat-data/bulk-action/manual_review/"
	bulk_mark_false_positive                      = "ingestion/threat-data/bulk-action/false_positive/"
	bulk_unmark_false_positive                    = "ingestion/threat-data/bulk-action/un_false_positive/"
	bulk_update_analyst_tlp                       = "ingestion/threat-data/bulk-action/analyst_tlp/"
	bulk_update_analyst_score                     = "ingestion/threat-data/bulk-action/analyst_score/"
	bulk_deprecate                                = "ingestion/threat-data/bulk-action/deprecate/"
	bulk_undeprecate                              = "ingestion/threat-data/bulk-action/un_deprecate/"
	bulk_add_watchlist                            = "ingestion/threat-data/bulk-action/watchlist/"
	bulk_remove_watchlist                         = "ingestion/threat-data/bulk-action/un_watchlist/"
	bulk_add_relation                             = "ingestion/threat-data/bulk-action/add_relation/"

	ThreatDataListBulkActionAddTag = `{
   "type":"object",
   "properties":{
      "all_objects":{
         "type":"boolean",
         "description" : "This is a flag which is always set to 'true' "
      },
      "objects":{
         "type":"array",
         "description" : "This is a list of objects to which tags are to be added.",
         "items":{
               "type":"object",
               "properties":{
                  "type":{
                     "type":"string",
                     "description" : "This is the type(indicator, report etc) of threat data object"
                  },
                  "ids":{
                     "type":"array",
                     "description" : "This is the list of threat data object_id of this specifc object type",
                     "items":{
                           "type":"string",
                           "description" : "This is the threat data object_id"
                        }
                  }
               },
               "required":[
                  "type",
                  "ids"
               ]
            }
      },
      "data":{
         "type":"object",
         "properties":{
            "reason":{
               "type":"string",
               "description" : "This is the reason for updating the tag of threat data object"
            },
            "tag_id":{
               "type":"array",
               "description" : "This is a list of tag ids",
               "items":{
                     "type":"string",
                     "description" : "This the tag id"
                  }
            }
         },
         "required":[
            "reason",
            "tag_id"
         ]
      }
   },
   "required":[
      "all_objects",
      "objects",
      "data"
   ]
}
`

	ThreatDataListBulkMarkIndicatorAllowed = `
   {
      "type":"object",
      "description" : "This is the payload used for bulk marking as indicator allowed.",
      "properties":{
         "object_type":{
            "type":"string",
            "description" : "This must be passed as indicator. Indicators can only be marked as indicator allowed."
         },
         "object_ids":{
            "type":"array",
            "description" : "This is the list of the object_id which must be passed.",
            "items":
               {
                  "type":"string",
                  "description" : "This is the object_id of the object"
               }
         }
      },
      "required":[
         "object_type",
         "object_ids"
      ]
   }`

	ThreatDataListBulkUnMarkIndicatorAllowed = `{
      "type":"object",
      "description" : "This is the payload used for bulk un-marking as indicator allowed.",
      "properties":{
         "object_type":{
            "type":"string",
            "description" : "This must be passed as indicator. Indicators can only be marked/unmarked as indicator allowed."
         },
         "object_ids":{
            "type":"array",
            "description" : "This is the list of the object_id which must be passed.",
            "items":
               {
                  "type":"string",
                  "description" : "This is the object_id of the object"
               }
         }
      },
      "required":[
         "object_type",
         "object_ids"
      ]
   }`

	ThreatDataListBulkAddTask = `
	{
	"type": "object",
	"description" : "This is the payload for adding a task in bulk to a number of threat data object",
	"properties": {
	"all_objects": {
		"type": "boolean",
		"description" : "This must be always set to true"
		},
	"objects": {
		"type": "array",
		"description" : "This is the list of the object_id seggregated with the object type",
		"items": {
			"type": "object",
			"properties": {
				"type": {
				"type": "string",
				"description" : "This is the type of the object eg: threat-actor, malware etc"
				},
				"ids": {
				"type": "array",
				"description" :"This is the list of the object ids",
				"items": 
					{
					"type": "string",
					"description" : "This is the object id"
					}
				}
			},
			"required": [
				"type",
				"ids"
			]
			}
		},
		"data": {
		"type": "object",
		"description" : "This is information of the task being assigned",
		"properties": {
			"priority": {
			"type": "string",
			"description" : "Priority of the task. Can have values only from 'high', 'medium' and 'low'"
			},
			"status": {
			"type": "string",
			"description" : "This is the status of the task. As this is creating new task so always use the value 'not_started' for this key"
			},
			"assignee": {
			"type": "string",
			"description" : "This is the user_id of the assignee  of the task. Fetch the user_id from the user listing tool"
			},
			"text": {
			"type": "string",
			"description" : "This is the Actual text for task, what needs to done. It has a character limit of 2000, so be precise."
			},
			"deadline": {
			"type": "integer",
			"description" : "Deadline for the task assigned. It should be a future date in epoch. Use tool to fetch the future date"
			}
		},
		"required": [
			"priority",
			"meta_data",
			"status",
			"assignee",
			"text",
			"deadline"
		]
		}
	},
	"required": [
		"all_objects",
		"objects",
		"data"
	]
	}
`

	ThreatDataListBulkManualReview = `{
	"type": "object",
	"description" : "This is the payload for adding threat data object in bulk for manual review.",
	"properties": {
		"all_objects": {
		"type": "boolean",
		"description" : "This is a flag which is always set to 'true' "
		},
		"objects":{
			"type":"array",
			"description" : "This is a list of objects which are to be added for manual review.",
			"items":{
				"type":"object",
				"properties":{
					"type":{
						"type":"string",
						"description" : "This is the type(indicator, report etc) of threat data object"
					},
					"ids":{
						"type":"array",
						"description" : "This is the list of threat data object_id of this specifc object type",
						"items":{
							"type":"string",
							"description" : "This is the threat data object_id"
							}
					}
				},
				"required":[
					"type",
					"ids"
				]
				}
		},
		"data": {
		"type": "object",
		"properties": {
			"reason": {
			"type": "string",
			"description" : "This is the reason for adding threat data objects for manual review"
			},
			"is_under_review": {
			"type": "boolean",
			"description" : "This is a flag which is always set to 'true' "
			}
		},
		"required": [
			"reason",
			"is_under_review"
		]
		}
	},
	"required": [
		"all_objects",
		"objects",
		"data"
	]
	}`

	ThreatDataListBulkMarkFalsePositive = `{
	"type": "object",
	"description" : "This is the payload used for bulk marking indicator as false postive.",
	"properties": {
		"object_type": {
		"type": "string",
		"description" : "This must be passed as indicator. Indicators can only be marked as false postive."

		},
		"object_ids":{
			"type":"array",
			"description" : "This is the list of the object_id which must be passed.",
			"items":
				{
					"type":"string",
					"description" : "This is the object_id of the object"
				}
		},
		"data": {
		"type": "object",
		"properties": {
			"reason": {
			"type": "string",
			"description" : "This is the reason for marking threat data objects as false positive"

			},
			"is_false_positive": {
			"type": "boolean",
			"description" : "This must be always set to true for marking"
			}
		},
		"required": [
			"reason",
			"is_false_positive"
		]
		}
	},
	"required": [
		"object_type",
		"object_ids",
		"data"
	]
	}`

	ThreatDataListBulkUnMarkFalsePositive = `{
	"type": "object",
	"description" : "This is the payload used for bulk marking indicator as false postive.",
	"properties": {
		"object_type": {
		"type": "string",
		"description" : "This must be passed as indicator. Indicators can only be marked as false postive."

		},
		"object_ids":{
			"type":"array",
			"description" : "This is the list of the object_id which must be passed.",
			"items":
				{
					"type":"string",
					"description" : "This is the object_id of the object"
				}
		},
		"data": {
		"type": "object",
		"properties": {
			"reason": {
			"type": "string",
			"description" : "This is the reason for unmarking threat data objects from false postive"

			},
			"is_false_positive": {
			"type": "boolean",
			"description" : "This must be always set to false for unmarking"
			}
		},
		"required": [
			"reason",
			"is_false_positive"
		]
		}
	},
	"required": [
		"object_type",
		"object_ids",
		"data"
	]
	}`

	ThreatDataListBulkUpdateAnalystTLP = `{
	"type": "object",
	"description" : "This is the payload for updating the analyst TLP of threat data object in bulk.",
	"properties": {
		"all_objects": {
		"type": "boolean",
		"description" : "This is a flag which is always set to 'true' "
		},
		"objects":{
			"type":"array",
			"description" : "This is a list of threat data objects to be updated with analyst tlp.",
			"items":{
				"type":"object",
				"properties":{
					"type":{
						"type":"string",
						"description" : "This is the type(indicator, report etc) of threat data object"
					},
					"ids":{
						"type":"array",
						"description" : "This is the list of threat data object_id of this specifc object type",
						"items":{
							"type":"string",
							"description" : "This is the threat data object_id"
							}
					}
				},
				"required":[
					"type",
					"ids"
				]
				}
		},
		"data": {
		"type": "object",
		"properties": {
			"reason": {
			"type": "string",
			"description" : "This is the reason for updating analyst tlp of threat data objects."
			},
			"analyst_tlp": {
			"type": "string",
			"description" : "This is the value TLP, it can be RED, AMBER_STRICT, AMBER, GREEN, CLEAR, NONE"
			}
		},
		"required": [
			"reason",
			"analyst_tlp"
		]
		}
	},
	"required": [
		"all_objects",
		"objects",
		"data"
	]
	}`

	ThreatDataListBulkUpdateAnalystScore = `{
	"type": "object",
	"description" : "This is the payload for updating the analyst score of threat data object in bulk.",
	"properties": {
		"all_objects": {
		"type": "boolean",
		"description" : "This is a flag which is always set to 'true' "
		},
		"objects":{
			"type":"array",
			"description" : "This is a list of threat data objects to be updated with analyst score.",
			"items":{
				"type":"object",
				"properties":{
					"type":{
						"type":"string",
						"description" : "This is the type(indicator, report etc) of threat data object"
					},
					"ids":{
						"type":"array",
						"description" : "This is the list of threat data object_id of this specifc object type",
						"items":{
							"type":"string",
							"description" : "This is the threat data object_id"
							}
					}
				},
				"required":[
					"type",
					"ids"
				]
				}
		},
		"data": {
		"type": "object",
		"properties": {
			"reason": {
			"type": "string",
			"description" : "This is the reason for updating analyst score of threat data objects."
			},
			"analyst_score": {
			"type": "string",
			"description" : "This is the value analyst score, it must be between 0-100"
			}
		},
		"required": [
			"reason",
			"analyst_score"
		]
		}
	},
	"required": [
		"all_objects",
		"objects",
		"data"
	]
	}`

	ThreatDataListBulkDeprecate = `{
	"type": "object",
	"description" : "This is the payload used for deprecating indicators.",
	"properties": {
		"object_type": {
		"type": "string",
		"description" : "This must be passed as indicator. Indicators can only be deprecated."

		},
		"object_ids":{
			"type":"array",
			"description" : "This is the list of the object_id which must be passed.",
			"items":
				{
					"type":"string",
					"description" : "This is the object_id of the object"
				}
		},
		"data": {
		"type": "object",
		"properties": {
			"reason": {
			"type": "string",
			"description" : "This is the reason for deprecating threat data objects."

			},
			"is_deprecated": {
			"type": "boolean",
			"description" : "This must be always set to true for deprecating."
			}
		},
		"required": [
			"reason",
			"is_deprecated"
		]
		}
	},
	"required": [
		"object_type",
		"object_ids",
		"data"
	]
	}`

	ThreatDataListBulkUnDeprecate = `{
	"type": "object",
	"description" : "This is the payload used for un-deprecating indicators.",
	"properties": {
		"object_type": {
		"type": "string",
		"description" : "This must be passed as indicator. Indicators can only be deprecated."

		},
		"object_ids":{
			"type":"array",
			"description" : "This is the list of the object_id which must be passed.",
			"items":
				{
					"type":"string",
					"description" : "This is the object_id of the object"
				}
		},
		"data": {
		"type": "object",
		"properties": {
			"reason": {
			"type": "string",
			"description" : "This is the reason for un-deprecating threat data objects."

			},
			"is_deprecated": {
			"type": "boolean",
			"description" : "This must be always set to false for un-deprecating."
			},
			"undeprecate_until":{
			"type" : "integer",
			"description" : "This is epoch time which is must be greater than the current time."
			}
		},
		"required": [
			"reason",
			"is_deprecated"
		]
		}
	},
	"required": [
		"object_type",
		"object_ids",
		"data"
	]
	}`

	ThreatDataListBulkAddWatchlist = `{
	"type": "object",
	"description" : "This is the payload used adding threat data to watchlist.",
	"properties": {
		"all_objects": {
		"type": "boolean",
		"description" : "This must be always set to true"
		},
		"objects": {
		"type": "array",
		"description" : "This is the list of the object_id seggregated with the object type",
		"items": {
			"type": "object",
			"properties": {
				"type": {
				"type": "string",
				"description" : "This is the type of the object eg: threat-actor, malware etc"
				},
				"ids": {
				"type": "array",
					"description" :"This is the list of the object ids",
				"items": {
					"type": "string",
					"description" : "This is the object id"
					}
				}
			},
			"required": [
				"type",
				"ids"
			]
			}
		},
		"data": {
		"type": "object",
		"description" : "This is the list of values(names) of the threat data object",
		"properties": {
			"name": {
			"type": "array",
			"items": {
				"type": "string",
				"description" : "This is value(name) fo the threat data object"
				}
			}
		},
		"required": [
			"name"
		]
		}
	},
	"required": [
		"all_objects",
		"objects",
		"data"
	]
	}`

	ThreatDataListBulkRemoveWatchlist = `{
	"type": "object",
	"description" : "This is the payload used to reomve a number of threat data from watchlist.",
	"properties": {
		"all_objects": {
		"type": "boolean",
		"description" : "This must be always set to true"
		},
		"objects": {
		"type": "array",
		"description" : "This is the list of the object_id seggregated with the object type",
		"items": {
			"type": "object",
			"properties": {
				"type": {
				"type": "string",
				"description" : "This is the type of the object eg: threat-actor, malware etc"
				},
				"ids": {
				"type": "array",
					"description" :"This is the list of the object ids",
				"items": {
					"type": "string",
					"description" : "This is the object id"
					}
				}
			},
			"required": [
				"type",
				"ids"
			]
			}
		},
		"data": {
		"type": "object",
		"description" : "This is the list of values(names) of the threat data object",
		"properties": {
			"name": {
			"type": "array",
			"items": {
				"type": "string",
				"description" : "This is value(name) fo the threat data object"
				}
			}
		},
		"required": [
			"name"
		]
		}
	},
	"required": [
		"all_objects",
		"objects",
		"data"
	]
	}`

	ThreatDataListBulkAddRelation = `{
	"type": "object",
	"description" : "This is the payload used to add relation to a number of threat data objects in CTIX.",
	"properties": {
		"all_objects": {
		"type": "boolean",
		"description" : "This must be always set to true"
		},
		"objects": {
		"type": "array",
		"description" : "This is the list of the object_id seggregated with the object type",
		"items": {
			"type": "object",
			"properties": {
				"type": {
				"type": "string",
				"description" : "This is the type of the object eg: threat-actor, malware etc"
				},
				"ids": {
				"type": "array",
				"description" :"This is the list of the object ids",
				"items": 
					{
					"type": "string",
					"description" : "This is the object id"
					}
				}
			},
			"required": [
				"type",
				"ids"
			]
			}
		},
		"data": {
		"type": "object",
		"description" : "This the details of the target threat data object to which all the mentioned object will be related.",
		"properties": {
			"target": {
			"type": "object",
			"properties": {
				"id": {
				"type": "string",
				"description" : "This is object_id of the target threat data object"
				},
				"name": {
				"type": "string",
				"description" : "This is name of the target threat data object"
				},
				"sub_type": {
				"type": "string",
				"description" : "This is sub_type of the target threat data object. It can be null if object doesn't have this value"
				},
				"type": {
				"type": "string",
				"description" : "This is type of the target threat data object. eg-malware"
				}
			},
			"required": [
				"id",
				"name",
				"sub_type",
				"type"
			]
			},
			"relationship_type": {
			"type": "string",
			"description" :"This is the relation type between the objects, which must be a value from the available relation types list"
			}
		},
		"required": [
			"target",
			"relationship_type"
		]
		}
	},
	"required": [
		"all_objects",
		"objects",
		"data"
	]
	}`
)

// This function returns an action map for each threat data bulk actions
func GetThreatDataBulkActionsMapping() map[string]map[string]string {
	action_map := map[string]map[string]string{
		"ThreatDataListBulkActionAddTag": {
			"endpoint":         bulk_update_add_tag_endpoint,
			"tool_name":        "threat-data-list-bulk-action-add-tag",
			"tool_description": "This tool adds the specified tags to multiple threat data objects in CTIX.",
			"schema":           ThreatDataListBulkActionAddTag,
		},
		"ThreatDataListBulkMarkIndicatorAllowed": {
			"endpoint":  bulk_update_mark_indicator_allowed_endpoint,
			"tool_name": "threat-data-list-bulk-mark-indicator-allowed",
			"tool_description": `Use this to mark a list of indicators as indicator allowed in CTIX.
      							!!-> Always use this tool to mark the indicator as indicator allowed, if its asked to mark from threat data listing`,
			"schema": ThreatDataListBulkMarkIndicatorAllowed,
		},
		"ThreatDataListBulkUnMarkIndicatorAllowed": {
			"endpoint":         bulk_update_unmark_indicator_allowed_endpoint,
			"tool_name":        "threat-data-list-bulk-unmark-indicator-allowed",
			"tool_description": `Use this to unmark/remove a list of indicators from indicator allowed in CTIX.`,
			"schema":           ThreatDataListBulkUnMarkIndicatorAllowed,
		},
		"ThreatDataListBulkManualReview": {
			"endpoint":         bulk_add_manual_review,
			"tool_name":        "threat-data-list-bulk-manual-review",
			"tool_description": `This tool adds number of threat data objects in CTIX for manual review.`,
			"schema":           ThreatDataListBulkManualReview,
		},
		"ThreatDataListBulkMarkFalsePositive": {
			"endpoint":         bulk_mark_false_positive,
			"tool_name":        "threat-data-list-bulk-mark-false-positive",
			"tool_description": `This tool marks number of threat data objects in CTIX as false positive.`,
			"schema":           ThreatDataListBulkMarkFalsePositive,
		},

		"ThreatDataListBulkUnMarkFalsePositive": {
			"endpoint":         bulk_unmark_false_positive,
			"tool_name":        "threat-data-list-bulk-unmark-false-positive",
			"tool_description": `This tool unmarks number of threat data objects in CTIX from false positive.`,
			"schema":           ThreatDataListBulkUnMarkFalsePositive,
		},
		"ThreatDataListBulkUpdateAnalystTLP": {
			"endpoint":         bulk_update_analyst_tlp,
			"tool_name":        "threat-data-list-bulk-update-analyst-tlp",
			"tool_description": `This tool updates analyst tlp in bulk for a number threat data objects in CTIX.`,
			"schema":           ThreatDataListBulkUpdateAnalystTLP,
		},
		"ThreatDataListBulkUpdateAnalystScore": {
			"endpoint":         bulk_update_analyst_score,
			"tool_name":        "threat-data-list-bulk-update-analyst-score",
			"tool_description": `This tool updates analyst score in bulk for a number threat data objects in CTIX.`,
			"schema":           ThreatDataListBulkUpdateAnalystScore,
		},
		"ThreatDataListBulkDeprecate": {
			"endpoint":         bulk_deprecate,
			"tool_name":        "threat-data-list-bulk-deprecate",
			"tool_description": `This tool deprecates number of threat data objects in CTIX.`,
			"schema":           ThreatDataListBulkDeprecate,
		},
		"ThreatDataListBulkUnDeprecate": {
			"endpoint":         bulk_undeprecate,
			"tool_name":        "threat-data-list-bulk-undeprecate",
			"tool_description": `This tool un-deprecates number of threat data objects in CTIX.`,
			"schema":           ThreatDataListBulkUnDeprecate,
		},
		"ThreatDataListBulkAddWatchlist": {
			"endpoint":         bulk_add_watchlist,
			"tool_name":        "threat-data-list-bulk-add-watchlist",
			"tool_description": `This tool add number of threat data objects to watchlist in CTIX.`,
			"schema":           ThreatDataListBulkAddWatchlist,
		},
		"ThreatDataListBulkRemoveWatchlist": {
			"endpoint":         bulk_remove_watchlist,
			"tool_name":        "threat-data-list-bulk-remove-watchlist",
			"tool_description": `This tool removes number of threat data objects from watchlist in CTIX.`,
			"schema":           ThreatDataListBulkRemoveWatchlist,
		},
		"ThreatDataListBulkAddRelation": {
			"endpoint":         bulk_add_relation,
			"tool_name":        "threat-data-list-bulk-add-relation",
			"tool_description": `This tool add a relation to a number of threat data objects in CTIX.`,
			"schema":           ThreatDataListBulkAddRelation,
		},
	}
	return action_map
}
