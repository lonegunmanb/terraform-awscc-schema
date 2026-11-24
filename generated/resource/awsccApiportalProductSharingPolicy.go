package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalProductSharingPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "description": "JSON IAM-style resource-based policy document. Size limit raised to 409,600 chars per RAM's recommended public-policy onboarding guidance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_id": {
        "description": "The identifier of the parent Product. Singleton key for this CFN type.",
        "description_kind": "plain",
        "required": true,
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
    "description": "Resource Type definition for AWS::APIPortal::ProductSharingPolicy. Singleton sub-property of an AWS::APIPortal::Product capturing the resource-based sharing policy attached via RAM.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApiportalProductSharingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalProductSharingPolicy), &result)
	return &result
}
