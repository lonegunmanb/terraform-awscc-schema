package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsDestination = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destination_name": {
        "computed": true,
        "description": "The name of the destination resource",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_policy": {
        "computed": true,
        "description": "An IAM policy document that governs which AWS accounts can create subscription filters against this destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource",
        "description_kind": "plain",
        "type": "string"
      },
      "target_arn": {
        "computed": true,
        "description": "The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Logs::Destination",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLogsDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsDestination), &result)
	return &result
}
