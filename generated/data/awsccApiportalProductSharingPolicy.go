package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalProductSharingPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "JSON IAM-style resource-based policy document. Size limit raised to 409,600 chars per RAM's recommended public-policy onboarding guidance.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_id": {
        "computed": true,
        "description": "The identifier of the parent Product. Singleton key for this CFN type.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the product sharing policy.",
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
    "description": "Data Source schema for AWS::APIPortal::ProductSharingPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApiportalProductSharingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalProductSharingPolicy), &result)
	return &result
}
