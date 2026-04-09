package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBcmpricingcalculatorBillScenario = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the bill scenario.",
        "description_kind": "plain",
        "type": "string"
      },
      "bill_interval": {
        "computed": true,
        "description": "The time period covered by the bill scenario",
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
        "description": "The unique identifier of the bill scenario",
        "description_kind": "plain",
        "type": "string"
      },
      "cost_category_group_sharing_preference_arn": {
        "computed": true,
        "description": "The ARN of the cost category group sharing preference",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the bill scenario was created",
        "description_kind": "plain",
        "type": "string"
      },
      "expires_at": {
        "computed": true,
        "description": "The timestamp when the bill scenario expires",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failure_message": {
        "computed": true,
        "description": "The failure message if the bill scenario failed",
        "description_kind": "plain",
        "type": "string"
      },
      "group_sharing_preference": {
        "computed": true,
        "description": "The group sharing preference for the bill scenario",
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
      "name": {
        "computed": true,
        "description": "The name of the bill scenario",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the bill scenario",
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
      }
    },
    "description": "Resource Type definition for AWS::BcmPricingCalculator::BillScenario",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBcmpricingcalculatorBillScenarioSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBcmpricingcalculatorBillScenario), &result)
	return &result
}
