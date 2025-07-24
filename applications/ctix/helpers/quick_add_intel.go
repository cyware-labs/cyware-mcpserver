package helpers

const (
	Quick_add_intel_schema = `{
  "type": "object",
   "description" : "This is the payload for creating intel in CTIX using quick add intel flow",
  "properties": {
    "context": {
      "type": "string",
      "description": "The value of this is always QUICK_ADD_INTEL_FLOW for quick add intel module"
    },
    "metadata": {
      "type": "object",
      "description": "This is metadata for all the SDO, point to note here is that the description is applied only to report object even if apply to all is set to true",
      "properties": {
        "tlp": {
          "type": "string",
          "description": "This is the TLP which is applied to all of the objects created. Value of this is string from this list (RED, AMBER_STRICT, AMBER, GREEN, CLEAR, NONE)"
        },
        "default_marking_definition": {
          "type": "string",
          "description": "This is the TLP which is applied to all of the objects created. Value of this is string from this list (RED, AMBER_STRICT, AMBER, GREEN, CLEAR, NONE)"
        },
        "marking_config": {
          "type": "string",
          "description": "This is the marking config, having value alway as tlp2"
        },
        "is_apply_all": {
          "type": "boolean",
          "description": "This is a flag which represents if the metadata should be applied to all of the object"
        },
        "confidence": {
          "type": "integer",
          "description": "This is the source confidence score value which ranges from 0-100, if not given take 100"
        },
        "description": {
          "type": "string",
          "description": "This is description of the report object"
        }
      },
      "required": [
        "tlp",
        "default_marking_definition",
        "marking_config",
        "is_apply_all",
        "confidence"
      ]
    },
    "indicators": {
      "type": "object",
      "description": "These are the indicators SDO which will be created",
      "properties": {
        "ipv4-addr": {
          "type": "string",
          "description": "this is a list of ipv4 separated by newline"
        },
        "ipv6-addr": {
          "type": "string",
          "description": "this is a list of ipv6 separated by newline"
        },
        "domain": {
          "type": "string",
          "description": "this is a list of domain separated by newline"
        },
        "url": {
          "type": "string",
          "description": "this is a list of ulr separated by newline"
        },
        "email": {
          "type": "string",
          "description": "this is a list of email addresses separated by newline"
        },
        "md5": {
          "type": "string",
          "description": "this is a list of md5 separated by newline"
        },
        "sha1": {
          "type": "string",
          "description": "this is a list of sha1 separated by newline"
        },
        "sha224": {
          "type": "string",
          "description": "this is a list of sha224 separated by newline"
        },
        "sha256": {
          "type": "string",
          "description": "this is a list of sha256 separated by newline"
        },
        "sha512": {
          "type": "string",
          "description": "this is a list of sha512 separated by newline"
        },
        "sha384": {
          "type": "string",
          "description": "this is a list of sha384 separated by newline"
        },
        "ssdeep": {
          "type": "string",
          "description": "this is a list of ssdeep separated by newline"
        },
        "autonomous-system": {
          "type": "string",
          "description": "this is a list of autonomous-system separated by newline"
        },
        "windows-registry-key": {
          "type": "string",
          "description": "this is a list of windows-registry-key separated by newline"
        }
      }
    },
    "title": {
      "type": "string"
    },
    "sdos": {
      "type": "object",
      "description": "These are other type SDO(stix domain objects) other than indicators",
      "properties": {
        "vulnerability": {
          "type": "string",
          "description": "this is a list of vulnerability separated by newline"
        },
        "intrusion-set": {
          "type": "string",
          "description": "this is a list of intrusion-set separated by newline"
        },
        "malware": {
          "type": "string",
          "description": "this is a list of malware separated by newline"
        },
        "campaign": {
          "type": "string",
          "description": "this is a list of campaign separated by newline"
        },
        "threat-actor": {
          "type": "string",
          "description": "this is a list of threat-actor separated by newline"
        },
        "attack-pattern": {
          "type": "string",
          "description": "this is a list of attack-pattern separated by newline"
        },
        "incident": {
          "type": "string",
          "description": "this is a list of incidents separated by newline"
        },
        "course-of-action": {
          "type": "string",
          "description": "this is a list of course-of-action separated by newline"
        },
        "identity": {
          "type": "string",
          "description": "this is a list of identity separated by newline"
        },
        "tool": {
          "type": "string",
          "description": "this is a list of tool separated by newline"
        },
        "infrastructure": {
          "type": "string",
          "description": "this is a list of infrastructure separated by newline"
        },
        "location": {
          "type": "object",
          "description": "this is the location object",
          "properties": {
            "type": {
              "type": "string",
              "description": "This is type of location, here its country"
            },
            "values": {
              "type": "array",
              "description": "this is a list of countries",
              "items": {
                "type": "object",
                "properties": {
                  "value": {
                    "type": "string",
                    "description": "this is the country code for the country for example IN for india"
                  },
                  "label": {
                    "type": "string",
                    "description": "this is the actual country name"
                  }
                },
                "required": [
                  "value",
                  "label"
                ]
              }
            }
          },
          "required": [
            "type",
            "values"
          ]
        },
        "malware-analysis": {
          "type": "string",
          "description": "this is a list of malware-analysis separated by newline"
        }
      }
    },
    "observables": {
      "type": "object",
      "description": "This is observables which can be create by quick add intel flow.",
      "properties": {
        "artifact": {
          "type": "string",
          "description": "this is a list of artifact separated by 2 newlines. The value should base64 value of any data"
        },
        "directory": {
          "type": "string",
          "description": "this is a list of valid directories separated by newline"
        },
        "mac-addr": {
          "type": "string",
          "description": "this is a list of mac-add separated by newline"
        },
        "email-message": {
          "type": "string",
          "description": "this is a list of email-messages separated by 2 newlines"
        },
        "mutex": {
          "type": "string",
          "description": "this is a list of mutex separated by newline"
        },
        "network-traffic": {
          "type": "string",
          "description": "this is a list of network-traffic separated by newline"
        },
        "process": {
          "type": "string",
          "description": "this is a list of process separated by newline"
        },
        "software": {
          "type": "string",
          "description": "this is a list of software separated by newline"
        },
        "user-account": {
          "type": "string",
          "description": "this is a list of user-account separated by newline"
        },
        "x509-certificate": {
          "type": "string",
          "description": "this is a list of valid x509 certificates having valid expiry date, separated by 2 newlines"
        },
        "file": {
          "type": "string",
          "description": "this is a list of file separated by newline, eg adb.txt"
        }
      }
    },
    "create_intel_feed": {
      "type": "boolean",
      "description": "this is a flag which is always set to true for the intel creation"
    }
  },
  "required": [
    "context",
    "metadata",
    "title",
    "create_intel_feed"
  ]
}`
)
