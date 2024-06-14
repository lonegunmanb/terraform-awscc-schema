package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineMonitor = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
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
      "identity_center_application_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_center_instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "monitor_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subdomain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Deadline::Monitor",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDeadlineMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineMonitor), &result)
	return &result
}
