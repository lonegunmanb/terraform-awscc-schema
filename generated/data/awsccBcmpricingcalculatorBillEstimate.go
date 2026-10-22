package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBcmpricingcalculatorBillEstimate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the bill estimate.",
        "description_kind": "plain",
        "type": "string"
      },
      "bill_estimate_id": {
        "computed": true,
        "description": "The unique identifier of the bill estimate",
        "description_kind": "plain",
        "type": "string"
      },
      "bill_interval": {
        "computed": true,
        "description": "The time period covered by the bill estimate",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "start": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "bill_scenario_id": {
        "computed": true,
        "description": "The unique identifier of the bill scenario this estimate is based on",
        "description_kind": "plain",
        "type": "string"
      },
      "cost_category_group_sharing_preference_arn": {
        "computed": true,
        "description": "The ARN of the cost category group sharing preference",
        "description_kind": "plain",
        "type": "string"
      },
      "cost_category_group_sharing_preference_effective_date": {
        "computed": true,
        "description": "The effective date of the cost category group sharing preference",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the bill estimate was created",
        "description_kind": "plain",
        "type": "string"
      },
      "expires_at": {
        "computed": true,
        "description": "The timestamp when the bill estimate expires",
        "description_kind": "plain",
        "type": "string"
      },
      "group_sharing_preference": {
        "computed": true,
        "description": "The group sharing preference for the bill estimate",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the bill estimate",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the bill estimate",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::BcmPricingCalculator::BillEstimate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBcmpricingcalculatorBillEstimateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBcmpricingcalculatorBillEstimate), &result)
	return &result
}
