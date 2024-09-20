package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectAgentStatus = `{
  "block": {
    "attributes": {
      "agent_status_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the agent status.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the status.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_order": {
        "computed": true,
        "description": "The display order of the status.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "last_modified_region": {
        "computed": true,
        "description": "Last modified region.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "Last modified time.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description": "The name of the status.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reset_order_number": {
        "computed": true,
        "description": "A number indicating the reset order of the agent status.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "state": {
        "description": "The state of the status.",
        "description_kind": "plain",
        "required": true,
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
        "description": "The type of agent status.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Connect::AgentStatus",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectAgentStatusSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectAgentStatus), &result)
	return &result
}
