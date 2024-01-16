package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFmsNotificationChannel = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "sns_role_name": {
        "description": "A resource ARN.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sns_topic_arn": {
        "description": "A resource ARN.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Designates the IAM role and Amazon Simple Notification Service (SNS) topic that AWS Firewall Manager uses to record SNS logs.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccFmsNotificationChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFmsNotificationChannel), &result)
	return &result
}
