package helpers

const Execute_playbook_schema = `{
  "type": "object",
  "description" : "This is the payload used to execute a playbook.",
  "properties": {
    "input_type": {
      "type": "integer",
      "description" : "This represents the input type. Pass this as 3 if input is required for playbook else pass 4 if its asked to run on the previous input."
    },
    "input": {
      "type": "string",
      "description" : "This is the input to the playboook. It must be json which is passed as a string. Input can decided based on the playbook details."
    },
    "pbhash": {
      "type": "string",
      "description" : "This is the hash of the playbook which can be fetched from the playbook details."
    }
  },
  "required": [
    "input_type",
    "input",
    "pbhash"
  ]
}`
