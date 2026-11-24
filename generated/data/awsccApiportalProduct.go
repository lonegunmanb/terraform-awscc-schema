package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalProduct = `{
  "block": {
    "attributes": {
      "api_selection_configuration": {
        "computed": true,
        "description": "Configuration controlling how APIs are selected for the product.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "api_selection_mode": {
              "computed": true,
              "description": "The mode used to select APIs for the product.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of the product.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the product.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_arn": {
        "computed": true,
        "description": "The ARN of the product.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_id": {
        "computed": true,
        "description": "The product identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the product.",
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
        "description": "The timestamp when the product was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::APIPortal::Product",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApiportalProductSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalProduct), &result)
	return &result
}
