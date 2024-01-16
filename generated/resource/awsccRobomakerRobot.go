package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRobomakerRobot = `{
  "block": {
    "attributes": {
      "architecture": {
        "description": "The target architecture of the robot.",
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      },
      "greengrass_group_id": {
        "description": "The Greengrass group id.",
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
      "name": {
        "computed": true,
        "description": "The name for the robot.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRobomakerRobotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRobomakerRobot), &result)
	return &result
}
