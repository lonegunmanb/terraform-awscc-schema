package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBcmpricingcalculatorWorkloadEstimate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the workload estimate.",
        "description_kind": "plain",
        "type": "string"
      },
      "cost_currency": {
        "computed": true,
        "description": "The currency of the estimated cost.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the workload estimate was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "expires_at": {
        "computed": true,
        "description": "The timestamp when the workload estimate will expire.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failure_message": {
        "computed": true,
        "description": "An error message if the workload estimate failed.",
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
        "description": "The name of the workload estimate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rate_timestamp": {
        "computed": true,
        "description": "The timestamp of the pricing rates used for the estimate.",
        "description_kind": "plain",
        "type": "string"
      },
      "rate_type": {
        "computed": true,
        "description": "The type of pricing rates used for the estimate.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the workload estimate.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "total_cost": {
        "computed": true,
        "description": "The total estimated cost for the workload.",
        "description_kind": "plain",
        "type": "number"
      },
      "workload_estimate_id": {
        "computed": true,
        "description": "The unique identifier of the workload estimate.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::BcmPricingCalculator::WorkloadEstimate",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBcmpricingcalculatorWorkloadEstimateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBcmpricingcalculatorWorkloadEstimate), &result)
	return &result
}
