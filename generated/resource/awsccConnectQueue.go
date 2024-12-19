package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectQueue = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the queue.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hours_of_operation_arn": {
        "description": "The identifier for the hours of operation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_contacts": {
        "computed": true,
        "description": "The maximum number of contacts that can be in the queue before it is considered full.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "The name of the queue.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "outbound_caller_config": {
        "computed": true,
        "description": "The outbound caller ID name, number, and outbound whisper flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "outbound_caller_id_name": {
              "computed": true,
              "description": "The caller ID name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "outbound_caller_id_number_arn": {
              "computed": true,
              "description": "The caller ID number.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "outbound_flow_arn": {
              "computed": true,
              "description": "The outbound whisper flow to be used during an outbound call.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "outbound_email_config": {
        "computed": true,
        "description": "The outbound email address ID.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "outbound_email_address_id": {
              "computed": true,
              "description": "The email address connect resource ID.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "queue_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "quick_connect_arns": {
        "computed": true,
        "description": "The quick connects available to agents who are working the queue.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description": "The status of the queue.",
        "description_kind": "plain",
        "optional": true,
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
      },
      "type": {
        "computed": true,
        "description": "The type of queue.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Connect::Queue",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectQueue), &result)
	return &result
}
