package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRobomakerRobotApplicationVersion = `{
  "block": {
    "attributes": {
      "application": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "application_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "current_revision_id": {
        "computed": true,
        "description": "The revision ID of robot application.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::RoboMaker::RobotApplicationVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRobomakerRobotApplicationVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRobomakerRobotApplicationVersion), &result)
	return &result
}
