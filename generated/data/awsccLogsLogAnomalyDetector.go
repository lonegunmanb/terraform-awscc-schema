package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsLogAnomalyDetector = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "Account ID for owner of detector",
        "description_kind": "plain",
        "type": "string"
      },
      "anomaly_detector_arn": {
        "computed": true,
        "description": "ARN of LogAnomalyDetector",
        "description_kind": "plain",
        "type": "string"
      },
      "anomaly_detector_status": {
        "computed": true,
        "description": "Current status of detector.",
        "description_kind": "plain",
        "type": "string"
      },
      "anomaly_visibility_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "creation_time_stamp": {
        "computed": true,
        "description": "When detector was created.",
        "description_kind": "plain",
        "type": "number"
      },
      "detector_name": {
        "computed": true,
        "description": "Name of detector",
        "description_kind": "plain",
        "type": "string"
      },
      "evaluation_frequency": {
        "computed": true,
        "description": "How often log group is evaluated",
        "description_kind": "plain",
        "type": "string"
      },
      "filter_pattern": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time_stamp": {
        "computed": true,
        "description": "When detector was lsat modified.",
        "description_kind": "plain",
        "type": "number"
      },
      "log_group_arn_list": {
        "computed": true,
        "description": "List of Arns for the given log group",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Logs::LogAnomalyDetector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLogsLogAnomalyDetectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsLogAnomalyDetector), &result)
	return &result
}
