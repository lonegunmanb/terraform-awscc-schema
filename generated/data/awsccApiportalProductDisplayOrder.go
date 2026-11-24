package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalProductDisplayOrder = `{
  "block": {
    "attributes": {
      "contents": {
        "computed": true,
        "description": "Sectioned arrangement of ProductApiPage ARNs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "product_api_page_arns": {
              "computed": true,
              "description": "Ordered ProductApiPage ARNs shown under this section. Required; use an empty list for a placeholder section.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "section_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "overview_page_arn": {
        "computed": true,
        "description": "ARN of the ProductPage rendered as the overview/landing page.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_id": {
        "computed": true,
        "description": "The identifier of the parent Product. Singleton key for this CFN type.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_page_arns": {
        "computed": true,
        "description": "Top-level ProductPage ARNs in display order.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::APIPortal::ProductDisplayOrder",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApiportalProductDisplayOrderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalProductDisplayOrder), &result)
	return &result
}
