package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIottwinmakerScene = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the scene.",
        "description_kind": "plain",
        "type": "string"
      },
      "capabilities": {
        "computed": true,
        "description": "A list of capabilities that the scene uses to render.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "content_location": {
        "description": "The relative path that specifies the location of the content definition file.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "creation_date_time": {
        "computed": true,
        "description": "The date and time when the scene was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the scene.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "generated_scene_metadata": {
        "computed": true,
        "description": "A key-value pair of generated scene metadata for the scene.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "scene_id": {
        "description": "The ID of the scene.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scene_metadata": {
        "computed": true,
        "description": "A key-value pair of scene metadata for the scene.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "update_date_time": {
        "computed": true,
        "description": "The date and time of the current update.",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_id": {
        "description": "The ID of the scene.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::IoTTwinMaker::Scene",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIottwinmakerSceneSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIottwinmakerScene), &result)
	return &result
}
