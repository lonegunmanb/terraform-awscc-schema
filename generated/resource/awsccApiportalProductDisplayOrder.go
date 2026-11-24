package resource

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
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "section_name": {
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "overview_page_arn": {
        "computed": true,
        "description": "ARN of the ProductPage rendered as the overview/landing page.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "product_id": {
        "description": "The identifier of the parent Product. Singleton key for this CFN type.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_page_arns": {
        "computed": true,
        "description": "Top-level ProductPage ARNs in display order.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Resource Type definition for AWS::APIPortal::ProductDisplayOrder. Singleton sub-property of an AWS::APIPortal::Product capturing its display layout (overview page, top-level ProductPage ordering, and sectioned ProductApiPage arrangement).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApiportalProductDisplayOrderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalProductDisplayOrder), &result)
	return &result
}
