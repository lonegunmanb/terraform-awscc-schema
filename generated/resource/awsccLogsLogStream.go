package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsLogStream = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_group_name": {
        "description": "The name of the log group where the log stream is created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_stream_name": {
        "computed": true,
        "description": "The name of the log stream. The name must be unique wihtin the log group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Logs::LogStream",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsLogStreamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsLogStream), &result)
	return &result
}
