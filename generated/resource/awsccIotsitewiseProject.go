package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewiseProject = `{
  "block": {
    "attributes": {
      "asset_ids": {
        "computed": true,
        "description": "The IDs of the assets to be associated to the project.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_id": {
        "description": "The ID of the portal in which to create the project.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_arn": {
        "computed": true,
        "description": "The ARN of the project.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_description": {
        "computed": true,
        "description": "A description for the project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project_id": {
        "computed": true,
        "description": "The ID of the project.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_name": {
        "description": "A friendly name for the project.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the project.",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::IoTSiteWise::Project",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotsitewiseProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewiseProject), &result)
	return &result
}
