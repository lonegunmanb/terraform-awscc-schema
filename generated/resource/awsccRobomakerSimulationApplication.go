package resource

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
        "description": "The name of the simulation application.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rendering_engine": {
        "computed": true,
        "description": "The rendering engine for the simulation application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description": "The name of the rendering engine.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version": {
              "description": "The version of the rendering engine.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "robot_software_suite": {
        "description": "The robot software suite used by the simulation application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description": "The name of the robot software suite.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version of the robot software suite.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "simulation_software_suite": {
        "description": "The simulation software suite used by the simulation application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description": "The name of the simulation software suite.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version of the simulation software suite.",
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
        "description": "The sources of the simulation application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "architecture": {
              "description": "The target processor architecture for the application.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_bucket": {
              "description": "The Amazon S3 bucket name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_key": {
              "description": "The s3 object key.",
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

func AwsccRobomakerSimulationApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRobomakerSimulationApplication), &result)
	return &result
}
