package resource

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
        "description": "The name of the destination resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_policy": {
        "computed": true,
        "description": "An IAM policy document that governs which AWS accounts can create subscription filters against this destination.",
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
      "role_arn": {
        "description": "The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_arn": {
        "description": "The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::Logs::Destination resource specifies a CloudWatch Logs destination. A destination encapsulates a physical resource (such as an Amazon Kinesis data stream) and enables you to subscribe that resource to a stream of log events.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsDestination), &result)
	return &result
}
