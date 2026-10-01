package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworksecuritymanagerRule = `{
  "block": {
    "attributes": {
      "configuration": {
        "computed": true,
        "description": "The rule configuration as a JSON string.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "firewall_type": {
        "computed": true,
        "description": "The type of firewall for this rule.",
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
      "rule_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_description": {
        "computed": true,
        "description": "A description of the rule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_id": {
        "computed": true,
        "description": "The unique identifier of the rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_name": {
        "description": "The name of the rule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_type": {
        "computed": true,
        "description": "The type of rule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the rule.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "version": {
        "computed": true,
        "description": "The version number of the rule.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::NetworkSecurityManager::Rule. Creates and manages a Network Security Manager rule.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworksecuritymanagerRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworksecuritymanagerRule), &result)
	return &result
}
