package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotResourceSpecificLogging = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_level": {
        "computed": true,
        "description": "The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_id": {
        "computed": true,
        "description": "Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_name": {
        "computed": true,
        "description": "The target name.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_type": {
        "computed": true,
        "description": "The target type. Value must be THING_GROUP, CLIENT_ID, SOURCE_IP, PRINCIPAL_ID, or EVENT_TYPE.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoT::ResourceSpecificLogging",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotResourceSpecificLoggingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotResourceSpecificLogging), &result)
	return &result
}
