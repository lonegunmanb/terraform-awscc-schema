package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRobomakerRobotApplicationVersion = `{
  "block": {
    "attributes": {
      "application": {
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "AWS::RoboMaker::RobotApplicationVersion resource creates an AWS RoboMaker RobotApplicationVersion. This helps you control which code your robot uses.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRobomakerRobotApplicationVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRobomakerRobotApplicationVersion), &result)
	return &result
}
