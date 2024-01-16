package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApsRuleGroupsNamespace = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The RuleGroupsNamespace ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "data": {
        "description": "The RuleGroupsNamespace data.",
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
      "name": {
        "description": "The RuleGroupsNamespace name.",
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
      },
      "workspace": {
        "description": "Required to identify a specific APS Workspace associated with this RuleGroupsNamespace.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "RuleGroupsNamespace schema for cloudformation.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApsRuleGroupsNamespaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApsRuleGroupsNamespace), &result)
	return &result
}
