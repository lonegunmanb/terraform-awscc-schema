package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIottwinmakerSyncJob = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the SyncJob.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date_time": {
        "computed": true,
        "description": "The date and time when the sync job was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of SyncJob.",
        "description_kind": "plain",
        "type": "string"
      },
      "sync_role": {
        "computed": true,
        "description": "The IAM Role that execute SyncJob.",
        "description_kind": "plain",
        "type": "string"
      },
      "sync_source": {
        "computed": true,
        "description": "The source of the SyncJob.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "update_date_time": {
        "computed": true,
        "description": "The date and time when the sync job was updated.",
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
    "description": "Data Source schema for AWS::IoTTwinMaker::SyncJob",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIottwinmakerSyncJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIottwinmakerSyncJob), &result)
	return &result
}
