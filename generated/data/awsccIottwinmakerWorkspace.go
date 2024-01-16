package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role": {
        "computed": true,
        "description": "The ARN of the execution role associated with the workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_location": {
        "computed": true,
        "description": "The ARN of the S3 bucket where resources associated with the workspace are stored.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of key-value pairs to associate with a resource.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "The ID of the workspace.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoTTwinMaker::Workspace",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIottwinmakerWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIottwinmakerWorkspace), &result)
	return &result
}
