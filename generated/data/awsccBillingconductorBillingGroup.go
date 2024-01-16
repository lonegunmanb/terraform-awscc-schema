package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBillingconductorBillingGroup = `{
  "block": {
    "attributes": {
      "account_grouping": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_associate": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "linked_account_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "arn": {
        "computed": true,
        "description": "Billing Group ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "computation_preference": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "pricing_plan_arn": {
              "computed": true,
              "description": "ARN of the attached pricing plan",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
      "primary_account_id": {
        "computed": true,
        "description": "This account will act as a virtual payer account of the billing group",
        "description_kind": "plain",
        "type": "string"
      },
      "size": {
        "computed": true,
        "description": "Number of accounts in the billing group",
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
    "description": "Data Source schema for AWS::BillingConductor::BillingGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBillingconductorBillingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBillingconductorBillingGroup), &result)
	return &result
}
