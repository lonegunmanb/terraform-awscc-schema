package data

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
        "computed": true,
        "description": "The name for the backup.",
        "description_kind": "plain",
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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "table_name": {
        "computed": true,
        "description": "The name of the table to back up.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DynamoDB::Backup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDynamodbBackupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDynamodbBackup), &result)
	return &result
}
