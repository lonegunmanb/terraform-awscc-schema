package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectInstance = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "An instanceArn is automatically generated on creation based on instanceId.",
        "description_kind": "plain",
        "type": "string"
      },
      "attributes": {
        "description": "The attributes for the instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_resolve_best_voices": {
              "computed": true,
              "description": "Boolean flag which enables AUTO_RESOLVE_BEST_VOICES on an instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "contact_lens": {
              "computed": true,
              "description": "Boolean flag which enables CONTACT_LENS on an instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "contactflow_logs": {
              "computed": true,
              "description": "Boolean flag which enables CONTACTFLOW_LOGS on an instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "early_media": {
              "computed": true,
              "description": "Boolean flag which enables EARLY_MEDIA on an instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "inbound_calls": {
              "description": "Mandatory element which enables inbound calls on new instance.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "outbound_calls": {
              "description": "Mandatory element which enables outbound calls on new instance.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "use_custom_tts_voices": {
              "computed": true,
              "description": "Boolean flag which enables USE_CUSTOM_TTS_VOICES on an instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "created_time": {
        "computed": true,
        "description": "Timestamp of instance creation logged as part of instance creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
        "computed": true,
        "description": "Existing directoryId user wants to map to the new Connect instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_management_type": {
        "description": "Specifies the type of directory integration for new instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_alias": {
        "computed": true,
        "description": "Alias of the new directory created as part of new instance creation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "An instanceId is automatically generated on creation and assigned as the unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_status": {
        "computed": true,
        "description": "Specifies the creation status of new instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_role": {
        "computed": true,
        "description": "Service linked role created as part of instance creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Connect::Instance",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectInstance), &result)
	return &result
}
