package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWellarchitectedReviewTemplate = `{
  "block": {
    "attributes": {
      "description": {
        "description": "The review template description.",
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
      "lenses": {
        "description": "The lenses applied to the review template.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "notes": {
        "computed": true,
        "description": "The notes associated with the review template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description": "The owner of the review template.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the review template.",
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
      },
      "template_arn": {
        "computed": true,
        "description": "The review template ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_name": {
        "description": "The name of the review template.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_status": {
        "computed": true,
        "description": "The latest status of the review template.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time the review template was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates a review template for the Well-Architected Tool.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWellarchitectedReviewTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWellarchitectedReviewTemplate), &result)
	return &result
}
