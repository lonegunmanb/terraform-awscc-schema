package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCeCostCategory = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Cost category ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "default_value": {
        "computed": true,
        "description": "The default value for the cost category",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_start": {
        "computed": true,
        "description": "ISO 8601 date time with offset format",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rules": {
        "description": "JSON array format of Expression in Billing and Cost Management API",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "split_charge_rules": {
        "computed": true,
        "description": "Json array format of CostCategorySplitChargeRule in Billing and Cost Management API",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the cost category.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
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
    "description": "Resource Type definition for AWS::CE::CostCategory. Cost Category enables you to map your cost and usage into meaningful categories. You can use Cost Category to organize your costs using a rule-based engine.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCeCostCategorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCeCostCategory), &result)
	return &result
}
