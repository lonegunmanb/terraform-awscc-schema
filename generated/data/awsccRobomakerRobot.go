package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRobomakerRobot = `{
  "block": {
    "attributes": {
      "architecture": {
        "computed": true,
        "description": "The target architecture of the robot.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fleet": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the fleet.",
        "description_kind": "plain",
        "type": "string"
      },
      "greengrass_group_id": {
        "computed": true,
        "description": "The Greengrass group id.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name for the robot.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::RoboMaker::Robot",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRobomakerRobotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRobomakerRobot), &result)
	return &result
}
