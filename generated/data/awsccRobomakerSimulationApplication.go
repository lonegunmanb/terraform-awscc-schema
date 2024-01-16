package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRobomakerSimulationApplication = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "current_revision_id": {
        "computed": true,
        "description": "The current revision id.",
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
        "description": "The name of the simulation application.",
        "description_kind": "plain",
        "type": "string"
      },
      "rendering_engine": {
        "computed": true,
        "description": "The rendering engine for the simulation application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of the rendering engine.",
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version of the rendering engine.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "robot_software_suite": {
        "computed": true,
        "description": "The robot software suite used by the simulation application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of the robot software suite.",
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version of the robot software suite.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "simulation_software_suite": {
        "computed": true,
        "description": "The simulation software suite used by the simulation application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of the simulation software suite.",
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version of the simulation software suite.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "sources": {
        "computed": true,
        "description": "The sources of the simulation application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "architecture": {
              "computed": true,
              "description": "The target processor architecture for the application.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_bucket": {
              "computed": true,
              "description": "The Amazon S3 bucket name.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description": "The s3 object key.",
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
    "description": "Data Source schema for AWS::RoboMaker::SimulationApplication",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRobomakerSimulationApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRobomakerSimulationApplication), &result)
	return &result
}
