package helpers

const Execute_actions_of_app_schema = `{
  "type": "object",
  "description" : "Payload to execute an action of the app",
  "properties": {
    "apphash": {
      "type": "string",
	  "description": "The unique apphash of app"
    },
    "actionid": {
      "type": "string",
	  "description" : "actionid of the specific action, can be fetched from the details of the action"
    },
    "instance": {
      "type": "string",
	  "description" : "slug of the instance configured, can be fetched from getting instances configured in the app"
    },
    "input": {
      "type": "object",
	  "description" : "Input which is required for the action to execute. Input can be generated using the input structure provided in the action details. "
    }
  },
  "required": [
    "apphash",
    "actionid",
    "instance",
    "input"
  ]
}`
