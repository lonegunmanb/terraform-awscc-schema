package data

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
        "computed": true,
        "description": "The ARN of the dataset group that the filter belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "filter_arn": {
        "computed": true,
        "description": "The ARN of the filter.",
        "description_kind": "plain",
        "type": "string"
      },
      "filter_expression": {
        "computed": true,
        "description": "The filter expression that defines which items are included or excluded from recommendations.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_date_time": {
        "computed": true,
        "description": "The time at which the filter was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the filter.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Personalize::Filter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPersonalizeFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPersonalizeFilter), &result)
	return &result
}
