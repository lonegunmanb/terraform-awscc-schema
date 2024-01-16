package resource

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
        "optional": true,
        "type": "string"
      },
      "environment": {
        "computed": true,
        "description": "The URI of the Docker image for the robot application.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The name of the robot application.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "robot_software_suite": {
        "description": "The robot software suite used by the robot application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description": "The name of robot software suite.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version of robot software suite.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "sources": {
        "computed": true,
        "description": "The sources of the robot application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "architecture": {
              "description": "The architecture of robot application.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_bucket": {
              "description": "The Arn of the S3Bucket that stores the robot application source.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_key": {
              "description": "The s3 key of robot application source.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
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
    "description": "This schema is for testing purpose only.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRobomakerRobotApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRobomakerRobotApplication), &result)
	return &result
}
