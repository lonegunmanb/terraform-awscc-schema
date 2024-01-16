package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRobomakerRobotApplication = `{
  "block": {
    "attributes": {
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
      "environment": {
        "computed": true,
        "description": "The URI of the Docker image for the robot application.",
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
        "description": "The name of the robot application.",
        "description_kind": "plain",
        "type": "string"
      },
      "robot_software_suite": {
        "computed": true,
        "description": "The robot software suite used by the robot application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of robot software suite.",
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version of robot software suite.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "sources": {
        "computed": true,
        "description": "The sources of the robot application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "architecture": {
              "computed": true,
              "description": "The architecture of robot application.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_bucket": {
              "computed": true,
              "description": "The Arn of the S3Bucket that stores the robot application source.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description": "The s3 key of robot application source.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
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
    "description": "Data Source schema for AWS::RoboMaker::RobotApplication",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRobomakerRobotApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRobomakerRobotApplication), &result)
	return &result
}
