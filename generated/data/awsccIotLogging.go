package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotLogging = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
        "description_kind": "plain",
        "type": "string"
      },
      "default_log_level": {
        "computed": true,
        "description": "The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
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
        "description": "The ARN of the role that allows IoT to write to Cloudwatch logs.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoT::Logging",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotLoggingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotLogging), &result)
	return &result
}
