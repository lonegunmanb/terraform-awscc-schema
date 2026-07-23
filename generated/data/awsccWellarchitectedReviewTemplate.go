package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWellarchitectedReviewTemplate = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The review template description.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "lenses": {
        "computed": true,
        "description": "The lenses applied to the review template.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "notes": {
        "computed": true,
        "description": "The notes associated with the review template.",
        "description_kind": "plain",
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
      },
      "template_arn": {
        "computed": true,
        "description": "The review template ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_name": {
        "computed": true,
        "description": "The name of the review template.",
        "description_kind": "plain",
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
    "description": "Data Source schema for AWS::WellArchitected::ReviewTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWellarchitectedReviewTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWellarchitectedReviewTemplate), &result)
	return &result
}
