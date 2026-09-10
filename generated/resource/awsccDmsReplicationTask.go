package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsReplicationTask = `{
  "block": {
    "attributes": {
      "cdc_start_position": {
        "computed": true,
        "description": "Indicates when you want a change data capture (CDC) operation to start. Use either CdcStartPosition or CdcStartTime to specify when you want a CDC operation to start. Specifying both values results in an error.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cdc_start_time": {
        "computed": true,
        "description": "Indicates the start time for a change data capture (CDC) operation. Use either CdcStartTime or CdcStartPosition to specify when you want a CDC operation to start. Specifying both values results in an error.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cdc_stop_position": {
        "computed": true,
        "description": "Indicates when you want a change data capture (CDC) operation to stop. The value can be either server time or commit time.",
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
      "migration_type": {
        "description": "The migration type.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_instance_arn": {
        "description": "The Amazon Resource Name (ARN) of a replication instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_task_arn": {
        "computed": true,
        "description": "The ARN of the ReplicationTask. Also serves the purpise of Primary Identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_task_identifier": {
        "computed": true,
        "description": "An identifier for the replication task.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_task_settings": {
        "computed": true,
        "description": "Overall settings for the task, in JSON format",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_identifier": {
        "computed": true,
        "description": "A friendly name for the resource identifier at the end of the EndpointArn response parameter that is returned in the created Endpoint object.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_endpoint_arn": {
        "description": "An Amazon Resource Name (ARN) that uniquely identifies the source endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "table_mappings": {
        "description": "The table mappings for the task, in JSON format.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag key",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag value",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target_endpoint_arn": {
        "description": "An Amazon Resource Name (ARN) that uniquely identifies the target endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "task_data": {
        "computed": true,
        "description": "Supplemental information that the task requires to migrate the data for certain source and target endpoints.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::DMS::ReplicationTask",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDmsReplicationTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsReplicationTask), &result)
	return &result
}
