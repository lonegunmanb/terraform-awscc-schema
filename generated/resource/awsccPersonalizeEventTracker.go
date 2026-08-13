package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPersonalizeEventTracker = `{
  "block": {
    "attributes": {
      "dataset_group_arn": {
        "description": "The Amazon Resource Name (ARN) of the dataset group that receives the event data.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_tracker_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the event tracker.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name for the event tracker.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the event tracker.",
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
      },
      "tracking_id": {
        "computed": true,
        "description": "The ID of the event tracker. Include this ID in requests to the PutEvents API.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Personalize::EventTracker",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPersonalizeEventTrackerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPersonalizeEventTracker), &result)
	return &result
}
