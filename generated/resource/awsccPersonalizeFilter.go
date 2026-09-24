package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPersonalizeFilter = `{
  "block": {
    "attributes": {
      "creation_date_time": {
        "computed": true,
        "description": "The time at which the filter was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "dataset_group_arn": {
        "description": "The ARN of the dataset group that the filter belongs to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filter_arn": {
        "computed": true,
        "description": "The ARN of the filter.",
        "description_kind": "plain",
        "type": "string"
      },
      "filter_expression": {
        "description": "The filter expression that defines which items are included or excluded from recommendations.",
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
      "last_updated_date_time": {
        "computed": true,
        "description": "The time at which the filter was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the filter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the filter.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to associate with the filter.",
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "A recommendation filter that defines which items are included or excluded from recommendations.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPersonalizeFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPersonalizeFilter), &result)
	return &result
}
