package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveInputSecurityGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Unique ARN of Input Security Group",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "input_security_group_id": {
        "computed": true,
        "description": "The Id of the Input Security Group",
        "description_kind": "plain",
        "type": "string"
      },
      "inputs": {
        "computed": true,
        "description": "The list of inputs currently using this Input Security Group.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "A collection of key-value pairs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "whitelist_rules": {
        "computed": true,
        "description": "Whitelist rules and their sync status",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr": {
              "computed": true,
              "description": "The IPv4 CIDR that's whitelisted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::MediaLive::InputSecurityGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMedialiveInputSecurityGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveInputSecurityGroup), &result)
	return &result
}
