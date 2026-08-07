package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftSnapshotSchedule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the snapshot schedule.",
        "description_kind": "plain",
        "type": "string"
      },
      "associated_cluster_count": {
        "computed": true,
        "description": "The number of clusters associated with the schedule.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_definitions": {
        "computed": true,
        "description": "The definition of the snapshot schedule. The definition is made up of schedule expressions, for example \"cron(30 12 *)\" or \"rate(12 hours)\".",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "schedule_description": {
        "computed": true,
        "description": "The description of the snapshot schedule.",
        "description_kind": "plain",
        "type": "string"
      },
      "schedule_identifier": {
        "computed": true,
        "description": "A unique identifier for the snapshot schedule. Only alphanumeric characters are allowed.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An optional set of tags for the snapshot schedule.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key, or name, for the resource tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the resource tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Redshift::SnapshotSchedule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRedshiftSnapshotScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftSnapshotSchedule), &result)
	return &result
}
