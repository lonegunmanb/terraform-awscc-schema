package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBillingconductorBillingGroup = `{
  "block": {
    "attributes": {
      "account_grouping": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_associate": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "linked_account_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "arn": {
        "computed": true,
        "description": "Billing Group ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "computation_preference": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "pricing_plan_arn": {
              "description": "ARN of the attached pricing plan",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "Latest modified timestamp in UNIX epoch time format",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "primary_account_id": {
        "description": "This account will act as a virtual payer account of the billing group",
        "description_kind": "plain",
        "required": true,
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
      }
    },
    "description": "A billing group is a set of linked account which belong to the same end customer. It can be seen as a virtual consolidated billing family.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBillingconductorBillingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBillingconductorBillingGroup), &result)
	return &result
}
