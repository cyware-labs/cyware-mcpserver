package helpers

const (
	Create_tag_schema = `{
	"type": "object",
	"properties": {
		"name": {
			"type": "string",
			"description": "This is a comma separated list of names of tag, which should be created."
		},
		"tag_type": {
			"type": "string",
			"description": "This is the type of the tag. It can be any of 'system' (internal tag), 'user' (created by user), 'privileged' (tags having restriction and allowed specific user groups). We can't create tag of source category. And when privillege category is used then we must pass the user group key. If the user doesn't specify the tag category, then use 'user' as the category."
		},
		"user_groups": {
			"type": "array",
			"description" : "This is a list of user groups that must be provided only when the tag category is 'privileged'. It must not be included for any other category. User group details can be fetched using the User Group Listing API.",
			"items": {
			"type": "object",
			"properties": {
				"id": {
				"type": "string",
				"description": "This is the user group id."
				},
				"name": {
				"type": "string",
				"description": "This is the user group name."
				},
				"description": {
				"type": "string",
				"description": "This is the user group description."
				},
				"created_by": {
				"type": "object",
				"description": "This is the details about who has created the group.",
				"properties": {
					"id": {
					"type": "string",
					"description": "This is the user id who has created the user group."
					},
					"first_name": {
					"type": "string",
					"description": "This is the user first name who has created the user group."
					},
					"last_name": {
					"type": "string",
					"description": "This is the user last name who has created the user group."
					},
					"email": {
					"type": "string",
					"description": "This is the user email who has created the user group."
					},
					"is_active": {
					"type": "boolean",
					"description": "This is the user active status who has created the user group."
					},
					"contact_number": {
					"type": "string",
					"description": "This is the user contact number who has created the user group."
					},
					"country_code": {
					"type": "string",
					"description": "This is the user country code who has created the user group."
					},
					"is_read_only": {
					"type": "boolean",
					"description": "This readonly flag for the user who has created the user group."
					}
				},
				"required": [
					"id",
					"first_name",
					"last_name",
					"email",
					"is_active",
					"contact_number",
					"country_code",
					"is_read_only"
				]
				},
				"is_editable": {
				"type": "boolean",
				"description": "This is the flag which says if user group is editable."
				},
				"is_active": {
				"type": "boolean",
				"description": "This is the flag which says if user group is active."
				},
				"created": {
				"type": "integer",
				"description": "This is date when the user group was created."
				},
				"permission_count": {
				"type": "integer",
				"description": "This is the number of permission which this user group has."
				},
				"user_count": {
				"type": "integer",
				"description": "This is number of user which are associated to this user group."
				},
				"is_read_only": {
				"type": "boolean",
				"description": "This is the flag which says if user group is read only user group."
				},
				"is_default": {
				"type": "boolean",
				"description": "This is the flag which says if user group is a default user group."
				},
				"saml_associated_groups": {
				"type": "array",
				"description": "This a list of SAML group associated the user group.",
				"items": 
					{
					"type": "string",
					"description": "This is the name of the SAML group"
					}
				}
			},
			"required": [
				"id",
				"name",
				"description",
				"created_by",
				"is_editable",
				"is_active",
				"created",
				"permission_count",
				"user_count",
				"is_read_only",
				"is_default",
				"saml_associated_groups"
			]
			}
		},
		"action": {
			"type": "string",
			"description" : "Action must be 'add' for tag creation."
		}
	},
	"required": [
		"name",
		"tag_type",
		"user_groups",
		"action"
	]
}
`
)
