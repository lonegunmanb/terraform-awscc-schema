package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotLogging = `{
  "block": {
    "attributes": {
      "account_id": {
        "description": "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_log_level": {
        "description": "The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "description": "The ARN of the role that allows IoT to write to Cloudwatch logs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Logging Options enable you to configure your IoT V2 logging role and default logging level so that you can monitor progress events logs as it passes from your devices through Iot core service.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotLoggingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotLogging), &result)
	return &result
}
