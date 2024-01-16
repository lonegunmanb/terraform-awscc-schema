package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBillingconductorPricingPlan = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Pricing Plan ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "Creation timestamp in UNIX epoch time format",
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "Latest modified timestamp in UNIX epoch time format",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pricing_rule_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "size": {
        "computed": true,
        "description": "Number of associated pricing rules",
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::BillingConductor::PricingPlan",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBillingconductorPricingPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBillingconductorPricingPlan), &result)
	return &result
}
