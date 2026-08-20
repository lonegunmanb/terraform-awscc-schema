package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDynamodbBackup = `{
  "block": {
    "attributes": {
      "backup_arn": {
        "computed": true,
        "description": "The ARN associated with the backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_creation_date_time": {
        "computed": true,
        "description": "The time at which the backup was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_id": {
        "computed": true,
        "description": "The identifier portion of the backup ARN (server-generated).",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_name": {
        "description": "The name for the backup.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "backup_size_bytes": {
        "computed": true,
        "description": "The size of the backup in bytes.",
        "description_kind": "plain",
        "type": "number"
      },
      "backup_status": {
        "computed": true,
        "description": "The current state of the backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_type": {
        "computed": true,
        "description": "The type of backup (USER, SYSTEM, or AWS_BACKUP).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_name": {
        "description": "The name of the table to back up.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Creates an on-demand backup of a DynamoDB table.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDynamodbBackupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDynamodbBackup), &result)
	return &result
}
