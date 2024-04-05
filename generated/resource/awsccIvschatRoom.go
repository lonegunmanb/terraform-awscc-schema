package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvschatRoom = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Room ARN is automatically generated on creation and assigned as the unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging_configuration_identifiers": {
        "computed": true,
        "description": "Array of logging configuration identifiers attached to the room.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "maximum_message_length": {
        "computed": true,
        "description": "The maximum number of characters in a single message.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "maximum_message_rate_per_second": {
        "computed": true,
        "description": "The maximum number of messages per second that can be sent to the room.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "message_review_handler": {
        "computed": true,
        "description": "Configuration information for optional review of messages.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fallback_result": {
              "computed": true,
              "description": "Specifies the fallback behavior if the handler does not return a valid response, encounters an error, or times out.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uri": {
              "computed": true,
              "description": "Identifier of the message review handler.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "name": {
        "computed": true,
        "description": "The name of the room. The value does not need to be unique.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "room_id": {
        "computed": true,
        "description": "The system-generated ID of the room.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource type definition for AWS::IVSChat::Room.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIvschatRoomSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvschatRoom), &result)
	return &result
}
