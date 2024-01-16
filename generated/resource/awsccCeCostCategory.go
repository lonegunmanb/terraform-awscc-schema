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
      }
    },
    "description": "Cost Category enables you to map your cost and usage into meaningful categories. You can use Cost Category to organize your costs using a rule-based engine.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCeCostCategorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCeCostCategory), &result)
	return &result
}
