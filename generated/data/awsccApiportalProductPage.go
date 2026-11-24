package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalProductPage = `{
  "block": {
    "attributes": {
      "display_content": {
        "computed": true,
        "description": "The displayed content of the product page.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "body": {
              "computed": true,
              "description": "Markdown-formatted body of the product page.",
              "description_kind": "plain",
              "type": "string"
            },
            "title": {
              "computed": true,
              "description": "The title of the product page.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_id": {
        "computed": true,
        "description": "The identifier of the parent Product.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_page_arn": {
        "computed": true,
        "description": "The ARN of the product page.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_page_id": {
        "computed": true,
        "description": "The product page identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the product page.",
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
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the product page was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::APIPortal::ProductPage",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApiportalProductPageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalProductPage), &result)
	return &result
}
