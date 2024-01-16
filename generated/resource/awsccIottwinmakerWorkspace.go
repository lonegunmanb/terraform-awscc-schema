package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIottwinmakerWorkspace = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date_time": {
        "computed": true,
        "description": "The date and time when the workspace was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the workspace.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "role": {
        "description": "The ARN of the execution role associated with the workspace.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_location": {
        "description": "The ARN of the S3 bucket where resources associated with the workspace are stored.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of key-value pairs to associate with a resource.",
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
        "description": "The ID of the workspace.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::IoTTwinMaker::Workspace",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIottwinmakerWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIottwinmakerWorkspace), &result)
	return &result
}
