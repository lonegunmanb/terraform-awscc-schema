package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2PortalProduct = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A description of the portal product.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The name of the portal product as it appears in a published portal.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified": {
        "computed": true,
        "description": "The timestamp when the portal product was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_product_arn": {
        "computed": true,
        "description": "The ARN of the portal product.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_product_id": {
        "computed": true,
        "description": "The portal product identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The collection of tags associated with the portal product.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
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
    "description": "Represents a portal product in Amazon API Gateway V2, which is a logical grouping of APIs that can be published in a developer portal.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayv2PortalProductSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2PortalProduct), &result)
	return &result
}
